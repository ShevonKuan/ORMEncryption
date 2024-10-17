/*
Copyright © 2024 Shevon Kwan drshevonkuan@gmail.com

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
	"ORMEncrption/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().StringP("input", "i", "", "Input file path (required)")
	encryptCmd.MarkFlagRequired("input")
	encryptCmd.Flags().StringP("output", "o", "", "Output file path. If not specified, the default is ./orms_core_config_encrypted.xml")
}

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt orms_core_config.xml",
	Long: `This command is used to encrypt an XML configuration file (orms_core_config.xml) using AES/GCM and Base64 encoding.
The encrypted content will be saved in a specified output file.

Example:
$ encrypt -i /path/to/orms_core_config.xml -o ./encrypted.xml
`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath, _ := cmd.Flags().GetString("input")
		outputFilePath, _ := cmd.Flags().GetString("output")

		data, err := utils.ReadFile(inputFilePath)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			os.Exit(1)
		}

		encryptedData, err := utils.Encrypt(utils.GetSha256Key(), data)
		if err != nil {
			fmt.Println("Error encrypting data:", err)
			os.Exit(1)
		}

		if outputFilePath == "" {
			outputFilePath = "orms_core_config_encrypted.xml"
		}
		err = utils.WriteFile(outputFilePath, encryptedData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		} else {
			fmt.Printf("Encrypted data written to %s\n", outputFilePath)
		}
	},
}
