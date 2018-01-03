网易蜂巢开源工具 `comb` 新鲜出炉
====================

## comb 是什么
[`comb`](https://github.com/bingohuang/comb) 是一款用来管理  [`网易蜂巢`](http://c.163.com) 在线资源的 CLI 工具，[`Golang`](http://golang.org) 编写，基于 [cloudcomb-go-sdk](https://github.com/bingohuang/cloudcomb-go-sdk) 和蜂巢 [OpenAPI](https://c.163.com/wiki/index.php?title=OpenAPI%E4%BB%8B%E7%BB%8D) 开发，使用简单，支持在三大平台上运行，包括Mac、Linux和Windows。

## comb 做什么
`comb` 能方便的查看和管理  [`网易蜂巢`](http://c.163.com)的线上资源，包括用户认证（已完成）、容器管理（已完成）、集群管理、镜像管理以及秘钥管理等。同时能直接通过 `bash` 等脚本和 CI 平台结合，实现资源的自动化分配和部署，集成蜂巢的资源管理功能。

## comb 怎么用
首选阅读该项目的 [`README`](https://github.com/bingohuang/comb) 文档，概要的介绍了 `comb` 的功能以及编译、安装和使用方式。运行 `comb` 或 `comb -h` 能查看到详细的帮助文档。
```
☁  comb [master] ⚡ comb -h
NAME:
   comb - is a tool to manage CloudComb resources base on cloudcomb-go-sdk.

USAGE:
   comb [global options] command [command options] [arguments...]

VERSION:
   1.0.1 darwin/amd64 go1.6.2

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
接着你还可以继续执行子命令的帮助查询，例如：
```
☁  comb [master] ⚡ comb co -h
NAME:
   comb container - Operate containers in CLoudComb

USAGE:
   comb container [command options] [arguments...]

OPTIONS:
   -i  List all containers' images.
   -a  List all containers.
   -f  Get specified container's flow.
   -c  Create container.
   -u  Update container.
   -r  Restart container.
   -t  Tag container.
   -d  Delete container.
```
你可以直接下载 `comb` 的可执行文件，将其配置到环境变量当中。当前主要支持三大平台64位操作系统，我打包编译后上传到了 `NOS`(网易对象存储)当中。32位平台如有需要可以[`联系我`](mailto:me@bingohuang.com)或者自行编译

- [x] [Mac 64](http://nos.126.net/comb/comb_darwin_amd64)
- [x] [Linux 64](http://nos.126.net/comb/comb_linux_amd64)
- [x] [Windows 64](http://nos.126.net/comb/comb_windows_amd64.exe)

注意：下载完成之后，可以在命令行下，修改文件名，为文件加上可执行权限，并放入到你的环境变量中，我以Mac为例，演示如下：
```
☁  Downloads  mv comb_darwin_amd64 comb
☁  Downloads  chmod u+x comb
☁  Downloads  mv comb /usr/local/bin
```

或者你也可以自行编译打包，请参考工程中 [`README`](https://github.com/bingohuang/comb) 文档

一般来说，你首先需要通过 `auth` 获取 `token`，有效期为 24 小时，AppKey 和 AppSecret 请在 [`网易蜂巢控制台`](https://c.163.com/dashboard) 的 [API菜单栏](https://c.163.com/dashboard#/m/account/api/)下获取。

`comb` 认证方式如下：
```
☁  comb [master] ⚡ comb auth [AppKey] [AppSercet] 
```
或者
```
☁  comb [master] ⚡ comb auth
AppKey: xxx
AppSercet: xxx
```
接下来你就可以通过 `comb`的其他命令来查看和操作蜂巢资源，比如查看你所有创建的容器，可以用 `comb container -a` 或者 `comb co -a` 或者 `comb co`（co 是 contaienr 的缩写），如下：
```
☁  comb [master] ⚡ comb co
{"total":1,"containers":[{"id":274320,"name":"cc2048","desc":"","status":"restart_succ","replicas":1,"bandwidth":100,"charge_type":1,"spec_id":1,"created_at":"2016-05-18T07:59:46Z","updated_at":"2016-06-28T10:54:31Z","public_ip":"106.2.110.163","private_ip":"10.166.224.44","ssh_key_ids":null,"image_id":20314,"use_public_network":1,"network_charge_type":2,"env_var":{}}]} 
```
查看特定容器的信息(274320为以上显示的容器 ID），借助python工具可以格式化 JSON 格式：
```
☁  comb [master] ⚡ comb co 274320 | python -mjson.tool
{
    "bandwidth": 100,
    "charge_type": 1,
    "created_at": "2016-05-18T07:59:46Z",
    "desc": "",
    "env_var": {},
    "id": 274320,
    "image_id": 20314,
    "name": "cc2048",
    "network_charge_type": 2,
    "private_ip": "10.166.224.44",
    "public_ip": "106.2.110.163",
    "replicas": 1,
    "spec_id": 1,
    "ssh_key_ids": null,
    "status": "restart_succ",
    "updated_at": "2016-06-28T10:54:31Z",
    "use_public_network": 1
}
```
查看 `comb` 的帮助文档，看到 `comb` 共支持五个子命令：
```
COMMANDS:
     auth           Auth in CloudComb with app key, app secret
     container, co  Operate containers in CloudComb
     cluster,   cu  Operate clusters in CloudComb
     repositry, re  Operate repositrys in CloudComb
     secretkey, sk  Operate secret keys in CloudComb
```
当前已经完整支持 `auth` 及 `container`，其他命令支持部分选项，比如 `-a`，完整的支持将在后续继续完善。

## 为什么开发 comb
1. 物尽其用，推广网易蜂巢。网易蜂巢开发的 OpenAPI 文档详细，使用方便，希望能将其推广到各个语言、各个平台上用起来。有兴趣者欢迎给来提 PR 或 Issue，或者开发更多语言版本的SDK及应用工具，并[`告知我们`](mailto:cloudcomb@188.com)。
2. 玩起来方便。网易蜂巢 定位于做一个专业的容器云平台，为开发、测试、运维者提供资源和平台，加快研发流程。`comb` 方便大家在命令行界面下，查看和操作蜂巢资源的使用情况，玩起来顺手，同时还能方便的和 CI 平台（比如 Jenkins）结合。
3. 上手 Golang 开发。这是我的第一个 Golang 实践项目，我认为 `Learning by doing` 是最快的学习方式。

## 为什么取名 comb
 [`网易蜂巢`](http://c.163.com) 的英文名为 `cloudcomb`，comb 既有`蜂巢`的意思，又能表示为`梳子` 。像梳子一样梳理网易蜂巢的资源，故取名 'comb'。

## 继续完善 comb 
当前 comb 主要完成了整体CLI框架的搭建、用户认证以及容器管理的内容。架构非常简单，分为 SDK 和 CLI 两个部分，SDK 主要负责了封装 OpenAPI 的实现和细节，提供给其他客户端调用，CLI 就是基于 SDK，提供命令行的操作方式。

从功能上说，comb 还需要继续完成 `集群管理、镜像管理以及秘钥管理`等内容，如果后续OpenAPI有更新，也需要继续跟进最新的内容。

从使用上说，comb 是一款 CLI 工具，借助 `命令行+选项+参数` 来操作。其中复杂的输入参数是需要转换为JSON格式，输入范例请参考 OpenAPI 文档，从体验上来说略有不足;复杂输出也以 JSON 格式为主，当前 JSON 格式未做格式化，有待后续完善。如果希望输出格式化的 JSON，一种折中的方式是通过python工具做个转换：
```
☁  comb [master] ⚡ comb co | python -mjson.tool
```

从工程上说，comb 是一款开源软件，托管到 [`Github`](https://github.com/bingohuang/comb) 上，采用 Apache License 2.0 协议。代码尽量规范，注释尽量清晰，README、CHANGELOG和ROADMAP齐全，同时书写了单元测试，后续将加入 `travis` 做持续构建。

总的来说，`comb` 还有不少需要完善的地方，具体可以参见项目的 [`ROADMAP`](https://github.com/bingohuang/comb/blob/master/ROADMAP.md)，欢迎任何人给 `comb` 提 `issue` 或 `PR`，我将会一一查看并回复，`PR` 审核通过将会合入主代码分支。

心动不如行动，既可以练手 Golang，还能更好的熟悉容器云平台的运作，保不准你还从此走上了 `'Golang+云平台开发大神'`之路哦。在此之前，还请大神给 `comb` 点个赞吧！([请猛戳此处](https://github.com/bingohuang/comb))

## 附上相关地址：
* comb 工具： https://github.com/bingohuang/comb
* Golang SDK:  https://github.com/bingohuang/cloudcomb-go-sdk
* 网易蜂巢 OpenAPI：https://c.163.com/wiki/index.php?title=OpenAPI%E4%BB%8B%E7%BB%8D
* 网易蜂巢 平台： http://c.163.com

