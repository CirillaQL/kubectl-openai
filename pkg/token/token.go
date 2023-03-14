package token

import "github.com/CirillaQL/kubectl-openai/pkg/log"

func SaveToken(token, filepath string) error {
	if token == "" {
		log.Fatalf("Failed to set token. Token is empty")
	}
	return nil
}
