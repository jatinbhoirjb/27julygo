package mystring
import "fmt"
import "strings"
func Low(s string){
s=strings.ToLower(s)
fmt.Println(s)
}

func High(s string) {
s = strings.ToUpper(s)
fmt.Println(s)
}