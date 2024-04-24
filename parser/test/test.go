package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"silver-data/parser-go/lexing"
)

/*

defaults for:
	folder
	input filter
	ouptut formats

operations:
	test
	regenerate output

test:
	calculate a set set based upon folders/filters/formats
	read each file
		generate AST
		stringify
		record comparison against existing output
	return test results

regenerate output:
	calculate a set set based upon folders/filters/formats
	read each input file
		generate AST
		stringify
		write result
	return output rewrites


*/

/*
func test() {
	//
}
*/

var testFolder = "./test/file/"
var testFilename = "factorial.agl"

func GetFiles() []string {
	testFilePath := fmt.Sprintf("%v%v", testFolder, testFilename)
	return []string{testFilePath}
}

// generate output for an array of files
func GenerateOutput(fileArray []string) {

	for index, thisFile := range fileArray {
		fmt.Printf("%v: %v \n", index, thisFile)

		input := ReadFile(thisFile)

		abstractSyntaxTree := lexing.LexDocument(input)
		outputFilepath := fmt.Sprintf("%s.ast.json", thisFile)

		json, err := json.MarshalIndent(abstractSyntaxTree, "", "	")
		if err != nil {
			log.Fatalf(err.Error())
		}
		WriteFile(outputFilepath, string(json))
	}
}

func ReadFile(filepath string) string {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v \n", err)
	}
	defer file.Close()

	fmt.Printf("File opened: %v \n", filepath)

	// Read file
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}
	fmt.Printf("File read: %v \n", filepath)

	return string(content)
}

func WriteFile(filepath string, content string) {
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v \n", err)
	}
	fmt.Printf("File written: %v \n", filepath)
}
