package utils

import (
	"errors"
	"github.com/ozoncp/ocp-service-api/internal/models"
)

//SplitToBatches исходный слайс конвертируется в слайс слайсов - чанки одинкового размера (кроме последнего)
func SplitToBatches(originalSlice []models.Service, batchSize int) [][]models.Service {
	// TODO: обрабатывать batchSize меньше 1
	var resultingSlice [][]models.Service

	totalLength := len(originalSlice)
	for i := 0; i < totalLength; i += batchSize {
		var lastElementIdx = i + batchSize
		if lastElementIdx > totalLength {
			lastElementIdx = totalLength
		}
		batch := originalSlice[i:lastElementIdx]
		resultingSlice = append(resultingSlice, batch)
	}
	return resultingSlice
}

//ReverseMapKeysAndValues происходит конвертация отображения (“ключ-значение“) в отображение (“значение-ключ“)
func ReverseMapKeysAndValues(entities []models.Service) (map[uint64]models.Service, error) {
	reversedMap := map[uint64]models.Service{}

	for _, v := range entities {
		currId := v.Id
		_, ok := reversedMap[currId] // ключ есть в новом словаре
		if !ok {
			return nil, errors.New("Duplicate key")
		}

		reversedMap[v.Id] = v
	}
	return reversedMap, nil
}

func contains(slice []string, value string) bool {
	for _, a := range slice {
		if a == value {
			return true
		}
	}
	return false
}

//FilterOut Фильтрация по захордкоженному списку
func FilterOut(original []string) []string {
	// TODO: лучше использовать аналог Set
	restrictedValues := []string{"Marv", "Intor", "Kreks"}
	var filteredValues []string
	for _, v := range original {
		if !contains(restrictedValues, v) {
			filteredValues = append(filteredValues, v)
		}
	}
	return filteredValues
}
