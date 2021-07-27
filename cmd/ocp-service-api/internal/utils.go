package utils

//SplitToBatches исходный слайс конвертируется в слайс слайсов - чанки одинкового размера (кроме последнего)
func SplitToBatches(originalSlice []string, batchSize int) [][]string {
	// TODO: обрабатывать batchSize меньше 1
	var resultingSlice [][]string

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
func ReverseMapKeysAndValues(originalMap map[string]string) map[string]string {

	reversedMap := map[string]string{}
	for k, v := range originalMap {
		reversedMap[v] = k
	}
	return reversedMap
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
