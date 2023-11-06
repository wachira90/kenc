package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wachira90/kenc/utils"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

var outFilePath string
var isSave bool

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode string in the data fields to base64.",
	Long:  `encode string in the data fields to base64. The output will print to the stdout by default.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires input file argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		file := utils.OpenFile(filePath)

		kConfig, err := yaml.Parse(string(file))
		if err != nil {
			fmt.Println("Unable to parse yaml")
			os.Exit(1)
		}
		dataMap := kConfig.GetDataMap()

		if kind := kConfig.GetKind(); kind != `Secret` {
			var isContinue string
			fmt.Printf("%s is not a secret kind. Do you want to continue encoding? (Y)es, (N)o: ", filePath)
			_, err = fmt.Scan(&isContinue)
			if err != nil {
				fmt.Println("Unable to get answer")
				os.Exit(1)
			}
			if isContinue == "N" || isContinue == "n" {
				os.Exit(0)
			}
		}

		for key, value := range dataMap {
			dataMap[key] = utils.ToBase64(value)
		}
		kConfig.SetDataMap(dataMap)
		strConfig, err := kConfig.String()
		if err != nil {
			fmt.Println("Unable to get string from config")
			os.Exit(1)
		}

		if !isSave && len(outFilePath) == 0 {
			fmt.Println(strConfig)
		} else {
			if len(outFilePath) == 0 {
				outFilePath = filePath
			}
			utils.WriteToFile(outFilePath, strConfig)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().StringVarP(&outFilePath, "out", "o", "", "Write the output to this file path")
	encodeCmd.Flags().BoolVarP(&isSave, "save", "s", false, "Save the output to the same file")
}
