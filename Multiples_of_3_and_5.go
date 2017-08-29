package main

import "fmt"

func main(){
  var ( /* A way of defining Multiple Variables */
    sum = 0
    five = 5
    three = 3
  )
  /* the i := is also a way of automatically defining Variables */
  for i := 0; i < 1000; i++ {
    if i % three == 0{
      sum += i

    } else if i % five == 0 {
      sum += i
    }
  }
  fmt.Println(sum)
}
