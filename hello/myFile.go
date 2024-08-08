package main

import (
	"fmt"
	"io"
	"os"
)

func WriteToFile() {
	content := "Let write something cool!"
	f, err := os.Create("./fromString.txt")

	if err != nil {
		fmt.Println("Error creating a file - fromString.txt: ", err)
		return
	}

	length, err := io.WriteString(f, content)

	if err != nil {
		fmt.Println("Error writing to file - fromString.txt: ", err)
		return
	}

	fmt.Printf("Wrote a file with %v characters\n", length)

	defer f.Close()
}

func ReadFromFile(fileName string) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error writing to file - fromString.txt: ", err)
		return
	}

	fmt.Println("Text from file: ", string(data))
}
