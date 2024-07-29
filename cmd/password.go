package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random passwords",
	Long: `Generate random passwords with customizable options

for example:
passwordGen generate -l 12 -d -s`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Lenght of the generated password.")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in passwords.")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in passwords.")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, err := cmd.Flags().GetInt("length")
	if err != nil {
		fmt.Println("Error getting length flag:", err)
		return
	}

	isDigits, err := cmd.Flags().GetBool("digits")
	if err != nil {
		fmt.Println("Error getting digits flag:", err)
		return
	}

	isSpecialChars, err := cmd.Flags().GetBool("special-chars")
	if err != nil {
		fmt.Println("Error getting special-chars flag:", err)
		return
	}

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!*#-=?@';:,.<>"
	}

	password, err := generateRandomString(charset, length)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}

func generateRandomString(charset string, length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}
	return string(result), nil
}
