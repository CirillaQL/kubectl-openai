/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/CirillaQL/kubectl-openai/pkg/log"
	"github.com/CirillaQL/kubectl-openai/pkg/token"
	"github.com/spf13/cobra"
)

var openAIToken string
var tokenPath string

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Set or Get your OpenAI Token",
	Long:  `Set or Get your OpenAI Token,`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info(openAIToken)
		err := token.SaveToken(openAIToken, tokenPath)
		if err != nil {
			log.Fatalf("save token into file failed. error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenCmd.PersistentFlags().String("foo", "", "A help for foo")
	tokenCmd.Flags().StringVarP(&tokenPath, "path", "p", "~/.kube/kubectl-openai", "custom file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	tokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
