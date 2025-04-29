@echo off
setlocal enabledelayedexpansion

echo 📦 Installing ghosthub...

if not exist "build\ghosthub-windows-amd64.exe" (
    echo ❌ Binary not found: build\ghosthub-windows-amd64.exe
    echo Run build.bat first
    exit /b 1
)

if not exist "%USERPROFILE%\AppData\Local\Microsoft\WindowsApps" (
    mkdir "%USERPROFILE%\AppData\Local\Microsoft\WindowsApps"
)

copy /Y "build\ghosthub-windows-amd64.exe" "%USERPROFILE%\AppData\Local\Microsoft\WindowsApps\ghosthub.exe"

echo ✅ Installation completed!
echo 🚀 Use 'ghosthub' to start

echo %PATH% | find /C /I "%USERPROFILE%\AppData\Local\Microsoft\WindowsApps" > nul
if errorlevel 1 (
    echo ⚠️ Directory not in PATH. Add it manually:
    echo %USERPROFILE%\AppData\Local\Microsoft\WindowsApps
)

pause