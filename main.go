package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	particleSize = 2
)

// Particle represents a single point in the firework explosion.
type Particle struct {
	x, y     float64
	vx, vy   float64
	lifespan int
	color    color.Color
}

// Firework is the main rocket that explodes.
type Firework struct {
	x, y     float64
	vx, vy   float64
	exploded bool
	lifespan int
	color    color.Color
}

// Game holds the state of the simulation.
type Game struct {
	fireworks []*Firework
	particles []*Particle
}

// NewGame initializes the game state.
func NewGame() *Game {
	return &Game{
		fireworks: []*Firework{},
		particles: []*Particle{},
	}
}

// Update proceeds the game state.
func (g *Game) Update() error {
	// Create new fireworks periodically
	if rand.Intn(30) == 0 {
		g.fireworks = append(g.fireworks, createFirework())
	}

	// Update fireworks
	for i := len(g.fireworks) - 1; i >= 0; i-- {
		f := g.fireworks[i]
		if !f.exploded {
			f.vy += 0.1 // gravity
			f.x += f.vx
			f.y += f.vy

			if f.vy >= 0 { // Explode at the apex of its trajectory
				f.exploded = true
				for j := 0; j < 100; j++ {
					g.particles = append(g.particles, createParticle(f.x, f.y, f.color))
				}
				g.fireworks = append(g.fireworks[:i], g.fireworks[i+1:]...)
			}
		}
	}

	// Update particles
	for i := len(g.particles) - 1; i >= 0; i-- {
		p := g.particles[i]
		p.vx *= 0.98 // friction
		p.vy += 0.08 // gravity
		p.x += p.vx
		p.y += p.vy
		p.lifespan--

		if p.lifespan <= 0 {
			g.particles = append(g.particles[:i], g.particles[i+1:]...)
		}
	}

	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Darken the screen slightly each frame for a trailing effect
	screen.Fill(color.RGBA{0, 0, 0, 40})

	// Draw fireworks
	for _, f := range g.fireworks {
		ebitenutil.DrawRect(screen, f.x, f.y, particleSize, particleSize, f.color)
	}

	// Draw particles
	for _, p := range g.particles {
		alpha := uint8(math.Max(0, float64(p.lifespan)*255/100))
		c := p.color.(color.RGBA)
		c.A = alpha
		ebitenutil.DrawRect(screen, p.x, p.y, particleSize, particleSize, c)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
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

// createFirework creates a new firework rocket.
func createFirework() *Firework {
	return &Firework{
		x:      rand.Float64() * screenWidth,
		y:      screenHeight,
		vx:     rand.Float64()*4 - 2,
		vy:     -rand.Float64()*5 - 8,
		exploded: false,
		color:  randomColor(),
	}
}

// createParticle creates a new explosion particle.
func createParticle(x, y float64, c color.Color) *Particle {
	angle := rand.Float64() * 2 * math.Pi
	speed := rand.Float64() * 4
	return &Particle{
		x:        x,
		y:        y,
		vx:       math.Cos(angle) * speed,
		vy:       math.Sin(angle) * speed,
		lifespan: rand.Intn(60) + 60, // Lasts 1 to 2 seconds
		color:    c,
	}
}

// randomColor generates a random bright color.
func randomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Intn(128) + 127),
		G: uint8(rand.Intn(128) + 127),
		B: uint8(rand.Intn(128) + 127),
		A: 255,
	}
}