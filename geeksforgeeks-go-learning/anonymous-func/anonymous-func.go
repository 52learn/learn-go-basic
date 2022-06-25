package main

import "fmt"

  
 // Passing anonymous function
 // as an argument 
func GFG(i func(p, q string)string){
 fmt.Println(i ("Geeks", "for"))
   
}


func main(){
	// Anonymous function
	func(){
		fmt.Println("Welcome! to GeeksforGeeks")
	}()
	
	
   // Assigning an anonymous 
   // function to a variable
   value := func(){
      fmt.Println("Welcome! to GeeksforGeeks")
  }
  value()
  
  
  // Passing arguments in anonymous function
  func(ele string){
      fmt.Println(ele)
  }("GeeksforGeeks")
  
  
    value2:= func(p, q string) string{
        return p + q + "Geeks"
    }
    GFG(value2)
	
	value3 := GFG3()
    fmt.Println(value3("Welcome ", "to "))
}


// Returning anonymous function 
 func GFG3() func(i, j string) string{
     myf := func(i, j string)string{
          return i + j + "GeeksforGeeks"
     }
    return myf
 }
    