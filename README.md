# Srun

[![Build Status](https://travis-ci.org/vouv/srun.svg?branch=master)](https://travis-ci.org/vouv/srun) [![Go Report Card](https://goreportcard.com/badge/github.com/vouv/srun)](https://goreportcard.com/report/github.com/vouv/srun) ![License](https://img.shields.io/packagist/l/doctrine/orm.svg) [![GoDoc](https://godoc.org/github.com/vouv/srun?status.svg)](https://godoc.org/github.com/vouv/srun/core) [![Donate](https://img.shields.io/badge/%24-donate-ff69b4.svg)](https://github.com/vouv/donate)

> A efficient client for BIT campus network

北京理工大学校园网命令行登录工具
- 支持linux、maxOS、windows
- 基于Go语言实现

## Update Log

2020.9.6

- 修复一些bug
- 移除不用的移动联通登录模式

2019.11.16

- 更新安装方式
- 优化项目api与项目结构

2019.9.10

- 修改优化登录逻辑
- 修复一些bug

2019.1.3
- 实现无缓冲输入密码（在macOS上测试通过）
- 修复宿舍无法登录移动网的bug

2018.11.24
- 增加登出功能
- 增加查询流量和余额功能

2018.9.1
- 实现登录与设置账号的功能


## Install

1. go get

如果已经[安装并配置GO环境](https://golang.google.cn/doc/install), 执行如下命令即可

```bash
go get -u -v github.com/vouv/srun/cmd/srun
```

开始使用
```bash
$GOPATH/bin/srun config
```

2. Download [Release](https://github.com/vouv/srun/releases/latest)

下载并开始使用
```bash
# linux
curl -L -o srun https://github.com/vouv/srun/releases/latest/download/srun-linux
chmod +x srun
./srun config
```

## Usage

Usage: `srun [OPTIONS] COMMAND`

### Show Help

```
$ srun -h

Usage:	srun [OPTIONS] COMMAND

A efficient client for BIT campus network

Options:
  -h        Show help
  -d        Show debug message
  -v        Print version information and quit

Commands:
  login     Login Srun
  logout    Logout Srun
  info      Get Srun Info
  update    Update srun
  config    Set Username and Password


```

### Config

```
$ srun config
设置校园网账号:
>vouv
设置校园网密码:
>
INFO[2019-12-02 21:30:49] 账号密码已被保存

```

### Login

```
$ srun login
INFO[2019-12-02 21:31:17] 尝试登录
...
```

### Extra - 查看余额
```
$ srun info
已用流量: 54,418.87M
已用时长: 366小时38分48秒
账户余额: ￥19.68
```


### Contribute

> 要求先安装好golang环境 go version > 1.10

先克隆项目

```
$ git clone https://github.com/vouv/srun && cd srun
```

macOS下编译

```bash
$ make
```
或
```bash
$ make darwin
```

Windows下编译
```bash
$ make windows
```

Linux下编译
```bash
$ make linux
```

编译好的可执行文件在bin文件夹中

### About

主要功能与原理

- 本地保存账号到`$HOME/.srun/account.json`（对安全性有疑问的请自行看代码）
- 使用账号快速登录校园网，环境支持的情况下也可以一键登录



### Thanks to

- [beego](https://github.com/astaxie/beego)
- [goquery](https://github.com/PuerkitoBio/goquery)




