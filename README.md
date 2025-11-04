# Jump Jump Game 🎮

[English](#english) | [中文](#中文)

A challenging 2D platform jumping game built with Go and Ebiten game engine, similar to the popular "Jump Jump" gameplay mechanics.

---

## English

### 🎯 Game Description

Jump Jump Game is an engaging 2D platformer where players control a character jumping between platforms to achieve the highest possible score. The game features realistic physics, combo systems, special platforms, and smooth animations.

### ✨ Features

- 🎮 **Physics Engine**: Realistic gravity and jump physics simulation
- 🏆 **Charge System**: Hold mouse to charge power, release to jump
- 🎯 **Combo System**: Land precisely in platform center for combo multipliers
- 🌈 **Special Platforms**: Various colored platforms with different score bonuses
- 🎪 **Dynamic Effects**: Smooth character animations and camera following
- 📱 **Cross-platform**: Support for multiple operating systems
- 🎵 **Visual Feedback**: Color-coded platforms and responsive controls

### 🎮 How to Play

#### Controls
- **Left Mouse Button**: Hold to charge power, release to jump
- **R Key**: Restart game after game over
- **Wall Bouncing**: Character automatically bounces off screen edges

#### Platform Types
- **Normal Platform** (Gray): Basic scoring platform
- **Music Box Platform** (Pink): +30 points bonus
- **Convenience Store Platform** (Green): +15 points bonus
- **Rubik's Cube Platform** (Yellow): +10 points bonus
- **Manhole Platform** (Brown): +5 points bonus

#### Scoring System
- Basic jump: +1 point
- Precise landing (platform center): Combo multiplier increases
- Special platforms: Additional bonus points
- Combo system: Consecutive center hits multiply your score

### 🚀 Quick Start

#### Prerequisites
- Go 1.23.3 or higher
- OpenGL 2.1+ compatible graphics card

#### Installation
```bash
# Clone the repository
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo

# Download dependencies
go mod download

# Run the game
go run .
```

#### Building
```bash
# Build for current platform
go build -o jump-game

# Cross-compile for Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# Cross-compile for macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac

# Build for Web (experimental)
GOOS=js GOARCH=wasm go build -o game.wasm
```

### 🏗️ Project Structure

```
.
├── main.go           # Main game logic and physics
├── embed.go          # Embedded game assets
├── go.mod            # Go module file
├── go.sum            # Dependency checksums
├── images/           # Game image assets
│   ├── player.png    # Player character sprite
│   ├── bullet.png    # Bullet sprite (reserved)
│   └── enemy.png     # Enemy sprite (reserved)
└── testfile/         # Test files
    └── test.go
```

### 🛠️ Technical Stack

- **Go**: Main programming language
- **Ebiten**: 2D game engine for Go
- **go:embed**: Static asset embedding
- **golang.org/x/image**: Image processing utilities

### 🎯 Core Game Mechanics

#### Physics System
- **Gravity**: Constant downward force applied to player
- **Jump Power**: Variable jump strength based on charge time
- **Wall Bouncing**: Automatic direction reversal at screen boundaries
- **Collision Detection**: AABB (Axis-Aligned Bounding Box) collision system

#### Platform Generation
- **Procedural Generation**: Platforms are generated dynamically as player progresses
- **Difficulty Scaling**: Platform spacing increases with height
- **Special Platform Distribution**: 20% chance for special colored platforms

### 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### 🔗 Resources

- [Ebiten](https://ebiten.org/) - Excellent Go 2D game engine
- [Go](https://golang.org/) - Powerful programming language

---

## 中文

### 🎯 游戏介绍

这是一个充满挑战性的跳跳游戏，玩家需要控制角色在各种平台间跳跃，获取尽可能高的分数。游戏具有物理引擎、连击系统和特殊平台等丰富玩法。

### ✨ 特性

- 🎮 **物理引擎**: 真实的重力和跳跃物理模拟
- 🏆 **蓄力跳跃**: 按住鼠标蓄力，松开后跳跃
- 🎯 **连击系统**: 精准着陆在平台中心可获得连击加分
- 🌈 **特殊平台**: 多种颜色的特殊平台，提供不同的分数奖励
- 🎪 **动态效果**: 流畅的角色动画和相机跟随
- 📱 **跨平台**: 支持多种操作系统
- 🎵 **视觉反馈**: 颜色编码平台和响应式控制

### 🎮 游戏玩法

#### 操作方式
- **鼠标左键**: 按住蓄力，松开跳跃
- **R键**: 游戏结束后重新开始
- **墙壁反弹**: 角色自动从屏幕边缘反弹

#### 平台类型
- **普通平台** (灰色): 基础分数平台
- **音乐盒平台** (粉红色): +30 分奖励
- **便利店平台** (绿色): +15 分奖励
- **魔方平台** (黄色): +10 分奖励
- **井盖平台** (棕色): +5 分奖励

#### 评分系统
- 基础跳跃: +1 分
- 精准着陆 (平台中心): 连击倍数递增
- 特殊平台: 额外奖励分数
- 连击系统: 连续命中中心位置倍增得分

### 🚀 快速开始

#### 安装要求
- Go 1.23.3 或更高版本
- 支持 OpenGL 2.1+ 的图形卡

#### 安装步骤
```bash
# 克隆仓库
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo

# 下载依赖
go mod download

# 运行游戏
go run .
```

#### 构建
```bash
# 构建当前平台
go build -o jump-game

# 交叉编译 Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# 交叉编译 macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac

# 构建 Web 版本 (实验性)
GOOS=js GOARCH=wasm go build -o game.wasm
```

### 🏗️ 项目结构

```
.
├── main.go           # 主游戏逻辑和物理引擎
├── embed.go          # 嵌入式游戏资源
├── go.mod            # Go 模块文件
├── go.sum            # 依赖校验文件
├── images/           # 游戏图片资源
│   ├── player.png    # 玩家角色精灵
│   ├── bullet.png    # 子弹精灵 (预留)
│   └── enemy.png     # 敌人精灵 (预留)
└── testfile/         # 测试文件
    └── test.go
```

### 🛠️ 技术栈

- **Go**: 主编程语言
- **Ebiten**: Go 的 2D 游戏引擎
- **go:embed**: 静态资源嵌入
- **golang.org/x/image**: 图像处理工具

### 🎯 核心游戏机制

#### 物理系统
- **重力系统**: 对玩家施加恒定向下的力
- **跳跃力度**: 基于蓄力时间的可变跳跃强度
- **墙壁反弹**: 在屏幕边界自动方向反转
- **碰撞检测**: AABB (轴对齐边界框) 碰撞系统

#### 平台生成
- **程序化生成**: 随着玩家进度动态生成平台
- **难度递增**: 平台间距随高度增加
- **特殊平台分布**: 20% 概率生成特殊颜色平台

### 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 📄 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

### 🔗 相关资源

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go](https://golang.org/) - 强大的编程语言

---

## 📞 Contact / 联系方式

If you encounter any issues or have suggestions, please contact us through GitHub Issues.
如果遇到问题或有建议，请通过 GitHub Issues 联系我们。

**Enjoy the game! / 享受游戏！** 🎮