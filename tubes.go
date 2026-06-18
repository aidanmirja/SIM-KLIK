package main

import "fmt"

const MAX = 100

type Pasien struct {
	ID              string
	Nama            string
	Umur            int
	Alamat          string
	JumlahKunjungan int
}

type Layanan struct {
	ID    string
	Nama  string
	Harga int
}

type Kunjungan struct {
	IDKunjungan string
	IDPasien    string
	IDLayanan   string
	Tanggal     string
	TotalBiaya  int
}

var pasien [MAX]Pasien
var layanan [MAX]Layanan
var kunjungan [MAX]Kunjungan

var nPasien int
var nLayanan int
var nKunjungan int

func main() {
	for {
		var pilih int

		fmt.Println("\n===== SIM-KLIK =====")
		fmt.Println("1. Kelola Pasien")
		fmt.Println("2. Kelola Layanan")
		fmt.Println("3. Kelola Kunjungan")
		fmt.Println("4. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			menuPasien()
		case 2:
			menuLayanan()
		case 3:
			menuKunjungan()
		case 4:
			menuStatistik()
		case 0:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPasien() {
	for {
		var pilih int

		fmt.Println("\n=== MENU PASIEN ===")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Ubah Pasien")
		fmt.Println("3. Hapus Pasien")
		fmt.Println("4. Tampil Pasien")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahPasien()
		case 2:
			ubahPasien()
		case 3:
			hapusPasien()
		case 4:
			tampilPasien()
		case 5:
			cariPasienSequential()
		case 6:
			cariPasienBinary()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuLayanan() {
	for {
		var pilih int

		fmt.Println("\n=== MENU LAYANAN ===")
		fmt.Println("1. Tambah Layanan")
		fmt.Println("2. Ubah Layanan")
		fmt.Println("3. Hapus Layanan")
		fmt.Println("4. Tampil Layanan")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahLayanan()
		case 2:
			ubahLayanan()
		case 3:
			hapusLayanan()
		case 4:
			tampilLayanan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuKunjungan() {
	for {
		var pilih int

		fmt.Println("\n=== MENU KUNJUNGAN ===")
		fmt.Println("1. Tambah Kunjungan")
		fmt.Println("2. Riwayat Kunjungan Rentang Tanggal")
		fmt.Println("3. Urut Berdasarkan Tanggal (Selection Sort)")
		fmt.Println("4. Urut Berdasarkan Biaya (Insertion Sort)")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahKunjungan()
		case 2:
			riwayatKunjungan()
		case 3:
			tampilUrutTanggal()
		case 4:
			tampilUrutBiaya()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuStatistik() {

	for {

		var pilih int

		fmt.Println("\n=== MENU STATISTIK ===")
		fmt.Println("1. Statistik Kunjungan Harian")
		fmt.Println("2. Total Kunjungan Periode")
		fmt.Println("3. Layanan Paling Diminati")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			statistikKunjunganHarian()

		case 2:
			totalKunjunganPeriode()

		case 3:
			layananPalingDiminati()

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahPasien() {
	if nPasien >= MAX {
		fmt.Println("Data pasien penuh.")
		return
	}

	fmt.Println("\n=== TAMBAH PASIEN ===")

	fmt.Print("ID      : ")
	fmt.Scan(&pasien[nPasien].ID)

	fmt.Print("Nama    : ")
	fmt.Scan(&pasien[nPasien].Nama)

	fmt.Print("Umur    : ")
	fmt.Scan(&pasien[nPasien].Umur)

	fmt.Print("Alamat  : ")
	fmt.Scan(&pasien[nPasien].Alamat)

	pasien[nPasien].JumlahKunjungan = 0

	nPasien++

	fmt.Println("Data pasien berhasil ditambahkan.")
}

func tampilPasien() {
	if nPasien == 0 {
		fmt.Println("Belum ada data pasien.")
		return
	}

	fmt.Println("\n=== DATA PASIEN ===")

	for i := 0; i < nPasien; i++ {
		fmt.Printf("%d. %s | %s | %d | %s | Kunjungan:%d\n",
			i+1,
			pasien[i].ID,
			pasien[i].Nama,
			pasien[i].Umur,
			pasien[i].Alamat,
			pasien[i].JumlahKunjungan)
	}
}

func ubahPasien() {
	var id string

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	for i := 0; i < nPasien; i++ {
		if pasien[i].ID == id {

			fmt.Print("Nama Baru   : ")
			fmt.Scan(&pasien[i].Nama)

			fmt.Print("Umur Baru   : ")
			fmt.Scan(&pasien[i].Umur)

			fmt.Print("Alamat Baru : ")
			fmt.Scan(&pasien[i].Alamat)

			fmt.Println("Data berhasil diubah.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func hapusPasien() {
	var id string

	fmt.Print("ID Pasien : ")
	fmt.Scan(&id)

	for i := 0; i < nPasien; i++ {

		if pasien[i].ID == id {

			for j := i; j < nPasien-1; j++ {
				pasien[j] = pasien[j+1]
			}

			nPasien--

			fmt.Println("Data berhasil dihapus.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func tambahLayanan() {
	if nLayanan >= MAX {
		fmt.Println("Data layanan penuh.")
		return
	}

	fmt.Println("\n=== TAMBAH LAYANAN ===")

	fmt.Print("ID Layanan : ")
	fmt.Scan(&layanan[nLayanan].ID)

	fmt.Print("Nama       : ")
	fmt.Scan(&layanan[nLayanan].Nama)

	fmt.Print("Harga      : ")
	fmt.Scan(&layanan[nLayanan].Harga)

	nLayanan++

	fmt.Println("Data layanan berhasil ditambahkan.")
}

func tampilLayanan() {
	if nLayanan == 0 {
		fmt.Println("Belum ada data layanan.")
		return
	}

	fmt.Println("\n=== DATA LAYANAN ===")

	for i := 0; i < nLayanan; i++ {
		fmt.Printf("%d. %s | %s | Rp%d\n",
			i+1,
			layanan[i].ID,
			layanan[i].Nama,
			layanan[i].Harga)
	}
}

func ubahLayanan() {
	var id string

	fmt.Print("ID Layanan : ")
	fmt.Scan(&id)

	for i := 0; i < nLayanan; i++ {

		if layanan[i].ID == id {

			fmt.Print("Nama Baru  : ")
			fmt.Scan(&layanan[i].Nama)

			fmt.Print("Harga Baru : ")
			fmt.Scan(&layanan[i].Harga)

			fmt.Println("Data berhasil diubah.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func hapusLayanan() {
	var id string

	fmt.Print("ID Layanan : ")
	fmt.Scan(&id)

	for i := 0; i < nLayanan; i++ {

		if layanan[i].ID == id {

			for j := i; j < nLayanan-1; j++ {
				layanan[j] = layanan[j+1]
			}

			nLayanan--

			fmt.Println("Data berhasil dihapus.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func cariPasien(id string) int {
	for i := 0; i < nPasien; i++ {
		if pasien[i].ID == id {
			return i
		}
	}
	return -1
}

func cariLayanan(id string) int {
	for i := 0; i < nLayanan; i++ {
		if layanan[i].ID == id {
			return i
		}
	}
	return -1
}

// TAMBAH KUNJUNGAN

func tambahKunjungan() {

	if nKunjungan >= MAX {
		fmt.Println("Data kunjungan penuh")
		return
	}

	var idPasien string
	var idLayanan string

	fmt.Println("\n=== TAMBAH KUNJUNGAN ===")

	fmt.Print("ID Kunjungan : ")
	fmt.Scan(&kunjungan[nKunjungan].IDKunjungan)

	fmt.Print("Tanggal (YYYY-MM-DD) : ")
	fmt.Scan(&kunjungan[nKunjungan].Tanggal)

	fmt.Print("ID Pasien : ")
	fmt.Scan(&idPasien)

	idxPasien := cariPasien(idPasien)

	if idxPasien == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}

	fmt.Print("ID Layanan : ")
	fmt.Scan(&idLayanan)

	idxLayanan := cariLayanan(idLayanan)

	if idxLayanan == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}

	kunjungan[nKunjungan].IDPasien = idPasien
	kunjungan[nKunjungan].IDLayanan = idLayanan
	kunjungan[nKunjungan].TotalBiaya =
		layanan[idxLayanan].Harga

	pasien[idxPasien].JumlahKunjungan++

	nKunjungan++

	fmt.Println("Kunjungan berhasil ditambahkan")
}

// RIWAYAT RENTANG TANGGAL

func riwayatKunjungan() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	fmt.Println("\n=== RIWAYAT KUNJUNGAN ===")

	ada := false

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			fmt.Printf("%s | %s | %s | %s | Rp%d\n",
				kunjungan[i].IDKunjungan,
				kunjungan[i].Tanggal,
				kunjungan[i].IDPasien,
				kunjungan[i].IDLayanan,
				kunjungan[i].TotalBiaya)

			ada = true
		}
	}

	if !ada {
		fmt.Println("Tidak ada data")
	}
}

// SELECTION SORT TANGGAL

func selectionSortTanggal(data *[MAX]Kunjungan, n int) {

	for i := 0; i < n-1; i++ {

		min := i

		for j := i + 1; j < n; j++ {

			if data[j].Tanggal <
				data[min].Tanggal {

				min = j
			}
		}

		temp := data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

func tampilUrutTanggal() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	var temp [MAX]Kunjungan
	var nTemp int

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			temp[nTemp] = kunjungan[i]
			nTemp++
		}
	}

	selectionSortTanggal(&temp, nTemp)

	fmt.Println("\n=== TRANSAKSI TERURUT BERDASARKAN TANGGAL ===")

	for i := 0; i < nTemp; i++ {

		fmt.Printf("%s | %s | Rp%d\n",
			temp[i].IDKunjungan,
			temp[i].Tanggal,
			temp[i].TotalBiaya)
	}
}

// INSERTION SORT BIAYA

func insertionSortBiaya(data *[MAX]Kunjungan, n int) {

	for i := 1; i < n; i++ {

		key := data[i]
		j := i - 1

		for j >= 0 &&
			data[j].TotalBiaya > key.TotalBiaya {

			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}

func tampilUrutBiaya() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	var temp [MAX]Kunjungan
	var nTemp int

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			temp[nTemp] = kunjungan[i]
			nTemp++
		}
	}

	insertionSortBiaya(&temp, nTemp)

	fmt.Println("\n=== TRANSAKSI TERURUT BERDASARKAN BIAYA ===")

	for i := 0; i < nTemp; i++ {

		fmt.Printf("%s | %s | Rp%d\n",
			temp[i].IDKunjungan,
			temp[i].Tanggal,
			temp[i].TotalBiaya)
	}
}

// SEQUENTIAL SEARCH

func cariPasienSequential() {

	if nPasien == 0 {
		fmt.Println("Data pasien kosong")
		return
	}

	var id string

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	for i := 0; i < nPasien; i++ {

		if pasien[i].ID == id {

			fmt.Println("\n=== DATA DITEMUKAN ===")
			fmt.Println("ID        :", pasien[i].ID)
			fmt.Println("Nama      :", pasien[i].Nama)
			fmt.Println("Umur      :", pasien[i].Umur)
			fmt.Println("Alamat    :", pasien[i].Alamat)
			fmt.Println("Kunjungan :", pasien[i].JumlahKunjungan)

			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// SORT PASIEN BY ID

func sortPasienByID() {

	for i := 0; i < nPasien-1; i++ {

		min := i

		for j := i + 1; j < nPasien; j++ {

			if pasien[j].ID < pasien[min].ID {
				min = j
			}
		}

		temp := pasien[i]
		pasien[i] = pasien[min]
		pasien[min] = temp
	}
}

// BINARY SEARCH

func cariPasienBinary() {

	if nPasien == 0 {
		fmt.Println("Data pasien kosong")
		return
	}

	sortPasienByID()

	var id string

	fmt.Print("Masukkan ID Pasien : ")
	fmt.Scan(&id)

	low := 0
	high := nPasien - 1

	for low <= high {

		mid := (low + high) / 2

		if pasien[mid].ID == id {

			fmt.Println("\n=== DATA DITEMUKAN ===")
			fmt.Println("ID        :", pasien[mid].ID)
			fmt.Println("Nama      :", pasien[mid].Nama)
			fmt.Println("Umur      :", pasien[mid].Umur)
			fmt.Println("Alamat    :", pasien[mid].Alamat)
			fmt.Println("Kunjungan :", pasien[mid].JumlahKunjungan)

			return

		} else if id < pasien[mid].ID {

			high = mid - 1

		} else {

			low = mid + 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// STATISTIK HARIAN

func statistikKunjunganHarian() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	fmt.Println("\n=== STATISTIK KUNJUNGAN HARIAN ===")

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			jumlah := 0

			for j := 0; j < nKunjungan; j++ {

				if kunjungan[j].Tanggal ==
					kunjungan[i].Tanggal {

					jumlah++
				}
			}

			sudah := false

			for k := 0; k < i; k++ {

				if kunjungan[k].Tanggal ==
					kunjungan[i].Tanggal {

					sudah = true
				}
			}

			if !sudah {

				fmt.Printf("%s : %d kunjungan\n",
					kunjungan[i].Tanggal,
					jumlah)
			}
		}
	}
}

// TOTAL KUNJUNGAN

func totalKunjunganPeriode() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	total := 0

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			total++
		}
	}

	fmt.Println("\nTotal Kunjungan :", total)
}

// LAYANAN PALING DIMINATI

func layananPalingDiminati() {

	if nKunjungan == 0 {
		fmt.Println("Belum ada data kunjungan")
		return
	}

	var awal string
	var akhir string

	fmt.Print("Tanggal Awal  : ")
	fmt.Scan(&awal)

	fmt.Print("Tanggal Akhir : ")
	fmt.Scan(&akhir)

	var counter [MAX]int

	for i := 0; i < nKunjungan; i++ {

		if kunjungan[i].Tanggal >= awal &&
			kunjungan[i].Tanggal <= akhir {

			for j := 0; j < nLayanan; j++ {

				if kunjungan[i].IDLayanan ==
					layanan[j].ID {

					counter[j]++
				}
			}
		}
	}

	maxIdx := 0

	for i := 1; i < nLayanan; i++ {

		if counter[i] > counter[maxIdx] {

			maxIdx = i
		}
	}

	fmt.Println("\n=== LAYANAN PALING DIMINATI ===")
	fmt.Println("ID Layanan   :", layanan[maxIdx].ID)
	fmt.Println("Nama Layanan :", layanan[maxIdx].Nama)
	fmt.Println("Harga        :", layanan[maxIdx].Harga)
	fmt.Println("Dipakai      :", counter[maxIdx], "kali")
}