package main
import "fmt"

func main() {
  c:= 'a'
  switch {
  case '0' <= c && c <= '9':
    fmt.Print("%c은 숫자입니다.",c)
  case 'a' <= c && c <= 'z':
    fmt.Print("%c은 소문자입니다.",c)
  case 'A' <= c && c <= 'Z':
    fmt.Println("%c은 대문자입니다.",c)
  }
}
