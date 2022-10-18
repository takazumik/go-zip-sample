package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
)

func main() {
	buffer := new(bytes.Buffer)
	if err := MakeZipFile(buffer); err != nil {
		return
	}
}

func MakeZipFile(buffer *bytes.Buffer) error {
	zipWriter := zip.NewWriter(buffer)
	file, err := zipWriter.Create("sample.txt")
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{"id", "name", "address", "birthday", "gender"})
	csvWriter.Flush()
	if err := zipWriter.Close(); err != nil {
		return err
	}
	return nil
}
