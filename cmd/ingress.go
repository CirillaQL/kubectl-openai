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

const ingressQuestion = "I will give you a kubernetes' service's detail, please help me to analyze it"

var ingressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Use chatGPT to help analyse ingress status",
	Long:  "Use chatGPT to help analyse ingress status",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		name := args[0]

		if len(name) <= 0 {
			log.Logger.Error("Failed to get ingress namespace")
		}
		client, err := client.LoadClient()
		if err != nil {
			log.Logger.Fatalf("Failed to load k8s client, error: %+v", err)
			panic(err)
		}

		ingress, err := client.NetworkingV1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			log.Logger.Fatalf("Failed to get Ingress: %s, Error: %v", name, err)
			panic(err.Error())
		}
		ingressStr := fmt.Sprint(ingress)
		tokenString, err := token.ReadToken(TokenPath)
		if err != nil {
			panic(err)
		}
		util.LoadingStart()
		result, err := openai.Ask(tokenString, ingressQuestion+"\n"+ingressStr)
		util.LoadingStop()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(ingressCmd)

	ingressCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Ingress's namespace")
}
