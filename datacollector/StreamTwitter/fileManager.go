package main

import (
	"encoding/json"
	"os"
)

func readjsonIntMap(path string) map[int64]bool {
	var data []int64
	readjsonFile(path, &data)
	result := make(map[int64]bool)
	for i := 0; i < len(data); i++ {
		result[data[i]] = true
	}
	return result
}

func readjsonFile(path string, outData interface{}) {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&outData)

	if err != nil {
		panic(err)
	}
}
