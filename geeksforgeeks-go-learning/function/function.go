
// Go program to illustrate the
// use of function
package main
import "fmt"

func area(length,width int)int{
	area:=length*width
	return area;
}

func swap(a,b int)int{
	var o int
	o=a
	a=b
	b=o
	return o
}

func swapByRef(a,b *int)int{
	var o int
	o=*a
	*a=*b
	*b=o
	return o
}

func main(){
	fmt.Printf("area of rectangle is :%d \n",area(3,4))
	
	
	var p int=10
	var q int=20
	fmt.Printf("p=%d ,q=%d \n",p,q)
	// call by values
	swap(p,q)
	fmt.Printf("p=%d ,q=%d \n",p,q)
	
	// call by reference
	swapByRef(&p,&q)
	fmt.Printf("p=%d ,q=%d \n",p,q)
	
}