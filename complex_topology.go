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
	done := make(chan struct{})
	doneGroups := make(chan bool,4)
	doneSubGroups := make(chan bool,4)
	progress := make(chan string)
	wordsCounter := make(chan string)
	// words counting group 1
	counterAccumulator := make(chan int, WORKERS) // workers
	doneCounter := make(chan bool, WORKERS) // leader
	// words processing group 2
	processorAccumulator := make(chan string, WORKERS)
	doneProcessor := make(chan bool, WORKERS)
	processor := make(chan string, WORKERS)
	// words encoding group 3
	encoderAccumulator := make(chan string, WORKERS)
	doneEncoder := make(chan bool, WORKERS)
	marshaller := make(chan string, WORKERS)
	encoder := make(chan string, WORKERS)
	// words encoding group 4(supervised)
	encoderXMLAccumulator := make(chan string, WORKERS)
	doneEncoderXML := make(chan bool, WORKERS)
	marshallerXML := make(chan string, WORKERS)
	encoderXML := make(chan string, WORKERS)
	progressXML := make(chan string)

	// STAND ALONE WORKER
	// process words worker
	go func() {
		defer close(progress)
		fmt.Printf(" * Reader started")
		for word := range wordsGenerator(data) {
			progress<- word
		}
		fmt.Println()
		fmt.Println(" * Reader terminated")
	}()


	// STAND ALONE WORKER
	// progress monitor worker
	go func() {
		defer close(wordsCounter)
		fmt.Println(" * Progress sarted")
		i := 0
		for word := range progress {
			i++
			fmt.Printf(".")
			wordsCounter<- word
		}
		fmt.Println()
		fmt.Println(" * Progress termitated")
	}()

	// STAND ALONE WORKER
	// simple nested worker
	// counter worker
	// Distributor(router,switching station or so on...)
	go func() {
		fmt.Printf(" * Counter started")
		// child worker
		go func () {
			defer close(processorAccumulator)
			defer close(counterAccumulator)
			defer close(encoderAccumulator)
			defer close(progressXML)
			fmt.Println(" * Counter child")
			n := 0;
			for word := range wordsCounter {
				//fmt.Printf("{%s} ", word)
				n++
				// queue words, cycle for similar copy for every worker
			//	for i := 0; i < cap(processorAccumulator); i++ {
					processorAccumulator<- word
					encoderAccumulator<- word
					progressXML<- word
			//	}
			}
			// queue count result
			for marker := 0; marker < cap(counterAccumulator); marker++ {
				// markup for differ elements
				counterAccumulator<- n + marker
			}
			fmt.Println()
			fmt.Println(" * Counter child terminated")
		}()
		fmt.Println(" * Counter terminated")
	}()

	// ---------- GROUP-------------------
	// WORKERS POOL
	// group or workers for processing queue
	// they dequeue values randomly
	go func() {
		for i := 0; i < cap(counterAccumulator); i++ {
			// Accountant worker
			go func(id int) {
				fmt.Printf(" * Accountant worker %d started\n", id)
				for wordsCount := range counterAccumulator {
					// worker's task
					fmt.Printf(" > Total %d: %d\n", id, wordsCount)
				}
				doneCounter<- true
				fmt.Println()
				fmt.Printf(" * Accountant worker %d terminated\n", id)
			}(i) // create worker ID
		}
	}()
	// LEADER
	// group sycronizer(leader)
	go func() {
		fmt.Println(" * Group leader started")
		workersDoneCount := 0
		for range doneCounter {
			workersDoneCount++
			fmt.Printf(" > G1 workers %d done\n", workersDoneCount)
			if workersDoneCount >= WORKERS {
				// stop cycle and terminate group
				close(doneCounter)
				fmt.Println()
				fmt.Println(" ! Done counting words")
			}
		}
		// done goroup signale
		doneGroups<- true
		fmt.Println(" * Group leader terminated")
	}()
	// ----------END GROUP-------------------

	// ---------- GROUP-------------------
	// nested comlex structure
	// starter parent worker
	go func() {
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func() {
			for i := 0; i < cap(processorAccumulator); i++ {
				// Word processor worker
				go func(id int) {
					//defer close(processed)
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range processorAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						processor<- fmt.Sprintf("{%s}", word)
					}
					// done when queue is empty
					doneProcessor<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(i) // create worker ID
			}
		}()
		// SNGLE WORKER
		// writer worker
		go func() {
			fmt.Println(" * Writer started")
			fout, _ := os.Create("./output.txt")
			defer fout.Close()
			for processedWord := range processor {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "%s\n", processedWord)
			}
			fmt.Fprintf(fout, "\n")
			fmt.Println(" * Writer terminated")
		}()
		go func() {
			// LEADER
			// group sycronizer(lider)
			go func() {
				defer close(processor)
				fmt.Println(" * Group 2 leader started")
				workersDoneCount := 0
				for range doneProcessor {
					workersDoneCount++
					fmt.Printf(" > G2 workers %d done\n", workersDoneCount)
					if workersDoneCount >= WORKERS {
						// stop cycle and terminate group
						close(doneProcessor)
						// done goroup signale
						doneGroups<- true
						fmt.Println()
						fmt.Println(" ! Done processing words")
					}
				}
				fmt.Println(" * Group 2 leader terminated")
			}()
		}()
	}()
	// ----------END GROUP-------------------


	// nested comlex structure
	// starter parent worker
	go func() {
		// ---------- GROUP-------------------
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func() {
			for i := 0; i < cap(encoderAccumulator); i++ {
				// Word processor worker
				go func(id int) {
					//defer close(processed)
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range encoderAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						marshaller<- fmt.Sprintf("%s", word)
					}
					// done when queue is empty
					doneEncoder<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(i) // create worker ID
			}
		}()
		// SNGLE WORKER
		// marshaller worker
		go func() {
			fmt.Println(" * Marshaller JSON started")
			for marshalledWord := range marshaller {
					// worker's task
					encoder<- fmt.Sprintf("'%s': %s,\n", marshalledWord, marshalledWord)
			}
			fmt.Println(" * Marshaller JSON terminated")
		}()
		// SINGLE WORKER
		// writer worker
		go func() {
			fmt.Println(" * Writer JSON started")
			fout, _ := os.Create("./output.json")
			defer fout.Close()
			fmt.Fprintf(fout, "{\n")
			for encoderWord := range encoder {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "\t%s", encoderWord)
			}
			fmt.Fprintf(fout, "}\n")
			fmt.Println(" * Writer JSON terminated")
		}()
		go func() {
			// LEADER
			// group sycronizer(lider)
			go func() {
				defer close(marshaller)
				defer close(encoder)
				fmt.Println(" * Group 3 leader started")
				workersDoneCount := 0
				for range doneEncoder {
					workersDoneCount++
					fmt.Printf(" > G2 workers %d done\n", workersDoneCount)
					if workersDoneCount >= WORKERS {
						// stop cycle and terminate group
						close(doneEncoder)
						// done goroup signale
						doneGroups<- true
						fmt.Println()
						fmt.Println(" ! Done marshalling words")
					}
				}
				fmt.Println(" * Group 3 leader terminated")
			}()
		}()
		// ----------END GROUP-------------------
	}()


	// ------------------------
	// Subpervised subtree in a supervised tree
	//
	// STAND ALONE WORKER
	// progress monitor #2 worker	
	go func() {
		defer close(encoderXMLAccumulator)
		fmt.Println(" * Progress XML sarted")
		i := 0
		for word := range progressXML {
			i++
			fmt.Printf("%s","'")
			encoderXMLAccumulator<- word
		}
		fmt.Println()
		fmt.Println(" * Progress XML termitated")
	}()

	// ---------- GROUP-------------------
	// nested comlex structure
	// starter parent worker
	go func() {
		// WORKERS POOL
		// group or workers for processing queue
		// they dequeue values randomly
		go func() {
			for i := 0; i < cap(encoderXMLAccumulator); i++ {
				// Word processor worker
				go func(id int) {
					//defer close(processed)
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range encoderXMLAccumulator {
						// worker's task
						//fmt.Printf(" > processor %d: %s\n", id, word)
						marshallerXML<- fmt.Sprintf("%s", word)
					}
					// done when queue is empty
					doneEncoderXML<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(i) // create worker ID
			}
		}()
		// SINGLE WORKER
		// marshaller worker
		go func() {
			fmt.Println(" * Marshaller XML started")
			for marshalledWord := range marshallerXML {
					// worker's task
					encoderXML<- fmt.Sprintf("<record value=\"%s\" name=\"%s\"/>\n", marshalledWord, marshalledWord)
			}
			fmt.Println(" * Marshaller XML terminated")
		}()
		// SINGLE WORKER
		// writer worker
		go func() {
			fmt.Println(" * Writer XML started")
			fout, _ := os.Create("./output.xml")
			defer fout.Close()
			fmt.Fprintf(fout, "<table name=\"Records\">\n")
			for encoderWord := range encoderXML {
					//fmt.Println(" -> put: ", processedWord)
					fmt.Fprintf(fout, "\t%s", encoderWord)
			}
			fmt.Fprintf(fout, "</table>\n")
			fmt.Println(" * Writer XML terminated")
		}()
		// LIDER
		// group sycronizer(lider)
		go func() {
			defer close(marshallerXML)
			defer close(encoderXML)
			fmt.Println(" * Group 4 leader started")
			workersDoneCount := 0
			for range doneEncoderXML {
				workersDoneCount++
				fmt.Printf(" > G4 workers %d done\n", workersDoneCount)
				if workersDoneCount >= WORKERS {
					// stop cycle and terminate group
					close(doneEncoderXML)
					// done goroup signale
					doneSubGroups<- true
					fmt.Println()
					fmt.Println(" ! Done marshalling words")
				}
			}
			fmt.Println(" * Group 4 leader terminated")
		}()
	}()
	// ----------END GROUP-------------------

	// SUPERVISOR
	// subtree groups sycronizer(Supervisor)
	go func() {
		fmt.Println(" * Supervisor 2 started")
		groupsDoneCount := 0
		for range doneSubGroups {
			groupsDoneCount++
			fmt.Printf(" # group %d done\n", groupsDoneCount)
			if groupsDoneCount >= 1 { // for all groups
				// stop cycle and terminate all supervised groups tree
				close(doneSubGroups)
				fmt.Println()
				fmt.Println(" ! Done supervised subtree")
			}
			// signal everything done in sub group
			//TODO: defer func() { doneGroups<- true }()
			doneGroups<- true
		}
		fmt.Println(" * Supervisor 2 terminated")
	}()
	// ------------------END Supervised subtree-----------

	// ROOT SUPERVISOR
	// All groups sycronizer(Supervisor)
	go func() {
		// signal everything done
		defer close(done)
		fmt.Println(" * Supervisor 1 started")
		groupsDoneCount := 0
		for range doneGroups {
			groupsDoneCount++
			fmt.Printf(" # group %d done\n", groupsDoneCount)
			if groupsDoneCount >= 4 { // for all groups
				// stop cycle and terminate all supervised groups tree
				close(doneGroups)
				fmt.Println()
				fmt.Println(" ! Done supervised tree")
			}
		}
		fmt.Println(" * Supervisor 1 terminated")
	}()

	// controller-sycronizer
	select {
	case <-done:
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
  go func() {
    defer close(outChan)
    for _, line := range data {
      words := strings.Split(line, " ")
      for _, word := range words {
        outChan<- word
      }
    }
  }()
  return outChan
}

