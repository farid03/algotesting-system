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

	aResult, bResult := strings.TrimSpace(aOut.String()), strings.TrimSpace(bOut.String())
	testOutputString := fmt.Sprintf("сase:\n%sa: %s\nb: %s\n\n", testCase, aResult, bResult)

	if aResult != bResult {
		resultChannel <- "false " + testOutputString
	} else {
		resultChannel <- "true " + testOutputString // TODO избавиться от записи в файл успешно пройденных тестов
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
		strings.TrimSpace(str)
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
