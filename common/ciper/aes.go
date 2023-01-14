package ciper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

//https://pyrasis.com/book/GoForTheReallyImpatient/Unit53/02

var cipherBlock cipher.Block

var cipherKey string

// LoadCipherKey set AES cipher key, AES cipher key is private variable
func LoadCipherKey(mode string) {
	if cipherKey != "" {
		return
	}

	switch mode {
	case enum.DevMode:
		// .env 로 설정한 환경변수 값보다 시스템에 설정한 환경변수 값이 우선적용된다. 따라서 보안을 위해 dev 일 때만 .env 파일에서 환경변수관리하자
		if err := godotenv.Load(enum.DevModeENVFilePath); err != nil {
			log.Fatal("Error loading .env file ::", err.Error())
		}
		setCipherKey(enum.CipherDevKey)
	case enum.ProdMode:
		setCipherKey(enum.CipherProdKey)
	}
}

func setCipherKey(keyName string) {
	cipherKey = os.Getenv(keyName)
}

func LoadCipherBlock() {
	if cipherKey == "" {
		log.Fatal("please call LoadCipherKey() function")
	}

	block, err := aes.NewCipher([]byte(cipherKey)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		log.Fatal(err.Error())
	}
	cipherBlock = block
}

func GetCipherBlock() cipher.Block {
	return cipherBlock
}

func AESEncrypt(block cipher.Block, plainText []byte) (string, error) {
	if mod := len(plainText) % aes.BlockSize; mod != 0 { // 블록 크기의 배수가 되어야함
		padding := make([]byte, aes.BlockSize-mod) // 블록 크기에서 모자라는 부분을
		plainText = append(plainText, padding...)  // 채워줌
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
	iv := cipherText[:aes.BlockSize]                         // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {  // 랜덤 값을 초기화 벡터에 넣어줌
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)               // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText) // 암호화 블록 모드 인스턴스로 암호화

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(b cipher.Block, cipherText string) (string, error) {
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	if len(decodedCipherText)%aes.BlockSize != 0 { // 블록 크기의 배수가 아니면 리턴
		return "", errors.New("the length of the encrypted data must be a multiple of the block size")
	}

	iv := decodedCipherText[:aes.BlockSize]               // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	decodedCipherText = decodedCipherText[aes.BlockSize:] // 부분 슬라이스로 암호화된 데이터를 가져옴

	plainText := make([]byte, len(decodedCipherText)) // 평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(b, iv)             // 암호화 블록과 초기화 벡터를 넣어서 복호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(plainText, decodedCipherText)    // 복호화 블록 모드 인스턴스로 복호화

	removedPaddingText := bytes.Trim(plainText, "\x00")

	return string(removedPaddingText), nil
}
