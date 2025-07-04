# ggvim

使用 vim 指令移动绿色 P 来躲避红色 G 从而练习 vim。

## 安装
linux 用户：
```
curl -L -o ggvim https://github.com/GGboya/ggvim/releases/download/v1.0.0/ggvim-v1.0.0-linux-amd64
sudo chmod +x ggvim
sudo mv ggvim /usr/local/bin/
```

macOS 用户
```
curl -L -o ggvim https://github.com/GGboya/ggvim/releases/download/v1.0.0/ggvim-v1.0.0-macos
sudo chmod +x ggvim
sudo mv ggvim /usr/local/bin/
```

windows 用户
```
curl -L https://github.com/GGboya/ggvim/releases/download/v1.0.0/ggvim.exe
执行 ggvim.exe
```

## 使用

键入 `ggvim` 来开始游戏


## 游玩方法

游戏方法和吃豆人类似，移动绿色 P 来躲避红色 G。

P 有两个障碍：
1. 不能走到墙上，可以通过指令穿墙
2. 踩到蓝色 ~，结束游戏

**实现的指令**

| 按键 | 操作 |
| --- | --- |
| h   | 左移 |
| j   | 下移 |
| k   | 上移 |
| l   | 右移 |
| w   | 移到下一个 word 的开头 |
| W   | 移到下一个 WORD 的开头 |
| e   | 移到下一个 word 的结尾 |
| E   | 移到下一个 WORD 的结尾 |
| b   | 移到上一个 word 的开头 |
| B   | 移到上一个 WORD 的开头 |
| $   | 移到行尾 |
| 0   | 移到硬行首 |
| ^   | 移到软行首 |
| gg  | 移到第一行第一个非空字符 |
| G   | 移到最后一行第一个非空字符 |
| Ctrl C | 退出游戏 |

## 开发日志

- 20250516
修复 gg, G 的逻辑，增加难度系数，用户可选 easy, mid, hard
