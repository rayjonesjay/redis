// resp package deals with serializing and deserializing the data
package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+' // represents data as string
	ERROR   = '-' // represent error
	INTEGER = ':' // represents data as int
	BULK    = '$' // represents data as long string
	ARRAY   = '*' // represents data as array or list
)

type Input struct {
	ty_pe string  // will be used to determine the data type carried by input
	str   string  // holds the value of string received simple strings
	num   int     // num holds the integers received integers
	bulk  string  // bulk is used to store string received from bulk strings
	array []Input // array holds all the arrays received from arrays
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

// readLine reads the line from the buffer
/* Logic used:
we read one byte at a time until we reach '\r'
which indicates the end of a line. Then we return the line without the last 2 bytes
which are \r\n and the number of bytes in the line
*/
func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		// ReadByte() reads one byte at a time
		b, err := r.reader.ReadByte()

		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)

		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {

	line, n, err := r.readLine()

	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Resp)  readArray() (Input,error){
	in := Input{}

	in.ty_pe = "array"

	// read length of array
	length , _, err := r.readInteger()
	if err != nil {
		return in, err
	}

	// for each line parse and read the value
	in.array = make([]Input, 0)
	for i := 0; i < length; i++ {
		val , err := r.Read()
		if err != nil {
			return in, err
		}

		// append parsed value to array
		in.array = append(in.array, val)
	}
	return in,nil
}


func (r *Resp) readBulk() (Input,error) {
	in := Input{}

	in.ty_pe = "bulk"

	len, _, err := r.readInteger()
	if err != nil {
		return in, err
	}

	bulk := make([]byte, len)

	r.reader.Read(bulk)
	in.bulk = string(bulk)

	// read the trailing CRLF
	r.readLine()

	return in,nil
}
func (r *Resp) Read() (Input, error) {
	ty_pe, err := r.reader.ReadByte()

	if err != nil {
		return Input{}, err
	}

	switch ty_pe {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("unknown type: %v", string(ty_pe))
		return Input{}, nil
	}
}
