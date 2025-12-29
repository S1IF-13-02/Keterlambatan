package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i<= N; i++ {
		for j := 1; j <= i; j++ {
			if j > 1 {
				fmt.Print(" ")
			}
			fmt.Print(i)
		}
		fmt.Print()
	}
}