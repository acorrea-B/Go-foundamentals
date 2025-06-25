package utilities

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) float64 {
	data, _ := os.ReadFile(fileName)
	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)
	return value
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
