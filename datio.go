//
// Data IO in Go
//
package main

import(
	"fmt"
	"os"
	"io"
	"strings"
)


type alphaReader string


type bettaReader struct {
	src io.Reader
}

func NewBettaReader(source io.Reader) *bettaReader {
	return &bettaReader{source}
}

func (a *bettaReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count, err := a.src.Read(p) // p has now soucce data
	if err != nil {
		return count, err
	}
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p [i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	} // eof for
	return count, io.EOF
}


// The stream of data is represented as a slice of bytes ([]byte)
// that can be accessed for reading or writing.
type Reader interface {
	Read(p []byte) (n int, err error)
}

func (a alphaReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(a); i++ {
		if (a[i] >= 'A' && a[i] <= 'Z') || (a[i] >= 'a' && a[i] <= 'z') {
			p[i] = a[i]
		}
		count++
	}
	return count, io.EOF
}


type Writer interface {
	Write(P []byte) (n int, err error)
}

type channelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *channelWriter {
	return &channelWriter {
		Channel: make(chan byte, 1024), // create byte channel of 1024 bytes size
	}
}

func (c *channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	go func() {
		defer close(c.Channel) // when done
		for _, b := range p {
			c.Channel<- b
		}
	}()
	return len(p), nil
}


func main() {
	// io.Reader interface
	// alpha
	str := alphaReader("Hello! Where is the sun?")
	io.Copy(os.Stdout, &str)
	fmt.Println()
	// betta
	str2 := strings.NewReader("hello! Where is the Sun?")
	betta := NewBettaReader(str2)
	io.Copy(os.Stdout, betta)
	fmt.Println()
	// The advantages of this approach may not be obvious at first.
	// However, by using an io.Reader as the underlying data source 
	// the alphaReader type is capable of reading from any reader implementation. 
	// For instance, the following code snippet shows how the bettaReader 
	// type can now be combined with an os.File to filter out non-alphabetic
	// characters from a file.
	file, _ := os.Open("./hollogo.go")
	betta = NewBettaReader(file)
	io.Copy(os.Stdout, betta)
	fmt.Println()

	// The io.Writer interface
	//
	cw := NewChannelWriter()
	go func() {
		fmt.Fprint(cw, "Stream me!")
	}()
	// the serialized bytes, queued in the channel, are consumed using a
	// for...range statement as they are successively printed
	for c := range cw.Channel {
		fmt.Printf("consumer: %c\n", c)
	}

	// The following snippet shows another example where the content of 
	// a file is serialized over a channel using the same channelWriter. 
	// In this implementation, an io.File value and io.Copy function are used
	// to source the data instead of the fmt.Fprint function
	cw2 := NewChannelWriter()
	file2, err := os.Open("./hollogo.go")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	_,err = io.Copy(cw2, file2)
	if err != nil {
		fmt.Println("Error copying: ", err)
		os.Exit(1)
	}

	// conusme channel
	for c := range cw2.Channel {
		fmt.Printf("consumer2: %c\n", c)
	}
}
