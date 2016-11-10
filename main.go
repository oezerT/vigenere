package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "unicode"
)

func main() {
  fmt.Println("de (d) or encode (e)?: ")
  var a string
  fmt.Scanln(&a)

  if strings.Compare(a, "d") == 0 {
    fmt.Println("enter code: ")
    code := readFromConsole()

    fmt.Println("enter key: ")
    key := readFromConsole()

    fmt.Println("plaintext: ", decode(code, key))
  } else {
    fmt.Println("enter plaintext: ")
    plaintext := readFromConsole()

    fmt.Println("enter key: ")
    key := readFromConsole()

    fmt.Println("code: ", encode(plaintext, key))
  }
}

func createShiftedAlphabet(shiftIndex int) (alphabet []string) {
  for i := 0; i < 26; i++ {
    alphabet = append(alphabet, string(i+97))
  }
  x, a := alphabet[:shiftIndex], alphabet[shiftIndex:]
  return append(a, x...)
}

func readFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println(err)
    return " "
  }
	input = strings.TrimRightFunc(input, unicode.IsSpace)
  input = strings.Replace(input, " ", "", -1)
  return strings.ToLower(input)
}
