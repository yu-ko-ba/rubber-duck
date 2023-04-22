package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
  s := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("> ")
    s.Scan()
    fmt.Println("🦆 " + reply())
  }
}

func reply() (s string) {
  n := rand.Intn(3)
  switch n {
  case 0:
    s = "ガァ"
  case 1:
    s = "グァ"
  case 2:
    s = "グワァ"
  }
  return
}
