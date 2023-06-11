package cmd

import (
	"context"
	"fmt"
	"github.com/CirillaQL/kubectl-openai/pkg/client"
	"github.com/CirillaQL/kubectl-openai/pkg/log"
	"github.com/CirillaQL/kubectl-openai/pkg/openai"
	"github.com/CirillaQL/kubectl-openai/pkg/token"
	"github.com/CirillaQL/kubectl-openai/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const serviceQuestion = "I will give you a kubernetes' service's detail, please help me to analyze it"

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Use chatGPT to help analyse service status",
	Long:  "Use chatGPT to help analyse service status",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		name := args[0]

		if len(name) <= 0 {
			log.Logger.Error("Failed to get pod namespace")
		}
		client, err := client.LoadClient()
		if err != nil {
			log.Logger.Fatalf("Failed to load k8s client, error: %+v", err)
			panic(err)
		}

		service, err := client.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			log.Logger.Fatalf("Failed to get Service: %s, Error: %v", name, err)
			panic(err.Error())
		}
		podStr := fmt.Sprint(service)
		tokenString, err := token.ReadToken(TokenPath)
		if err != nil {
			panic(err)
		}
		util.LoadingStart()
		result, err := openai.Ask(tokenString, serviceQuestion+"\n"+podStr)
		util.LoadingStop()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Pod's namespace")
}
