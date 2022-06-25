// Go program to illustrate the
// concept of variadic function
package main
 
import(
    "fmt"
    "strings"
)

// Variadic function to join strings
func joinstr(element ... string)string{
	return strings.Join(element,"-")
}

func main(){
  // zero argument
   fmt.Println(joinstr())
    
   // multiple arguments
   fmt.Println(joinstr("GEEK", "GFG"))
   fmt.Println(joinstr("Geeks", "for", "Geeks"))
   fmt.Println(joinstr("G", "E", "E", "k", "S"))
   
   // pass a slice in variadic function
   element:=[]string{"hellow","world"}
   fmt.Println(joinstr(element...))
}