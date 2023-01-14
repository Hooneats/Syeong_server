package ciper

import (
	"crypto/aes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestGenerateEncrypteAndDecrypt(t *testing.T) {
	testText := "keep_moving_forward_for_future"
	ENVFPath := "../../.env"
	keyName := "WEMIXON_DEV_KEY"

	if err := godotenv.Load(ENVFPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	AESDEVKey := os.Getenv(keyName)
	log.Println("AES Encrypt key is ::", AESDEVKey)

	block, err := aes.NewCipher([]byte(AESDEVKey)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		log.Fatal(err.Error())
	}

	encryptedText, err := AESEncrypt(block, []byte(testText))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Encrypted text is ::", encryptedText)

	decryptedText, err := AESDecrypt(block, encryptedText)
	if err != nil {
		return
	}

	log.Println("Decrypted text is ::", decryptedText)
}
