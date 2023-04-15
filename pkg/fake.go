package pkg

import (
	"encoding/json"

	"github.com/ahnsv/schemesis/types"
	"github.com/brianvoe/gofakeit/v6"
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
