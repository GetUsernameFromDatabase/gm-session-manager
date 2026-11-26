@echo off
cd /d "%~dp0"
gm-session-manager.exe --player
if errorlevel 1 pause
