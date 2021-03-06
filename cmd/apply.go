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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply subfuction",
	Long:  `Create and Apply Manifest similarly to "kubectl apply -f": `,
	Run: func(cmd *cobra.Command, args []string) {

		clientset, err := buildClient(cfgFile)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}

		ns, _ := cmd.Flags().GetString("namespace")

		file, _ := cmd.Flags().GetString("file")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("File reading error", err)
			os.Exit(1)
		}

		deployoyment, _ := cmd.Flags().GetBool("deployoyment")
		if deployoyment == true {
			createDeploymentFromYaml(clientset, data, ns)
		}

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	applyCmd.Flags().StringP("namespace", "n", "", "namespace")
	applyCmd.Flags().BoolP("deployoyment", "d", false, "test deploy")
	applyCmd.Flags().StringP("file", "f", "", "file path")
}
