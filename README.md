# Supernova_CN
真正的 Shellcode 加密器（对比原版新增加了  CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20 等加密方式和汉化帮助信息）感谢原作者 [@nickvourd](https://github.com/nickvourd) 的 [Supernova](https://github.com/nickvourd/Supernova) 以及他们的共同编写者。

原始版本：https://github.com/nickvourd/Supernova


## 介绍
Supernova 是一个开源的 Golang 工具，它使用户能够**方便地**加密原始 shellcode。此外，它还提供将加密的 shellcode 自动转换为与各种编程语言兼容的格式，并提供了**输出对应的解密代码**(引导模式)，包括:
- C
- C#
- Rust
- Nim
- Go
- Raw

它支持多种不同的密码，包括:

ROT、XOR、RC4、AES、CHACHA20、Base64XOR、Base64XOR、Base64RC4、Base64AES、Base64CHACHA20

# 能力表

Supernova 是用跨平台语言 Golang 编写的，可以在 Windows、 Linux 和 MacOS（未测试） 系统上使用。

该工具使用选定的密码和语言生成解密函数，同时还提供如何有效使用它的说明。下面是能力表

| 格式 | ROT  | XOR  | RC4  | AES  | CHACHA20 | BASE64XOR | BASE64RC4 | BASE64AES | BASE64CHACHA20 |
| ---- | ---- | ---- | ---- | ---- | -------- | --------- | --------- | --------- | -------------- |
| C    | ✔️    | ✔️    | ✔️    | ✔️    | ❌        | ❌         | ❌         | ❌         | ❌              |
| C#   | ✔️    | ✔️    | ✔️    | ✔️    | ❌        | ❌         | ❌         | ❌         | ❌              |
| Rust | ✔️    | ✔️    | ✔️    | ✔️    | ❌        | ❌         | ❌         | ❌         | ❌              |
| Nim  | ✔️    | ✔️    | ✔️    | ✔️    | ❌        | ❌         | ❌         | ❌         | ❌              |
| Go   | ✔️    | ✔️    | ✔️    | ✔️    | ✔️        | ✔️         | ✔️         | ✔️         | ✔️              |
| Raw  | ✔️    | ✔️    | ✔️    | ✔️    | ✔️        | ✔️         | ✔️         | ✔️         | ✔️              |

✔️表示以支持

❌表示暂未支持



## 许可证

本工具使用 [![License: MIT](https://img.shields.io/badge/MIT-License-yellow.svg)](LICENSE) 许可证。

## 安装

要安装Supernova，请运行以下命令，或者下载[发行版本](https://github.com/yutianqaq/Supernova_CN/releases)。
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

Supernova v1.2.0 - 真正的Shellcode加密器。
Supernova是一个开源工具，受MIT许可证保护。
由@nickvourd、@0xvm、@Papadope9和@yutianqaq用<3编写...
原版请访问https://github.com/nickvourd/Supernova了解更多信息...
汉化版本https://github.com/yutianqaq/Supernova_CN

Usage of Suprenova.exe:
  -d    开启 Debug 模式
  -enc string
        Shellcode加密方式 (例如, ROT, XOR, RC4, AES, CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20)
  -guide
        开启引导模式（输出解密代码）
  -i string
        64 位原始格式 Shellcode 的路径
  -k int
        加密的密钥长度 (default 1)
  -lang string
        转换(Raw, Nim, Rust, C, CSharp, Go)格式的 Shellcode
  -o string
        输出到文件
  -v string
        Shellcode 的变量名称 (default "shellcode")
  -version
        展示 Supernova 当前的版本
```

# 使用方式

## 仅加密 Shellcode - XOR

这将输出 key 长度为 10 的多字节 xor shellcode。

```
msfvenom -p windows/x64/exec CMD="calc.exe" -f raw -o calc.bin

./Supernova -enc xor -i calc.bin -k 10 -lang go -o calc_xor10.bin


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

[+] Payload size: 276 bytes

[+] Converted payload to go language

[+] Generated XOR key: 0x86, 0xbe, 0xa8, 0x2e, 0x2b, 0x21, 0x42, 0xa6, 0x3a, 0x56

[+] Encrypted payload with xor

[+] Save encrypted shellcode file to /root/Tools-dev/Test/calc_xor10.bin

```

## 加密并输出解密方式 - XOR

追加 -guide 参数将输出加密后的 Shellcode 并输出加密方式。

```
./Supernova -enc xor -i calc.bin -k 10 -lang go -o calc_xor10.bin -guide


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

[+] Payload size: 276 bytes

[+] Converted payload to go language

[+] Generated XOR key: 0xf8, 0x4f, 0x2c, 0x15, 0xd6, 0xcd, 0x30, 0x9f, 0xa0, 0xd1

[+] Encrypted payload with xor

[+] XOR decrytpion function has been saved to Program.go file

[+] Save encrypted shellcode file to /root/Tools-dev/Test/calc_xor10.bin

```



**Program.go**

```
┌──(root㉿kali)-[~/Tools-dev/Test]
└─#go mod init xor10

┌──(root㉿kali)-[~/Tools-dev/Test]
└─# go build Program.go                  

┌──(root㉿kali)-[~/Tools-dev/Test]
└─# ./Program 
XOR  Decrypted Payload:
0xFC, 0x48, 0x83, 0xE4, 0xF0, 0xE8, 0xC0, 0x00, 0x00, 0x00, 0x41, 0x51, 0x41, 0x50, 0x52, 0x51, 0x56, 0x48, 0x31, 0xD2, 0x65, 0x48, 0x8B, 0x52, 0x60, 0x48, 0x8B, 0x52, 0x18, 0x48, 0x8B, 0x52, 0x20, 0x48, 0x8B, 0x72, 0x50, 0x48, 0x0F, 0xB7, 0x4A, 0x4A, 0x4D, 0x31, 0xC9, 0x48, 0x31, 0xC0, 0xAC, 0x3C, 0x61, 0x7C, 0x02, 0x2C, 0x20, 0x41, 0xC1, 0xC9, 0x0D, 0x41, 0x01, 0xC1, 0xE2, 0xED, 0x52, 0x41, 0x51, 0x48, 0x8B, 0x52, 0x20, 0x8B, 0x42, 0x3C, 0x48, 0x01, 0xD0, 0x8B, 0x80, 0x88, 0x00, 0x00, 0x00, 0x48, 0x85, 0xC0, 0x74, 0x67, 0x48, 0x01, 0xD0, 0x50, 0x8B, 0x48, 0x18, 0x44, 0x8B, 0x40, 0x20, 0x49, 0x01, 0xD0, 0xE3, 0x56, 0x48, 0xFF, 0xC9, 0x41, 0x8B, 0x34, 0x88, 0x48, 0x01, 0xD6, 0x4D, 0x31, 0xC9, 0x48, 0x31, 0xC0, 0xAC, 0x41, 0xC1, 0xC9, 0x0D, 0x41, 0x01, 0xC1, 0x38, 0xE0, 0x75, 0xF1, 0x4C, 0x03, 0x4C, 0x24, 0x08, 0x45, 0x39, 0xD1, 0x75, 0xD8, 0x58, 0x44, 0x8B, 0x40, 0x24, 0x49, 0x01, 0xD0, 0x66, 0x41, 0x8B, 0x0C, 0x48, 0x44, 0x8B, 0x40, 0x1C, 0x49, 0x01, 0xD0, 0x41, 0x8B, 0x04, 0x88, 0x48, 0x01, 0xD0, 0x41, 0x58, 0x41, 0x58, 0x5E, 0x59, 0x5A, 0x41, 0x58, 0x41, 0x59, 0x41, 0x5A, 0x48, 0x83, 0xEC, 0x20, 0x41, 0x52, 0xFF, 0xE0, 0x58, 0x41, 0x59, 0x5A, 0x48, 0x8B, 0x12, 0xE9, 0x57, 0xFF, 0xFF, 0xFF, 0x5D, 0x48, 0xBA, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x48, 0x8D, 0x8D, 0x01, 0x01, 0x00, 0x00, 0x41, 0xBA, 0x31, 0x8B, 0x6F, 0x87, 0xFF, 0xD5, 0xBB, 0xF0, 0xB5, 0xA2, 0x56, 0x41, 0xBA, 0xA6, 0x95, 0xBD, 0x9D, 0xFF, 0xD5, 0x48, 0x83, 0xC4, 0x28, 0x3C, 0x06, 0x7C, 0x0A, 0x80, 0xFB, 0xE0, 0x75, 0x05, 0xBB, 0x47, 0x13, 0x72, 0x6F, 0x6A, 0x00, 0x59, 0x41, 0x89, 0xDA, 0xFF, 0xD5, 0x63, 0x61, 0x6C, 0x63, 0x2E, 0x65, 0x78, 0x65, 0x00
```





## 参考

- [D3Ext/maldev: Golang library for malware development](https://github.com/D3Ext/maldev)

- [nickvourd/Supernova: Real fucking shellcode encryption tool](https://github.com/nickvourd/Supernova)
