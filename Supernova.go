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
	"log"
	"os"
	"strings"
)

// Structure
type FlagOptions struct {
	outFile    string
	inputFile  string
	language   string
	encryption string
	// obfuscation string
	variable string
	key      int
	debug    bool
	guide    bool
	version  bool
}

// global variables
var (
	__version__         = "1.2.0"
	__license__         = "MIT"
	__author__          = "@nickvourd"
	__original_github__ = "https://github.com/nickvourd/Supernova"
	__github__          = "https://github.com/yutianqaq/Supernova_CN"
	__coauthors__       = [3]string{"@yutianqaq", "@Papadope9", "@0xvm"}
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
	inputFile := flag.String("i", "", "64 位原始格式 Shellcode 的路径")
	encryption := flag.String("enc", "", "Shellcode加密方式 (例如, ROT, XOR, RC4, AES, CHACHA20, B64XOR, B64RC4, B64AES, B64CHACHA20)")
	language := flag.String("lang", "", "转换(Raw, Nim, Rust, C, CSharp, Go)格式的 Shellcode")
	outFile := flag.String("o", "", "输出到文件")
	variable := flag.String("v", "shellcode", "Shellcode 的变量名称")
	debug := flag.Bool("d", false, "开启 Debug 模式")
	key := flag.Int("k", 1, "加密的密钥长度")
	version := flag.Bool("version", false, "展示 Supernova 当前的版本")
	guide := flag.Bool("guide", false, "开启引导模式（输出解密代码）")
	// obfuscation := flag.String("obf", "", "Shellcode obfuscation (i.e., IPv4, IPv6, MAC, UUID)")
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

	// Check for valid values of language argument
	foundLanguage := Arguments.ValidateArgument("lang", options.language, []string{"Nim", "Rust", "C", "CSharp", "Go", "Python", "RAW"})

	// Check if the encryption or obfuscation option is not used
	if options.encryption == "" {
		// && options.obfuscation == ""
		logger := log.New(os.Stderr, "[!] ", 0)
		logger.Fatal("请先选择加密选项！\n")
		// logger.Fatal("Please choose either the encryption option or the obfuscation option to proceed!\n")
	}

	// Check if encryption and obfuscation are used together
	// if options.encryption != "" && options.obfuscation != "" {
	//	logger := log.New(os.Stderr, "[!] ", 0)
	//	logger.Fatal("You cannot choose both the encryption and obfuscation options; please select only one!\n")
	//}

	// Call function named ConvertShellcode2Hex
	convertedShellcode, payloadLength := Converters.ConvertShellcode2Hex(rawShellcode, foundLanguage)

	// Print payload size and chosen language
	fmt.Printf("[+] Payload大小：%d bytes\n\n[+] 已将Payload转换为 %s 语言\n\n", payloadLength, foundLanguage)

	if options.debug {
		// Call function named ConvertShellcode2Template
		template := Converters.ConvertShellcode2Template(convertedShellcode, foundLanguage, payloadLength, options.variable)

		// Print original template
		fmt.Printf("[+] 原始Payload:\n\n%s\n\n", template)
	}

	// Encryption option is enable
	if options.encryption != "" {
		// Call function named ValidateArgument
		Arguments.ValidateArgument("enc", options.encryption, []string{
			"RAW", "XOR", "RC4", "AES", "ROT", "CHACHA20",
			"B64XOR", "B64RC4", "B64AES", "B64CHACHA20",
		})

		// Call function ValidateKeySize
		options.key = Arguments.ValidateKeySize(options.key, options.encryption)

		// Call function named DetectEncryption
		encryptedShellcode, encryptedLength, key, passphrase, iv := Encryptors.DetectEncryption(options.encryption, rawShellcode, options.key, foundLanguage)

		// Call function named ConvertShellcode2Template
		template := Converters.ConvertShellcode2Template(encryptedShellcode, foundLanguage, encryptedLength, options.variable)

		// Print encrypted template
		if encryptedLength < 500 && options.outFile == "" {
			fmt.Printf("[+] Payload 加密后 %s:\n\n%s\n\n", strings.ToUpper(options.encryption), template)
		} else if encryptedLength > 500 && options.outFile == "" {
			fmt.Printf("[+] Payload 加密后 %s(%d): \n\n", strings.ToUpper(options.encryption), encryptedLength)
			fmt.Printf("[!] Payload 尺寸太长，请使用 -o filename!\n")
		} else {
			fmt.Printf("[+] Payload 加密后 %s(%d bytes): \n\n", strings.ToUpper(options.encryption), encryptedLength)
		}

		// Guide option is enable
		if options.guide {
			Decryptors.DecryptorsTemplates(foundLanguage, options.encryption, options.variable, options.key, encryptedLength, encryptedShellcode, key, passphrase, iv)
		}

		// Outfile option is enable
		if options.outFile != "" {
			language := strings.ToLower(options.language)
			if language == "raw" {
				err := Output.SaveShellcodeToFile(template, options.outFile)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			} else {
				err := Output.SaveOutputToFile(template, options.outFile)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
		}
	}

	// Obfuscation option is enable
	// if options.obfuscation != "" {
	// Call function named ValidateArgument
	// Arguments.ValidateArgument("obf", options.obfuscation, []string{"IPv4", "IPv6", "MAC", "UUID"})
	// }
}
