# Flylog

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/flylog/flylog?color=peru)](https://github.com/flylog/flylog/releases/latest)
[![Released API docs](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/flylog/flylog)
[![Build](https://img.shields.io/github/workflow/status/flylog/flylog/Go.svg)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/flylog/flylog?color=pink)](https://goreportcard.com/report/github.com/flylog/flylog)
[![Lines of code](https://img.shields.io/tokei/lines/github/flylog/flylog.svg?color=beige)](#)
[![Downloads of releases](https://img.shields.io/github/downloads/flylog/flylog/total.svg?color=lavender)](https://github.com/flylog/flylog/releases/latest)
[![Languages](https://img.shields.io/github/languages/top/flylog/flylog.svg?color=yellow)](#)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/flylog/flylog)](#)
[![GPL3 licensed](https://img.shields.io/github/license/flylog/flylog.svg)](./LICENSE)

Flylog 是一个简单快速的开发日志模块。支持同步/异步输出，缓存日志批量写入支持多种输出方式控制台，文件，tcp等。

Flylog is a simple and fast development log module. Support synchronous / asynchronous output, cache log batch writing, and multiple output modes, console, file, TCP, etc.

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/flylog/flylog)

# Install
    
    go get -u github.com/flylog/flylog@latest

# ToDo

- [ ] 增加日志输出到TCP支持。 
- [ ] Console日志彩色输出由https://github.com/flylog/colorstyle 支持开发中 
- [ ] 日志异步缓存输出参考heka。 
- [ ] 优化lylog.properties配置项,参考Log4j配置管理。
- [ ] 日志输出格式自定义Formart 
- [ ] 增加http访问日志参考nginx。
