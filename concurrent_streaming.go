//
// Goroutines, channels and io.Readers/Writers
// more complex usage example
//
package main

import(
	"fmt"
	"os"
	"io"
//	"bufio"
//	"strings"
	"sync"
)


// Marshaller interface
type MrshallerSQL interface {
	Read(p []byte) (readBytes int, err error)
}
// 
type QuerySQL struct {
	source io.Reader // wrap an existing reader and use its stream
	text string
}
//
func NewQuerySQL(source io.Reader) *QuerySQL {
	return &QuerySQL{
		text: "",
		source: source,
	}
}
// Read method inmplementation
func (q *QuerySQL) Read(p []byte) (int, error) {
	fmt.Println("method *READ")
	if len(p) == 0 {
		return 0, nil
	}

	count, err := q.source.Read(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR", count, err)
		return count, err
	}
	//
	// NOTE: 
	// As a guideline, implementations of the io.Reader should return an error value of io.EOF 
	// when the reader has no more data to transfer into stream p .
	//
	return count, io.EOF
}



// serial application
func main() {
	/*
	// make originate srouce
	fin, err := os.Open("./test_case10.csv")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer fin.Close()

	query := bufio.NewReader(fin)
	tokenizer := bufio.NewReader(query)

	for stop := false; stop != true; {
		switch line, err := tokenizer.ReadString('\n'); err {
		case io.EOF: // on end of file
			stop = true
		case nil: // on success
			fmt.Print("|>", line)
		default: // on error
			fmt.Println("Error reading:", err)
			os.Exit(1)
		} // eof switch
	} // eof for
	*/



} // eof main


