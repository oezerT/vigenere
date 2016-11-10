package main

import "strings"

func encode(plaintext string, key string) string {
  var output []string

  for i, _ := range plaintext {
    shiftIndex := int(string(plaintext[i])[0] - 97)
    keyIndex := string(key[i%len(key)])[0] - 97

    shiftedAlphabet := createShiftedAlphabet(shiftIndex)
    output = append(output, shiftedAlphabet[keyIndex])

    if (i+1)%5 == 0{
      output = append(output, " ")
    }
  }
  return strings.Join(output, "")
}
