package main

import (
	"bytes"
	"dataprocessing"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var firstTestedProgramName string  // a
var secondTestedProgramName string // b

func run(testCase string, resultChannel chan string) {
	cmd := exec.Command(firstTestedProgramName)

	var aOut bytes.Buffer
	cmd.Stdout = &aOut
	cmd.Stdin = strings.NewReader(testCase)

	aErr := cmd.Run()
	if aErr != nil {
		log.Fatal(aErr)
	}

	cmd = exec.Command(secondTestedProgramName)
	var bOut bytes.Buffer
	cmd.Stdout = &bOut
	cmd.Stdin = strings.NewReader(testCase)

	bErr := cmd.Run()
	if bErr != nil {
		log.Fatal(bErr)
	}

	testOutputString := fmt.Sprintf("Case:\n%s\n Result: %s %s", testCase, aOut.String(), bOut.String())
	// TODO возникают лишние переводы строки для каждого %s

	if aOut.String() == bOut.String() { // TODO подумать над форматом вывода
		resultChannel <- "correct " + testOutputString
	} else {
		resultChannel <- "incorrect " + testOutputString
	}
}

// TODO цветной вывод в консоль для лучшей читаемости
func main() { // TODO поработать над разбиением на функции
	args := os.Args

	if len(args) != 4 { // TODO добавить --help
		fmt.Println("Invalid command-line arguments.\n" +
			"Type ./tester --help to help. ")
		os.Exit(1)
	}

	resultFile, err := os.Create("result.txt")
	if err != nil { // TODO добавить возможность задавать output-файл
		fmt.Println("Unable to create result file:", err)
		os.Exit(1)
	}
	defer resultFile.Close()

	testsSrcFile, err := os.Open(args[3])
	if err != nil {
		fmt.Println("Unable to open tests source file:", err)
		os.Exit(1)
	}

	firstTestedProgramName = args[1]
	secondTestedProgramName = args[2]
	testCases := dataprocessing.SplitTests(testsSrcFile)

	fmt.Println("The tests are ready to run.")

	var resultChannel = make(chan string)

	bufferedStringsCounter := 0 // буфферизация записи в файл
	var bufStr bytes.Buffer

	for _, testCase := range testCases {
		go run(testCase, resultChannel)

		bufStr.WriteString(<-resultChannel)
		bufferedStringsCounter++
		if bufferedStringsCounter == 1000 {
			resultFile.Write(bufStr.Bytes())
			bufStr = bytes.Buffer{}
			bufferedStringsCounter = 0
		}
	}

	if bufferedStringsCounter <= 1000 { // дочищаем буффер
		resultFile.Write(bufStr.Bytes())
		bufStr = bytes.Buffer{}
	}

	close(resultChannel)
	fmt.Println("The output is in \"result.txt\"")
}
