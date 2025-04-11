# air.ps1 - Windows runner for Sticky Notes (Go + HTMX)

Write-Host "Building for Windows..."

# Ensure tmp dir exists
if (-not (Test-Path -Path "./tmp")) {
    New-Item -ItemType Directory -Path "./tmp" | Out-Null
}

# Build the binary with .exe extension
go build -o ./tmp/main.exe ./cmd/web

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed."
    exit 1
}

Write-Host "Running with Air..."

# Ensure Air is installed
if (-not (Get-Command "air" -ErrorAction SilentlyContinue)) {
    Write-Host "Air is not installed. Run: go install github.com/air-verse/air@latest"
    exit 1
}

# Run air
air
