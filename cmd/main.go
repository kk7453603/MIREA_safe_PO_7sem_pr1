package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"task1/internal/dyscinfo"
	"task1/internal/file"
)

func main() {

	disk := dyscinfo.New()

	disk.GetLogicalDrivesInfo()

	err := file.CreateFileWithData("test5.txt", "check")
	if err != nil {
		log.Fatal(err)
	}

	err = file.DeleteFile("test5.txt")
	if err != nil {
		log.Fatal(err)
	}

	person := map[string]interface{}{
		"name":  "Alicedasadsadasdasdasdads",
		"email": "alice@example.comasdasads",
		"age":   28,
	}

	err = file.CreateJSONFile("person.json", person)
	if err != nil {
		log.Fatal(err)
	}

	err = file.ReadJSONFile("person.json")
	if err != nil {
		log.Fatal(err)
	}

	err = file.DeleteFile("person.json")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя для XML:")
	name, _ := reader.ReadString('\n')

	fmt.Println("Введите телефон для XML:")
	phone, _ := reader.ReadString('\n')

	type Contact struct {
		Name  string `xml:"name"`
		Phone string `xml:"phone"`
	}

	contact := Contact{
		Name:  name,
		Phone: phone,
	}

	err = file.CreateXMLFile("contact.xml", contact)
	if err != nil {
		log.Fatal(err)
	}

	err = file.ReadXMLFile("contact.xml")
	if err != nil {
		log.Fatal(err)
	}

	err = file.DeleteFile("contact.xml")
	if err != nil {
		log.Fatal(err)
	}

	err = file.CreateFileWithData("archive_test.txt", "This is some data for the archive")
	if err != nil {
		log.Fatal(err)
	}

	err = file.CreateZipArchive("archive.zip", "archive_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = file.UnzipFile("archive.zip")
	if err != nil {
		log.Fatal(err)
	}

	err = file.DeleteFile("archive_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = file.DeleteFile("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
}
