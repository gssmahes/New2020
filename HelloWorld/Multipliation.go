
//Program to display Multiplication Table of any number
package main
import "fmt"
func main(){
    var n int
    fmt.Print("Enter any Integer Number : ")
    fmt.Scan(&n)
    i:=1
	/*     For loop as a Go's While     */

	for {
		if i > 10 {
			break
		}
		fmt.Println(i," X ",n, " = ", i*n)
		i++
	}
 //     using normal for loop
	for i:=1; i<11;i++ {
		fmt.Println(i," X ",n, " = ", i*n)
	}
}