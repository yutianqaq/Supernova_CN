package main

import (
	"Supernova/Arguments"
	"Supernova/Converters"
	"Supernova/Decryptors"
	"Supernova/Encryptors"
	"Supernova/Output"
	"Supernova/Utils"
	"flag"
	"fmt"
	"strings"
)

// Structure
type FlagOptions struct {
	outFile    string
	inputFile  string
	language   string
	encryption string
	//obfuscation string
	variable string
	key      int
	debug    bool
	guide    bool
	version  bool
}

// global variables
var (
	__version__   = "1.0.0"
	__license__   = "MIT"
	__author__    = "@nickvourd"
	__original_github__    = "https://github.com/nickvourd/Supernova"
	__github__    = "https://github.com/yutianqaq/Supernova_CN"
	__coauthors__ = [3]string{"@yutianqaq", "@Papadope9", "@0xvm"}
)

var __ascii__ = `

███████╗██╗   ██╗██████╗ ███████╗██████╗ ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗ 
██╔════╝██║   ██║██╔══██╗██╔════╝██╔══██╗████╗  ██║██╔═══██╗██║   ██║██╔══██╗
███████╗██║   ██║██████╔╝█████╗  ██████╔╝██╔██╗ ██║██║   ██║██║   ██║███████║
╚════██║██║   ██║██╔═══╝ ██╔══╝  ██╔══██╗██║╚██╗██║██║   ██║╚██╗ ██╔╝██╔══██║
███████║╚██████╔╝██║     ███████╗██║  ██║██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
╚══════╝ ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

Supernova v%s - 真正的Shellcode加密器。
Supernova是一个开源工具，受%s许可证保护。
由%s、%s、%s和%s用<3编写...
原版请访问%s了解更多信息...
汉化版本%s

`

// Options function
func Options() *FlagOptions {
	inputFile := flag.String("i", "", "64位原始格式 Shellcode 的路径")
	encryption := flag.String("enc", "", "Shellcode加密方式 (例如, ROT, XOR, RC4, AES, CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20)")
	language := flag.String("lang", "", "转换(Nim, Rust, C, CSharp, Go)格式的Shellcode")
	outFile := flag.String("o", "", "输出文件名")
	variable := flag.String("v", "shellcode", "Shellcode 的变量名称")
	debug := flag.Bool("d", false, "开启 Debug 模式")
	key := flag.Int("k", 1, "加密的密钥长度")
	version := flag.Bool("version", false, "展示 Supernova 当前的版本")
	guide := flag.Bool("guide", false, "开启引导模式")
	flag.Parse()

	return &FlagOptions{outFile: *outFile, inputFile: *inputFile, language: *language, encryption: *encryption, variable: *variable, debug: *debug, key: *key, version: *version, guide: *guide}
}

// main function
func main() {
	// Print ascii
	fmt.Printf(__ascii__, __version__, __license__, __author__, __coauthors__[2], __coauthors__[1], __coauthors__[0], __original_github__, __github__)

	// Check GO version of the current system
	Utils.Version()

	// Retrieve command-line options using the Options function
	options := Options()

	// Check Arguments Length
	Arguments.ArgumentLength(options.version)

	// Check Version of tool
	Arguments.ShowVersion(__version__, options.version)

	// Call function named ArgumentEmpty
	Arguments.ArgumentEmpty(options.inputFile, 1)

	// Call function name ConvertShellcode2String
	rawShellcode, err := Converters.ConvertShellcode2String(options.inputFile)
	if err != nil {
		fmt.Println("[!] Error:", err)
		return
	}

	// Call function named ArgumentEmpty
	Arguments.ArgumentEmpty(options.language, 2)

	// Call function named ArgumentEmpty
	Arguments.ArgumentEmpty(options.encryption, 3)

	// Call function ValidateKeySize
	Arguments.ValidateKeySize(options.key, options.encryption)

	// Check for valid values of language argument
	foundLanguage := Arguments.ValidateArgument("lang", options.language, []string{"Nim", "Rust", "C", "CSharp", "Go"})

	// Call function named ConvertShellcode2Hex
	convertedShellcode, payloadLength := Converters.ConvertShellcode2Hex(rawShellcode, foundLanguage)

	// Print payload size and choosen language
	fmt.Printf("[+] Payload size: %d bytes\n\n[+] Converted payload to %s language\n\n", payloadLength, foundLanguage)

	if options.debug {
		// Call function named ConvertShellcode2Template
		template := Converters.ConvertShellcode2Template(convertedShellcode, foundLanguage, payloadLength, options.variable)

		// Print original template
		fmt.Printf("[+] The original payload:\n\n%s\n\n", template)
	}

	// Encryption option is enable
	if options.encryption != "" {
		// Call function named ValidateArgument
		Arguments.ValidateArgument("enc", options.encryption, []string{
			"XOR", "RC4", "AES", "ROT", "CHACHA20",
			"B64XOR", "B64RC4", "B64AES", "B64CHACHA20",
		})

		// Call function named DetectEncryption
		encryptedShellcode, encryptedLength, key, passphrase, iv := Encryptors.DetectEncryption(options.encryption, rawShellcode, options.key)

		// Call function named ConvertShellcode2Template
		template := Converters.ConvertShellcode2Template(encryptedShellcode, foundLanguage, encryptedLength, options.variable)

		// Print encrypted template
		fmt.Printf("[+] The encrypted payload with %s:\n\n%s\n\n", strings.ToLower(options.encryption), template)

		// Guide option is enable
		if options.guide {
			Decryptors.DecryptorsTemplates(foundLanguage, options.encryption, options.variable, options.key, encryptedLength, encryptedShellcode, key, passphrase, iv)
		}

		// Outfile option is enable
		if options.outFile != "" {
			err := Output.SaveOutputToFile(template, options.outFile)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
	}
}
