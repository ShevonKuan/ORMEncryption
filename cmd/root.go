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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ORMEncrption",
	Short: "A tool to encrypt and decrypt orms_core_config.xml for oppo and realme phones.",
	Long: `This application is used to encrypt and decrypt the orms_core_config.xml file, 
which is a core thermal control and scheduling configuration file for oppo and realme phones.

Example:
$ ORMEncrption encrypt -i /path/to/orms_core_config.xml -o ./encrypted.xml
$ ORMEncrption decrypt -i /path/to/encrypted.xml -o ./decrypted.xml
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
