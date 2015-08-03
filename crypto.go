package crypto

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"io"
)

type (
	IVFunc func([]byte, bool) []byte
)

// Encryption `key` must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes
func Encrypt(text []byte, key []byte, ivFunc IVFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	cfb := cipher.NewCFBEncrypter(block, ivFunc(cipherText, true))
	cfb.XORKeyStream(cipherText[aes.BlockSize:], text)
	return Encode64(cipherText), nil
}

// Decryption `key` must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes
func Decrypt(text []byte, key []byte, ivFunc IVFunc) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	text, err = Decode64(text)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("text too short")
	}
	iv := ivFunc(text, false)
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return text, nil
}

func IV(text []byte, random bool) []byte {
	iv := text[:aes.BlockSize]
	if random {
		io.ReadFull(rand.Reader, iv)
	}
	return iv
}

func Encode64(text []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(buf, text)
	return buf
}

func Decode64(text []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	n, err := base64.StdEncoding.Decode(buf, text)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}