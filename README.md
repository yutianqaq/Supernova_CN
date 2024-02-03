# Supernova_CN
真正的 Shellcode 加密器（对比原版新增加了  CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20 等加密方式和汉化帮助信息）感谢原作者 [@nickvourd](https://github.com/nickvourd) 的 [Supernova](https://github.com/nickvourd/Supernova) 以及他们的共同编写者

原始版本：https://github.com/nickvourd/Supernova


## Description
Supernova 是一个开源的 Golang 工具，它使用户能够安全地加密原始 shellcode。此外，它还提供将加密的 shellcode 自动转换为与各种编程语言兼容的格式，包括:
- C
- C#
- Rust
- Nim
- Go

它支持多种不同的密码，包括:
- ROT
- XOR
- RC4
- AES
- CHACHA20
- Base64XOR
- Base64XOR
- Base64RC4
- Base64AES
- Base64CHACHA20

此外，该工具使用选定的密码和语言生成解密函数，同时还提供如何有效使用它的说明。

Supernova 是用跨平台语言 Golang 编写的，可以在 Windows 和 Linux 系统上使用。

## 许可证

本工具使用 [![License: MIT](https://img.shields.io/badge/MIT-License-yellow.svg)](LICENSE) 许可证.

## 安装

要安装Supernova，请运行以下命令，或者下载[发行版本](https://github.com/yutianqaq/Supernova_CN/releases)
```
go build Supernova.go
```

### 命令行帮助

```


███████╗██╗   ██╗██████╗ ███████╗██████╗ ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗████╗  ██║██╔═══██╗██║   ██║██╔══██╗
███████╗██║   ██║██████╔╝█████╗  ██████╔╝██╔██╗ ██║██║   ██║██║   ██║███████║
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗██║╚██╗██║██║   ██║╚██╗ ██╔╝██╔══██║
███████║╚██████╔╝██║     ███████╗██║  ██║██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

Supernova v1.0.0 - 真正的Shellcode加密器。
Supernova是一个开源工具，受MIT许可证保护。
由@nickvourd、@0xvm、@Papadope9和@yutianqaq用<3编写...
原版请访问https://github.com/nickvourd/Supernova了解更多信息...
汉化版本https://github.com/yutianqaq/Supernova_CN

Usage of Suprenova.exe:
  -d    开启 Debug 模式
  -enc string
        Shellcode加密方式 (例如, ROT, XOR, RC4, AES, CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20)
  -guide
        开启引导模式
  -i string
        64位原始格式 Shellcode 的路径
  -k int
        加密的密钥长度 (default 1)
  -lang string
        转换(Nim, Rust, C, CSharp, Go)格式的Shellcode
  -o string
        输出文件名
  -v string
        Shellcode 的变量名称 (default "shellcode")
  -version
        展示 Supernova 当前的版本
```



## 参考

- [D3Ext/maldev: Golang library for malware development](https://github.com/D3Ext/maldev)

- [nickvourd/Supernova: Real fucking shellcode encryption tool](https://github.com/nickvourd/Supernova)
