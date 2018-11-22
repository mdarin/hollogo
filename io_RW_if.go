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
//                         Stream of bytes ([]byte)
//  Input with Reader:  
// Reads bytes into stram
//                                                      Output with Writer:
//                                                      writes bytes form stream

import(
	"fmt"
	"os"
	"io"
	"strings"
	"crypto/sha1"
	"crypto/md5"
	"compress/gzip"
)


// The io.Reader interface
//
// The io.Reader interface, as shown in the following listing, is simple.
// It consists of a single method, Read([]byte)(int, error) , 
// intended to let programmers implement code that reads data, 
// from an arbitrary source, and transfers it into the provided slice of bytes.
//
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
	fmt.Println("method: alpha.Read ::a ->", a)
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
// Chaining readers
// Chances are the standard library already has a reader that you can reuse – so it is common
// to wrap an existing reader and use its stream as the source for the new implementation.
//
type bettaSource struct {
	src io.Reader // wrap an existing reader and use its stream
}
// constructor
func NewBettaSource(source io.Reader) *bettaSource {
	return &bettaSource{source}
}
//
func (b *bettaSource) Read(p []byte) (int, error) {
	fmt.Println("method: betta.Read ::b.src ->", b.src)

	if len(p) == 0 {
		return 0, nil
	}

	// set up 'p' as a data source(target) to transfer into it
	// from data source(origin) b.src as io.Reader that initialized by string
	count, err := b.src.Read(p)
	if err != nil {
		return count, err
	}

	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			continue;
		} else {
			p[i] = 0
		}
	} // eof for
//
// NOTE: As a guideline, implementations of the io.Reader should return an error value of io.EOF 
// when the reader has no more data to transfer into stream p .
//
	return count, io.EOF
} // eof func



// The io.Writer interface
//
// The interface requires the implementation of a single method, Write(p []byte)(c int,e error) ,
// that copies data from the provided stream p and writes that data to a sink
// resource such as an in-memory structure, standard output, a file, a network connection,
// or any number of io.Writer implementations that come with the Go standard library.
//
type Writer interface {
	Write(p []byte) (n int, err error)
}
//
type channelSource struct {
	Channel chan byte
}
// constructor
func NewChannelSource() *channelSource {
	return &channelSource{Channel: make(chan byte, 1024)}
}
// writer method realization
func (c *channelSource) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	go func() {
		// close channel when done
		defer close(c.Channel)
		// read input stream p and sink it into the channel
		for _, b := range p {
			c.Channel <-b
		}
	}() // of goroutine
//	
// NOTE: The Write method returns the number of bytes copied from p 
// followed by an error value if any was encountered.
//
	return len(p), nil
} // eof func





//
// main driver
//
func main() {
	// reader
	alpha := alphaSource("Hello! Where is the Sun?")
	_, err := io.Copy(os.Stdout, &alpha)
	if err != nil {
		fmt.Println("Error copying:", err)
		os.Exit(1)
	}

	str := strings.NewReader("Hello! Where is the Sun?")
	betta := NewBettaSource(str)
	_, err = io.Copy(os.Stdout, betta)
	if err != nil {
		fmt.Println("Error copying:", err)
		os.Exit(1)
	}

	fmt.Println("Reading from file and writing into STDOUT")

	file_a, err := os.Open("./hollogo.go")
	if err != nil {
		fmt.Println("Unable to open file", err)
		os.Exit(1)
	}
	defer file_a.Close()

	betta = NewBettaSource(file_a)
	_, err = io.Copy(os.Stdout, betta)
	if err != nil {
		fmt.Println("Error copying:", err)
		os.Exit(1)
	}

	fmt.Println()

	// writer
	gamma := NewChannelSource()
	// craate a data source origin(Reader)
	go func() {
		fmt.Fprint(gamma, "Stream me!")
	}()

	// consume channel
	for c := range gamma.Channel {
		fmt.Printf("[%c]", c)
	}
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()

	writer := NewChannelSource()
	file_b, err := os.Open("./hollogo.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	defer file_b.Close()

	_, err = io.Copy(writer, file_b)
	if err != nil {
		fmt.Println("Error copying:", err)
		os.Exit(1)
	}
	// consume channel
	for c := range writer.Channel {
		fmt.Printf("[%c]", c)
	}
	fmt.Println()



	fin, _ := os.Open("./hollogo.go")
	defer fin.Close()
	fout, _ := os.Create("./teereader.gz")
	defer fout.Close()

	zip := gzip.NewWriter(fout)
	defer zip.Close()

	sha := sha1.New()

	data := io.TeeReader(fin, sha)
	io.Copy(zip, data)
	fmt.Printf("SHA1 hash %x\n", sha.Sum(nil))

	md := md5.New()

	data2 := io.TeeReader(io.TeeReader(fin, md), sha)
	io.Copy(zip, data2)

	// let's try to create more complex schema where 
	// we have got a FileReader as a sourceOrigin for DataTransfromer that is a sourceTarger for the FileReader
	// furthermore our DataTransformer is a sourceOrigin for another on component FileWriter which is the end of the our schema
	// fin = os.Open()
	// fout = os.Create()
	// r.FileReader(FileIn)
	// w.FileWriter(FileOut)
	// t.DataTransformer(r,w)

} // eof main
