package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var B[NMAX]int
	var C[NMAX]int
	var n, i, j, j1, jum, c, x int
	j = 0
	j1 = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		j = (10*j)+A[i]
	}
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&B[i])
		j1 = (10*j1)+B[i]
	}
	jum = j+j1
	n = jum
	for jum > 0{
		jum = jum/10
		c++
	}
	for i = 0; i < c; i++{
		x = n%10
		C[i] = x
		n = n/10
	}
	for i = 0; i < c; i++{
		fmt.Print(C[i], " ")
	}
	fmt.Println()
}