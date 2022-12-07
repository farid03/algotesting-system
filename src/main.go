package main

import (
	"dataprocessing"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 4 { // TODO добавить --help
		fmt.Println("Invalid command-line arguments.\n" +
			"Type ./tester --help to help. ")
		os.Exit(1)
	}

	testsSrcFile, err := os.Open(args[3])
	if err != nil {
		log.Println("Unable to open tests source file:", err)
		os.Exit(1)
	}
	testCases := dataprocessing.SplitTests(testsSrcFile)
	log.Println("The tests are ready to run. In total", len(testCases), "cases.")

	var firstProgramName = args[1]
	var secondProgramName = args[2]
	dataprocessing.StartTests(firstProgramName, secondProgramName, testCases, "result.txt")
	//TODO добавить возможность задавать файл для выходных данных
	log.Println("The output is in \"result.txt\".")
}

// TODO добавить режим быстрой прогонки тестов без записи правильных ответов в файл, только неправильные || DEFAULT
// TODO режим без записи любых результатов, только результат совпадения во всех тестах
