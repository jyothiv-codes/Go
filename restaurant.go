package main
import "fmt"
func main(){
	fmt.Println("Hey!")
	menu := map[string]int {
		"idli":20,
		"burger":90,
		"pizza":300,
		"dosa":60,
		"tea":20,
		"coffee":40,
	}
	menu["idli"]=12
	fmt.Println(menu)
	fmt.Println("Here is the current menu : ")
	for item, price := range menu {
		fmt.Println(item,price)
	}
	var choice int
	var add string 
			var price_add int
	fmt.Println("Enter a number based on the following: \n 1. Modify the price \n 2. Delete an item \n 3. Add a new item \n 4. Billing\n\n")
	fmt.Scanln(&choice)
	switch choice{
	case 1: var name string
			var price int 
			fmt.Println("Enter the item name: ")
			fmt.Scanln(&name)
			fmt.Println("Enter the price to be updated: ")
			fmt.Scanln(&price)
			menu[name]=price 
	case 2: var deleted string
			fmt.Println("Enter the item that you wish to delete: ")
			fmt.Scanln(&deleted)
			delete(menu,deleted)
	case 3: 
			fmt.Println("Enter the item that you wish to add: ")
			fmt.Scanln(&add)
			fmt.Println("Enter the price of the item: ")
			fmt.Scanln(&price_add)
			menu[add]=price_add

	case 4: var table int
			fmt.Println("Enter the table number: ")
			fmt.Scanln(&table)
			//fmt.Println("Enter the items ordered: ")



}//switch case closes here


for item, price := range menu {
	fmt.Println(item,price)
}
}