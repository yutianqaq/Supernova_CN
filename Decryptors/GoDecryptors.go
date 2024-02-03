package Decryptors

// go b64xor
var __go_b64xor__ = `
package main

import (
	"fmt"
	"encoding/base64"
)

func XOREncryption(shellcode []byte, key []byte) []byte {
	encrypted := make([]byte, len(shellcode))
	keyLen := len(key)

	for i := 0; i < len(shellcode); i++ {
		encrypted[i] = shellcode[i] ^ key[i%%keyLen]
	}

	return encrypted
}

func main() {
	%s := []byte{%s}
	key := []byte { %s }

	decryptedShellcodeBase64, _ := base64.StdEncoding.DecodeString(string(%s))
	decryptedShellcode := XOREncryption(decryptedShellcodeBase64, key)

	fmt.Print("B64XOR Decrypted Payload:\n")
	for i, b := range decryptedShellcode {
		fmt.Printf("0x%%02X", b)
		if i < len(decryptedShellcode)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

}
`

// go b64aes
var __go_b64aes__ = `
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// PKCS7Unpadding function
func PKCS7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("Invalid padding")
	}
	return data[:length-unpadding], nil
}

// AESDecryption function
func AESDecryption(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new CBC mode decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// Remove PKCS7 padding
	unpaddedData, err := PKCS7Unpadding(decrypted)
	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}

func main() {
	%s := []byte{%s}
	key := []byte{%s}
	iv := []byte{%s}

	// Decryption
	decryptedShellcodeBase64, _ := base64.StdEncoding.DecodeString(string(%s))
	decrypted, err := AESDecryption(key, iv, decryptedShellcodeBase64)
	if err != nil {
		panic(err)
	}

	fmt.Print("B64aes Decrypted Payload:\n")
	for i, b := range decrypted {
		fmt.Printf("0x%%02X", b)
		if i < len(decrypted)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
`

// go b64rc4
var __go_b64rc4__ = `
package main

import (
	"encoding/base64"

	"fmt"
)

type Rc4Context struct {
	i uint32
	j uint32
	s [256]uint8
}

func RC4Encryption(data []byte, key []byte) []byte {
	var s [256]byte
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + int(s[i]) + int(key[i%%len(key)])) %% 256
		s[i], s[j] = s[j], s[i]
	}

	encrypted := make([]byte, len(data))
	i, j := 0, 0
	for k := 0; k < len(data); k++ {
		i = (i + 1) %% 256
		j = (j + int(s[i])) %% 256
		s[i], s[j] = s[j], s[i]
		encrypted[k] = data[k] ^ s[(int(s[i])+int(s[j]))%%256]
	}

	return encrypted
}

func main() {
	%s := []byte{%s}

	key := []byte("%s")
	decryptedShellcodeBase64, _ := base64.StdEncoding.DecodeString(string(%s))
	decryptedShellcode := RC4Encryption(decryptedShellcodeBase64, key)

	fmt.Print("B64rc4 Decrypted Payload:\n")
	for i, b := range decryptedShellcode {
		fmt.Printf("0x%%02X", b)
		if i < len(decryptedShellcode)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

}
`

// go chacha20
var __go_chacha20__ = `
package main

import (
	"golang.org/x/crypto/chacha20poly1305"
	"fmt"
)

func Chacha20Decrypt(data []byte, key []byte) ([]byte) {
	aead, err := chacha20poly1305.NewX(key)


	nonceSize := aead.NonceSize()

	if len(data) < nonceSize {
		return nil
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		if err.Error() == "chacha20poly1305: message authentication failed" {
			return nil
		}

		return nil
	}

	return plaintext
}


func main() {
	%s := []byte{ %s }
	key := []byte { %s }

	decryptedShellcode := Chacha20Decrypt(%s, key)

	fmt.Print("Chacha20 Decrypted Payload:\n")
	for i, b := range decryptedShellcode {
		fmt.Printf("0x%%02X", b)
		if i < len(decryptedShellcode)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

}
`

// go b64chacha20
var __go_b64chacha20__ = `
package main

import (
	"golang.org/x/crypto/chacha20poly1305"
	"encoding/base64"

	"fmt"
)

func Chacha20Decrypt(data []byte, key []byte) ([]byte) {
	aead, err := chacha20poly1305.NewX(key)


	nonceSize := aead.NonceSize()

	if len(data) < nonceSize {
		return nil
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		if err.Error() == "chacha20poly1305: message authentication failed" {
			return nil
		}

		return nil
	}

	return plaintext
}


func main() {
	%s := []byte{ %s }
	key := []byte { %s }

	decryptedShellcodeBase64, _ := base64.StdEncoding.DecodeString(string(%s))
	decryptedShellcode := Chacha20Decrypt(decryptedShellcodeBase64, key)

	fmt.Print("B64chacha20 Decrypted Payload:\n")
	for i, b := range decryptedShellcode {
		fmt.Printf("0x%%02X", b)
		if i < len(decryptedShellcode)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

}
`