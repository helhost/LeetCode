package main

import "fmt"

// -------- SOLVE HERE --------

func longestValidParentheses(s string) int {
  best := 0

  for i := range s {
    canClose := 0
    curr := 0
    currBest := 0

    for j := i; j < len(s); j++ {
      curr++

      if s[j] == '(' {
        canClose++
        continue
      }

      canClose--

      if canClose < 0 {
        break
      }

      if canClose == 0 {
        currBest = curr
      }

    }

    if currBest > best {
      best = currBest
    }
  }

  return best
}

func main() {
  s := longestValidParentheses("())(()))(((()(()))))))")
  fmt.Println(s)
}
