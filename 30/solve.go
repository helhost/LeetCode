package main

import "fmt"

// -------- SOLVE HERE --------

func findSubstring(s string, words []string) []int {
  // Solve with fixed window chunked frequency

  if s == "" || len(words) == 0 { // early stopping
    return []int{}
  }

  L := len(words[0]) // L : length of words (all the same)
  K := len(words) // number of words
  W := L * K // window-size
  N := len(s)

  if W > N { // early stopping
    return []int{}
  }

  var result []int

  need := make(map[string]int) // a map to store how many of each word we need

  for _, w := range words {
    need[w] += 1
  }

  for i := 0; i <= N - W; i++ { // loop over every possible starting index

    seen := make(map[string]int) // a map to store how many of word we have seen
    ok := true

    for j := 0; j < K; j++ {

      start := i + j*L
      w := s[start:start + L] // slice for a word

      if need[w] == 0 { // if we don't need the word, the starting index is not correct
          ok = false
          break
      }

      seen[w] += 1

      if seen[w] > need[w] { // we have seen the word too many times
        ok = false
        break
      }
    }

    if ok && mapsEqual(need,seen) { // did not see any incorrect words, and we saw all the words needed
      result = append(result, i)
    }
  }
  return result
}

// Helper
func mapsEqual(a,b map[string]int) bool {
  if len(a) != len(b) {
    return false
  }

  for k := range a {
    if b[k] != a[k] {
      return false
    }
  }

  return true
}

func main() {
  // add test case here
  s := findSubstring("barfoothefoobarman", []string{"foo","bar"})
  fmt.Println(s)
}
