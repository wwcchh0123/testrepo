# Jump Jump Game

A fun platform jumping game built with Go and the Ebiten game engine.

## 📝 Description

This is a "Jump Jump" style platform game where players control a character that jumps between platforms. The goal is to jump as high as possible while collecting points and avoiding falling off the screen.

## 🎮 Game Features

- **Physics-based jumping**: Hold down mouse to charge jump power, release to jump
- **Wall bouncing**: Character bounces off screen edges
- **Special platforms**: Different colored platforms with bonus scoring
  - 🎵 Pink platforms (Music Box): 30 points
  - 🏪 Green platforms (Convenience Store): 15 points  
  - 🧩 Yellow platforms (Rubiks Cube): 10 points
  - 🕳️ Brown platforms (Manhole): 5 points
- **Combo system**: Landing near platform centers increases combo multiplier
- **Smooth camera**: Camera follows player movement
- **Game over and restart**: Press 'R' to restart when game ends

## 🚀 Getting Started

### Prerequisites

- Go 1.23.3 or later
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

2. Install dependencies:
```bash
go mod download
```

3. Run the game:
```bash
go run .
```

## 🎮 How to Play

1. **Launch the game** by running `go run .`
2. **Charge your jump** by holding down the left mouse button
3. **Release** to jump - longer charge = more powerful jump
4. **Aim for platform centers** to build up combo multipliers
5. **Try to reach special colored platforms** for bonus points
6. **Don't fall off the screen!**
7. **Press 'R'** to restart when game over

## 🏗️ Project Structure

```
.
├── main.go              # Main game logic and engine
├── embed.go             # Embedded image assets
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── images/              # Game sprites and assets
│   ├── player.png       # Player character sprite
│   ├── bullet.png       # Bullet sprite (unused)
│   └── enemy.png        # Enemy sprite (unused)
├── testfile/            # Test files
└── .github/             # GitHub workflows
    └── workflows/
        ├── claude.yml   # Claude AI workflow
        └── ci.yml       # CI/CD pipeline
```

## 🛠️ Technical Details

- **Language**: Go 1.23.3
- **Game Engine**: [Ebiten v2.8.8](https://github.com/hajimehoshi/ebiten)
- **Window Size**: 480x640 pixels
- **Dependencies**:
  - `github.com/hajimehoshi/ebiten/v2` - 2D game engine
  - `golang.org/x/image` - Image processing utilities

## 🎯 Game Mechanics

### Player Physics
- **Gravity**: 0.4 units downward acceleration
- **Charge Rate**: 0.1 per frame when holding mouse
- **Max Charge**: 15.0 units
- **Jump Angle**: 60 degrees for optimal trajectory

### Platform Generation
- Platforms are randomly positioned
- 20% chance for special colored platforms
- Minimum 80-pixel vertical spacing between platforms
- New platforms generate as player progresses upward

### Scoring System
- Base score: 1 point per platform
- Combo multiplier: x2 points when landing near center
- Special platform bonuses (see features section)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source. Feel free to use and modify as needed.

## 🎮 Screenshots

The game features a colorful jumping character navigating between platforms of different colors, each offering unique scoring opportunities.

---

**Have fun jumping!** 🦘✨