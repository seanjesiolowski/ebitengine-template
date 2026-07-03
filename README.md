# firstgame

A 2D game built in **Go** with the **[Ebiten](https://ebitengine.org/)** engine.
Right now it opens a window and draws a white square in the center, plus some
debug text — the foundation to build a real game on.

## Run it

```bash
go run .
```

A 640×480 window titled "first ebiten game" should open. Close the window to quit.

## Build a standalone executable

```bash
go build -o firstgame.exe .
```

## Run it in the browser (WebAssembly)

The same game code also runs in a browser — WASM is just a build target, not a
rewrite. One command builds it and serves it locally:

```powershell
.\serve-web.ps1
```

This compiles the game to `web/firstgame.wasm`, opens `http://localhost:8080`,
and serves it. Press **Ctrl+C** to stop the server.

Why a server at all? Browsers refuse to load `.wasm` from a `file://` path, so
the game must be served over `http://`. That's what `devserver/` provides.

## How the code is laid out

Everything lives in `main.go`, organized top-down so you can grab the big picture
before the details:

| Section      | What's there                                              |
| ------------ | --------------------------------------------------------- |
| `FILE MAP`   | Header comment — what the program does, and the mental model |
| `CONFIG`     | Named numbers: screen size, square size, centered position |
| `GAME STATE` | The `Game` struct — everything the game knows right now    |
| `GAME LOOP`  | `Update` (logic) → `Draw` (pixels) → `Layout` (resolution) |
| `STARTUP`    | `main()` — set up the window, build the game, run the loop |

### The Ebiten loop, in one breath

You hand Ebiten a `Game`; it runs a loop ~60 times per second:
**Update** (change state) → **Draw** (paint the screen). `Layout` tells it the
logical resolution to draw at. You write those methods; Ebiten calls them.

## Files

| File               | Purpose                                             |
| ------------------ | --------------------------------------------------- |
| `main.go`          | Stage manager: window + game loop                   |
| `player.go`        | The movable square (its data + how it moves/draws)  |
| `ebiten.png`       | Window icon (embedded into the build)               |
| `go.mod` / `go.sum`| Dependencies (Ebiten)                               |
| `serve-web.ps1`    | Build for the browser and serve it locally          |
| `web/index.html`   | The page that loads and runs the game in a browser  |
| `devserver/`       | Tiny local web server (dev only, its own package)   |

## Project conventions

See [`CLAUDE.md`](CLAUDE.md) for the code-style rules this project follows
(map-before-detail, visual sections, no magic numbers, small focused files).
