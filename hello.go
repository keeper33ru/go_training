package main

import (
	"fmt"
	"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {

			p[count] = ph[i]
			count++

		}
	}
	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+1(235) 58 547")
	phone2 := phoneReader("+2-3-458-74-6363")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)

	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(buffer)

}
