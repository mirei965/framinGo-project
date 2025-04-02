package framingo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

const (
	randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"
)
//RandomStringは、指定された長さ(n)のランダムな文字列を生成する関数
func (f *Framingo) RandomString(n int) string {
	s, r := make([]rune, n),[]rune(randomString)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

func (f *Framingo) CreateDirIfNotExist(path string) error {
	const mode = 0755 // モードを定義
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Framingo) CreateFileIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		// ファイル作成に失敗した場合、エラーを返す
		if err != nil {
			return err
		}
		// ファイルを閉じるためのdefer関数
		// deferは、関数が終了する直前に実行されるため、リソースリークを防止
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}

//暗号化と復号化
type Encryption struct {
	Key []byte
}

func (e *Encryption) Encrypt(text string) (string, error) {
	plainText := []byte(text)
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return base64.URLEncoding.EncodeToString(cipherText), nil
}
//復号化
func (e *Encryption) Decrypt(cryptText string) (string, error) {
	cipherText, _ := base64.URLEncoding.DecodeString(cryptText)

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
