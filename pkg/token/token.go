package token

import (
	"os"

	"github.com/CirillaQL/kubectl-openai/pkg/log"
)

func SaveToken(token, filepath string) error {
	if token == "" {
		log.Fatal("Failed to set token. Token is empty")
	}
	if filepath == "" {
		log.Fatal("Failed to get token filepath")
	}
	os.Remove(filepath)
	file, err := os.Create(os.ExpandEnv(filepath))
	if err != nil {
		log.Fatalf("Failed to set token. error: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(token + "\n")
	if err != nil {
		log.Fatalf("Failed to set token. error: %v", err)
	}
	return nil
}

func ReadToken(filepath string) (string, error) {
	content, err := os.ReadFile(os.ExpandEnv(filepath))
	if err != nil {
		log.Fatalf("Failed to load token. error: %v", err)
	}
	return string(content), nil
}
