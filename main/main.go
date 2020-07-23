package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	message string
}

func (businessError BusinessError) Error() string {
	return businessError.message
}

func testError1(i int) (string, error) {
	if i < 10 {
		return "Fail", BusinessError{message: "i can't < 0"}
	}
	return "Success", nil
}

func testError2(i int) (string, error) {
	if i > 10 {
		return "Fail", errors.New("i can't > 10")
	}
	return "Success", nil
}
func main() {
	str, err := testError1(9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = testError2(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
}
