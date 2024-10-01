package services

import (
	"encoding/json"
	"fmt"
	"os"
)

const DatabaseDirectory = "./database"
const DatabaseFormat = ".json"

func LoadDatabase[T any](databaseName string, databaseObject *T) error {

	databasePath := DatabaseDirectory + "/" + databaseName + DatabaseFormat

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		createDatabase(databasePath, databaseObject)
		return err
	}

	data, err := os.ReadFile(databasePath)
	if err != nil {
		panic("Error reading database file")
	}

	err = json.Unmarshal(data, databaseObject)
	if err != nil {
		panic("Error parsing database file")
	}

	return nil
}

func createDatabase[T any](databasePath string, databaseObject *T) {

	_, err := os.Stat(DatabaseDirectory)

	if os.IsNotExist(err) {
		err = os.MkdirAll(DatabaseDirectory, 0755)
		if err != nil {
			fmt.Errorf("failed to create directory: %w", err)
			return
		}
	}

	file, err := os.Create(databasePath)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	SaveDatabase(databasePath, databaseObject)
}

func SaveDatabase[T any](databaseName string, databaseObject *T) {

	databasePath := DatabaseDirectory + "/" + databaseName + DatabaseFormat

	jsonBytes, err := json.MarshalIndent(databaseObject, "", " ")
	if err != nil {
		panic("Error parsing database file")
	}

	err = os.WriteFile(databasePath, jsonBytes, 0644)
}
