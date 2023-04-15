package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/ahnsv/schemesis/pkg"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "schemesis",
		Short: "Generate fake data based on a BigQuery schema",
		Run:   runCmd,
	}
	inputJSON string
)

func init() {
	RootCmd.Flags().StringP("input-file", "i", "", "Input BigQuery schema JSON file path")
	RootCmd.Flags().StringVarP(&inputJSON, "input-json", "j", "", "Input BigQuery schema JSON (can use from shell script pipe)")
	RootCmd.Flags().StringP("config-file", "c", "", "Config YAML file path")
	RootCmd.Flags().IntP("num-data", "n", 10, "Number of fake data to generate (default 10)")
}

func runCmd(cmd *cobra.Command, args []string) {
	var (
		schema bigquery.Schema
	)

	inputFile, _ := cmd.Flags().GetString("input-file")
	inputJSON, _ := cmd.Flags().GetString("input-json")
	numberOfData, _ := cmd.Flags().GetInt("num-data")

	if inputFile != "" {
		bytes, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		inputJSON = string(bytes)
	}

	if inputJSON == "-" {
		inputJSON = ""
		// Read from stdin
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				inputJSON += scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
		}

	}

	err := json.Unmarshal([]byte(inputJSON), &schema)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes, err := pkg.CreateJsonSchema(schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := pkg.GenerateFakeData(jsonBytes, numberOfData)
	if err != nil {
		fmt.Println(err)
		return
	}
	output, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
