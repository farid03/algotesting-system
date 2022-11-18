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

	testOutputString := fmt.Sprintf("сase:\n%sa: %sb: %s\n", testCase, aOut.String(), bOut.String())

	if aOut.String() != bOut.String() {
		resultChannel <- "false " + testOutputString
	} else {
		resultChannel <- "true " + testOutputString // TODO избавиться от записи в файл успешно пройденных тестов
	}
}

func StartTests(firstTestedProgramName string, secondTestedProgramName string, testCases []string, outputFileName string) {
	const fileBufferSize = 1000
	var resultChannel = make(chan string)
	resultFile := fileworkers.CreateFile(outputFileName)

	// буфферизация записи в файл
	var bufStr bytes.Buffer

	for _, testCase := range testCases {
		go run(testCase, resultChannel, firstTestedProgramName, secondTestedProgramName)

		bufStr.WriteString(<-resultChannel)
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
