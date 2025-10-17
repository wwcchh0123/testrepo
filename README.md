# 🎮 Jump Jump Game

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23.3-00ADD8?style=for-the-badge&logo=go" alt="Go版本">
  <img src="https://img.shields.io/badge/Ebiten-v2.8.8-FF6B35?style=for-the-badge" alt="Ebiten版本">
  <img src="https://img.shields.io/badge/平台-跨平台-brightgreen?style=for-the-badge" alt="跨平台">
  <img src="https://img.shields.io/badge/许可证-MIT-blue?style=for-the-badge" alt="MIT许可证">
</p>

<p align="center">
  <strong>一个使用 Go 语言和 Ebiten 游戏引擎开发的现代化跳跃平台游戏</strong><br>
  灵感来源于经典的"跳一跳"玩法，融入了现代游戏设计元素
</p>

## ✨ 游戏特色

这是一个充满挑战性和趣味性的跳跃游戏，玩家需要精确控制角色在各种特殊平台间进行跳跃，通过技巧和策略获取最高分数。游戏集成了先进的物理引擎、动态难度调节和丰富的视觉效果。

## 🚀 核心特性

- 🎮 **精准物理引擎**: 真实的重力模拟和跳跃物理计算
- 🎯 **蓄力跳跃系统**: 按住鼠标/触摸屏蓄力，释放实现精准跳跃
- 🏆 **智能连击系统**: 精准落在平台中心可获得连击奖励和额外分数
- 🌈 **多样化特殊平台**: 音乐盒、便利店、魔方、井盖等特色平台提供不同奖励
- 🎨 **绚丽视觉效果**: 流畅的粒子系统、动画效果和平台着色
- 📱 **全平台兼容**: 支持触摸屏、鼠标操控，适配移动端和桌面端
- ⚡ **道具增强系统**: 速度提升、双倍积分、额外跳跃、慢动作等道具
- 🎯 **动态难度调节**: 根据得分自动调整游戏难度，保持挑战性

## 🎮 游戏玩法

### 🕹️ 操作方式
- **鼠标操控**: 按住鼠标左键蓄力，释放后跳跃
- **触摸操控**: 按住屏幕蓄力，释放后跳跃
- **重新开始**: 游戏结束后按 R 键重新开始

### 🏆 平台类型与得分系统
- **普通平台**: 基础平台，提供基本分数
- **音乐盒平台** (粉色): +30 分
- **便利店平台** (绿色): +15 分
- **魔方平台** (黄色): +10 分
- **井盖平台** (棕色): +5 分

### 💎 道具增强系统
- **速度提升道具** (红色): 跳跃力增强 50%，持续 10 秒
- **双倍积分道具** (金色): 得分翻倍，持续 15 秒
- **额外跳跃道具** (青色): 获得额外跳跃机会，持续 20 秒
- **慢动作道具** (紫色): 时间放慢效果，持续 8 秒

### 🎯 连击系统
- 基础得分: +1 分
- 落在平台中心 (精准降落): 连击加分递增
- 特殊平台: 额外奖励分数叠加

## 📋 系统要求

### 💻 开发环境
- **Go**: 1.23.3 或更高版本
- **操作系统**: 支持 OpenGL 2.1 或更高版本的图形卡

### 🔧 依赖库
- **Ebiten**: v2.8.8 - 优秀的 Go 2D 游戏引擎
- **golang.org/x/image**: 图像处理库

## 🚀 快速开始

### 📥 克隆项目
```bash
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo
```

### 📦 安装依赖
```bash
go mod download
```

### ▶️ 运行游戏
```bash
go run .
```

## 🔧 项目结构

```
.
├── main.go             # 主游戏逻辑
├── embed.go            # 游戏资源嵌入
├── go.mod              # Go 模块文件
├── go.sum              # 依赖验证文件
├── images/             # 游戏图片资源
│   ├── player.png      # 玩家角色图片
│   ├── bullet.png      # 子弹图片 (预留)
│   └── enemy.png       # 敌人图片 (预留)
├── testfile/           # 测试文件
│   └── test.go
└── .github/            # GitHub 工作流配置
    └── workflows/
        ├── ci.yml      # 持续集成
        └── claude.yml  # Claude 自动化
```

## 🛠️ 技术架构

### 💻 核心技术栈
- **Go**: 主要编程语言，提供出色的性能和并发能力
- **Ebiten**: 专业的 2D 游戏引擎，跨平台支持
- **go:embed**: 静态资源嵌入，简化部署过程
- **golang.org/x/image**: 高效图像处理

### 🏗️ 核心功能模块
- **Player**: 玩家角色控制和物理状态管理
- **Platform**: 平台生成、碰撞检测和特殊效果
- **Game**: 游戏状态管理和主循环逻辑
- **Physics**: 重力、跳跃和碰撞物理引擎
- **Particles**: 粒子效果系统，提升视觉体验
- **PowerUps**: 道具系统，增加游戏趣味性

## 🔨 构建与发布

### 🏗️ 构建开发版本
```bash
# 构建当前平台
go build -o jump-game

# 交叉编译 Windows
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# 交叉编译 macOS
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

### 🌐 构建 Web 版本 (实验性)
```bash
# 需要安装 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 🎯 游戏截图

游戏包含精美的像素风格图形和流畅的动画效果。角色会根据移动方向自动翻转，粒子系统会在跳跃和道具收集时创建华丽的视觉反馈。

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. **Fork** 这个仓库
2. **创建功能分支** (`git checkout -b feature/amazing-feature`)
3. **提交更改** (`git commit -m 'Add some amazing feature'`)
4. **推送到分支** (`git push origin feature/amazing-feature`)
5. **创建 Pull Request**

## 📄 许可证

本项目基于 **MIT** 许可证开源。详见 [LICENSE](LICENSE) 文件。

## 🔗 相关链接

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go 官网](https://golang.org/) - 强大的编程语言

## 📞 联系我们

如果遇到问题或有建议，请通过 GitHub Issues 联系我们。

---

**🎮 享受游戏！** 🏆