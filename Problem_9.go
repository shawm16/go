package main

import (
  "fmt"
  "math"
)

func main(){
  var(
    boolean = 1
  )
  var a float64 = 0
  var b float64 = 0
  for boolean != 0 {
    for a = 1; a < 1000; a++ {
      if a + b + math.Sqrt(a*a+b*b) == 1000{
        if (a*a) + (b*b) == ((math.Sqrt(a*a+b*b))*(math.Sqrt(a*a+b*b))) && a < b && b < math.Sqrt(a*a+b*b) {
          fmt.Println(a, "*", b, "*", math.Sqrt(a*a+b*b), "= 1000")
          boolean = 0
        }
      }
      for b = 1; b < 1000; b++ {
        if a + b + math.Sqrt(a*a+b*b) == 1000{
          if (a*a) + (b*b) == (math.Sqrt(a*a+b*b)*math.Sqrt(a*a+b*b)) && a < b && b < math.Sqrt(a*a+b*b) {
            fmt.Println(a, "*", b, "*", math.Sqrt(a*a+b*b), "= 1000")
            boolean = 0
          }
        }
      }
    }
  }

}
