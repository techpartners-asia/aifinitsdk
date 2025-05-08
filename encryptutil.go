package aifinitsdk

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

type EncryptUtil struct {
	merchantCode string
	secretKey    string
}

func NewEncryptUtil(merchantCode, secretKey string) *EncryptUtil {
	return &EncryptUtil{
		merchantCode: merchantCode,
		secretKey:    secretKey,
	}
}

func (e *EncryptUtil) pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func (e *EncryptUtil) decryptECB(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	decrypted := make([]byte, len(encrypted))
	for start := 0; start < len(encrypted); start += bs {
		block.Decrypt(decrypted[start:start+bs], encrypted[start:start+bs])
	}
	return decrypted, nil
}

func (e *EncryptUtil) encryptECB(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	plainText = e.pkcs5Padding(plainText, bs)
	encrypted := make([]byte, len(plainText))
	for start := 0; start < len(plainText); start += bs {
		block.Encrypt(encrypted[start:start+bs], plainText[start:start+bs])
	}
	return encrypted, nil
}

func (e *EncryptUtil) Encrypt(data string) (string, error) {
	plainText := []byte(data)
	key := []byte(e.secretKey)

	encrypted, err := e.encryptECB(plainText, key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (e *EncryptUtil) Decrypt(data string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	key := []byte(e.secretKey)

	// Ensure key is exactly 16 bytes
	if len(key) < 16 {
		key = append(key, make([]byte, 16-len(key))...)
	} else if len(key) > 16 {
		key = key[:16]
	}

	decrypted, err := e.decryptECB(encrypted, key)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
