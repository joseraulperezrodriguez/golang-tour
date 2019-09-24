package io

import (
	"fmt"
	"io"
	"strings"
)

//Reader1 test
func Reader1() {
	r := strings.NewReader("Hello Readers!")
	buffer := make([]byte, 8)

	for {
		readed, error := r.Read(buffer)
		if error == nil {
			fmt.Println("readed", readed, string(buffer))
		} else if error != nil {
			if error == io.EOF {
				break
			}
			fmt.Println(error.Error())
		}
	}

}
