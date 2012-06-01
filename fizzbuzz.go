package main 

import "fmt"
import "strconv"

func fizzbuzz(x int) string {
  if x % 3 == 0 && x % 5 == 0 {
    return "FizzBuzz"
  } else if x % 3 == 0 && x % 5 != 0 {
    return "Fizz"
  } else if x % 5 == 0 && x % 3 != 0 {
    return "Buzz"
  } 
    return strconv.FormatInt(int64(x), 10)
}

func main() { 
  for i := 0; i <= 100; i++ {
    fmt.Println(fizzbuzz(i)) 
  }
} 
