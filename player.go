// =============================================================================
// FILE MAP — player.go
// =============================================================================
// WHAT THIS FILE DOES
//   The player: a movable white square. It owns its own position, size, and
//   speed, and knows how to do two things each frame — update itself from the
//   keyboard, and draw itself onto the screen.
//
//   The point of giving the player its own file/type: Game no longer
//   micromanages a square. It just says "Player, update yourself" and
//   "Player, draw yourself." Each unit minds its own data.
//
// WHERE THINGS LIVE (read top-down)
//   PLAYER CONFIG  -> the player's named numbers (size, speed)
//   PLAYER STATE   -> the Player struct: what the player knows right now
//   BEHAVIOR       -> NewPlayer (build) / Center (place) / Update / Draw
// =============================================================================

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// =============================================================================
// PLAYER CONFIG — the player's own magic numbers, named
// =============================================================================

const (
	// The player square's edge length, in logical pixels.
	playerSize = 16

	// How far the player moves per tick (~60 ticks/sec, so ~120 px/sec).
	// Change this one number to make the player faster or slower.
	playerSpeed = 2
)

// =============================================================================
// PLAYER STATE — everything the player knows right now
// =============================================================================

type Player struct {
	image *ebiten.Image // the white square we draw
	x, y  float64       // current top-left position; changes as the player moves
}

// =============================================================================
// BEHAVIOR — build the player, then update/draw it each frame
// =============================================================================

// NewPlayer builds a player (its white square image). It doesn't set a position
// yet — the game calls Center once it knows the real window size (see Layout).
func NewPlayer() *Player {
	img := ebiten.NewImage(playerSize, playerSize)
	img.Fill(color.White)

	return &Player{image: img}
}

// Center places the player in the middle of a screen of the given size. It uses
// math (not a hard-typed coordinate), so it stays correct at any window size.
func (p *Player) Center(screenW, screenH int) {
	p.x = float64(screenW-playerSize) / 2
	p.y = float64(screenH-playerSize) / 2
}

// Update moves the player one tick based on WASD input, then keeps it inside
// the screen. While a key is held, IsKeyPressed is true every tick, so the
// player glides continuously. screenW/screenH are the current window size (from
// Layout), so the square stays in bounds no matter how big the window is.
func (p *Player) Update(screenW, screenH int) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.y -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.y += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x += playerSpeed
	}

	// Keep the square fully on screen. Its top-left corner can sit anywhere
	// from 0 up to (screen size − square size); beyond that the square would
	// poke past an edge, so pull it back to the edge.
	maxX := float64(screenW - playerSize)
	maxY := float64(screenH - playerSize)
	if p.x < 0 {
		p.x = 0
	}
	if p.x > maxX {
		p.x = maxX
	}
	if p.y < 0 {
		p.y = 0
	}
	if p.y > maxY {
		p.y = maxY
	}
}

// Draw paints the player at its current position, once per frame.
func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, op)
}
