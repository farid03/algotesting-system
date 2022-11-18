package fileworkers

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil { // TODO вынести в отдельную функцию обработку ошибок (дублирование)
		fmt.Println("Unable to create \""+name+"\" file:", err)
		log.Fatal(err)
	}

	return file
}

func OpenFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Unable to open \""+file.Name()+"\" file:", err)
		log.Fatal(err)
	}

	return file
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Unable to close \""+file.Name()+"\" file:", err)
		log.Fatal(err)
	}
}
