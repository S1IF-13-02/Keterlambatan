package main
import "fmt"

func main() {
	var N int
	fmt.Print("Masukan Angka: ")
	fmt.Scan(&N)

	angka := N
	total := 0

	for angka > 0 {
		total += angka % 10
		angka = angka / 10
	}
	fmt.Println("Hasil Penjumlahan :", total)
}