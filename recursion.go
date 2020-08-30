/*
learning recursion thru factorial and fibonacci

*/
package main
import "fmt"
func factorial(n int)int {
	if (n==1){
		return 1
	}
	return n*factorial(n-1)
}
func fibonacci(i int) int{
	 //fmt.Println(i)
	if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonacci(i-1) + fibonacci(i-2)
}
func main(){
	var n,i int
	fmt.Println("Enter a number (factorial): ")
	fmt.Scanln(&n)
	fmt.Println(factorial(n))
	fmt.Println("Enter a number (fibonacci): ")
	fmt.Scanln(&n)
	fmt.Println(fibonacci(n))
	for i = 0; i < 10; i++ {
      fmt.Printf("%d ", fibonacci(i))
   }
}