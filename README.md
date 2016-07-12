CloudComb CLI tool: `comb`
==========================


## Get Started

`comb` is a CLI tool for manage resources in [CloudComb](http://c.163.com) base on [cloudcomb-go-sdk](https://github.com/bingoHuang/cloudcomb-go-sdk). Support Mac, Linux and Windows.

We had better read `comb` or `comb -h` to get more instruments.

```
☁  comb [master] ⚡ comb -h
NAME:
   comb - is a tool to manage CloudComb resources base on cloudcomb-go-sdk.

USAGE:
   comb [global options] command [command options] [arguments...]

VERSION:
   0.0.3 darwin/amd64 go1.6.2

AUTHOR(S):
   Bingo Huang <me@bingohuang.com>

COMMANDS:
     auth           Auth in CloudComb with app key, app secret
     container, co  Operate containers in CloudComb
     cluster,   cu  Operate clusters in CloudComb
     repositry, re  Operate repositrys in CloudComb
     secretkey, sk  Operate secret keys in CloudComb

GLOBAL OPTIONS:
   --debug                      debug mode [$DEBUG]
   --log-level value, -l value  Log level (options: debug, info, warn, error, fatal, panic) (default: "info")
   --help, -h                   show help
   --version, -v                print the version
```

Then enjoy your journey with `comb` if you happen to find some command are helpful to you.

## Build and Install `comb`

### Build in current directory
```
git clone https://github.com/bingoHuang/comb.git
cd comb
export GOPATH=`pwd`
go get -d
go build

cp comb /usr/local/bin

```

Now `comb` is in your PATH.

### Build in $GOPATH with `gvm`(**Recommend**)

1. install [`gvm`](https://github.com/moovweb/gvm)
2. install `go1.6.2`(**Recommend**)
```
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.6.2
gvm use go1.6.2 --default
```

3. get and build `comb`
```
go get github.com/bingoHuang/comb
```
Now `comb` is in your `$GOPATH/bin`.

### Or you can download the `comb` executable binary file from blew:

- [x] [Mac 64](http://nos.126.net/comb/comb_darwin_amd64_1.0.1)
- [x] [Linux 64](http://nos.126.net/comb/comb_linux_amd64_1.0.1)
- [x] [Windows 64](http://nos.126.net/comb/comb_windows_amd64_1.0.1.exe)

## Usage

## Participating

You can contribute to `comb` in several different ways:

* To report a problem or request a feature, please feel free to file an issue.

* Of course, we welcome pull requests and patches. Setting up a local `comb` development environment and submitting PRs is described here.


## Copyright and license
Copyright © 2016. All rights reserved