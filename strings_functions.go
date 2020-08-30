/*
string functions
*/
/*
DOUBTS
1. D1
2. case sensitive string comparison


*/

package main
import ("fmt"
	str "strings")
func main(){
	var print=fmt.Println
	s:= "this is a string we are gonna use  "
	print(str.ToUpper(s))
	print(str.Index(s,"e"))
	print(str.Repeat(s,4))
	print(str.Split("abcde",""))
	print(str.Contains(s,"ti"))
	print(str.ContainsAny(s,"t"))
	print(str.Contains("hello",""))
	print(str.Count("test", ""))//D1
	print(str.Count("test", " "))
}