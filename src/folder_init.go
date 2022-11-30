package src

import (
	"fmt"
	"os"
	"strings"
)

func FolderInit() error {
	layers, err := ReadJsonFile("./src/data/layers.json")
	if err != nil {
		return err
	}
	photographers, err := ReadJsonFile("./src/data/photographers.json")
	if err != nil {
		return err
	}

	for i := range layers {
		layerFolderName := fmt.Sprintf("%s[%s]", strings.ReplaceAll(strings.ToLower(layers[i].Name), " ", "-"), strings.ToLower(layers[i].Character))
		// Mkdir a root folder of layer
		if err := os.MkdirAll(layerFolderName, os.FileMode(0777)); err != nil {
			return fmt.Errorf("error, can't mkdir %s with an error: %s", layerFolderName, err.Error())
		}

		// Mkdir foreach photographer
		for j := range photographers {
			// Mkdir a detail folder
			photoGrapherFolderName := fmt.Sprintf("%s/%s", layerFolderName, photographers[j].Character)
			retouchFolder := fmt.Sprintf("%s/%s/%s", layerFolderName, photographers[j].Character, "retouch")
			if err := os.MkdirAll(photoGrapherFolderName, os.FileMode(0777)); err != nil {
				return fmt.Errorf("error, can't mkdir %s with an error: %s", photoGrapherFolderName, err.Error())
			}
			if err := os.MkdirAll(retouchFolder, os.FileMode(0777)); err != nil {
				return fmt.Errorf("error, can't mkdir %s with an error: %s", retouchFolder, err.Error())
			}
		}

		// Mkdir for ref folder
		refFolder := fmt.Sprintf("%s/%s", layerFolderName, "ref")
		if err := os.MkdirAll(refFolder, os.FileMode(0777)); err != nil {
			return fmt.Errorf("error, can't mkdir %s with an error: %s", refFolder, err.Error())
		}
	}
	return nil
}
