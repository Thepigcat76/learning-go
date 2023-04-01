package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(inputMessage string, errorMessage string, outputMessage string) {
	fmt.Print(inputMessage)
	var reader = bufio.NewReader(os.Stdin) //Assigns reader variable that reads from terminal
	// ReadString will block until the delimiter is entered
	var input, err = reader.ReadString('\n')
	if err != nil { //Creates an error message
		fmt.Println(errorMessage, err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n") //Cleans up the output
	fmt.Println(outputMessage + input)
}