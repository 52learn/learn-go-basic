package main
import "fmt"

func main(){
	for i:=0;i<4;i++{
		fmt.Println("GeeksforGeeks,i:",i)
	}
	
	 
	flag:=true
	i:=0;
	for flag{
		fmt.Println("GeeksforGeeks,i:",i)
		i++
		if(i==10){
			flag=false
		}
	}
	
	// Here rvariable is a array
	rvariable:=[]string{"lily","kim","tom"}
    // i and j stores the value of rvariable
    // i store index number of individual string and
    // j store individual string of the given array

	for i,j:=range rvariable{
		fmt.Println(i,j)
	}
	
	for i:=range rvariable{
		fmt.Println(i)
	}
	
	// String as a range in the for loop
	for index,chr:=range "hello"{
		fmt.Println(index,chr)
	}
	
	// using maps
	mmap:=map[int]string{
		22:"grey",
		33:"green",
	}
	
	for key,value:=range mmap{
		fmt.Println("mmap key:",key," value:",value)
	}
	
	
	// using channel
	chnl:=make(chan int)
	go func(){
		chnl <-100
		chnl <- 1000
		chnl<-10000
		chnl<-100000
		close(chnl)
	}()
	fmt.Println("channel content:")
	for i:=range chnl{
		fmt.Println(i)
	}
	
	// empty string
	for index,chr:=range ""{
		fmt.Println("empty string")
		fmt.Println(index,chr)
	}
}