@echo off
setlocal

:: Paths
set "PYTHONW=C:\Users\inarr\AppData\Local\Programs\Python\Python313\pythonw.exe"
set "SCRIPT=C:\Users\inarr\OneDrive\Desktop\Remote Power Tool\RemotePowerTool.pyw"
set "WORKDIR=C:\Users\inarr\OneDrive\Desktop\Remote Power Tool\"
set "SHORTCUT=%AppData%\Microsoft\Windows\Start Menu\Programs\Startup\RemotePowerTool.lnk"

:: VBScript to create the shortcut
set "VBS=%temp%\MakeShortcut.vbs"
echo Set WshShell = CreateObject("WScript.Shell") > "%VBS%"
echo Set Shortcut = WshShell.CreateShortcut("%SHORTCUT%") >> "%VBS%"
echo Shortcut.TargetPath = "%PYTHONW%" >> "%VBS%"
echo Shortcut.Arguments = Chr(34) ^& "%SCRIPT%" ^& Chr(34) >> "%VBS%"
echo Shortcut.WorkingDirectory = "%WORKDIR%" >> "%VBS%"
echo Shortcut.WindowStyle = 7 >> "%VBS%"
echo Shortcut.Save >> "%VBS%"

:: Create shortcut
cscript //nologo "%VBS%"
del "%VBS%"

echo:
echo ✅ Shortcut created to auto-run your script on login!
echo Shortcut: %SHORTCUT%
pause
