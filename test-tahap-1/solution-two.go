package test_tahap_1

import "fmt"

func SolutionTwo() {
	var totalBelanja int
	var totalBayar int
	var kembalian int
	pecahanUang := []int{50000, 20000, 5000, 2000, 200, 100}

	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Pembeli membayar: Rp ")
	fmt.Scan(&totalBayar)

	kembalian = totalBayar - totalBelanja
	if kembalian < 0 {
		fmt.Println("False, kurang bayar")
		return
	}

	fmt.Println(fmt.Sprintf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d", kembalian, pembulatanUang(kembalian, 100)))

	fmt.Println("Pecahan uang: ")
	sisaUang := kembalian

	for _, pecahan := range pecahanUang {
		if sisaUang == 0 {
			break
		}
		hitungKembalianByPecahan(sisaUang, pecahan)
		sisaUang = sisaUang % pecahan
	}
}

func hitungKembalianByPecahan(sisaUang int, pecahan int) {
	jenisPecahan := "lembar"
	jumlah := jumlahPecahan(sisaUang, pecahan)

	if pecahan < 1000 {
		jenisPecahan = "koin"
	}

	if jumlah > 0 {
		fmt.Println(fmt.Sprintf("%d %s %d", jumlah, jenisPecahan, pecahan))
	}
}

func pembulatanUang(totalKembalian int, kelipatanPembulat int) int {
	sisa := totalKembalian % kelipatanPembulat
	if sisa != 0 {
		return totalKembalian - sisa
	}
	return totalKembalian
}

func jumlahPecahan(total int, pecahan int) int {
	return total / pecahan
}
