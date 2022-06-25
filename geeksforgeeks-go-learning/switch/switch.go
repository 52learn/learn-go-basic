package main
  
import "fmt"
  
func main() {
      
    // Switch statement with both 
    // optional statement, i.e, day:=4
    // and expression, i.e, day
    switch day:=4; day{
       case 1:
       fmt.Println("Monday")
       case 2:
       fmt.Println("Tuesday")
       case 3:
       fmt.Println("Wednesday")
       case 4:
       fmt.Println("Thursday")
       case 5:
       fmt.Println("Friday")
       case 6:
       fmt.Println("Saturday")
       case 7:
       fmt.Println("Sunday")
       default: 
       fmt.Println("Invalid")
   }
   var value int = 2
      
    // Switch statement without an     
    // optional statement and 
    // expression
   switch {
       case value == 1:
       fmt.Println("Hello")
       case value == 2:
       fmt.Println("Bonjour")
       case value == 3:
       fmt.Println("Namstay")
       default: 
       fmt.Println("Invalid")
   }
   var value2 string = "five"
  
    // Switch statement without default statement
    // Multiple values in case statement
   switch value2 {
       case "one":
       fmt.Println("C#")
       case "two", "three":
       fmt.Println("Go")
       case "four", "five", "six":
       fmt.Println("Java")
   }  

// concept of Type switch   
    var value3 interface{}
    switch q:= value3.(type) {
       case bool:
       fmt.Println("value3 is of boolean type")
       case float64:
       fmt.Println("value3 is of float64 type")
       case int:
       fmt.Println("value3 is of int type")
       default:
       fmt.Printf("value3 is of type: %T", q)
         
   }
}