//
// generator
//
package main

import(
	"fmt"
	"strings"
)

func main() {

 data := []string{
    "The yellow fish swims",
    "Blue is the water in a crystal glass",
    "Blue are the lilacs that grow in the grass",
    "Blue is delicious blueberry pie",
    "Blue are the sparkles in my cats eyes",
  }

  for word := range wordsGenerator(data) {
    fmt.Println(word)
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

