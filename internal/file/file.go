package file

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func CreateFileWithData(filename string, data string) error {
	file, err := os.Create(filename)
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return err
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	filedata, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}
	fmt.Printf("File data: %v\n", string(filedata))
	return nil
}

func CreateJSONFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("JSON File data: %v\n", string(fileData))

	return nil
}

func ReadJSONFile(filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("Read JSON data: %v\n", string(fileData))
	return nil
}

func CreateXMLFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	xmlData, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(xmlData)
	if err != nil {
		return err
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("XML File data: %v\n", string(fileData))

	return nil
}

func ReadXMLFile(filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("Read XML data: %v\n", string(fileData))
	return nil
}

func CreateZipArchive(zipFilename, fileToZip string) error {
	// Создаем архив
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToAdd, err := os.Open(fileToZip)
	if err != nil {
		return err
	}
	defer fileToAdd.Close()

	// Добавляем файл в архив
	info, err := fileToAdd.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToAdd)
	if err != nil {
		return err
	}

	fmt.Printf("File %v added to archive %v\n", fileToZip, zipFilename)
	return nil
}

func UnzipFile(zipFilename string) error {
	r, err := zip.OpenReader(zipFilename)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		fmt.Printf("Unzipping file: %v\n", file.Name)
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		if err != nil {
			return err
		}

		fmt.Printf("File content:\n%s\n", buf.String())
	}
	return nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	fmt.Printf("File %v deleted successfully\n", filename)
	return nil
}
