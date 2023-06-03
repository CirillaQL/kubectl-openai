/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/CirillaQL/kubectl-openai/pkg/log"
	"github.com/CirillaQL/kubectl-openai/pkg/util"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/CirillaQL/kubectl-openai/pkg/client"
	"github.com/CirillaQL/kubectl-openai/pkg/openai"
	"github.com/CirillaQL/kubectl-openai/pkg/token"
	"github.com/spf13/cobra"
)

var namespace string

const podQuestion = "I will give you a kubernetes' pod's detail, please help me to analyze it."
const podsQuestion = "I will give you some kubernetes' pods' details, please help me to analyze them."

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Use chatGPT to help analyse pod status",
	Long:  `Use chatGPT to help analyse pod status`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if len(name) <= 0 {
			log.Logger.Error("Failed to get pod namespace")
		}
		client, err := client.LoadClient()
		if err != nil {
			panic(err)
		}
		pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		// TODO: optimize pod struct to reduce prompts cost
		podStr := fmt.Sprint(pod)
		tokenString, err := token.ReadToken(TokenPath)
		if err != nil {
			panic(err)
		}
		util.LoadingStart()
		result, err := openai.Ask(tokenString, podQuestion+"\n"+podStr)
		util.LoadingStop()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Use chatGPT to help analyse pods status",
	Long:  `Use chatGPT to help analyse pods status`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace = args[0]
		if len(namespace) <= 0 {
			log.Logger.Error("Failed to get pods ")
		}
		client, err := client.LoadClient()
		if err != nil {
			panic(err)
		}
		pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		// TODO: optimize pods struct to reduce prompts cost
		podsStr := fmt.Sprint(pods)
		tokenString, err := token.ReadToken(TokenPath)
		if err != nil {
			panic(err)
		}
		util.LoadingStart()
		result, err := openai.Ask(tokenString, podsQuestion+"\n"+podsStr)
		util.LoadingStop()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	rootCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")
	podCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Pod's namespace")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
