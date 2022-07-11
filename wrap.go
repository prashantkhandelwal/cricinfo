package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Wrap[T any] struct {
	Data []T
}

func (w *Wrap[_]) Save(path string, format Format) (bool, error) {

	if len(strings.TrimSpace(path)) == 0 {
		return false, errors.New("Parameter: path cannot be empty")
	}
	dir, fname := filepath.Split(path)
	_, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Path does not exists. Creating...")
		err := os.MkdirAll(dir, os.ModeSticky|os.ModeDir|os.ModePerm)
		if err != nil {
			return false, errors.New("ERROR: Cannot create path due to permission - " + err.Error())
		}
	}
	if !strings.Contains(fname, ".json") {
		return false, errors.New("ERROR: Please provide a valid json file name.")
	}

	if format == JSON {
		json, err := json.MarshalIndent(&w.Data, "", "  ")
		if err != nil {
			return false, err
		}
		f, err := os.Create(path)
		if err != nil {
			return false, errors.New("ERROR: Cannot create file - " + err.Error())
		}

		_, err = f.Write(json)
		if err != nil {
			return false, errors.New("ERROR: Cannot write to the file - " + err.Error())
		}

		err = f.Close()
		if err != nil {
			return false, errors.New("ERROR: Cannot close the file - " + err.Error())
		}

		fmt.Printf("Saving File - %s\n", fname)
	}

	// if format == CSV {
	// 	// TODO Implement CSV file format.
	// }

	return true, nil
}
