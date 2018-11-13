//
// Data IO in Go
//
package main

import(
	"fmt"
	"os"
	"io"
	"strings"
	"bufio"
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


type metalloid struct {
	name string
	number int32
	weigth float64
}


//
// main driver
//
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


	// Working with files

	f1,err := os.Open("./datio.go") // reblace by existing and not existing filename :)
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer f1.Close()

	f2,err := os.Create("./datio.go.bkp")
	if err != nil {
		fmt.Println("Unable to create file: ", err)
		os.Exit(1)
	}
	defer f2.Close()

	n,err := io.Copy(f2,f1)
	if err != nil {
		fmt.Println("Failed to copy: ", err)
		os.Exit(1)
	}

	fmt.Printf("First: Copied %d bytes from %s to %s\n", n, f1.Name(), f2.Name())


	//	The os.OpenFile function take three parameters.
	f3, err := os.OpenFile("./datio.go", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer f3.Close()

	//* Unable to create file:  open ./datio.go.back: no such file or directory
	// if there is no file exists
	f4, err := os.OpenFile("./datio.go.back", os.O_WRONLY|os.O_RDONLY, 0666)
	if err != nil {
		//* fmt.Println("Unable to create file: ", err)
		//* os.Exit(1)
		// solved by creating one by mysefl :)
		// here is just an initialization(=) of existitng variables but not a short variable declaration(:=) of them
		f4, err = os.Create("./datio.go.back")
		if err != nil {
			fmt.Println("Unable to create file: ", err)
			os.Exit(1)
		}
	}
	defer f4.Close()

	n1, err := io.Copy(f4, f3)
	if err != nil {
		fmt.Println("Unable to copy: ", err)
		os.Exit(1)
	}

	fmt.Printf("Second: Copied %d bytes from %s to %s\n", n1, f3.Name(), f4.Name())


	// Files writing and reading
	rows := []string{
		"The quick brown fox",
		"jumps over the lazy dog",
	}

	fmt.Println("Crating the file")
	fout, err := os.Create("./filewrite.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()

	fmt.Println("Writing to the file")
	for _, row := range rows {
		fout.WriteString(row)
	}
	// finalize file
	fout.WriteString("\n")
	fmt.Println("Done!")

	// If, however, the source of your data is not text, you can write raw bytes directly to the file
	rawData := [][]byte{
		[]byte("The quick brown fox\n"),
		[]byte("jumps over the lazy dog\n"),
	}

	fmt.Println("Crating the file")
	fout1, err := os.Create("./filewrite.dat")
	if err != nil {
		fmt.Println("Unable to create file: ", err)
		os.Exit(1)
	}
	defer fout1.Close()

	fmt.Println("Writing to the file")
	for _, out := range rawData {
		fout1.Write(out)
	}
	fmt.Println("Done!")


	// As an io.Reader , reading from of the io.File type directly can be done using the Read
	// method. This gives access to the content of the file as a raw stream of byte slices. The
	// following code snippet reads the content of file ./concurrency.go as raw bytes assigned
	// to slice p up to 256-byte chunks at a time:
	fin, err := os.Open("./concurrency.go")
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer fin.Close()

	// raw bytes assigned to slice p up to 256-byte chunks at a time
	p := make([]byte, 256)
	for {
		n, err := fin.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
		fmt.Printf("\n---------Part len: %d bytes---------\n", n)
	}


	//
	// Standard input, output, and error
	//
	f5, err := os.Open("./datio.go")
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer f5.Close()

	n2, err := io.Copy(os.Stdout, f5)
	if err != nil {
		fmt.Println("Failed to copy: ", err)
		os.Exit(1)
	}
	fmt.Printf("Copied %d bytes from %s\n", n2, f5.Name())


// Formatted IO with fmt
// One of the most widely used packages for IO is fmt ( h t t p s : / / g o l a n g . o r g / p k g / f m t ). It
// comes with an amalgam of functions designed for formatted input and output. The most
// common usage of the fmt package is for writing to standard output and reading from
// standard input.	

	var metalloids = []metalloid{
		{"Boron", 5, 10.81},
		//...
		{"Polonium", 84, 209.0},
	}

	fmt.Println("Crating metalloids output file")
	file1, err := os.Create("./metalloids.txt")
	if err != nil {
		fmt.Println("Unable to create file: ", err)
		os.Exit(1)
	}
	defer file1.Close()

	fmt.Println("Writing both to file and to stdout")
	for _, m := range metalloids {
		// out to file
		fmt.Fprintf(file1, "%-10s %-10d %-10.3f\n", m.name, m.number, m.weigth)
		// out to STDOUT
		fmt.Printf("%-10s %-10d %-10.3f\n", m.name, m.number, m.weigth)
	}
	fmt.Println("Done!")

	// Reading from io.Reader

	var name, hasRing string
	var diam, moons int

	// read data from file
	dataIn, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer dataIn.Close()

	var stop bool = false
	for !stop {
		// scan and switch it in one string
		switch _, err := fmt.Fscanf(dataIn, "%s %d %d %s\n", &name, &diam, &moons, &hasRing); err {
		case io.EOF: stop = true // end of file then stop the cycle
		case nil: // on success show the file content
			fmt.Printf("%-10s %-10d %-6d %-6s\n", name, diam, moons, hasRing)
		default: // if error occured
			fmt.Println("Scan error: ", err)
			os.Exit(1)
		} // eof switch
	} // eof for

	// Reading from standard input
	var choice int
	fmt.Println("A squrare is what?")
	fmt.Print("Enter 1=qudrilateral 2=rectagonal\n>")

	//
	// NOTE!
	// bug? you can input 'als' and program will terminate with exit reason 1 
	// but furthermore it EXECUTE 'ls' command!
	// I suppose that you shouldn't usr the scanf in the real world...
	//
	n3, err := fmt.Scanf("%d", &choice)
	if n3 != 1 || err != nil {
		fmt.Println("Follow derections!")
		os.Exit(1)
	}
	if 1 == choice {
		fmt.Println("You are correct!")
	} else {
		fmt.Println("Wrong, Google it.")
	}


	//
	// Buffered IO
	// Most IO operations covered so far have been unbuffered. This implies that each read and
	// write operation could be negatively impacted by the latency of the underlying OS to handle
	// IO requests. Buffered operations, on the other hand, reduces latency by buffering data in
	// internal memory during IO operations.
	//

	rowsBuf := []string{
		"The quick borwn fox",
		"jumps over the lazy dog",
	}

	fmt.Println("Craating file")
	fout2, err := os.Create("./filewrite_buf.dat")
	fmt.Println("Creating writer")
	writer := bufio.NewWriter(fout2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout2.Close()
	fmt.Println("Writing to file")
	for _, row := range rowsBuf {
		writer.WriteString(row)
	}
	// finalize file
	writer.WriteString("\n")
	// and flush buffer
	writer.Flush()
	fmt.Println("Done")



	// read data from file
	planetesFile, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer planetesFile.Close()

	reader := bufio.NewReader(planetesFile)

	stop = false
	for !stop {
		// scan and switch it in one string
		switch line, err := reader.ReadString('\n'); err {
		case io.EOF: stop = true // end of file then stop the cycle
		case nil: // on success pring the line 
			fmt.Print("line: ", line)
		default: // if error occured
			fmt.Println("Error reading: ", err)
			os.Exit(1)
		} // eof switch
	} // eof for
} // eof main
