package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to to find file")
	}

	valueText := string(data)
	value, err_ := strconv.ParseFloat(valueText, 64)

	if err_ != nil {
		return 0.0, errors.New("Failed to parse stored value")
	}

	return value, nil
}
