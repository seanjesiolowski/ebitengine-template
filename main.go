// =============================================================================
// FILE MAP — firstgame (Ebiten)
// =============================================================================
// WHAT THIS PROGRAM DOES
//   Opens a window and runs the game loop. The thing you move with W/A/S/D
//   lives in its own file now — see player.go. This file owns the screen/window
//   setup and hands work off to the player.
//
// HOW EBITEN WORKS (the mental model)
//   You hand Ebiten a "Game" value, then it runs a loop ~60 times per second:
//       1. Update()  -> change game state (movement, input, physics)   [logic]
//       2. Draw()    -> paint the current state onto the screen         [pixels]
//       3. Layout()  -> report the logical resolution to draw at        [size]
//   You write those three methods; Ebiten calls them for you.
//
// WHERE THINGS LIVE (read top-down)
//   CONFIG      -> the window's starting size (the game's area follows it)
//   GAME STATE  -> the Game struct: everything the game "knows" right now
//   GAME LOOP   -> Update / Draw / Layout (delegates to the player)
//   STARTUP     -> main(): set up the window, build the game, run it
//   player.go   -> the player itself (its data + how it moves and draws)
//
// Tutorial reference: https://ebitengine.org/en/tour/hello_world.html
// =============================================================================

package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png" // registers the PNG decoder so image.Decode understands .png
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// =============================================================================
// CONFIG — all the "magic numbers" live here, named, so the rest reads clearly
// =============================================================================

const (
	// The window's starting size, in pixels. The game's drawing area follows
	// the window (see Layout), so there are no black bars: resize or maximize
	// the window — or open it in a browser — and the play area grows to fill it.
	windowWidth  = 640
	windowHeight = 480
)

// ebiten.png is baked into the .exe at build time and used as the window icon.
//
//go:embed ebiten.png
var iconBytes []byte

// =============================================================================
// GAME STATE — everything the game knows. Add fields here as the game grows.
// =============================================================================

type Game struct {
	player *Player // the thing you move with W/A/S/D (see player.go)

	// The current drawing size, in logical pixels. Layout keeps these equal to
	// the window's size so the game always fills the window. `centered` records
	// whether we've placed the player yet — we can only center it once we know
	// the real size, which Layout is the first to tell us.
	width, height int
	centered      bool
}

// =============================================================================
// GAME LOOP — the three methods Ebiten calls for you, ~60 times per second
// =============================================================================

// Update advances the game logic by one tick. The Game doesn't move anything
// itself — it just tells the player to update itself, handing over the current
// window size so the player can keep itself in bounds.
func (g *Game) Update() error {
	g.player.Update(g.width, g.height)
	return nil
}

// Draw paints the current state onto the screen, once per frame. Same idea:
// the Game draws its text, then asks the player to draw itself.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Move the square-player with WASD.")
	g.player.Draw(screen)
}

// Layout makes the game's drawing area match the window's size, so the game
// fills the window with no letterbox bars (Ebiten isn't scaling a fixed-size
// picture anymore). The size can change whenever the window is resized (or the
// browser tab is), so we also center the player the first time we're told the
// real size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.width, g.height = outsideWidth, outsideHeight

	if !g.centered {
		g.player.Center(g.width, g.height)
		g.centered = true
	}

	return g.width, g.height
}

// =============================================================================
// STARTUP — set up the window, build the game, then run the loop
// =============================================================================

func main() {
	// 1. Configure the window.
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("first ebiten game")
	ebiten.SetWindowIcon([]image.Image{loadIcon()})
	// Let the desktop window be resized/maximized; the game fills whatever size
	// it becomes (Layout handles that). This has no effect in the browser.
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// 2. Build the game's starting state. The player has no position yet — the
	//    game centers it in Layout, once it knows the real window size.
	g := &Game{
		player: NewPlayer(),
	}

	// 3. Hand the game to Ebiten and run until the window closes.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// loadIcon decodes the embedded ebiten.png into an image for use as the
// window icon. If it can't be decoded, the program can't start, so we stop.
func loadIcon() image.Image {
	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		log.Fatal(err)
	}
	return icon
}
