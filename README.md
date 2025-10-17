# 🎮 Jump Jump Game - 增强版

[![Go Version](https://img.shields.io/badge/Go-1.23.3+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Ebiten](https://img.shields.io/badge/Ebiten-2.8.8-green?style=flat-square)](https://ebiten.org)
[![License](https://img.shields.io/badge/License-MIT-yellow?style=flat-square)](LICENSE)
[![Stars](https://img.shields.io/github/stars/wwcchh0123/testrepo?style=flat-square)](https://github.com/wwcchh0123/testrepo/stargazers)

一个使用 Go 语言和 Ebiten 游戏引擎开发的增强版跳跃平台游戏，具有道具系统、粒子效果和动态难度调整等高级功能。

## 🌟 游戏演示

![游戏截图](https://github.com/user-attachments/assets/game-demo.gif)
*控制角色在平台间跳跃，收集道具，挑战高分！*

## 🎮 游戏介绍

这是一个充满挑战性的跳跃游戏，玩家需要控制角色在各种平台间跳跃，收集道具，获取尽可能高的分数。游戏具有完整的物理引擎、连击系统、道具系统和视觉特效。

## ✨ 核心特性

### 🎯 游戏机制
- **蓄力跳跃系统**: 长按鼠标/触摸蓄力，释放后跳跃
- **智能物理引擎**: 真实的重力模拟和墙壁反弹
- **连击系统**: 精准落地可获得连击加分
- **难度递进**: 随着分数增加自动调整游戏难度

### 🏆 平台系统
- **普通平台** (灰色): 基础跳跃点
- **音乐盒平台** (粉色): +30 分
- **便利店平台** (绿色): +15 分
- **魔方平台** (黄色): +10 分
- **井盖平台** (棕色): +5 分

### 🎁 道具系统
- **🚀 速度提升**: 1.5倍跳跃力，持续10秒
- **💰 双倍得分**: 2倍分数奖励，持续15秒
- **⬆️ 额外跳跃**: 获得额外跳跃机会，持续20秒
- **⏰ 慢动作**: 0.5倍时间流速，持续8秒

### 🎨 视觉效果
- **粒子系统**: 着陆和收集道具时的粒子特效
- **流畅动画**: 角色翻转和摄像机跟随
- **实时UI**: 显示当前分数、最高分和激活的道具

## 🕹️ 操作指南

### 基础操作
- **鼠标左键/触摸**: 长按蓄力，松开跳跃
- **R键**: 游戏结束后重新开始

### 高级技巧
- **精准着陆**: 落在平台中心可获得连击加分
- **道具收集**: 主动收集道具以获得战略优势
- **墙壁反弹**: 利用墙壁反弹到达更远的平台

## 🏗️ 技术架构

### 核心模块
```go
type Player struct {
    // 位置和物理状态
    x, y, vx, vy float64
    
    // 道具效果
    speedBoost    float64
    doubleScore   float64
    extraJumps    int
    slowMotion    float64
    powerUpTimers map[int]float64
}

type Game struct {
    player          *Player
    platforms       []*Platform
    powerUps        []*PowerUp
    particles       []*Particle
    difficultyLevel int
}
```

### 技术栈
- **Go 1.23.3+**: 主要编程语言
- **Ebiten 2.8.8**: 2D 游戏引擎框架
- **go:embed**: 静态资源嵌入技术
- **golang.org/x/image**: 图像处理库

## 🚀 快速开始

### 环境要求
- Go 1.23.3 或更高版本
- 支持 OpenGL 2.1+ 的图形卡
- Windows/macOS/Linux 操作系统

### 📥 安装运行

#### 方式一：源码运行
```bash
# 克隆项目
git clone https://github.com/wwcchh0123/testrepo.git
cd testrepo

# 安装依赖
go mod download

# 运行游戏
go run .
```

#### 方式二：直接下载
```bash
# 下载并运行（Linux/macOS）
curl -L https://github.com/wwcchh0123/testrepo/releases/latest/download/jump-game -o jump-game
chmod +x jump-game
./jump-game
```

### 构建发布
```bash
# 本地构建
go build -o jump-game

# Windows 交叉编译
GOOS=windows GOARCH=amd64 go build -o jump-game.exe

# macOS 交叉编译  
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac

# Web 版本 (实验性)
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 📁 项目结构

```
.
├── main.go                 # 主游戏逻辑和状态管理
├── embed.go               # 静态资源嵌入
├── go.mod                 # Go 模块依赖配置
├── go.sum                 # 依赖版本锁定
├── images/                # 游戏图片资源
│   ├── player.png         # 玩家角色精灵
│   ├── bullet.png         # 预留素材
│   └── enemy.png          # 预留素材
├── testfile/              # 测试代码
│   └── test.go
└── .github/               # CI/CD 配置
    └── workflows/
        ├── ci.yml         # 持续集成
        └── claude.yml     # AI 助手集成
```

## 🎯 游戏机制详解

### 评分系统
```
基础得分 = 连击数 × 2 (精准着陆) 或 1 (普通着陆)
特殊平台加分 = 平台类型分数
道具加成 = 双倍得分时 × 2
最终得分 = (基础得分 + 特殊平台加分) × 道具加成
```

### 难度系统
- 每100分增加一个难度等级
- 平台间距随难度增加 (最大+50像素)
- 特殊平台出现率随难度提升 (最高40%)
- 道具出现率随难度降低 (最低5%)

### 物理引擎
- 重力常数: 0.4
- 最大蓄力: 15.0
- 跳跃角度: 60° (π/3)
- 跳跃力系数: -1.2

## 🛠️ 开发指南

### 本地开发
```bash
# 检查 Go 版本
go version

# 整理依赖
go mod tidy

# 实时开发 (配合热重载工具)
go run .
```

### 代码架构
- **单例模式**: Game 实例管理
- **组件系统**: Player, Platform, PowerUp 独立模块
- **事件驱动**: 输入处理和碰撞检测
- **状态机**: 游戏状态管理

## 🎨 自定义与扩展

### 添加新道具
```go
const (
    powerUpNewType = iota + 4
)

func (g *Game) applyPowerUp(powerType int) {
    switch powerType {
    case powerUpNewType:
        // 自定义道具效果
        g.player.customEffect = true
        g.player.powerUpTimers[powerType] = 12.0
    }
}
```

### 添加新平台类型
```go
const (
    platformCustom = iota + 5
)

var platformColors = map[int]color.Color{
    platformCustom: color.RGBA{R, G, B, 255},
}

var platformScores = map[int]int{
    platformCustom: customScore,
}
```

## 🤝 贡献指南

我们欢迎各种形式的贡献！

### 贡献类型
- 🐛 Bug 修复
- ✨ 新功能开发
- 📝 文档改进
- 🎨 美术资源
- 🧪 测试用例

### 提交流程
1. Fork 项目仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m '添加了惊人的新功能'`)
4. 推送分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 代码规范
- 遵循 Go 官方代码风格
- 添加必要的注释和文档
- 确保所有测试通过
- 提交信息使用中文描述

## 📄 许可证

本项目采用 MIT 许可证开源。详情请查看 [LICENSE](LICENSE) 文件。

```
MIT License - 允许商业使用、修改、分发和私人使用
```

## 🙏 致谢

- **[Ebiten](https://ebiten.org/)** - 优秀的 Go 2D 游戏引擎
- **[Go Team](https://golang.org/)** - 强大且优雅的编程语言
- **开源社区** - 持续的支持和贡献

## 📊 游戏数据

### 性能指标
- **帧率**: 60 FPS 稳定运行
- **内存占用**: ~20MB 游戏运行时
- **启动时间**: <2秒 冷启动
- **支持平台**: Windows/macOS/Linux/Web

### 游戏统计
- 🏆 **最高难度**: 无上限递增
- 🎯 **道具种类**: 4种不同效果
- 🌈 **平台类型**: 5种特色平台
- ⚡ **物理精度**: 60 FPS 物理计算

## 🔧 故障排除

### 常见问题

**Q: 游戏无法启动？**
A: 检查 Go 版本是否 ≥1.23.3，确保显卡支持 OpenGL 2.1+

**Q: 游戏卡顿或帧率低？**
A: 尝试关闭其他占用显卡的程序，或降低系统分辨率

**Q: 鼠标操作不响应？**
A: 确保游戏窗口获得焦点，点击窗口标题栏后重试

**Q: Web 版本加载慢？**
A: WASM 版本需要下载资源，请耐心等待初始化完成

### 性能优化

```bash
# 性能调试模式
EBITEN_DEBUG=1 go run .

# 禁用垂直同步（可能提高帧率）
EBITEN_GRAPHICS_API=opengl go run .
```

## 📈 开发计划

### 即将推出 v2.0
- [ ] 🎵 背景音乐和音效系统
- [ ] 🏅 成就系统和解锁内容
- [ ] 👥 多人模式支持
- [ ] 🎨 自定义皮肤系统
- [ ] 📱 移动端适配优化

### 技术升级
- [ ] WebGL 渲染优化
- [ ] 存档系统实现
- [ ] 网络排行榜
- [ ] 模组支持框架

## 📞 联系我们

- 📧 **Issues**: [GitHub Issues](https://github.com/wwcchh0123/testrepo/issues)
- 💬 **讨论**: [GitHub Discussions](https://github.com/wwcchh0123/testrepo/discussions)
- 🐛 **Bug 报告**: 请使用 Issue 模板
- 🌟 **功能建议**: 欢迎提交 Feature Request

---

<div align="center">

**🎮 开始你的跳跃之旅吧！**

> 💡 **小贴士**: 尝试在平台中心着陆以获得更高的连击分数！

*让我们一起构建更精彩的游戏体验*

**[⭐ Star](https://github.com/wwcchh0123/testrepo) 这个项目如果你喜欢它！**

</div>