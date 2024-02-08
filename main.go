package main

import (
	"bytes"
	"fmt"
	"os"
	"syscall"

	"github.com/angelina612/file-encryption/filecrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file or decrypt to decrypt a file. Run 'go run . help' for more information.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("Simple file encryption tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\t go run . encrypt /path/to/file")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("\t encrypt\tEncrypt a file")
	fmt.Println("\t decrypt\tDecrypt a file")
	fmt.Println("\t help\t\tShow this help message")
	fmt.Println()
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. Run 'go run . help' for more information.")
		os.Exit(0)
	}

	filePath := os.Args[2]
	if !validateFile(filePath) {
		panic("File not found")
	}

	password := getPassword()
	fmt.Println("Encrypting file...")
	filecrypt.EncryptFile(filePath, password)
	fmt.Println("File encrypted successfully.")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. Run 'go run . help' for more information.")
		os.Exit(0)
	}

	filePath := os.Args[2]
	if !validateFile(filePath) {
		panic("File not found")
	}
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("\nDecrypting file...")
	filecrypt.DecryptFile(filePath, password)
	fmt.Println("File decrypted successfully.")
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\nConfirm password: ")
	confirmPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	if !validatePassword(password, confirmPassword) {
		fmt.Println("Passwords do not match.")
		return getPassword()
	}
	return password
}

func validatePassword(password []byte, confirmPassword []byte) bool {
	return bytes.Equal(password, confirmPassword)
}

func validateFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
