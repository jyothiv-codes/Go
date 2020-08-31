package main
import "fmt"
//type amount func(int,int)int
func amount (a int,b int)int {
	return a*b
}
func main(){
	//map
	 menu := map[string]int{
		"idli":20,
		"burger":90,
		"pizza":300,
		"dosa":60,
		"tea":20,
		"coffee":40,
	}
	fmt.Println("Hi, welcome to Laddoo restaurant! What would you like to have???")
	for item, price := range menu { 
        fmt.Println(item, price) 
    } 
    //list 
    cricket_scores := [10]int{14,65,33,98,56,81,82,38,22,91}
    fmt.Println(cricket_scores)
    var list [5]int = [5]int{10, 20, 30}
    fmt.Println(list)

    //user defined functions
    var total int 
    total=amount(40,5)
    fmt.Printf("The total is: %d \n",total)
 

}

