package main

import "fmt"

const mobilMAX int = 1024
const pabrikanMAX int = 1024

type Pabrikan struct {
	Nama         string
	JumlahProduk int
	Penjualan    int
}

type Mobil struct {
	TahunKeluar int
	Nama        string
	Pabrikan    string
	Penjualan   int
}

type listPabrikans [pabrikanMAX]Pabrikan
type listMobils [mobilMAX]Mobil

func main() {
	var LP listPabrikans
	var LM listMobils
	var n int

	prompt(&LP, &LM, &n)
}

/* Template */
func welcome() {
	fmt.Println("============ Selamat Datang ==============")
	fmt.Println("          Aplikasi Dealer Mobil           ")
	fmt.Println("          Algoritma Pemrograman           ")
	fmt.Println("------------------------------------------")
	fmt.Println()
}

/* Menjalankan perintah */
func prompt(LP *listPabrikans, LM *listMobils, n *int) {
	var input int

	for input != 13 {
		welcome()
		// Pabrikan
		fmt.Println("------ Pabrikan ------")
		fmt.Println("1. Lihat Daftar Pabrikan")
		fmt.Println("2. Lihat Penjualan Tertinggi")
		fmt.Println("3. Tambahkan Pabrikan")
		fmt.Println("4. Edit Pabrikan")
		fmt.Println("5. Hapus Pabrikan")
		fmt.Println("6. Cari Pabrikan")

		// Mobil
		fmt.Println("\n------ Mobil ------")
		fmt.Println("7. Lihat Daftar Mobil")
		fmt.Println("8. Lihat Penjualan Tertinggi")
		fmt.Println("9. Tambahkan Mobil")
		fmt.Println("10. Edit Mobil")
		fmt.Println("11. Hapus Mobil")
		fmt.Println("12. Cari Mobil")

		// Keluar Aplikasi
		fmt.Println("\n------ Lainnya ------")
		fmt.Println("13. Keluar")

		// Meminta Input
		fmt.Println()
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scan(&input)
		fmt.Println()
		fmt.Println()
		switch input {
		case 1:
			daftarPabrikan(LP, n)
		case 2:
			penjualanTertinggiPabrikan(LP, *n)
		case 3:
			tambahkanPabrikan(LP, LM, n)
		case 4:
			editPabrikan(LP, n)
		case 5:
			hapusPabrikan(LP, LM, n)
		case 6:
			cariPabrikan(LP, n)
		case 7:
			daftarMobil(LM, n)
		case 8:
			penjualanTertinggiMobil(LM, *n)
		case 9:
			tambahkanMobil(LM, LP, n)
		case 10:
			editMobil(LM, n)
		case 11:
			hapusMobil(LM, n)
		case 12:
			cariMobil(LM, n)
		case 13:
			fmt.Println("Terimakasih :>")
		}
	}
}

/* Sequential Search pabrikan */
func sequentialSearchPabrikan(LP *listPabrikans, n int, pabrikan string) int {
	for i := 0; i < n; i++ {
		if (*LP)[i].Nama == pabrikan {
			return i
		}
	}
	return -1
}

/* Binary Search pabrikan */
func binarySearchPabrikan(LP listPabrikans, n int, nama string) int {
	/* Mengembalikan indeks pabrikan jika ditemukan, Mengembalikan -1 jika pabrikan tidak ditemukan */
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if LP[mid].Nama == nama {
			return mid
		} else if LP[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

/* Insertion Sort Asc nama pabrikan */
func insertionAscendingNamaPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].Nama > LP[j].Nama {
			// Swap pabrikan jika nama sebelumnya lebih besar dari nama saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Desc nama pabrikan */
func insertionDescendingNamaPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].Nama < LP[j].Nama {
			// Swap pabrikan jika nama sebelumnya lebih kecil dari nama saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Asc jumlah pabrikan */
func insertionAscendingJumlahPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].JumlahProduk > LP[j].JumlahProduk {
			// Swap pabrikan jika Jumlah Produk sebelumnya lebih besar dari Jumlah Produk saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Desc jumlah pabrikan */
func insertionDescendingJumlahPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].Nama < LP[j].Nama {
			// Swap pabrikan jika Jumlah Produk sebelumnya lebih lebih dari Jumlah Produk saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Asc penjualan pabrikan */
func insertionAscendingPenjualanPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].Penjualan > LP[j].Penjualan {
			// Swap pabrikan jika Penjualan sebelumnya lebih besar dari Penjualan saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Desc penjualan pabrikan */
func insertionDescendingPenjualanPabrikan(LP *listPabrikans, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LP[j-1].Penjualan < LP[j].Penjualan {
			// Swap pabrikan jika Penjualan sebelumnya lebih kecil dari Jumlah Produk saat ini
			temp := LP[j]
			LP[j] = LP[j-1]
			LP[j-1] = temp
			j--
		}
	}
}

/* Binary Search mobil */
func binarySearchMobil(LM listMobils, n int, nama string) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if LM[mid].Nama == nama {
			return mid // Mengembalikan indeks pabrikan jika ditemukan
		} else if LM[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Mengembalikan -1 jika pabrikan tidak ditemukan
}

/* Insertion Sort Asc nama mobil */
func insertionAscendingNamaMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan nama secara menurun
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LM[j-1].Nama > LM[j].Nama {
			// Swap mobil jika nama sebelumnya lebih besar dari nama saat ini
			temp := LM[j]
			LM[j] = LM[j-1]
			LM[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Desc nama mobil */
func insertionDescendingNamaMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan nama secara menurun
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LM[j-1].Nama < LM[j].Nama {
			// Swap mobil jika nama sebelumnya lebih besar dari nama saat ini
			temp := LM[j]
			LM[j] = LM[j-1]
			LM[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Asc tahun mobil */
func insertionAscendingTahunMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan penjualan secara menaik
	for i := 1; i < n; i++ {
		idx := (*LM)[i]
		j := i - 1
		for j >= 0 && (*LM)[j].TahunKeluar > idx.TahunKeluar {
			(*LM)[j+1] = (*LM)[j]
			j--
		}
		(*LM)[j+1] = idx
	}
}

/* Insertion Sort Desc tahun mobil */
func insertionDescendingTahunMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan penjualan secara menurun
	for i := 1; i < n; i++ {
		idx := (*LM)[i]
		j := i - 1
		for j >= 0 && (*LM)[j].TahunKeluar < idx.TahunKeluar {
			(*LM)[j+1] = (*LM)[j]
			j--
		}
		(*LM)[j+1] = idx
	}
}

/* Insertion Sort Asc penjualan mobil */
func insertionAscendingPenjualanMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan penjualan secara menaik
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LM[j-1].Penjualan > LM[j].Penjualan {
			// Swap mobil jika nama sebelumnya lebih besar dari nama saat ini
			temp := LM[j]
			LM[j] = LM[j-1]
			LM[j-1] = temp
			j--
		}
	}
}

/* Insertion Sort Desc penjualan mobil */
func insertionDescendingPenjualanMobil(LM *listMobils, n int) {
	// Fungsi insertionSort untuk mengurutkan mobil berdasarkan nama secara menurun
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && LM[j-1].Penjualan < LM[j].Penjualan {
			// Swap mobil jika nama sebelumnya lebih besar dari nama saat ini
			temp := LM[j]
			LM[j] = LM[j-1]
			LM[j-1] = temp
			j--
		}
	}
}

/* Menambahkan pabrikan baru */
func tambahkanPabrikan(LP *listPabrikans, LM *listMobils, n *int) {
	var temp string
	var pabrikan Pabrikan

	if *n == pabrikanMAX {
		fmt.Println("Data telah penuh")
		return
	}

	fmt.Println("------ Tambahkan Pabrikan ------")

	fmt.Print("Nama: ")
	fmt.Scan(&pabrikan.Nama)

	// Periksa keberadaan pabrikan dengan nama yang sama
	idx := sequentialSearchPabrikan(LP, *n, pabrikan.Nama)
	if idx != -1 {
		fmt.Println("Pabrikan dengan nama tersebut sudah ada. Masukkan nama yang berbeda.")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	pabrikan.JumlahProduk = 0
	pabrikan.Penjualan = 0

	LP[*n] = pabrikan
	*n++

	fmt.Println("Pabrikan berhasil ditambahkan!")
	fmt.Println()

	// Tambahkan mobil dengan pabrikan yang sama ke daftar mobil
	fmt.Println("------ Tambahkan Mobil ------")

	var mobil Mobil
	fmt.Print("Nama Mobil: ")
	fmt.Scan(&mobil.Nama)
	fmt.Print("Tahun Keluar: ")
	fmt.Scan(&mobil.TahunKeluar)
	mobil.Pabrikan = pabrikan.Nama
	fmt.Print("Penjualan: ")
	fmt.Scan(&mobil.Penjualan)

	(*LM)[*n-1] = mobil

	// Update JumlahProduk and Penjualan of the corresponding pabrikan
	LP[*n-1].JumlahProduk++
	LP[*n-1].Penjualan += mobil.Penjualan

	fmt.Println("Mobil berhasil ditambahkan!")
	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menampilkan daftar pabrikan */
func daftarPabrikan(LP *listPabrikans, n *int) {
	var temp string
	fmt.Println("----- Daftar Pabrikan ------")

	if *n > 0 {
		var order string
		var sortBy string

		// Choose the order (ascending or descending)
		fmt.Print("Urutan (asc/desc): ")
		fmt.Scan(&order)
		fmt.Println()

		// Choose the sorting criteria
		fmt.Print("Sortir berdasarkan (nama/produk/penjualan): ")
		fmt.Scan(&sortBy)
		fmt.Println()

		switch sortBy {
		case "nama":
			if order == "asc" {
				insertionAscendingNamaPabrikan(LP, *n)
			} else {
				insertionDescendingNamaPabrikan(LP, *n)
			}
		case "produk":
			if order == "asc" {
				insertionAscendingJumlahPabrikan(LP, *n)
			} else {
				insertionDescendingJumlahPabrikan(LP, *n)
			}
		case "penjualan":
			if order == "asc" {
				insertionAscendingPenjualanPabrikan(LP, *n)
			} else {
				insertionDescendingPenjualanPabrikan(LP, *n)
			}
		default:
			fmt.Println("Invalid sorting criteria!")
			return
		}

		for i := 0; i < *n; i++ {
			pabrikan := (*LP)[i]
			fmt.Printf("%d. Nama: %s\tJumlah Produk: %d\tPenjualan: %d\n", i+1, pabrikan.Nama, pabrikan.JumlahProduk, pabrikan.Penjualan)
		}
	} else {
		fmt.Println("Daftar Pabrikan Kosong")
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menampilkan pabrikan dengan penjualan tertinggi */
func penjualanTertinggiPabrikan(LP *listPabrikans, n int) {
	var temp string

	fmt.Println("----- Penjualan Tertinggi Pabrikan -----")
	if n == 0 {
		fmt.Println("Daftar Pabrikan Kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	// Insertion sort Pabrikan
	if n > 0 {
		var order string

		// Pilih sortir by (ascending atau descending)
		fmt.Print("Urutan (asc/desc): ")
		fmt.Scan(&order)
		fmt.Println()

		if order == "asc" {
			insertionAscendingPenjualanPabrikan(LP, n)
		} else {
			insertionDescendingPenjualanPabrikan(LP, n)
		}
	}

	// Menampilkan 3 pabrikan dengan penjualan tertinggi
	maxPabrikan := 3
	if n < maxPabrikan {
		maxPabrikan = n
	}

	for i := 0; i < maxPabrikan; i++ {
		pabrikan := (*LP)[i]
		fmt.Printf("%d. Nama: %s\tJumlah Produk: %d\tPenjualan: %d\n", i+1, pabrikan.Nama, pabrikan.JumlahProduk, pabrikan.Penjualan)
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Mengubah data pabrikan */
func editPabrikan(LP *listPabrikans, n *int) {
	var nama, temp string
	var pabrikan Pabrikan

	fmt.Println("------ Edit Pabrikan ------")
	fmt.Print("Masukkan Nama Pabrikan: ")
	fmt.Scan(&nama)
	if *n == 0 {
		fmt.Println("Daftar pabrikan kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	insertionAscendingNamaPabrikan(LP, *n)
	idx := binarySearchPabrikan(*LP, *n, nama)

	if idx != -1 {
		pabrikan = LP[idx]
		fmt.Printf("\nNama: %s\nJumlah Produk: %d\nPenjualan: %d\n\n", pabrikan.Nama, pabrikan.JumlahProduk, pabrikan.Penjualan)
		fmt.Println("Lakukan perubahan pabrikan")
		fmt.Print("Nama Baru: ")
		fmt.Scan(&nama)
		fmt.Print("Jumlah Produk Baru: ")
		fmt.Scan(&LP[idx].JumlahProduk)
		fmt.Print("Penjualan Baru: ")
		fmt.Scan(&LP[idx].Penjualan)
		pabrikan.Nama = nama
		LP[idx] = pabrikan
	} else {
		fmt.Println("Data pabrikan tidak ditemukan")
	}

	fmt.Println("Pabrikan berhasil diubah!")
	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menghapus pabrikan */
func hapusPabrikan(LP *listPabrikans, LM *listMobils, n *int) {
	var temp string
	var index int

	fmt.Println("------ Hapus Pabrikan ------")
	if *n == 0 {
		fmt.Println("Daftar pabrikan kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	for i := 0; i < *n; i++ {
		pabrikan := LP[i]
		fmt.Printf("%d. Nama: %s\tJumlah Produk: %d\tPenjualan: %d\n", i+1, pabrikan.Nama, pabrikan.JumlahProduk, pabrikan.Penjualan)
	}

	loop := false

	for !loop {
		fmt.Print("Pilih nomor pabrikan yang ingin dihapus: ")
		fmt.Scan(&index)
		if index < 0 || index >= *n {
			index--
			fmt.Println("Nomor pabrikan tidak valid")
			fmt.Println("Masukkan apapun untuk mencoba lagi, atau tekan 'x' untuk kembali ke menu")
			fmt.Scan(&temp)
			if temp == "x" || temp == "X" {
				return
			}
		} else {
			loop = true
		}
	}

	// Update JumlahProduk and Penjualan of mobils associated with the pabrikan
	pabrikan := LP[index]
	for i := 0; i < *n; i++ {
		mobil := &LM[i]
		if mobil.Pabrikan == pabrikan.Nama {
			pabrikan.JumlahProduk--
			pabrikan.Penjualan -= mobil.Penjualan
		}
	}

	// Menghapus array pabrikan dari list pabrikan
	for i := index; i < *n-1; i++ {
		LP[i] = LP[i+1]
	}

	emptyPabrikan := Pabrikan{}
	LP[*n-1] = emptyPabrikan
	*n--

	fmt.Println("Pabrikan berhasil dihapus")
	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}


/* Mencari pabrikan */
func cariPabrikan(LP *listPabrikans, n *int) {
	var temp string
	var pabrikan Pabrikan
	var namaPabrikan string

	fmt.Println("-------- Cari Pabrikan --------")
	fmt.Print("Nama Pabrikan: ")
	fmt.Scan(&namaPabrikan)

	insertionAscendingNamaPabrikan(LP, *n)
	idx := sequentialSearchPabrikan(LP, *n, namaPabrikan)

	if idx != -1 {
		pabrikan = LP[idx]
		if pabrikan.Nama == namaPabrikan {
			fmt.Printf("\nNama: %s\nJumlah Produk: %d\nPenjualan: %d\n\n", pabrikan.Nama, pabrikan.JumlahProduk, pabrikan.Penjualan)
		}
	} else {
		fmt.Println("Data pabrikan tidak ditemukan!")
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menambahkan mobil baru */
func tambahkanMobil(LM *listMobils, LP *listPabrikans, n *int) {
	var temp string
	var mobil Mobil

	if *n == mobilMAX {
		fmt.Println("Data telah penuh")
	} else {
		fmt.Println("------ Tambahkan Mobil ------")

		fmt.Print("Nama: ")
		fmt.Scan(&mobil.Nama)
		fmt.Print("Tahun Keluar: ")
		fmt.Scan(&mobil.TahunKeluar)
		fmt.Print("Pabrikan: ")
		fmt.Scan(&mobil.Pabrikan)
		fmt.Print("Penjualan: ")
		fmt.Scan(&mobil.Penjualan)

		// Periksa apakah pabrikan sudah ada di daftar pabrikan menggunakan sequential search
		idx := sequentialSearchPabrikan(LP, *n, mobil.Pabrikan)

		if idx != -1 {
			LM[*n] = mobil
			*n++
		
			// Update jumlah produk and penjualan for the pabrikan
			(*LP)[idx].JumlahProduk++
			(*LP)[idx].Penjualan += mobil.Penjualan
		
			fmt.Println("Mobil berhasil ditambahkan!")
		} else {
			fmt.Println("Pabrikan tidak ditemukan! Silahkan tambahkan daftar pabrikan")
		} 		

		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
	}
}

/* Menampilkan daftar mobil berdsarkan kriteria */
func daftarMobil(LM *listMobils, n *int) {
	var temp string
	fmt.Println("----- Daftar Mobil ------")

	if *n > 0 {
		var order string
		var sortBy string

		// Choose the order (ascending or descending)
		fmt.Print("Urutan (asc/desc): ")
		fmt.Scan(&order)
		fmt.Println()

		// Choose the sorting criteria
		fmt.Print("Sortir berdasarkan (nama/tahun/penjualan): ")
		fmt.Scan(&sortBy)
		fmt.Println()

		switch sortBy {
		case "nama":
			if order == "asc" {
				insertionAscendingNamaMobil(LM, *n)
			} else {
				insertionDescendingNamaMobil(LM, *n)
			}
		case "tahun":
			if order == "asc" {
				insertionAscendingTahunMobil(LM, *n)
			} else {
				insertionDescendingTahunMobil(LM, *n)
			}
		case "penjualan":
			if order == "asc" {
				insertionAscendingPenjualanMobil(LM, *n)
			} else {
				insertionDescendingPenjualanMobil(LM, *n)
			}
		default:
			fmt.Println("Invalid sorting criteria!")
			return
		}

		for i := 0; i < *n; i++ {
			mobil := (*LM)[i]
			fmt.Printf("%d. Nama: %s\tTahun Keluar: %d\tPabrikan: %s\tPenjualan: %d\n", i+1, mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
		}
	} else {
		fmt.Println("Daftar Mobil Kosong")
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menampilkan 3 penjualan tertinggi mobil */
func penjualanTertinggiMobil(LM *listMobils, n int) {
	var temp string

	fmt.Println("----- Penjualan Tertinggi Mobil -----")
	if n == 0 {
		fmt.Println("Daftar Mobil Kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	// Insertion sort Mobil
	if n > 0 {
		var order string

		// Pilih sortir by (ascending atau descending)
		fmt.Print("Urutan (asc/desc): ")
		fmt.Scan(&order)
		fmt.Println()

		if order == "asc" {
			insertionAscendingPenjualanMobil(LM, n)
		} else {
			insertionDescendingPenjualanMobil(LM, n)
		}
	}

	// Menampilkan 3 pabrikan dengan penjualan tertinggi
	maxMobils := 3
	if n < maxMobils {
		maxMobils = n
	}

	for i := 0; i < maxMobils; i++ {
		mobil := LM[i]
		fmt.Printf("%d. Nama: %s\tTahun Keluar: %d\tPabrikan: %s\tPenjualan: %d\t\n", i+1, mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Mengedit mobil */
func editMobil(LM *listMobils, n *int) {
	var nama, temp string
	var mobil Mobil

	fmt.Println("------ Edit Mobil ------")
	fmt.Print("Masukan Nama Mobil: ")
	fmt.Scan(&nama)
	if *n == 0 {
		fmt.Println("Daftar mobil kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	insertionDescendingNamaMobil(LM, *n)
	idx := binarySearchMobil(*LM, *n, nama)

	if idx != -1 {
		mobil = LM[idx]
		fmt.Printf("Nama: %s\nTahun Keluar: %d\nPabrikan: %s\nPenjualan: %d\n\n", mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
		fmt.Println("Masukan Perubahan Mobil")
		fmt.Print("Nama Baru: ")
		fmt.Scan(&mobil.Nama)
		fmt.Print("Tahun Keluar Baru: ")
		fmt.Scan(&mobil.TahunKeluar)
		fmt.Print("Pabrikan Baru: ")
		fmt.Scan(&mobil.Pabrikan)
		fmt.Print("Penjualan Baru: ")
		fmt.Scan(&mobil.Penjualan)
		LM[idx] = mobil
		fmt.Println("Mobil berhasil diubah")
	} else {
		fmt.Println("Data Mobil Tidak Ditemukan")
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Menghapus mobil */
func hapusMobil(LM *listMobils, n *int) {
	var temp string
	var index int

	fmt.Println("------ Hapus Mobil ------")

	if *n == 0 {
		fmt.Println("Daftar mobil masih kosong")
		fmt.Println()
		fmt.Print("Masukkan apapun untuk melihat menu: ")
		fmt.Scan(&temp)
		fmt.Println()
		fmt.Println()
		return
	}

	for i := 0; i < *n; i++ {
		mobil := LM[i]
		fmt.Printf("%d. Nama: %s\tTahun Keluar: %d\tPabrikan: %s\tPenjualan: %d\t\n", i+1, mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
	}

	loop := false

	for !loop {
		fmt.Print("Pilih nomor mobil yang ingin di hapus: ")
		fmt.Scan(&index)
		if index < 0 || index >= *n {
			fmt.Println("Nomor mobil tidak valid!")
			fmt.Println("Masukan apapun untuk mencoba  lagi, atau tekan 'x' untuk kembali ke menu")
			fmt.Scan(&temp)
			if temp == "x" || temp == "X" {
				return
			}
		} else {
			loop = true
		}
	}

	// Memindahkan elemen-elemen setelah indeks yang akan dihapus ke posisi sebelum
	for i := index; i < *n-1; i++ {
		LM[i] = LM[i+1]
	}

	// Mengosongkan elemen terakhir pada array
	emptyMobil := Mobil{}
	LM[*n-1] = emptyMobil

	*n--

	fmt.Println("Mobil berhasil dihapus")
	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}

/* Mencari mobil */
func cariMobil(LM *listMobils, n *int) {
	var temp, nama string
	var mobil Mobil
	var namaMobil, namaPabrikan string

	fmt.Println("------ Cari Mobil ------")
	fmt.Print("Nama Pabrikan: ")
	fmt.Scan(&namaPabrikan)

	fmt.Print("Nama Mobil: ")
	fmt.Scan(&namaMobil)

	insertionDescendingNamaMobil(LM, *n)
	idx := binarySearchMobil(*LM, *n, nama)

	if idx != -1 {
		mobil = LM[idx]
		if mobil.Nama == namaMobil && mobil.Pabrikan == namaPabrikan {
			fmt.Printf("\nNama: %s\nTahun Keluar: %d\nPabrikan: %s\nPenjualan: %d\n\n", mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
		} else {
			fmt.Println("Data mobil tidak ditemukan!")
		}
	} else {
		fmt.Println("Data mobil tidak ditemukan!")
	}

	fmt.Println()
	fmt.Print("Masukkan apapun untuk melihat menu: ")
	fmt.Scan(&temp)
	fmt.Println()
	fmt.Println()
}
