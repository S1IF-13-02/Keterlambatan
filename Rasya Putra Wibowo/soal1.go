package main
import "fmt"

func main() {
	var pengeluaran, lamaJoin int

	fmt.Scan(&pengeluaran, &lamaJoin)

	if pengeluaran >= 5000 && lamaJoin >= 24 {
		fmt.Println("VVIP Backstage")
	} else if pengeluaran >= 2000 && lamaJoin >=12 {
		fmt.Println("VIP Soundcheck")
	} else if pengeluaran >=500 {
		fmt.Println("Festival Ground")
	} else {
		fmt.Println("Menonton dari Youtobe")
	}

}