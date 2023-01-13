package chiper

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGenerateEncrypteAndDecrypt(t *testing.T) {
	testText := "wemixon"

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	AESDEVKey := os.Getenv("WEMIXON_DEV_KEY")
	log.Println("AES Encrypt key is ::", AESDEVKey)

	LoadCipherKey("dev")
	LoadCipherBlock()

	encryptedText, err := AESEncrypt(CipherBlock, []byte(testText))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Encrypted text is ::", encryptedText)

	decryptedText, err := AESDecrypt(CipherBlock, encryptedText)
	if err != nil {
		return
	}

	//decryptedText = removePadding(decryptedText)

	log.Println("Decrypted text is ::", decryptedText)
}

func removePadding(decryptedText string) string {
	pt := []byte(decryptedText)
	padLength := int(pt[len(pt)-1])
	fmt.Println("####### ", pt[:len(pt)-padLength])
	return string(pt[:len(pt)-padLength])
}
