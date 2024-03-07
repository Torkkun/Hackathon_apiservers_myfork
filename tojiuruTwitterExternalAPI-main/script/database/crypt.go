package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"
	"unsafe"
)

var Block cipher.Block

func init() {
	var err error
	keystring := os.Getenv("KEY")
	key := []byte(keystring)
	Block, err = aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
}

func createIV(token []byte) (iv []byte, cipherToken []byte, err error) {
	cipherToken = make([]byte, aes.BlockSize+len(token))
	iv = cipherToken[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	return
}

func TokenEncrypt(tokenstring string) (entoken string) {
	token := []byte(tokenstring)
	iv, cipherToken, err := createIV(token)
	if err != nil {
		log.Fatalln(err)
	}
	// Encrypt
	encryptStream := cipher.NewCTR(Block, iv)
	encryptStream.XORKeyStream(cipherToken[aes.BlockSize:], token)
	// Base64 Encode
	entoken = base64.StdEncoding.EncodeToString(cipherToken)
	return
}

func TokenDecrypt(token string) (detoken string) {
	// Base64 Decode
	cipherToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		//testデータ用
		log.Println(err)
		detoken = token
		return
	}
	// Decrypt
	decryptedToken := make([]byte, len(cipherToken[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(Block, cipherToken[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedToken, cipherToken[aes.BlockSize:])
	detoken = *(*string)(unsafe.Pointer(&decryptedToken))
	return
}
