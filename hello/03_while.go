package main
import "fmt"

func main () {
  sum, i := 0, 0
  for {
    if i >=10 {
      break
    }
    sum += 1
    i++
  }
  fmt.Println(sum)
}
