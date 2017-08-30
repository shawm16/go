package main

/* This is a way to import more than one package */
import (
  "fmt"
  /*This is the math package that we need for math.Sqrt() */
  "math"
)

func main(){
  var(
    boolean = 1
  )
  /*The reason we define the variables as a float 64 is because math.Sqrt(FLOAT64) */
  var a float64 = 0
  var b float64 = 0
/*the boolean variable is just a true or false checker*/
  for boolean != 0 {
/*This is where the algorithm is put into play. Thanks Mar. */
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
