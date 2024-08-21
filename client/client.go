// client package contains functionality for handling client connections, 
// communication with the server 
package main 

import (
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"os"
)
func main() {
	input := "$5\r\nworld\r\n"

	reader := bufio.NewReader(strings.NewReader(input))

	dataType , _ := reader.ReadByte()

	if dataType != '$'{
		fmt.Println("invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size , _ := reader.ReadByte()
	s := reader.UnreadByte()
	fmt.Println(s)


	stringSize , _ := strconv.ParseInt(string(size), 10, 64)

	fmt.Println(stringSize)

	// consume the next two bytes /r/n
	reader.ReadByte()
	reader.ReadByte()

	text := make([]byte,stringSize)

	reader.Read(text)

	fmt.Println("hello world")

}

