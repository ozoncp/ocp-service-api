package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const COL_DELIMITER = ";"

// Читает csv файл, возвращает двумерный массив со значениями
func ReadCSVFile(path string) (configValues [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Error during close", closeErr)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		columns := strings.Split(row, COL_DELIMITER)
		configValues = append(configValues, columns)
	}

	return configValues, nil

}

func main() {
	fmt.Println("This is service api microservice.")
	values, err := ReadCSVFile("err path")
	if err != nil {
		fmt.Printf("Error loading config %s", err)
		os.Exit(1)
	}
	fmt.Println(values)
}
