package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {
	Errors()
	SentinelErrors()
}

func Errors() {
	//* Errors
	// Go handles errors by returning a value of type error as the last return value of a function.
	// When a function executes as expected, `nil` is returned as the error.
	// If something goes wrong, an error value is returned. The calling function then checks the
	// error return value by comparing it with `nil`, handling the error or returning an error of
	// its own.

	// error msgs should not be capitalized nor should they end with a punctuation or new line.
	// when an error value is returned, the other return values should be set to their zero values.

	numerator := 20
	denominator := 3

	remainder, mod, err := calcRemainderAndMod(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remainder, mod)

	// there are two ways to create an error from a string:
	// `errors.New("string")`: takes in a string and returns an error
	// `fmt.Errorf("%v string", i)`: this allows to include runtime values in the error msg
}

func calcRemainderAndMod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("denominator is 0")
	}

	return num / den, num % den, nil
}

func SentinelErrors() {
	//* Sentinel Errors
	// Some errors are meant to signify that processing cannot continue because of a problem
	// with the current state. We call these Sentinel Errors.

	// Sentinel errors are one of the few variables that are declared at the package level.
	// They start with `Err` with the exception of io.EOF

	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("told you so")
	}
}
