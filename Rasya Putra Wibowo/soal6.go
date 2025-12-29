package main
import "fmt"

func main() {
	var kapasitas int
	fmt.Scan(&kapasitas)

	total := 0

	for {
		var masuk int
		if _, err :=fmt.Scan(&masuk); err != nil {
			break
		}

		if total+masuk <= kapasitas {
			total += masuk
			fmt.Printf("Masuk %d liter. Total: %d\n", masuk, total)
		} else {
			fmt.Printf("Luber! Total tetap %d. Mobil sudah penuh, tidak bisa tambah %d liter.\n", total, masuk,)
			break
		}
	}
}