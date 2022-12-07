package dataprocessing

import (
	"bytes"
	"fileworkers"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func run(testCase string, resultChannel chan string, firstTestedProgramName string, secondTestedProgramName string) {
	cmd := exec.Command(firstTestedProgramName)

	var aOut bytes.Buffer
	cmd.Stdout = &aOut
	cmd.Stdin = strings.NewReader(testCase)

	aErr := cmd.Run()
	if aErr != nil {
		log.Printf("RE in program %s in case:\n%s", firstTestedProgramName, testCase)
		log.Fatal(aErr)
	}

	cmd = exec.Command(secondTestedProgramName)
	var bOut bytes.Buffer
	cmd.Stdout = &bOut
	cmd.Stdin = strings.NewReader(testCase)

	bErr := cmd.Run()
	if bErr != nil {
		log.Printf("RE in program %s in case:\n %s\n", firstTestedProgramName, testCase)
		log.Fatal(bErr)
	}

	firstResult, secondResult := strings.TrimSpace(aOut.String()), strings.TrimSpace(bOut.String())
	testOutputString := fmt.Sprintf("сase:\n%s%s: %s\n%s: %s\n\n",
		testCase, firstTestedProgramName, firstResult, secondTestedProgramName, secondResult)

	if firstResult != secondResult {
		resultChannel <- "false " + testOutputString
	} else {
		resultChannel <- "true " + testOutputString // запись успешно пройденных тестов
		//resultChannel <- "_" // для пропуска успешно пройденных тестов
	}
}

func StartTests(firstTestedProgramName string, secondTestedProgramName string, testCases []string, outputFileName string) {
	const fileBufferSize = 10000
	var resultChannel = make(chan string)
	resultFile := fileworkers.CreateFile(outputFileName)

	// буфферизация записи в файл
	var bufStr bytes.Buffer

	for _, testCase := range testCases {
		go run(testCase, resultChannel, firstTestedProgramName, secondTestedProgramName)
		str := <-resultChannel
		if str == "_" {
			continue
		}

		bufStr.WriteString(str)
		if bufStr.Len() >= fileBufferSize {
			resultFile.Write(bufStr.Bytes())
			bufStr.Reset()
		}
	}

	if bufStr.Len() != 0 { // дочищаем буффер
		resultFile.Write(bufStr.Bytes())
		bufStr.Reset()
	}

	fileworkers.CloseFile(resultFile)
	close(resultChannel)
}
