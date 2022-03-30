package main

import (
	"fmt"
	"os"
  "strings"
)
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
  arg := os.Args[1:2]
  word := strings.Join(arg, "")
  if(word == Reverse(word)){
      fmt.Println(word + " is a palindrome!")
      return
    }else{
      fmt.Println(word + " is not a palindrome!")
		}
	
}
