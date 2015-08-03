package crypto
import (
	"testing"
	"bytes"
	"crypto/aes"
)

func TestEncode64(t *testing.T) {
	text := []byte("Hello World")
	if !bytes.Equal(Encode64(text), []byte("SGVsbG8gV29ybGQ=")) {
		t.Fail()
	}
}

func BenchmarkEncode64(b *testing.B) {
	text := []byte("Hello")
	for i := 0; i < b.N; i++ {
		Encode64(text)
	}
}

func TestEncrypt(t *testing.T) {
	key := []byte("WelcomeToMyWorld")
	text := []byte("Hello")
	cipher, _ := Encrypt(text, key, mockIV);
	if !bytes.Equal(cipher, []byte("AAAAAAAAAAAAAAAAAAAAABk9D0Hr")) {
		t.Fail()
	}
}

func BenchmarkEncrypt(b *testing.B) {
	key := []byte("WelcomeToMyWorld")
	text := []byte("Hello")
	for i := 0; i < b.N; i++ {
		Encrypt(text, key, IV);
	}
}

func TestDecode64(t *testing.T) {
	text := []byte("SGVsbG8gV29ybGQ=")
	text, _ = Decode64(text)
	if !bytes.Equal(text, []byte("Hello World")) {
		t.Fail()
	}
}

func BenchmarkDecode64(b *testing.B) {
	text := []byte("SGVsbG8gV29ybGQ=")
	for i := 0; i < b.N; i++ {
		Decode64(text)
	}
}

func TestDecrypt(t *testing.T) {
	key := []byte("WelcomeToMyWorld")
	text := []byte("AAAAAAAAAAAAAAAAAAAAABk9D0Hr")
	cipher, _ := Decrypt(text, key, mockIV);
	if !bytes.Equal(cipher, []byte("Hello")) {
		t.Fail()
	}
}

func BenchmarkDecrypt(b *testing.B) {
	key := []byte("WelcomeToMyWorld")
	text := []byte("AAAAAAAAAAAAAAAAAAAAABk9D0Hr")
	for i := 0; i < b.N; i++ {
		Decrypt(text, key, IV);
	}
}

func mockIV(text []byte, random bool) []byte {
	return text[:aes.BlockSize]
}