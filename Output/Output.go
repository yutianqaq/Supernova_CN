package Output

import (
	"Supernova/Utils"
	"Supernova/Converters"
	"encoding/hex"
	"fmt"
	"os"
)

// SaveOutputToFile function
func SaveOutputToFile(outputData string, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the output data to the file
	_, err = file.WriteString(outputData)
	if err != nil {
		return err
	}

	// Call function named GetAbsolutePath
	absolutePath, err := Utils.GetAbsolutePath(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Printf("[+] Save encrypted shellcode to " + absolutePath + "\n\n")
	return nil
}

// PrintKeyDetails function
func PrintKeyDetails(key []byte) {
	for i, b := range key {
		// decimalValue := int(b)
		hexValue := fmt.Sprintf("%02x", b)
		// fmt.Printf("byte(0x%s) => %d", hexValue, decimalValue)
		fmt.Printf("0x%s", hexValue)
		if i < len(key)-1 {
			fmt.Printf(", ")
		}
	}

	fmt.Printf("\n\n")
}

// KeyDetailsFormatter function
func KeyDetailsFormatter(key []byte) string {
	var formattedKey string
	for i, b := range key {
		hexValue := fmt.Sprintf("%02x", b)
		formattedKey += "0x" + hexValue
		if i < len(key)-1 {
			formattedKey += ", "
		}
	}
	return formattedKey
}

// SaveShellcodeToFile function
func SaveShellcodeToFile(shellcode, filename string) error {
	// Removes Spaces and the "0x" prefix from the string
	shellcode = Converters.CleanShellcodeString(shellcode)

	// Decodes shellcode string into byte array
	data, err := hex.DecodeString(shellcode)
	if err != nil {
		return fmt.Errorf("Error decoding shellcode: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	absolutePath, err := Utils.GetAbsolutePath(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Printf("[+] Save encrypted shellcode file to " + absolutePath + "\n\n")
	return nil
}