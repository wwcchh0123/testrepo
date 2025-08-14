package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gravity      = 0.1
)

// Particle represents a single particle in a firework explosion.
type Particle struct {
	x, y   float64
	vx, vy float64
	color  color.Color
	life   int
}

// Firework represents a single firework, from launch to explosion.
type Firework struct {
	x, y      float64
	vx, vy    float64
	exploded  bool
	particles []*Particle
	color     color.Color
}

// NewFirework creates a new firework that will launch from the bottom of the screen.
func NewFirework(targetX float64) *Firework {
	return &Firework{
		x:        targetX,
		y:        screenHeight,
		vx:       (rand.Float64() - 0.5) * 2,
		vy:       -8 - rand.Float64()*4,
		exploded: false,
		color:    color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 0xff},
	}
}

// Update updates the firework's state.
func (f *Firework) Update() {
	if !f.exploded {
		f.x += f.vx
		f.y += f.vy
		f.vy += gravity

		if f.vy >= 0 {
			f.exploded = true
			for i := 0; i < 100; i++ {
				angle := rand.Float64() * 2 * math.Pi
				speed := rand.Float64() * 4
				f.particles = append(f.particles, &Particle{
					x:    f.x,
					y:    f.y,
					vx:   math.Cos(angle) * speed,
					vy:   math.Sin(angle) * speed,
					life: 100,
				})
			}
		}
	} else {
		for i := len(f.particles) - 1; i >= 0; i-- {
			p := f.particles[i]
			p.x += p.vx
			p.y += p.vy
			p.vy += gravity
			p.life--
			if p.life <= 0 {
				f.particles = append(f.particles[:i], f.particles[i+1:]...)
			}
		}
	}
}

// Draw draws the firework.
func (f *Firework) Draw(screen *ebiten.Image) {
	if !f.exploded {
		screen.Set(int(f.x), int(f.y), f.color)
	} else {
		for _, p := range f.particles {
			alpha := uint8(float64(p.life) / 100 * 255)
			r, g, b, _ := f.color.RGBA()
			c := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
			screen.Set(int(p.x), int(p.y), c)
		}
	}
}

// Game holds the state of the firework simulation.
type Game struct {
	fireworks []*Firework
}

// NewGame creates a new game instance.
func NewGame() *Game {
	return &Game{}
}

// Update updates the game state.
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, _ := ebiten.CursorPosition()
		g.fireworks = append(g.fireworks, NewFirework(float64(x)))
	}

	for i := len(g.fireworks) - 1; i >= 0; i-- {
		f := g.fireworks[i]
		f.Update()
		if f.exploded && len(f.particles) == 0 {
			g.fireworks = append(g.fireworks[:i], g.fireworks[i+1:]...)
		}
	}
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff}) // Black background
	for _, f := range g.fireworks {
		f.Draw(screen)
	}
}

// Layout returns the screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fireworks")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}