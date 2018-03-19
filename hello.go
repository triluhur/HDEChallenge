package main
import "fmt"

var tamp []int
var sum, y, flag int

func array (tamp []int, n int, p int){ 
	  y = 1
	  if n == 0 {
		 return
	  }	   
	  fmt.Scan(&tamp[flag])
	  if (tamp[flag] >= 0){
		  sum = sum + power(tamp[flag], p+1, p)	
	  }
      flag += 1
	  array(tamp, n-1, p)
}

func power(tamp int,v int, p int) int{ 
	 if(v > 1){
		y = y * tamp
		power(tamp, v-1, p)
	 } 
	 if p == 1{
		y = 1
	 } 
	 return y  
}

func main() {
	 var p, n int
	 _, err := fmt.Scan(&p)
	 if err != nil {
       panic(err)  }
	 fmt.Scan(&n)
	 tamp = make([]int, n)
	 array(tamp, n, p)
	 fmt.Println("output :", sum)	  
}