package main
  
import (
    "encoding/base64"
    "fmt"
)
  
func main() {
  
    StringToEncode := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  
    Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
    fmt.Println(Encoding)                                        
}