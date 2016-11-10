package main

import "strings"

func decode(code string, key string) string {
  var output []string

  for i, _ := range code {
    keyIndex := string(key[i%len(key)])[0] - 97
    shiftedAlphabet := createShiftedAlphabet(int(keyIndex))

    var plaintextIndex int

    for j, letter := range shiftedAlphabet{
      if letter == string(code[i]) {
        plaintextIndex = j
      }
    }

    alphabet := createShiftedAlphabet(0)
    output = append(output, alphabet[plaintextIndex])

    if (i+1)%5 == 0{
      output = append(output, " ")
    }
  }
  return strings.Join(output, "")
}
