package main

import "fmt"

func main(){
  var (
    sum = 0
    five = 5
    three = 3
  )

  for i := 0; i < 1000; i++ {
    if i % three == 0{
      sum += i

    } else if i % five == 0 {
      sum += i
    }
  }
  fmt.Println(sum)
}
