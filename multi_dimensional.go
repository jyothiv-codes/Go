/*
multidimensional arrays with and without user defined input
*/

package main

import "fmt"

func main() {
   /* an array with 5 rows and 2 columns*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
   fmt.Println("User input defined: ")
   var multi [3][2] int
   for i=0;i<3;i++{
   	for j=0;j<2;j++{
   		fmt.Scanln(&multi[i][j])
   	}
   }
   for i=0;i<3;i++{
   	for j=0;j<2;j++{
   		fmt.Print(multi[i][j]," ")
   	}
   	fmt.Println()
   }
}