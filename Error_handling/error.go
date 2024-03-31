package main

// import the errors package
import (
	"errors"
	"fmt"
)

func main() {

	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Programiz" {
		fmt.Println(myError)
	}

}
