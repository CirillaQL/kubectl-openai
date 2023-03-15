/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/CirillaQL/kubectl-openai/pkg/client"
	"github.com/spf13/cobra"
)

var namespace string
var name string

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]
		client, err := client.LoadClient()
		if err != nil {
			panic(err)
		}
		pod, err := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(pod)
	},
}

func init() {
	rootCmd.AddCommand(podCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")
	podCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Pod's namespace")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
