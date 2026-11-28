#!/usr/bin/perl
# Normalizes SQL from stdin so comment-only and whitespace-only edits compare equal (update.sh's migration gate).
# Strings and "quoted identifiers" come first in the pattern, so a -- or /* inside one stays.
use strict;
use warnings;

my $sql = do { local $/; <STDIN> } // '';
# A run of whitespace and comments is one space, so a removed comment leaves no double space.
$sql =~ s{('(?:[^']|'')*'|"(?:[^"]|"")*")|(?:\s|/\*.*?\*/|--[^\n]*)+}{defined $1 ? $1 : ' '}gse;
$sql =~ s/^ | $//g;
print $sql, "\n";
