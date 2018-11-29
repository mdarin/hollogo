//
// Syncronization via channels and communication 
// complex topology of workers
//
package main

import(
	"fmt"
	"strings"
	"os"
	"time"
)

//FIXME:
// Check if a Go Channel is open or closed 
// https://gist.github.com/iamatypeofwalrus/84b6c7d946a6a4143a1d
// panic: send on closed channel

//
// You should restart number of times to see different rusult
//
// you may adjust this value to activate timeout case in the select statement
// it should be less then current value
// it's a boundary value now
const TIMEOUT = 80000

const WORKERS = 3

func main() {

	data := []string{
		"The yellow fish swims",
		"Blue is the water in a crystal glass",
		"Blue are the lilacs that grow in the grass",
		"Blue is delicious blueberry pie",
		"Blue are the sparkles in my cats eyes",
	}

	type Communications struct {
		done chan struct{}
		doneGroups chan bool
		doneSubGroups chan bool
		progress chan string
		wordsCounter chan string
		// words counting group 1
		counterAccumulator chan int // workers
		doneCounter chan bool // leader
		// words processing group 2
		processorAccumulator chan string
		doneProcessor chan bool
		processor chan string
		// words encoding group 3
		encoderAccumulator chan string
		doneEncoder chan bool
		marshaller chan string
		encoder chan string
		// words encoding group 4(supervised)
		encoderXMLAccumulator chan string
		doneEncoderXML chan bool
		marshallerXML chan string
		encoderXML chan string
		progressXML chan string
	}

	var com Communications = Communications{
		done: make(chan struct{}),
		doneGroups: make(chan bool,4),
		doneSubGroups:  make(chan bool,4),
		progress: make(chan string),
		wordsCounter: make(chan string),
		// words counting group 1
		counterAccumulator:  make(chan int, WORKERS), // workers
		doneCounter: make(chan bool, WORKERS), // leader
		// words processing group 2
		processorAccumulator: make(chan string, WORKERS),
		doneProcessor: make(chan bool, WORKERS),
		processor: make(chan string, WORKERS),
		// words encoding group 3
		encoderAccumulator: make(chan string, WORKERS),
		doneEncoder: make(chan bool, WORKERS),
		marshaller: make(chan string, WORKERS),
		encoder: make(chan string, WORKERS),
		// words encoding group 4(supervised)
		encoderXMLAccumulator: make(chan string, WORKERS),
		doneEncoderXML: make(chan bool, WORKERS),
		marshallerXML: make(chan string, WORKERS),
		encoderXML: make(chan string, WORKERS),
		progressXML: make(chan string),
	}

	// STAND ALONE WORKER
	// process words worker
	go func(c *Communications) {
		defer close(c.progress)
		defer func() {
			//
			// To be able to recover from an unwinding panic sequence, 
			// the code must make a deferred call to the recover function.
			//
			if r := recover(); r != nil {
				fmt.Println("[W] Reader FAULT")
			}
		}()
		fmt.Printf(" * Reader started")
		for word := range wordsGenerator(data) {
			c.progress<- word
		}
		fmt.Println()
		fmt.Println(" * Reader terminated")
	}(&com)


	// STAND ALONE WORKER
	// progress monitor worker
	go func(c *Communications) {
		defer close(c.wordsCounter)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[W] Progress FAULT")
			}
		}()
		fmt.Println(" * Progress sarted")
		i := 0
		for word := range c.progress {
			i++
			fmt.Printf(".")
			(c.wordsCounter)<- word
		}
		fmt.Println()
		fmt.Println(" * Progress termitated")
	}(&com)

	// STAND ALONE WORKER
	// simple nested worker
	// counter worker
	// Distributor(router,switching station or so on...)
	go func(c *Communications) {
		fmt.Printf(" * Counter started")
		// child worker
		go func (c *Communications) {
			defer close(c.processorAccumulator)
			defer close(c.counterAccumulator)
			defer close(c.encoderAccumulator)
			defer close(c.progressXML)
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Counter child FAULT")
				}
			}()
			fmt.Println(" * Counter child")
			n := 0;
			for word := range c.wordsCounter {
				//fmt.Printf("{%s} ", word)
				n++
				// queue words, cycle for similar copy for every worker
			//	for i := 0; i < cap(processorAccumulator); i++ {
					c.processorAccumulator<- word
					c.encoderAccumulator<- word
					c.progressXML<- word
			//	}
			}
			// queue count result
			for marker := 0; marker < cap(c.counterAccumulator); marker++ {
				// markup for differ elements
				(c.counterAccumulator)<- n + marker
			}
			fmt.Println()
			fmt.Println(" * Counter child terminated")
		}(c)
		fmt.Println(" * Counter terminated")
	}(&com)

	// ---------- GROUP-------------------
	// WORKERS POOL
	// group or workers for processing queue
	// they dequeue values randomly
	go func(c *Communications) {
		for i := 0; i < cap(c.counterAccumulator); i++ {
			// Accountant worker
			go func(c *Communications, id int) {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("[W] Accauntant worker FAULT")
					}
				}()
				fmt.Printf(" * Accountant worker %d started\n", id)
				for wordsCount := range c.counterAccumulator {
					// worker's task
					fmt.Printf(" > Total %d: %d\n", id, wordsCount)
				}
				(c.doneCounter)<- true
				fmt.Println()
				fmt.Printf(" * Accountant worker %d terminated\n", id)
			}(c,i) // create worker ID
		}
	}(&com)
	// LEADER
	// group sycronizer(leader)
	go func(c *Communications) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[W] Group 1 leader FAULT")
			}
		}()
		fmt.Println(" * Group 1 leader started")
		workersDoneCount := 0
		for range c.doneCounter {
			workersDoneCount++
			fmt.Printf(" > G1 workers %d done\n", workersDoneCount)
			if workersDoneCount >= WORKERS {
				// stop cycle and terminate group
				close(c.doneCounter)
				fmt.Println()
				fmt.Println(" ! Done counting words")
			}
		}
		// done goroup signale
		(c.doneGroups)<- true
		fmt.Println(" * Group 1 leader terminated")
	}(&com)
	// ----------END GROUP-------------------

	// ---------- GROUP-------------------
	// nested comlex structure
	// starter parent worker
	go func(c *Communications) {
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func(c *Communications) {
			for i := 0; i < cap(c.processorAccumulator); i++ {
				// Word processor worker
				go func(c *Communications, id int) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Println("[W] Word processor worker FAULT")
						}
					}()
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range c.processorAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						(c.processor)<- fmt.Sprintf("{%s}", word)
					}
					// done when queue is empty
					(c.doneProcessor)<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(c,i) // create worker ID
			}
		}(c)
		// SNGLE WORKER
		// writer worker
		go func(c *Communications) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Writer worker FAULT")
				}
			}()
			fmt.Println(" * Writer started")
			fout, _ := os.Create("./output.txt")
			defer fout.Close()
			for processedWord := range c.processor {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "%s\n", processedWord)
			}
			fmt.Fprintf(fout, "\n")
			fmt.Println(" * Writer terminated")
		}(c)
		go func(c *Communications) {
			// LEADER
			// group sycronizer(lider)
			go func(c *Communications) {
				defer close(c.processor)
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("[W] Group 2 leader FAULT")
					}
				}()
				fmt.Println(" * Group 2 leader started")
				workersDoneCount := 0
				for range c.doneProcessor {
					workersDoneCount++
					fmt.Printf(" > G2 workers %d done\n", workersDoneCount)
					if workersDoneCount >= WORKERS {
						// stop cycle and terminate group
						close(c.doneProcessor)
						// done goroup signale
						c.doneGroups<- true
						fmt.Println()
						fmt.Println(" ! Done processing words")
					}
				}
				fmt.Println(" * Group 2 leader terminated")
			}(c)
		}(c)
	}(&com)
	// ----------END GROUP-------------------


	// nested comlex structure
	// starter parent worker
	go func(c *Communications) {
		// ---------- GROUP-------------------
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func(c *Communications) {
			for i := 0; i < cap(c.encoderAccumulator); i++ {
				// Word processor worker
				go func(c *Communications, id int) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Println("[W] Word processor FAULT")
						}
					}()
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range c.encoderAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						(c.marshaller)<- fmt.Sprintf("%s", word)
					}
					// done when queue is empty
					(c.doneEncoder)<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(c,i) // create worker ID
			}
		}(c)
		// SNGLE WORKER
		// marshaller worker
		go func(c *Communications) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Marshaller JSON FAULT")
				}
			}()
			fmt.Println(" * Marshaller JSON started")
			for marshalledWord := range c.marshaller {
					// worker's task
					(c.encoder)<- fmt.Sprintf("'%s': %s,\n", marshalledWord, marshalledWord)
			}
			fmt.Println(" * Marshaller JSON terminated")
		}(c)
		// SINGLE WORKER
		// writer worker
		go func(c *Communications) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Writer JSON FAULT")
				}
			}()
			fmt.Println(" * Writer JSON started")
			fout, _ := os.Create("./output.json")
			defer fout.Close()
			fmt.Fprintf(fout, "{\n")
			for encoderWord := range c.encoder {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "\t%s", encoderWord)
			}
			fmt.Fprintf(fout, "}\n")
			fmt.Println(" * Writer JSON terminated")
		}(c)
		go func(c *Communications) {
			// LEADER
			// group sycronizer(lider)
			go func(c *Communications) {
				defer close(c.marshaller)
				defer close(c.encoder)
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("[W] Group 3 leader FAULT")
					}
				}()
				fmt.Println(" * Group 3 leader started")
				workersDoneCount := 0
				for range c.doneEncoder {
					workersDoneCount++
					fmt.Printf(" > G2 workers %d done\n", workersDoneCount)
					if workersDoneCount >= WORKERS {
						// stop cycle and terminate group
						close(c.doneEncoder)
						// done goroup signale
						(c.doneGroups)<- true
						fmt.Println()
						fmt.Println(" ! Done marshalling words")
					}
				}
				fmt.Println(" * Group 3 leader terminated")
			}(c)
		}(c)
		// ----------END GROUP-------------------
	}(&com)


	// ------------------------
	// Subpervised subtree in a supervised tree
	//
	// STAND ALONE WORKER
	// progress monitor #2 worker	
	go func(c *Communications) {
		defer close(c.encoderXMLAccumulator)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[W] Progress XML FAULT")
			}
		}()
		fmt.Println(" * Progress XML sarted")
		i := 0
		for word := range c.progressXML {
			i++
			fmt.Printf("%s","'")
			(c.encoderXMLAccumulator)<- word
		}
		fmt.Println()
		fmt.Println(" * Progress XML termitated")
	}(&com)

	// ---------- GROUP-------------------
	// nested comlex structure
	// starter parent worker
	go func(c *Communications) {
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func(c *Communications) {
			for i := 0; i < cap(c.encoderXMLAccumulator); i++ {
				// Word processor worker
				go func(c *Communications, id int) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Println("[W] Word process worker FAULT")
						}
					}()
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range c.encoderXMLAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						(c.marshallerXML)<- fmt.Sprintf("%s", word)
					}
					// done when queue is empty
					(c.doneEncoderXML)<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(c,i) // create worker ID
			}
		}(c)
		// SINGLE WORKER
		// marshaller worker
		go func(c *Communications) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Marshaler XML FAULT")
				}
			}()
			fmt.Println(" * Marshaller XML started")
			for marshalledWord := range c.marshallerXML {
					// worker's task
					(c.encoderXML)<- fmt.Sprintf("<record value=\"%s\" name=\"%s\"/>\n", marshalledWord, marshalledWord)
			}
			fmt.Println(" * Marshaller XML terminated")
		}(c)
		// SINGLE WORKER
		// writer worker
		go func(c *Communications) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Writer XML FAULT")
				}
			}()
			fmt.Println(" * Writer XML started")
			fout, _ := os.Create("./output.xml")
			defer fout.Close()
			fmt.Fprintf(fout, "<table name=\"Records\">\n")
			for encoderWord := range c.encoderXML {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "\t%s", encoderWord)
			}
			fmt.Fprintf(fout, "</table>\n")
			fmt.Println(" * Writer XML terminated")
		}(c)
		// LIDER
		// group sycronizer(lider)
		go func(c *Communications) {
			defer close(c.marshallerXML)
			defer close(c.encoderXML)
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("[W] Group 4 leader FAULT")
				}
			}()
			fmt.Println(" * Group 4 leader started")
			workersDoneCount := 0
			for range c.doneEncoderXML {
				workersDoneCount++
				fmt.Printf(" > G4 workers %d done\n", workersDoneCount)
				if workersDoneCount >= WORKERS {
					// stop cycle and terminate group
					close(c.doneEncoderXML)
					// done goroup signale
					(c.doneSubGroups)<- true
					fmt.Println()
					fmt.Println(" ! Done marshalling words")
				}
			}
			fmt.Println(" * Group 4 leader terminated")
		}(c)
	}(&com)
	// ----------END GROUP-------------------

	// SUPERVISOR
	// subtree groups sycronizer(Supervisor)
	go func(c *Communications) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[W] Supervisor 2 FAULT")
			}
		}()
		fmt.Println(" * Supervisor 2 started")
		groupsDoneCount := 0
		for range c.doneSubGroups {
			groupsDoneCount++
			fmt.Printf(" # group %d done\n", groupsDoneCount)
			if groupsDoneCount >= 1 { // for all groups
				// stop cycle and terminate all supervised groups tree
				//close(c.doneSubGroups)
				// signal everything done in sub group
				(c.doneGroups)<- true
				fmt.Println()
				fmt.Println(" ! Done supervised subtree")
			}
		}
		fmt.Println(" * Supervisor 2 terminated")
	}(&com)
	// ------------------END Supervised subtree-----------

	// ROOT SUPERVISOR
	// All groups sycronizer(Supervisor)
	go func(c *Communications) {
		// signal everything done
		defer close(c.done)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[W] Supervisor 1 FAULT")
			}
		}()
		fmt.Println(" * Supervisor 1 started")
		groupsDoneCount := 0
		for range c.doneGroups {
			groupsDoneCount++
			fmt.Printf(" # group %d done\n", groupsDoneCount)
			if groupsDoneCount >= 4 { // for all groups
				// stop cycle and terminate all supervised groups tree
				close(c.doneGroups)
				// close supervised sub tree
				close(c.doneSubGroups)
				fmt.Println()
				fmt.Println(" ! Done supervised tree")
			}
		}
		fmt.Println(" * Supervisor 1 terminated")
	}(&com)

	// controller-sycronizer
	select {
	case <-(com.done):
		fmt.Println()
		fmt.Println()
		fmt.Println(" ! Done")
	case <-time.After(TIMEOUT * time.Microsecond):
//	case <-time.After(TIMEOUT * time.Millisecond):
		fmt.Println()
		fmt.Println()
		fmt.Println(" ! Sorry, took too long to count")
	}

} // eof main


func wordsGenerator(data []string) <-chan string {
  outChan := make(chan string)
  go func(c chan string) {
    defer close(c)
    for _, line := range data {
      words := strings.Split(line, " ")
      for _, word := range words {
        c<- word
      }
    }
  }(outChan)
  return outChan
}

