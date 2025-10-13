# Jump Jump Game

[English](#english) | [中文](#中文)

## English

A challenging platform jumping game built with Go and the Ebiten game engine, inspired by the popular "Jump Jump" mobile game mechanics.

### Game Description

This is an engaging physics-based jumping game where players control a character jumping between various platforms to achieve the highest possible score. The game features realistic physics, combo systems, and special platforms with unique scoring mechanics.

### Features

- 🎮 **Physics Engine**: Realistic gravity and jumping physics simulation
- 🎯 **Charge Jump**: Hold left mouse button to charge power, release to jump
- 🏆 **Combo System**: Land precisely in platform center for combo bonuses
- 🌈 **Special Platforms**: Multiple colored special platforms with different score rewards
- 🎨 **Smooth Animation**: Fluid character animation and camera following
- 📱 **Cross-Platform**: Supports multiple operating systems

### Special Platform Types

- **Music Box Platform** (Pink): +50 points
- **Convenience Store Platform** (Green): +15 points  
- **Rubik's Cube Platform** (Yellow): +20 points
- **Manhole Platform** (Brown): +10 points

### Scoring System

- Basic jump: +1 point
- Center landing (combo zone): +50 points
- Special platforms: Additional points based on platform type

### System Requirements

- Go 1.23.3 or higher
- OpenGL 2.1 or higher graphics support

### Quick Start

#### Clone Repository
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

#### Install Dependencies
```bash
go mod download
```

#### Run Game
```bash
go run .
```

### Game Controls

- **Mouse Left Click**: Hold to charge jump power, release to jump
- **R Key**: Restart game (when game over)

### Project Structure

```
.
├── main.go           # Main game logic
├── embed.go          # Game asset embedding
├── go.mod            # Go module dependencies
├── go.sum            # Module verification file
├── images/           # Game image assets
│   ├── player.png    # Player character sprite
│   ├── bullet.png    # Bullet sprite (reserved)
│   └── enemy.png     # Enemy sprite (reserved)
└── testfile/         # Test files
    └── test.go
```

### Technology Stack

#### Core Technologies
- **Go**: Main programming language
- **Ebiten**: 2D game engine
- **go:embed**: Static asset embedding
- **golang.org/x/image**: Image processing

#### Key Components
- **Player**: Player character control and physics state
- **Platform**: Platform collision detection and special effects
- **Game**: Game state management and main loop
- **Physics**: Gravity, jumping, and collision physics

### Development & Building

#### Development Environment Setup
```bash
# Ensure Go version
go version

# Install dependencies
go mod tidy

# Run development version
go run .
```

#### Building for Distribution
```bash
# Build current platform
go build -o jump-game

# Cross-compile for Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# Cross-compile for macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

#### Building Web Version (Experimental)
```bash
# Requires Ebiten Web support
GOOS=js GOARCH=wasm go build -o game.wasm
```

### Game Screenshots

The game features beautiful pixel art graphics and smooth fluid animations. The character automatically turns and flips direction while bouncing between platforms, creating an engaging visual experience.

### Contributing

Welcome contributions! Please follow these steps:

1. Fork this repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add some amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Create Pull Request

### License

This project is based on MIT License. See [LICENSE](LICENSE) file for details.

### References

- [Ebiten](https://ebiten.org/) - Excellent Go 2D game engine
- [Go Community](https://golang.org/) - Powerful programming language

### Contact

If you have questions or suggestions, please contact us through GitHub Issues.

---

## 中文

一个使用 Go 语言和 Ebiten 游戏引擎开发的跳跃平台游戏，类似于流行的"跳一跳"手机游戏玩法。

### 游戏介绍

这是一个充满挑战性的物理跳跃游戏，玩家需要控制角色在各种平台间跳跃，获取尽可能高的分数。游戏具有真实的物理引擎、连击系统和特殊平台等丰富玩法。

### 游戏特性

- 🎮 **物理引擎**: 真实的重力和跳跃物理模拟
- 🎯 **蓄力跳跃**: 按住鼠标左键蓄力，释放后跳跃
- 🏆 **连击系统**: 精准落在平台中心可获得连击加分
- 🌈 **特殊平台**: 多种颜色的特殊平台，提供不同的分数奖励
- 🎨 **流畅动画**: 流畅的角色动画和镜头跟随效果
- 📱 **跨平台**: 支持多种操作系统

### 特殊平台类型

- **音乐盒平台** (粉色): +50 分
- **便利店平台** (绿色): +15 分
- **魔方平台** (黄色): +20 分
- **井盖平台** (棕色): +10 分

### 评分系统

- 基础跳跃: +1 分
- 中心落地 (连击区域): +50 分
- 特殊平台: 根据平台类型获得额外分数

### 系统要求

- Go 1.23.3 或更高版本
- 支持 OpenGL 2.1 或更高版本的图形卡

### 快速开始

#### 克隆项目
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

#### 安装依赖
```bash
go mod download
```

#### 运行游戏
```bash
go run .
```

### 游戏操作

- **鼠标左键**: 按住蓄力，松开跳跃
- **R 键**: 重新开始游戏（游戏结束时）

### 项目结构

```
.
├── main.go           # 主游戏逻辑
├── embed.go          # 游戏资源嵌入
├── go.mod            # Go 模块依赖
├── go.sum            # 依赖验证文件
├── images/           # 游戏图像资源
│   ├── player.png    # 玩家角色图片
│   ├── bullet.png    # 子弹图片 (预留)
│   └── enemy.png     # 敌人图片 (预留)
└── testfile/         # 测试文件
    └── test.go
```

### 技术栈

#### 核心技术
- **Go**: 主编程语言
- **Ebiten**: 2D 游戏引擎
- **go:embed**: 静态资源嵌入
- **golang.org/x/image**: 图像处理

#### 核心功能模块
- **Player**: 玩家角色控制和物理状态
- **Platform**: 平台碰撞检测和特殊效果
- **Game**: 游戏状态管理和主循环
- **Physics**: 重力、跳跃和碰撞物理

### 开发与构建

#### 开发环境设置
```bash
# 确保 Go 版本
go version

# 安装依赖
go mod tidy

# 运行开发版本
go run .
```

#### 构建发布版本
```bash
# 构建当前平台
go build -o jump-game

# 交叉编译 Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# 交叉编译 macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

#### 构建 Web 版本 (实验性)
```bash
# 需要 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

### 游戏截图

游戏采用精美的像素艺术风格和流畅的动画效果。角色会在平台间弹跳时自动转向和翻转方向，营造引人入胜的视觉体验。

### 贡献代码

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

### 参考资料

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go 社区](https://golang.org/) - 强大的编程语言

### 联系方式

如果有问题或建议，请通过 GitHub Issues 联系我们。

**享受游戏！** 🎮