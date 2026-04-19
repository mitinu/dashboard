package excel

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashing(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", path, err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла %s: %w", path, err)
	}

	return hash.Sum(nil), nil
}
