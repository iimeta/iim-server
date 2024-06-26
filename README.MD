<div align=center>

<img src="https://iim.ai/public/images/logo.png?t=20231207" width="128"/>

# 智元 IIM

[![Release](https://img.shields.io/github/v/release/iimeta/iim-server?color=blue)](https://github.com/iimeta/iim-server/releases)
[![License](https://img.shields.io/static/v1?label=license&message=MIT&color=green)](https://github.com/iimeta/iim-server/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/iimeta/iim-server)](https://goreportcard.com/report/github.com/iimeta/iim-server)

</div>

`智元 IIM` 是一款开源的网页版即时聊天系统, 同时拥有AI聊天对话功能, 本项目基于[Lumen IM](https://github.com/gzydong/go-chat)使用[GoFrame](https://github.com/gogf/gf)框架开发

## 功能介绍
- 支持 WebSocket 通信
- 支持单聊及群聊场景
- 支持聊天消息类型有: 文本、代码块、图片及其它类型文件等
- 支持聊天消息撤回、删除或批量删除、转发消息（逐条转发、合并转发）及群投票功能
- 支持编写个人笔记
- 支持ChatGPT、Midjourney、文心一言、讯飞星火、通义千问等AI助手功能

## 项目预览

- 地址： [https://iim.ai](https://iim.ai)
- 账号： user1@iim.ai 或 user2@iim.ai
- 密码： 123456

## 快速上手
0. 需提前安装好Redis和MongoDB
1. 到 [Releases](https://github.com/iimeta/iim-client/releases) 下载最新版本
2. 解压后, 进入 iim-client/manifest/config 目录, 把 config.example.yaml 拷贝或修改文件名为 config.yaml 
3. 按 config.yaml 中配置项的注释, 修改对应参数
4. Linux环境, 到 iim-client/bin 目录下, 执行: chmod +777 ./*
5. Linux启动服务, 执行: ./start.sh
6. Windows启动服务, 双击: start.bat
7. 访问: http://ip:8000/ 如: http://localhost:8000/

## Docker部署
### [传送门](https://github.com/iimeta/iim-server/tree/docker)

## 开源地址
| 仓库 | 前端                              | 后端                                 | API                               |
| ---- | --------------------------------- | ------------------------------------ | --------------------------------- |
| 主库 | https://github.com/iimeta/iim-web | https://github.com/iimeta/iim-client | https://github.com/iimeta/iim-api |
| 码云 | https://gitee.com/iimeta/iim-web  | https://gitee.com/iimeta/iim-client  | https://gitee.com/iimeta/iim-api  |

## 作者
<img src="https://iim.ai/public/images/Author.png?t=20231207" width="320" alt="作者"/>

## 交流群
<img src="https://iim.ai/public/images/WeChatGroup.jpg?t=20240510" width="320" alt="交流群"/>


![Star History Chart](https://api.star-history.com/svg?repos=iimeta/iim-server&type=Date)
