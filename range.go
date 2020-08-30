/*
range keyword
we might learn a little about arrays also here, more of a revision kind
*/
package main
import "fmt"
func main(){
	list := []int {1,2,3}
	fmt.Println(list)
	empty := []int{}
	fmt.Println(empty)
	for i:= range list {
		fmt.Println(i,list[i])
	}
	for i:= range list{
		fmt.Println(list[i],&list[i])
	}
}