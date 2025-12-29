package main
import "fmt"

func main() {
	var pilihan, jumlah int

	namaProduk := []string{
		"Little Trees",
		"Lap Microfiber",
		"Cover Steer",
		"Sponge Cuci Mobil",
	}

	hargaProduk := []int{
		35000,
		25000,
		150000,
		10000,
	}
	
	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	for i := 0; i < len(namaProduk); i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, namaProduk[i], hargaProduk[i])
	}

	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukan jumlah beli: ")
	fmt.Scan(&jumlah)

	index := pilihan - 1
	total := hargaProduk[index] * jumlah 

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", namaProduk[index])
	fmt.Println("Harga : Rp", hargaProduk[index])
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total : Rp", total )
}	