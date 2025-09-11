# Jump Jump Game 🎮

一个使用 Go 语言和 Ebiten 游戏引擎开发的跳跃平台游戏，类似于"跳一跳"的经典玩法。

![Game Language](https://img.shields.io/badge/language-Go-blue.svg)
![Game Engine](https://img.shields.io/badge/engine-Ebiten-green.svg)
![License](https://img.shields.io/badge/license-MIT-yellow.svg)

## 📖 游戏介绍

这是一个充满挑战性的跳跃游戏，玩家需要控制角色在各种平台间跳跃，获取尽可能高的分数。游戏具有真实的物理引擎、连击系统和多种特殊平台，为玩家带来丰富的游戏体验。

## ✨ 游戏特性

- 🎯 **蓄力跳跃系统**: 按住鼠标左键蓄力，控制跳跃距离和高度
- 🏆 **连击评分机制**: 精准落在平台中心可获得连击加分
- 🌈 **特殊平台系统**: 5种不同颜色的平台，提供不同的分数奖励
- 🎨 **流畅动画效果**: 角色动画和摄像机跟随系统
- 📱 **跨平台支持**: 支持 Windows、macOS、Linux
- ⚡ **实时物理引擎**: 真实的重力、碰撞和跳跃物理模拟

## 🎮 游戏玩法

### 操作控制
- **鼠标左键**: 按住蓄力，松开跳跃
- **R键**: 游戏结束后重新开始

### 平台类型与分数
| 平台类型 | 颜色 | 分数奖励 | 特殊效果 |
|---------|------|----------|----------|
| 普通平台 | 灰色 | +1 分 | 基础得分 |
| 音乐盒平台 | 粉色 | +30 分 | 最高奖励 |
| 便利店平台 | 绿色 | +15 分 | 中等奖励 |
| 魔方平台 | 黄色 | +10 分 | 普通奖励 |
| 井盖平台 | 棕色 | +5 分 | 基础奖励 |

### 评分系统
- **基础跳跃**: +1 分
- **精准落地** (平台中心): 连击倍数递增 (×2, ×3, ×4...)
- **特殊平台**: 根据平台类型获得额外奖励分数

## 🚀 快速开始

### 系统要求
- Go 1.23.3 或更高版本
- 支持 OpenGL 2.1 或更高版本的图形卡

### 安装运行

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

## 🏗️ 项目结构

```
.
├── main.go           # 主游戏逻辑和引擎
├── embed.go          # 游戏资源嵌入文件
├── go.mod           # Go 模块依赖配置
├── go.sum           # 依赖校验文件
├── images/          # 游戏图片资源目录
│   ├── player.png   # 玩家角色精灵图
│   ├── bullet.png   # 子弹图片 (预留扩展)
│   └── enemy.png    # 敌人图片 (预留扩展)
└── testfile/        # 测试文件目录
    └── test.go
```

## 🛠️ 技术实现

### 技术栈
- **[Go](https://golang.org/)**: 主要编程语言，提供出色的性能和跨平台支持
- **[Ebiten v2.8.8](https://ebiten.org/)**: 专业的 2D 游戏引擎
- **go:embed**: 静态资源嵌入，无需外部文件依赖
- **golang.org/x/image**: 图像处理和字体渲染

### 核心模块
- **Player**: 玩家角色控制、物理状态管理和碰撞检测
- **Platform**: 平台系统、类型管理和特殊效果
- **Game**: 游戏状态管理、主循环和UI渲染
- **Physics**: 重力系统、跳跃力学和碰撞响应

### 游戏机制
- **蓄力系统**: 鼠标按压时间转换为跳跃力度
- **物理引擎**: 重力加速度 0.4，最大蓄力 15.0
- **碰撞检测**: AABB 包围盒碰撞检测算法
- **摄像机系统**: 平滑跟随玩家移动

## 🔧 开发与构建

### 开发环境设置
```bash
# 检查 Go 版本
go version

# 整理依赖
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

# 交叉编译 Linux
GOOS=linux GOARCH=amd64 go build -o jump-game-linux
```

### 构建 Web 版本 (实验性)
```bash
# 需要安装 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 🎯 游戏攻略

### 基础技巧
1. **掌握蓄力时机**: 观察平台距离，合理控制蓄力时间
2. **瞄准平台中心**: 精准落地可获得连击奖励
3. **优先特殊平台**: 粉色和绿色平台提供最高分数
4. **保持节奏**: 连续精准跳跃可快速提升分数

### 高分策略
- 连击是获得高分的关键，尽量瞄准每个平台的中心
- 特殊平台出现时优先选择，即使路径稍远也值得
- 学会控制跳跃角度，利用墙壁反弹到达更远的平台

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 开发规范
- 代码风格遵循 Go 官方规范
- 提交信息使用英文，格式清晰
- 新功能需要添加相应的测试
- 确保代码通过 `go fmt` 和 `go vet` 检查

## 📝 更新日志

### v1.0.0 (当前版本)
- ✅ 基础跳跃游戏机制
- ✅ 物理引擎和碰撞检测
- ✅ 5种特殊平台类型
- ✅ 连击和评分系统
- ✅ 角色动画和摄像机跟随
- ✅ 跨平台支持

### 计划功能
- 🔄 音效和背景音乐
- 🔄 成就系统
- 🔄 在线排行榜
- 🔄 更多平台类型和道具
- 🔄 关卡模式

## 📄 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- [Ebiten](https://ebiten.org/) - 出色的 Go 2D 游戏引擎
- [Go Team](https://golang.org/) - 强大而优雅的编程语言
- 所有为项目贡献代码和建议的开发者

## 📞 联系方式

- 🐛 **问题反馈**: [GitHub Issues](https://github.com/wwcchh0123/testrepo/issues)
- 💬 **功能建议**: [GitHub Discussions](https://github.com/wwcchh0123/testrepo/discussions)
- 📧 **其他联系**: 通过 GitHub 联系项目维护者

---

<div align="center">

**🎮 享受游戏，挑战高分！**

Made with ❤️ using Go and Ebiten

[⭐ Star](https://github.com/wwcchh0123/testrepo) | [🍴 Fork](https://github.com/wwcchh0123/testrepo/fork) | [📝 Issues](https://github.com/wwcchh0123/testrepo/issues)

</div>