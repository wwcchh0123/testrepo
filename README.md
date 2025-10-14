# Jump Jump Game / 跳跳游戏 🎮

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.23.3-blue.svg)
![Ebiten](https://img.shields.io/badge/Ebiten-v2.8.8-green.svg)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux%20%7C%20Web-lightgrey.svg)

一个使用 Go 语言和 Ebiten 游戏引擎开发的跳跃平台游戏，类似于"跳一跳"的玩法。

*An exciting platform jumping game built with Go and Ebiten engine, similar to "Jump Jump" gameplay.*

</div>

## 🎯 Game Description / 游戏介绍

这是一个充满挑战性的跳跃游戏，玩家需要控制角色在各种平台间跳跃，获取尽可能高的分数。游戏具有物理引擎、连击系统和特殊平台等丰富玩法。

*This is a challenging jumping game where players control a character jumping between various platforms to achieve the highest possible score. The game features physics engine, combo system, and special platforms.*

### 🎮 Key Features / 核心特性

- 🎯 **物理引擎** / **Physics Engine**: 真实的重力和跳跃物理模拟
- 🏆 **连击系统** / **Combo System**: 按住蓄力释放，精准跳跃获取连击加分  
- 🌈 **特殊平台** / **Special Platforms**: 流行歌曲中心可获得连击加分
- 💎 **道具系统** / **Power-ups**: 多种道具提供不同的增益效果
- ⚡ **动态效果** / **Dynamic Effects**: 流畅的角色动画和粒子特效
- 📈 **难度递增** / **Progressive Difficulty**: 随着分数增加，游戏难度逐渐提升
- 🎨 **视觉特效** / **Visual Effects**: 支持多种操作系统

### 🎪 Game Controls / 游戏控制

#### 操作方式 / Controls
- **蓄力跳跃** / **Charge Jump**: 按住蓄力，松开跳跃
- **重新开始** / **Restart**: 游戏结束后按 'R' 键重新开始

#### 平台类型 / Platform Types
- **普通平台** / **Normal Platform** (灰色): 基础分数
- **音乐盒平台** / **Music Box Platform** (粉红色): +30 分
- **便利店平台** / **Convenience Store Platform** (绿色): +15 分  
- **魔方平台** / **Rubik's Cube Platform** (黄色): +10 分
- **井盖平台** / **Manhole Platform** (棕色): +5 分

#### 评分系统 / Scoring System
- **基础跳跃**: +1 分
- **精准跳跃** (平台中心): 连击倍数递增
- **特殊平台**: 额外奖励分数

## 🚀 Quick Start / 快速开始

### 📋 Prerequisites / 前置要求

- Go 1.23.3 或更高版本 / *Go 1.23.3 or higher*
- 支持 OpenGL 2.1 或更高版本的图卡 / *Graphics card supporting OpenGL 2.1 or higher*

### 🛠️ Installation / 安装

#### 克隆项目 / Clone Repository
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

#### 安装依赖 / Install Dependencies
```bash
go mod download
```

#### 运行游戏 / Run Game
```bash
go run .
```

## 🏗️ Build / 构建

### 构建开发版本 / Build Development Version
```bash
# 确保 Go 版本
go version

# 安装依赖
go mod tidy

# 运行开发版本
go run .
```

### 构建发布版本 / Build Release Version
```bash
# 构建当前平台
go build -o jump-game

# 交叉编译 Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# 交叉编译 macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

### 构建 Web 版本 (实验性) / Build Web Version (Experimental)
```bash
# 需要安装 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 📁 Project Structure / 项目结构

```
.
├── main.go           # 主游戏逻辑
├── embed.go          # 游戏资源嵌入
├── go.mod            # Go 模块文件
├── go.sum            # 依赖校验文件
├── images/           # 游戏图片资源
│   ├── player.png    # 玩家角色图片
│   ├── bullet.png    # 子弹图片 (预留)
│   └── enemy.png     # 敌人图片 (预留)
└── testfile/         # 测试文件
    └── test.go
```

## 🔧 Technical Implementation / 技术实现

### 使用的技术栈 / Technology Stack
- **Go**: 主编程语言
- **Ebiten**: 2D 游戏引擎
- **go:embed**: 静态资源嵌入
- **golang.org/x/image**: 图像处理

### 核心功能模块 / Core Feature Modules
- **Player**: 玩家角色控制和物理状态
- **Platform**: 平台类型和碰撞检测  
- **Game**: 游戏状态管理和主循环
- **Physics**: 重力、跳跃和碰撞物理引擎

## 🎮 Game Screenshots / 游戏截图

游戏包含精美的像素风格图形和流畅的动画特效。角色会根据移动方向自动翻转，平台根据类型显示不同颜色。

*The game features beautiful pixel-style graphics and smooth animation effects. Characters automatically flip based on movement direction, and platforms display different colors based on their types.*

## 🤝 Contributing / 贡献指南

欢迎贡献代码！请遵循以下步骤：

*Contributions are welcome! Please follow these steps:*

1. Fork 这个仓库 / *Fork this repository*
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📄 License / 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

*This project is open source under the MIT License. See the [LICENSE](LICENSE) file for details.*

## 🔗 Links / 相关链接

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go 团队](https://golang.org/) - 强大的编程语言

## 💬 Support / 支持

如有问题或建议，请通过 GitHub Issues 联系我们。

*If you have any questions or suggestions, please contact us through GitHub Issues.*

---

**享受游戏！** 🎮 / ***Enjoy the game!*** 🎮