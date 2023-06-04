package util

import (
	"encoding/json"
	"os"
)

func SaveToLocal(data any, outpath string) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f, err := os.Create(outpath)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
