package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ahnsv/schemesis/pkg"

	"cloud.google.com/go/bigquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [input file]")
		return
	}
	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var schema bigquery.Schema
	err = json.Unmarshal(bytes, &schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBytes, err := pkg.CreateJsonSchema(schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := pkg.GenerateFakeData(jsonBytes, 3)
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
