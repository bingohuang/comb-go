CloudComb CLI tool: `comb`
==========================


## Get Started

`comb` is a CLI tool for manage resources in [CloudComb](http://c.163.com) base on [cloudcomb-go-sdk](https://github.com/bingoHuang/cloudcomb-go-sdk). Mac, Linux, Windows supported.

We had better read `comb -h` to get more instruments.

```
☁  cloudcomb-go-cli [master] ⚡ comb -h
NAME:
   comb - is a tool for manage resources in [CloudComb](http://c.163.com)
        base on [cloudcomb-go-sdk](https://github.com/bingoHuang/cloudcomb-go-cli)

USAGE:
   comb [global options] command [command options] [arguments...]

VERSION:
   0.0.3 darwin/amd64 go1.6.2

AUTHOR(S):
   Bingo Huang <bingo@xbing.me>

COMMANDS:
     auth            Auth in CloudComb with app key, app secret
     container, co   Container related API
     cluster, cl     Cluster related API
     repository, re  Repository related API
     secretkey, se   Sercet key related API

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Then enjoy your journey with `comb` if you happen to find some command are helpful to you.

## Build and Install `comb`
```
git clone https://github.com/bingoHuang/cloudcomb-go-cli.git
cd cloudcomb-go-cli
export GOPATH=`pwd`
go get -d
go build -o comb

cp comb /usr/local/bin

```

Now `comb` is in your PATH.

Or you can download the `comb` executable binary file from blew:

- [x] [Mac 64位](http://nos.126.net/cloudadmin/comb)
- [ ] [Linux 64位](#)
- [ ] [Linux 32位](#)
- [ ] [Windows 64位](#)
- [ ] [Windows 32位](#)

## Usage

## Participating

You can contribute to `comb` in several different ways:

* To report a problem or request a feature, please feel free to file an issue.

* Of course, we welcome pull requests and patches. Setting up a local `comb` development environment and submitting PRs is described here.


## Copyright and license
Copyright © 2016. All rights reserved