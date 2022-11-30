package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Rayato159/apex-cosplay-private-2022/src/models"
)

func ReadJsonFile(path string) ([]models.User, error) {
	// Open a file
	jsonFile, err := os.Open(path)
	if err != nil {
		defer jsonFile.Close()
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println(path)

	// Read a data
	data, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(data))

	// Parse a byte data into the object
	obj := make([]models.User, 0)
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return make([]models.User, 0), fmt.Errorf("error, can't parse a layer data into the object")
	}
	fmt.Println("success, data has been loaded")

	return obj, nil
}
