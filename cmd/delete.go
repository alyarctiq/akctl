/*
Copyright © 2020 akctl aly.khimji@arctiq.ca

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete subfuction",
	Long:  `delete subfuction similar to "kubectl delete ": `,
	Run: func(cmd *cobra.Command, args []string) {

		clientset, err := buildClient(cfgFile)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}

		ns, _ := cmd.Flags().GetString("namespace")
		if ns == "" {
			fmt.Println("namespace has not been declared use: '-n <nanespace>")
			os.Exit(1)
		}
		deployment, _ := cmd.Flags().GetString("deployment")
		if deployment != "" {
			deleteDeployment(clientset, deployment, ns)
			os.Exit(0)
		}
		pod, _ := cmd.Flags().GetString("pod")
		if pod != "" {
			deletePod(clientset, pod, ns)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("namespace", "n", "", "namespace")
	deleteCmd.Flags().StringP("deployment", "d", "", "delete deployment  <name of deployment>")
	deleteCmd.Flags().StringP("pod", "p", "", "delete pod <name of deployment>")
}
