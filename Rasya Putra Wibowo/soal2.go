package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	totalYes := 0

	for i := 0; i < N; i++ {
		var skor int 
		fmt.Scan(&skor)

		if skor > 75 {
			fmt.Println("Juri bilang YES")
			totalYes++
		} else {
			fmt.Println("Juri bilang NO")
		}
	}
	fmt.Println("Total YES:", totalYes)
}