/*
package main
import "fmt"
func main(){
	fmt.Println("hello world!")
}*/
//常量

/*const(
	pi=3.1415
	e=2.7182
)
*/
package main
import "fmt"

/*const(
	n1=10
	n2
	n3
)*/
const(
	n1 = iota//0
	n2
	n3=100
	n4)
const n5=iota//0
/*const(
	_=iota
	kb=1<<(10*iota)//1<<10向左移一位相当于 * 2,2^10 = 1024
	ab=1<<(10*iota)//1<<20,2^20
	cb=1<<(10*iota)//1<<30,2^30
	db=1<<(10*iota)//1<<40,2^40
	fb=1<<(10*iota)//1<<50,2^50 
	
)*/
const(
	a,b=iota+1,iota+2//iota=0,a=1,b=2
	c,d //iota =1,a=2,b=3
	e,f //iota=2,a=3,b=4
)
func main()  {
fmt.Println(a,b,c,d,f)	
}

 