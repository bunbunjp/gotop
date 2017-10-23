# gotop
[![version](https://img.shields.io/badge/version-v1.0.0-orange.svg)](https://github.com/bunbunjp/gotop/releases/tag/v1.0.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/bunbunjp/gotop)](https://goreportcard.com/report/github.com/bunbunjp/gotop)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bunbunjp/gotop/blob/master/LICENSE)
[![platform](https://img.shields.io/badge/platform-linux_64%20%7C%20OSX_64-lightgray.svg)]()

The gotop is a system monitoring dashboard for terminal console.   
![](./gotop_movie.gif)   
Developed with reference to the [gtop](https://github.com/aksakalli/gtop), I got inspiration hints. Thank you [@aksakalli](https://github.com/aksakalli).   
The gotop is reimplement gtop with golang. So gotop is lightweight and quick.   
gtop looks better, gotop does looks bad but it is light.   


# Install
```shell
go get github.com/bunbunjp/gotop
```

# Usage
exec `gotop` command.
- c : Process list sort as CPU.
- p : Process list sort as ProcessID.
- m : Process list sort as Memory usage.
- ↑↓ : Process list scroll up or down.

# Setup for developer
```shell
$ curl https://glide.sh/get | sh # need glide
$ git clone ...
$ cd $GOPATH/src/github.com/forkuser/gotop # please clone here
$ glide install -v
$ go run main.go # launch
```

# Requirements
- OSX, Linux supported
- go >= 1.9


# todo
- [ ] プロセスを選択してKillを実行する
- [ ] 高さの動的調整
- [ ] Windows対応
- [ ] 長時間稼働し続けてもメモリ使用量が増えないように考慮
