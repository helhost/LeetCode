package main

import "fmt"
import "math"

// -------- SOLVE HERE --------

func nextPermutation(nums []int) {
    n := len(nums)
    if n < 2 { return }

    // find pivot: first index from right where nums[i] < nums[i+1]
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }

    // if pivot exists, find rightmost element > nums[i] and swap
    if i >= 0 {
        j := n - 1
        for j > i && nums[j] <= nums[i] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }

    // reverse the suffix to get the smallest tail
    for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
        nums[l], nums[r] = nums[r], nums[l]
    }
}


func main() {
  // add test case here
  s := []int{1,2,4,5,3}
  nextPermutation(s)
  fmt.Println(s)
}
