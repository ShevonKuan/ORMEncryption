package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	_ "github.com/alecthomas/chroma/v2"
)

const (
	ivLength     = 12 // GCM IV length
	tagLengthBit = 128
	keyLength    = 16
	algorithm    = "AES/GCM/NoPadding"
)

// encrypt 加密数据并返回Base64编码的字符串
func Encrypt(key []byte, rawData string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, ivLength)
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	encrypted := aesgcm.Seal(nil, nonce, []byte(rawData), nil)

	buffer := append([]byte{byte(ivLength)}, nonce...)
	buffer = append(buffer, encrypted...)

	return base64.StdEncoding.EncodeToString(buffer), nil
}

// decrypt 解密Base64编码的字符串并返回解密后的数据
func Decrypt(key []byte, encryptedData string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	ivLength := int(decodedBytes[0])
	nonce := decodedBytes[1 : 1+ivLength]
	ciphertext := decodedBytes[1+ivLength:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// getSha256Key 计算SHA-256哈希并返回前16字节作为密钥
func GetSha256Key() []byte {
	input := "ORMS" // this is utf8 strings
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return []byte(hex.EncodeToString(hasher.Sum(nil))[:keyLength])
}
