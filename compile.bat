@echo off
setlocal enabledelayedexpansion

:: ============================================================
:: compile.bat - Windows equivalent of compile.sh
:: Requires: Go, Git, certutil (built-in on all Windows)
:: Usage: compile.bat <File\or\Folder\To\Compile> <VERSION>
:: Example: compile.bat .\cmd\balafetch\ 0.5.0
:: ============================================================

if "%~2"=="" (
    echo Usage: %~nx0 ^<File\or\Folder\To\Compile^> ^<VERSION^>
    echo Note: use .\\ for a folder
    exit /b 1
)

set "APP_NAME=balafetch"
set "COMPILE_PATH_TARGET=%~1"
set "VERSION=%~2"
set "TARGET_FOLDER=dist"
set "CHECKSUMS_FILE=%TARGET_FOLDER%\checksums.txt"
set "DATA_MODULE_PATH=github.com/gitmobkab/balafetch/internal/data"
set /a CURRENT=1
set /a TOTAL=10

for /f "delims=" %%i in ('git rev-parse --short HEAD') do set "COMMIT_HASH=%%i"
if errorlevel 1 ( echo ERROR: failed to get git commit hash & exit /b 1 )

:: ISO 8601 UTC build time via pwsh (PowerShell 7)
for /f "delims=" %%i in ('pwsh -NoProfile -Command "(Get-Date).ToUniversalTime().ToString(\"yyyy-MM-ddTHH:mm:ssZ\")"') do set "BUILD_TIME=%%i"
if "%BUILD_TIME%"=="" ( echo ERROR: failed to get build time - is pwsh installed? & exit /b 1 )

call :InfoLog "Creating %TARGET_FOLDER% folder"
if not exist "%TARGET_FOLDER%" mkdir "%TARGET_FOLDER%"

call :InfoLog "Cleaning %CHECKSUMS_FILE%"
type nul > "%CHECKSUMS_FILE%"

call :InfoLog "=== Starting compilation of %COMPILE_PATH_TARGET% with version %VERSION% ==="
echo.

call :Build windows amd64 || exit /b 1
call :Build windows arm64 || exit /b 1
call :Build linux   amd64 || exit /b 1
call :Build linux   arm64 || exit /b 1
call :Build linux   arm   || exit /b 1
call :Build darwin  amd64 || exit /b 1
call :Build darwin  arm64 || exit /b 1
call :Build freebsd arm64 || exit /b 1
call :Build freebsd amd64 || exit /b 1
call :Build freebsd arm   || exit /b 1

call :SuccessLog "Compilation to all platforms successful"
call :SuccessLog "All checksums have been written, check %CHECKSUMS_FILE%"
exit /b 0


:: ============================================================
:Build  %1=GOOS  %2=GOARCH
:: ============================================================
set "GOOS=%~1"
set "GOARCH=%~2"
set "HASH="

set "OUTPUT_FILE=%APP_NAME%-%GOOS%-%GOARCH%"
if "%GOOS%"=="windows" set "OUTPUT_FILE=%OUTPUT_FILE%.exe"

call :InfoLog "Compiling %OUTPUT_FILE% from %COMPILE_PATH_TARGET%"

go build -ldflags "-X %DATA_MODULE_PATH%.Version=%VERSION% -X %DATA_MODULE_PATH%.CommitHash=%COMMIT_HASH% -X %DATA_MODULE_PATH%.BuildTime=%BUILD_TIME%" -o "%TARGET_FOLDER%\%OUTPUT_FILE%" %COMPILE_PATH_TARGET%
if errorlevel 1 (
    echo ERROR: build failed for %GOOS%/%GOARCH%
    exit /b 1
)

call :SuccessLog "Compiling Done (%CURRENT%/%TOTAL%)"
set /a CURRENT+=1

:: SHA256 via certutil (built-in on all Windows).
:: skip=1 skips the "SHA256 hash of..." header; "if not defined" stops after the hash line.
:: Format: "<hash>  <filename>" (two spaces, text-mode) ? matches ChecksumValidator in go-selfupdate.
for /f "skip=1 delims=" %%H in ('certutil -hashfile "%TARGET_FOLDER%\%OUTPUT_FILE%" SHA256') do (
    if not defined HASH set "HASH=%%H"
)
set "HASH=%HASH: =%"
set "CHECKSUM_LINE=%HASH%  %OUTPUT_FILE%"
echo %CHECKSUM_LINE%
echo ________________
>>"%CHECKSUMS_FILE%" echo %CHECKSUM_LINE%
exit /b 0


:: ============================================================
:InfoLog
echo [INFO]    %~1
exit /b 0

:SuccessLog
echo [SUCCESS] %~1
exit /b 0
