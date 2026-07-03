// =============================================================================
// FILE MAP — devserver
// =============================================================================
// WHAT THIS PROGRAM DOES
//   Serves the web/ folder over http://localhost so a browser can load and run
//   the game. Browsers refuse to run .wasm loaded from a file:// path, so we
//   need a real (if tiny) web server. This is a *development* tool, not part of
//   the game — that's why it lives in its own folder / package.
//
// WHY ITS OWN FOLDER
//   The game is `package main` (main.go). This is a second, separate `main`.
//   Two mains can't share one folder, so this gets its own. Run it with:
//       go run ./devserver
//
// WHERE THINGS LIVE (read top-down)
//   CONFIG  -> the address to listen on and the folder to serve
//   SERVE   -> wire a file server to that folder, fixing the .wasm MIME type
//   STARTUP -> main(): print the URL, then serve until you stop it (Ctrl+C)
// =============================================================================

package main

import (
	"log"
	"net/http"
	"strings"
)

// =============================================================================
// CONFIG — where we listen and what we serve
// =============================================================================

const (
	listenAddr = "localhost:8080" // the address you open in the browser
	webDir     = "web"            // the folder holding index.html + the .wasm
)

// =============================================================================
// SERVE — a plain file server, with one fix for WebAssembly
// =============================================================================

// wasmMIME wraps a handler so that .wasm files are served with the exact
// content type browsers require ("application/wasm"). Without this,
// WebAssembly.instantiateStreaming in index.html refuses to start the game.
func wasmMIME(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		next.ServeHTTP(w, r)
	})
}

// =============================================================================
// STARTUP — announce the URL, then serve until stopped
// =============================================================================

func main() {
	files := http.FileServer(http.Dir(webDir))

	log.Printf("serving %q at http://%s  (press Ctrl+C to stop)", webDir, listenAddr)
	if err := http.ListenAndServe(listenAddr, wasmMIME(files)); err != nil {
		log.Fatal(err)
	}
}
