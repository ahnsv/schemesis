package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/ahnsv/schemesis/types"
	"github.com/brianvoe/gofakeit/v6"
	"gopkg.in/yaml.v2"
)

func GenerateFakeData(schemaBytes []byte, length int) ([]map[string]interface{}, error) {
	var schema types.JSONSchema
	err := json.Unmarshal(schemaBytes, &schema)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	for i := 0; i < length; i++ {
		item := make(map[string]interface{})
		for propertyName, propertySchema := range schema.Properties {
			propertyType := propertySchema.(map[string]interface{})["type"].(string)
			switch propertyType {
			case "string":
				item[propertyName] = gofakeit.Word()
			case "number":
				item[propertyName] = gofakeit.Float64()
			case "integer":
				item[propertyName] = gofakeit.Int64()
			case "boolean":
				item[propertyName] = gofakeit.Bool()
			}
		}
		data = append(data, item)
	}

	return data, nil
}

func GenerateFakeDataV2(configFile string, length int) ([]map[string]interface{}, error) {
	// Read the configuration file
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %v", err)
	}

	// Parse the configuration data
	var config types.Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("failed to parse configuration data: %v", err)
	}

	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate fake data for each field
	data := make([]map[string]interface{}, length)
	for i := 0; i < length; i++ {
		record := make(map[string]interface{})
		for _, field := range config.Fields {
			value, err := generateFieldValue(field, record) // TODO: generate field value based on config
			if err != nil {
				return nil, fmt.Errorf("failed to generate value for field %s: %v", field.Name, err)
			}
			record[field.Name] = value
		}
		data[i] = record
	}

	return data, nil
}
