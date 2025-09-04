# 跳一跳小游戏 (Jump Jump Game)

一个使用 Go 语言和 Ebiten 游戏引擎开发的跳一跳小游戏。玩家需要控制角色在不同的平台间跳跃，获得尽可能高的分数。

## 游戏特色

### 🎮 核心玩法
- **蓄力跳跃**: 按住鼠标左键蓄力，松开后角色会根据蓄力时间进行跳跃
- **物理引擎**: 真实的重力和弹跳物理效果
- **墙壁反弹**: 角色碰到屏幕边缘会反弹并改变方向
- **精准着陆**: 着陆在平台中心可以获得连击奖励

### 🏆 计分系统
- **基础分数**: 每次成功着陆获得 1 分
- **连击奖励**: 精准着陆（距离平台中心 < 1/4 平台宽度）可以累积连击，连击数 × 2 分
- **特殊平台奖励**: 不同类型的平台提供额外分数
  - 🎵 音乐盒平台：+40 分（粉色）
  - 🏪 便利店平台：+15 分（绿色）
  - 🧩 魔方平台：+10 分（黄色）
  - 🕳️ 井盖平台：+10 分（棕色）

### 🎨 视觉效果
- 流畅的摄像头跟随
- 彩色平台系统
- 角色朝向自动调整
- 实时显示蓄力条和分数

## 技术架构

### 🔧 核心技术
- **语言**: Go 1.23.3
- **游戏引擎**: [Ebiten v2.8.8](https://github.com/hajimehoshi/ebiten)
- **图像处理**: golang.org/x/image
- **资源嵌入**: Go embed 特性

### 📁 项目结构
```
.
├── main.go          # 主要游戏逻辑
├── embed.go         # 图像资源嵌入
├── images/          # 游戏图片资源
│   ├── player.png   # 玩家角色图片
│   ├── bullet.png   # 子弹图片（预留）
│   └── enemy.png    # 敌人图片（预留）
├── testfile/        # 测试文件
│   └── test.go
├── go.mod          # Go 模块配置
├── go.sum          # 依赖锁定文件
└── README.md       # 项目说明文档
```

### 🏗️ 核心组件

#### Player（玩家）
- 位置和速度管理
- 蓄力跳跃机制
- 方向控制
- 碰撞检测

#### Platform（平台）
- 5 种不同类型的平台
- 颜色区分和分数奖励
- 随机生成算法
- 碰撞区域计算

#### Game（游戏管理）
- 游戏状态管理
- 摄像头控制
- 分数和连击系统
- 游戏结束判定

## 安装与运行

### 📋 环境要求
- Go 1.23.3 或更高版本
- 支持 OpenGL 的系统

### 🚀 快速开始

1. **克隆仓库**
   ```bash
   git clone https://github.com/wwcchh0123/testrepo.git
   cd testrepo
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **运行游戏**
   ```bash
   go run .
   ```

### 🎮 操作方式
- **鼠标左键按住**: 蓄力
- **鼠标左键释放**: 跳跃
- **R 键**: 游戏结束后重新开始

## 游戏截图

游戏窗口尺寸：480 × 640 像素

## 开发说明

### 🔄 游戏循环
1. **Update()**: 处理输入、物理计算、碰撞检测
2. **Draw()**: 渲染游戏画面、UI 元素
3. **Layout()**: 管理屏幕布局

### ⚙️ 关键参数
```go
const (
    screenWidth      = 480      // 屏幕宽度
    screenHeight     = 640      // 屏幕高度
    gravity          = 0.4      // 重力加速度
    maxCharge        = 15.0     // 最大蓄力值
    jumpPowerMultiplier = -1.2  // 跳跃力度倍数
)
```

### 🧪 测试
项目包含测试文件目录 `testfile/`，可以运行：
```bash
go test ./testfile
```

## 贡献指南

欢迎提交 Issue 和 Pull Request！

### 🐛 报告问题
- 请详细描述问题和重现步骤
- 包含您的操作系统和 Go 版本信息

### 💡 功能建议
- 新的平台类型和特效
- 改进游戏平衡性
- 增加音效和背景音乐
- 添加存档和排行榜功能

## 许可证

本项目遵循开源许可证，具体请查看 LICENSE 文件。

## 更新日志

- **v1.0**: 基础跳一跳游戏功能
  - 蓄力跳跃机制
  - 特殊平台系统
  - 连击和计分系统
  - 物理引擎集成

## 支持项目 (Support This Project)

如果这个项目对您有帮助，请考虑给予支持！您的支持是我们持续改进和开发的动力。

If this project has been helpful to you, please consider supporting it! Your support motivates us to keep improving and developing.

### 🌟 如何支持 (How to Support)

- **⭐ 给项目点星**: 在 GitHub 上为项目点 Star，让更多人发现这个项目
- **🐛 报告问题**: 发现 Bug 或有改进建议？请创建 Issue
- **💻 贡献代码**: 欢迎提交 Pull Request，一起完善项目
- **📢 分享推广**: 将项目分享给朋友和社区，扩大影响力
- **📝 完善文档**: 帮助我们改进文档和教程

### 💝 特别感谢 (Special Thanks)

感谢所有为项目贡献代码、提出建议、报告问题的开发者们！正是有了你们的参与，这个项目才能不断进步。

Thanks to all the developers who have contributed code, suggestions, and bug reports! It is your participation that makes this project continuously improve.

### 🤝 加入我们 (Join Us)

- 关注项目获取最新动态
- 加入讨论，分享您的想法和建议  
- 成为贡献者，让项目变得更好

Follow the project for updates, join discussions, share your ideas, and become a contributor to make this project even better!

---

**享受游戏时光！🎉**

**Have fun gaming! 🎉**