package main

import "fmt"

func main()  {
  foo := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
  index := binarySearch(foo, 8)
  fmt.Println(index)
}

func binarySearch(bar []int, value int) int{
  // dota2's references sorry
  top := len(bar) -1
  bot := 0

  for {
    mid := (top + bot) / 2
    
    if value == bar[mid] {
      return mid
    } else if value < bar[mid] {
      top = mid - 1
    } else if value > bar[mid] {
      bot = mid + 1
    }
  }
}
