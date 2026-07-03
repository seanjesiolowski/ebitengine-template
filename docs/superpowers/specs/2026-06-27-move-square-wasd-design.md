# Move the square with WASD — design

**Date:** 2026-06-27
**Status:** Approved
**File touched:** `main.go` only

## Goal

Make the white square move with the keyboard. This is the first feature where
`Update()` does real work and where the square's position becomes *state the game
remembers and changes*, rather than a fixed constant. It teaches the core
input → state → draw cycle.

## Decisions

- **Keys:** WASD only (W up, S down, A left, D right).
- **Edges:** the square may slide off-screen and disappear. Clamping is a
  deliberately separate next step, so this change stays focused on movement alone.
- **Speed:** `moveSpeed = 2` pixels per tick (~120 px/sec at 60 ticks/sec).
- **Position type:** `float64`, because `GeoM.Translate` takes float64 and it
  leaves room for fractional/diagonal speeds later without a type change.

## Changes to `main.go`

1. **CONFIG** — rename `squareX`/`squareY` to `startX`/`startY` (now just the
   *starting* spot). Add `moveSpeed = 2`.

2. **GAME STATE** — add live position to `Game`:
   ```go
   type Game struct {
       square *ebiten.Image
       x, y   float64 // current top-left position; changes as the player moves
   }
   ```

3. **GAME LOOP — `Update()`** — read keys, adjust position:
   ```go
   if ebiten.IsKeyPressed(ebiten.KeyW) { g.y -= moveSpeed }
   if ebiten.IsKeyPressed(ebiten.KeyS) { g.y += moveSpeed }
   if ebiten.IsKeyPressed(ebiten.KeyA) { g.x -= moveSpeed }
   if ebiten.IsKeyPressed(ebiten.KeyD) { g.x += moveSpeed }
   ```
   Holding a key moves continuously. No new imports — `IsKeyPressed` and the
   `Key*` constants are in the already-imported `ebiten` package.

4. **`Draw()`** — translate by `g.x, g.y` instead of the old constants.

5. **`main()`** — set the start: `x: startX, y: startY`.

## Verification

After the change: `go fmt ./...`, then `go vet ./...`, then `go run .` and confirm
the square starts centered and glides with W/A/S/D.
