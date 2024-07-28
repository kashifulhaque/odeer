@echo off

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% neq 0 (
  echo Go is not installed. Please install Go and try again.
  exit /b 1
)

REM Get OS name (always "windows" for Windows)
set OS_NAME=windows

REM Get architecture
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
  set OS_ARCH=amd64
) else if "%PROCESSOR_ARCHITECTURE%"=="ARM64" (
  set OS_ARCH=arm64
) else (
  echo Unsupported architecture: %PROCESSOR_ARCHITECTURE%
  exit /b 1
)

REM Build the Go program
set GOOS=%OS_NAME%
set GOARCH=%OS_ARCH%
go build -o odeer.exe cmd\odeer\main.go

echo Build complete. Output: odeer.exe