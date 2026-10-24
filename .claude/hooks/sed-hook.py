#!/usr/bin/env python3
"""
Claude Code PreToolUse hook (matcher: Bash).

Auto-approves a command only when it is clearly safe:
  - it contains a sed, and
  - it is a single command or a plain | pipeline, and
  - every part is either a read-only sed or a read-only basic command.

Anything else: no output, so Claude Code does what it would do without the hook.
"""

import json
import re
import shlex
import sys

READONLY = {"cat", "head", "tail", "wc", "grep", "rg", "cut", "tr", "nl", "echo"}
# sed script commands that write files (w, W) or run a shell (e), incl. s///w and s///e
WRITEY = re.compile(r"(^|[;{}\n]|[0-9$/]\s*|\s)[wWe](\s|$)|/[gpIiMm0-9]*[we](\s|$)")


def safe_sed(args):
    scripts, files = [], []
    take_script = False
    for a in args:
        if take_script:
            scripts.append(a)
            take_script = False
        elif a in ("-e", "--expression"):
            take_script = True
        elif a in ("-n", "-E", "-r", "--quiet", "--silent", "-z", "-s", "-u"):
            pass
        elif a.startswith("-"):
            return False  # -i, -f, --in-place, combined flags...: not clearly safe
        else:
            files.append(a)
    if not scripts and files:
        scripts.append(files[0])
    return bool(scripts) and not any(WRITEY.search(s) for s in scripts)


def clearly_safe(cmd):
    cmd = re.sub(r"\s2>(&1|/dev/null)", "", cmd)
    if re.search(r"[;&<>`]|\$\(|\|\|", cmd):  # chaining, redirects, substitution
        return False
    parts = cmd.split("|")
    saw_sed = False
    for part in parts:
        try:
            words = shlex.split(part)
        except ValueError:
            return False
        if not words:
            return False
        if words[0] == "sed":
            if not safe_sed(words[1:]):
                return False
            saw_sed = True
        elif words[0] not in READONLY:
            return False
    return saw_sed


data = json.load(sys.stdin)
cmd = (data.get("tool_input") or {}).get("command", "")
if clearly_safe(cmd):
    print(
        json.dumps(
            {
                "hookSpecificOutput": {
                    "hookEventName": "PreToolUse",
                    "permissionDecision": "allow",
                    "permissionDecisionReason": "read-only sed",
                }
            }
        )
    )
