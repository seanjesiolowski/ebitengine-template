# =============================================================================
# serve-web.ps1 — build the game for the browser, then serve it locally
# =============================================================================
# WHAT THIS DOES (three steps, top-down)
#   1. BUILD  -> compile the game to web/firstgame.wasm (target: js/wasm)
#   2. SYNC   -> copy Go's wasm_exec.js so it matches this exact Go version
#   3. SERVE  -> start the local server and open the game in your browser
#
# USAGE
#   .\serve-web.ps1          then leave it running; Ctrl+C stops the server.
#
# NOTE
#   Desktop still works unchanged: `go run .` opens the native window.
# =============================================================================

$ErrorActionPreference = "Stop"

# --- 1. BUILD: compile to WebAssembly ---------------------------------------
# GOOS/GOARCH pick the browser target. We scope them to this build only so the
# rest of your shell keeps building for desktop normally.
Write-Host "Building web/firstgame.wasm ..." -ForegroundColor Cyan
$env:GOOS = "js"
$env:GOARCH = "wasm"
try {
    go build -o web/firstgame.wasm .
}
finally {
    Remove-Item Env:\GOOS, Env:\GOARCH
}

# --- 2. SYNC: copy Go's runtime glue so it matches this Go version ----------
$wasmExec = Join-Path (go env GOROOT) "lib\wasm\wasm_exec.js"
Copy-Item $wasmExec "web\wasm_exec.js" -Force
Write-Host "Synced wasm_exec.js" -ForegroundColor Cyan

# --- 3. SERVE: open the browser, then run the local server ------------------
Start-Process "http://localhost:8080"
Write-Host "Opening http://localhost:8080 ..." -ForegroundColor Green
go run ./devserver
