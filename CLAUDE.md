# CLAUDE.md — guide for Claude Code

This file is read at the start of every session. It keeps building straightforward
and low-error by stating how this project is built, structured, and styled.

## What this is

`firstgame` — a 2D game built in **Go** with the **Ebiten** engine
(`github.com/hajimehoshi/ebiten/v2`). Currently: a window with a centered white
square + debug text. The owner is **learning Ebiten/Go**, so prefer clear,
explained code over clever code.

## Commands

```bash
go run .              # build + run the game (a window opens)
go build -o firstgame.exe .   # produce the standalone .exe
go vet ./...          # catch likely mistakes
go fmt ./...          # format (run before finishing any change)
```

## Architecture (the one-screen map)

Two files, each organized top-down into labeled sections. `main.go` is the stage
manager (screen/window + the game loop); it delegates the moving thing to the
player, which lives in its own file.

```
main.go — the stage manager
  FILE MAP    header comment — the macro picture before any detail
  CONFIG      named constants for the screen + window
  GAME STATE  the Game struct (just holds a *Player) — what the game knows now
  GAME LOOP   Update / Draw / Layout — delegates to the player
  STARTUP     main() — set up window, build game (NewPlayer), run loop

player.go — the player (a movable square)
  FILE MAP       header comment
  PLAYER CONFIG  the player's own constants (size, speed)
  PLAYER STATE   the Player struct (image, x, y)
  BEHAVIOR       NewPlayer (build + center) / Update (move) / Draw (paint)
```

Ebiten's model: you give it a `Game`; it calls `Update` → `Draw` ~60x/sec, and
`Layout` for the logical resolution. `Game` talks to the player only through
`NewPlayer` / `Update` / `Draw` — the player's whole insides stay private to it.

## Code style (matches the owner's way of reading — follow these)

The owner is a **visual-spatial, top-down thinker**: macro-map first, then detail.
Write code the same way.

1. **Map before detail.** Every file starts with a header comment explaining what
   it does and where things live. Reading the top should give the whole shape.
2. **Visual sections.** Group code into labeled sections with `// ===` dividers so
   the eye can jump to a region instead of reading line-by-line.
3. **No magic numbers.** Any meaningful number is a named constant, grouped in a
   CONFIG section near the concept it belongs to (screen/window in `main.go`,
   player size/speed in `player.go`). Derive values with math (e.g. centering)
   instead of hand-typing results.
4. **Small, focused files.** Keep `main.go` lean. When a clear new concept appears
   (input, a player, a level), give it its own file rather than growing one blob.
   Don't pre-split before there's a real concept to hold.
5. **Explain the why.** Comments say *why* and what a section's role is — the owner
   is learning. Avoid noise that just restates the code.
6. **Behavior stays predictable.** Don't change game behavior while refactoring.
   After any change: `go fmt ./...`, then `go vet ./...`, then `go run .` to confirm.

## When adding a new concept (the growth path)

Today everything fits in `main.go`. As the game grows, split by concept into its
own file in this package (still `package main`), each with its own FILE MAP header:
e.g. `player.go`, `input.go`, `level.go`. Revisit a package/folder structure only
when files get large — not before.
