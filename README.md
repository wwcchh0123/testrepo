# Jump Jump Game / 跳一跳小游戏

一个使用 Go 语言和 Ebiten 游戏引擎开发的跳跃平台游戏，类似于"跳一跳"的玩法。玩家需要控制角色在各种平台间跳跃，收集道具并获取尽可能高的分数。

## 🎮 游戏特色

- **🎯 精准跳跃**: 真实的物理引擎和重力模拟
- **🌟 连击系统**: 精确着陆获得连击加成
- **🎨 多样平台**: 多种特殊平台类型，每种都有独特的分数奖励
- **⚡ 道具系统**: 四种不同的增益道具
- **💫 视觉特效**: 粒子效果和流畅的动画
- **📊 难度调节**: 动态难度调整系统
- **🏆 高分记录**: 本地最高分记录保存

## 🎪 游戏内容

### 平台类型
- **普通平台**: 基础得分平台
- **音乐盒平台** (粉色): +20 分
- **便利店平台** (绿色): +50 分  
- **魔方平台** (黄色): +10 分
- **井盖平台** (棕色): +20 分

### 道具系统
- **⚡ 速度提升**: 跳跃力度增强
- **💰 双倍得分**: 得分加倍效果
- **🦘 额外跳跃**: 获得额外跳跃机会
- **⏰ 慢动作**: 时间减缓效果

### 评分机制
- 基础跳跃: +1 分
- 精准着陆 (平台中心): 连击加成
- 特殊平台: 额外分数奖励
- 道具效果: 分数增益加成

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

## 🎮 游戏操作

- **鼠标左键** / **触摸屏幕**: 蓄力跳跃
  - 按住蓄力，松开释放
  - 蓄力时间决定跳跃距离和高度
- **R 键**: 游戏结束后重新开始

## 📂 项目结构

```
.
├── main.go           # 主游戏逻辑
├── embed.go          # 游戏资源嵌入
├── go.mod            # Go 模块依赖
├── go.sum            # 依赖校验文件
├── images/           # 游戏图片资源
│   ├── player.png    # 玩家角色图片
│   ├── bullet.png    # 子弹图片 (预留)
│   └── enemy.png     # 敌人图片 (预留)
└── testfile/         # 测试文件
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
- **Platform**: 平台系统和碰撞检测
- **Game**: 游戏状态管理和主循环
- **Physics**: 重力、跳跃和碰撞物理引擎

## 🏗 构建发布

### 构建当前平台
```bash
go build -o jump-game
```

### 交叉编译 Windows
```bash
GOOS=windows GOARCH=amd64 go build -o jump-game.exe
```

### 交叉编译 macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o jump-game-mac
```

### 构建 Web 版本 (实验性)
```bash
# 需要安装 Ebiten 的 Web 支持
GOOS=js GOARCH=wasm go build -o game.wasm
```

## 🎯 游戏截图

游戏包含精美的像素风格图形和流畅的动画效果。角色会根据跳跃方向自动翻转，粒子效果会在着陆时产生视觉反馈。

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📄 许可证

本项目基于 MIT 许可证开源。详见 [LICENSE](LICENSE) 文件。

## 🔗 相关链接

- [Ebiten](https://ebiten.org/) - 优秀的 Go 2D 游戏引擎
- [Go 官网](https://golang.org/) - 强大的编程语言

## 💡 技术亮点

### 物理引擎特性
- 重力模拟和自由落体运动
- 墙面反弹和方向控制
- 精确的 AABB 碰撞检测

### 游戏性创新
- 动态难度调整：分数越高，平台间距越大
- 连击系统：精准着陆获得更高分数
- 道具增益：多种临时能力提升
- 视觉反馈：着陆和收集道具的粒子特效

---

**享受游戏！** 🎮 如果遇到问题或有建议，请通过 GitHub Issues 联系我们。