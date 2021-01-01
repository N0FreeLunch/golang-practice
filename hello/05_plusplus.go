// this is error code
package main
import "fmt"

func post () {
  sum,i := 0, 0
  for i < 10 {
    sum += i++
  }
  fmt.Println(sum)
}

func prev () {
  sum, i := 0, 0
  for i < 10 {
    sum += i
    i++
  }
}

func main() {
  post()
  prev()
}
