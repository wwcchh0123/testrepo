# Jump Jump Game

一个使用 Go 语言和 Ebiten 游戏引擎开发的趣味跳跃平台游戏，类似于"跳一跳"的玩法。

## 🎮 游戏介绍

这是一个充满挑战性的跳跃游戏，玩家需要控制角色在各种平台间跳跃，获取尽可能高的分数。游戏具有物理引擎、连击系统和特殊平台等丰富功能。

## ✨ 特性

- 🎯 **物理引擎**：真实的重力和跳跃物理模拟
- 🎮 **蓄力跳跃**：按住按键积蓄力量，释放后跳跃
- 🌟 **连击系统**：连续准确着陆可获得加分
- 🎨 **特殊平台**：多种特殊平台提供不同的分数奖励
- 💎 **动态难度**：游戏随着进度逐渐增加挑战性
- 📱 **跨平台支持**：支持多种操作系统

## 🎯 游戏玩法

### 操作方式
- **鼠标左键**：按住蓄力，释放跳跃
- **R键**：游戏结束后重新开始

### 平台类型
- **普通平台**：基础平台，提供基本分数
- **音乐盒平台** (音响符号)：+30 分
- **便利店平台** (便利店)：+15 分  
- **魔方平台** (魔方)：+10 分
- **橙色平台** (井盖)：+5 分

### 评分系统
- 基础跳跃：+1 分
- 粒子发射区 (平台中心)：连击倍数递增
- 特殊平台：额外奖励分数

## 🛠 安装要求

- Go 1.23.3 或更高版本
- 支持 OpenGL 2.1 或更高版本的图形卡

## 🚀 快速开始

### 克隆项目
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

### 安装依赖
```bash
go mod download
```

### 运行游戏
```bash
go run .
```

## 📁 项目结构

```
.
├── main.go              # 主游戏逻辑
├── embed.go             # 游戏资源嵌入
├── go.mod               # Go 模块依赖
├── go.sum               # 依赖校验文件
├── images/              # 游戏图片资源
│   ├── player.png       # 玩家角色图片
│   ├── bullet.png       # 子弹图片 (预留)
│   └── enemy.png        # 敌人图片 (预留)
└── testfile/            # 测试文件
    └── test.go
```

## 🔧 技术实现

### 使用的技术栈
- **Go**: 主编程语言
- **Ebiten**: 2D 游戏引擎
- **go:embed**: 静态资源嵌入
- **golang.org/x/image**: 图像处理

### 核心功能模块
- **Player**: 玩家角色控制和物理状态
- **Platform**: 平台类型和碰撞检测
- **Game**: 游戏状态管理和主循环
- **Physics**: 重力、跳跃和碰撞物理引擎

## 🔨 构建发布

### 构建当前平台
```bash
# 确保 Go 版本
go version

# 安装依赖
go mod tidy

# 运行开发版本
go run .
```

### 构建发布版本
```bash
# 构建当前平台
go build -o jump-game

# 交叉编译 Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# 交叉编译 macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

### 构建 Web 版本 (实验性)
```bash
# 需要安装 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 🎨 游戏截图

游戏包含精美的像素风格图形和流畅的动画效果。角色会根据移动方向自动翻转，平台会随机生成跳跃机会。

## 📄 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

## 🙋 联系方式

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go 官网](https://golang.org/) - 强大的编程语言

## 🔗 相关链接

如果遇到问题或建议，请通过 GitHub Issues 联系我们。

---

**享受游戏！** 🎮