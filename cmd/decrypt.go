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
	"ORMEncryption/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringP("input", "i", "", "Input file path (required)")
	decryptCmd.MarkFlagRequired("input")
	decryptCmd.Flags().StringP("output", "o", "", "Output file path. If not specified, the default is ./orms_core_config_decrypted.xml")
	decryptCmd.Flags().BoolP("no-output", "n", false, "Do not write to a file and print the decrypted content in the terminal instead")
}

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt orms_core_config_decrypted.xml",
	Long: `This command is used to decrypt an AES/GCM encrypted and Base64 encoded XML configuration file (orms_core_config_decrypted.xml).
The decrypted content can be saved in a specified output file or printed directly into the terminal if no output file is provided.

Example:
$ decrypt -i /path/to/orms_core_config.xml -o ./decrypted.xml
`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath, _ := cmd.Flags().GetString("input")
		outputFlag, _ := cmd.Flags().GetBool("no-output")

		data, err := utils.ReadFile(inputFilePath)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			os.Exit(1)
		}

		decryptedData, err := utils.Decrypt(utils.GetSha256Key(), data)
		if err != nil {
			fmt.Println("Error decrypting data:", err)
			os.Exit(1)
		}

		if outputFlag {
			utils.PrintDecryptedData(decryptedData)
		} else {
			outputFilePath, _ := cmd.Flags().GetString("output")
			if outputFilePath == "" {
				outputFilePath = "orms_core_config_decrypted.xml"
			}
			err = utils.WriteFile(outputFilePath, decryptedData)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(1)
			} else {
				fmt.Printf("Decrypted data written to %s\n", outputFilePath)
			}
		}
	},
}
