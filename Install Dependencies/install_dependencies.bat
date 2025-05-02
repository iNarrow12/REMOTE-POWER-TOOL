@echo off
REM ----------------------------------------
REM Remote PC Control Tool — Dependency Installer
REM ----------------------------------------

REM 1) Upgrade pip
echo Upgrading pip...
python -m pip install --upgrade pip

if errorlevel 1 (
  echo Failed to upgrade pip. Please ensure Python is on your PATH.
  pause & exit /b 1
)

REM 2) Install requirements
echo Installing Python packages from requirements.txt...
pip install -r requirements.txt

if errorlevel 1 (
  echo Failed to install some packages.
  pause & exit /b 1
)

REM 3) Confirm nircmd.exe is present
if not exist "%~dp0nircmd.exe" (
  echo WARNING: nircmd.exe not found in this folder.
  echo Download it from https://www.nirsoft.net/utils/nircmd.html and place it here.
  pause
) else (
  echo nircmd.exe found.
)

echo.
echo All set! You can now run the Remote PC Control tool.
pause