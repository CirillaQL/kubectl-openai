package token

import (
	"os"

	"github.com/CirillaQL/kubectl-openai/pkg/log"
)

func SaveToken(token, filepath string) error {
	if token == "" {
		log.Logger.Fatal("Failed to set token. Token is empty")
	}
	if filepath == "" {
		log.Logger.Fatal("Failed to get token filepath")
	}
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	file, err := os.Create(os.ExpandEnv(filepath))
	if err != nil {
		log.Logger.Fatalf("Failed to set token. error: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(token)
	if err != nil {
		log.Logger.Fatalf("Failed to set token. error: %v", err)
	}
	return nil
}

func ReadToken(filepath string) (string, error) {
	content, err := os.ReadFile(os.ExpandEnv(filepath))
	if err != nil {
		log.Logger.Fatalf("Failed to load token. error: %v", err)
	}
	return string(content), nil
}
