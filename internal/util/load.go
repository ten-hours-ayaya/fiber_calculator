package util

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ten-hours-ayaya/fiber-calculator/internal/model"
)

func LoadFromFile(path string) (*model.Input, error) {
	if _, err := os.Stat(path); err != nil {
		log.Fatal("File not found - ", err)
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read the file - ", err)
		return nil, err
	}
	var input *model.Input
	if err := json.Unmarshal(data, &input); err != nil {
		log.Fatal("Cannot unmarshal json - ", err)
		return nil, err
	}
	return input, nil
}
