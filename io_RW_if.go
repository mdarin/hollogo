//
// IO with readers and writers
// yet anothor exercise with reader and writer
//
package main

//
// Go models data input and output as a stream that flows from sources to targets.
//
// Data resources, such as files, networked connections, or even some in-memory objects, 
// can be modeled as streams of bytes from which data can be read or written to, 
// as illustrated in the following figure:
//
// (data source(origin)) ===|> [][][]...[][][] ===|> (data source(target))
//               						Stream of bytes ([]byte)
//  Input with Reader:  
// Reads bytes into stram
//                                                      Output with Writer:
//                                                      writes bytes form stream

import(
	"fmt"
	"os"
	"io"
)


// The io.Reader interface
//
// The io.Reader interface, as shown in the following listing, is simple. It consists of a single
// method, Read([]byte)(int, error) , intended to let programmers implement code that
// reads data, from an arbitrary source, and transfers it into the provided slice of bytes.
type Reader interface {
	Read(p []byte) (n int, err error)
}

type alphaSource string
//
// alphaSource Read method implementation
// The Read method returns the total number of bytes transferred into 
// the provided slice and an error value (if necessary). 
//
func (a alphaSource) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(a); i++ {
		if (a[i] >= 'A' && a[i] <= 'Z') || (a[i] >= 'a' && a[i] <= 'z') {
			p[i] = a[i]
		}
		count++
	} // eof for
//
// NOTE: As a guideline, implementations of the io.Reader should return an error value of io.EOF 
// when the reader has no more data to transfer into stream p .
//
	return count, io.EOF
} // eof func



//
// main driver
//
func main() {
	src := alphaSource("Hello! Where is the Sun?")
	io.Copy(os.Stdout, &src)
	fmt.Println()

} // eof main
