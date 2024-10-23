package main 

import (
	"fmt"
	"strings"
)

const (
	maxBarang    = 100
	maxTransaksi = 100
)

type Barang struct {
	ID    int
	Nama  string
	Harga float64
	Stock int
}

type Transaksi struct {
	ID         int
	BarangID   int
	Jumlah     int
	HargaTotal float64
}

var barangList [maxBarang]Barang
var transaksiList [maxTransaksi]Transaksi
var jumlahBarang int
var jumlahTransaksi int
var nextBarangID int = 1
var nextTransaksiID int = 1

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Edit Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Lihat Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Tambah Transaksi")
		fmt.Println("7. Edit Transaksi")
		fmt.Println("8. Hapus Transaksi")
		fmt.Println("9. Lihat Transaksi")
		fmt.Println("10. Lihat laporan")
		fmt.Println("11. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			editBarang()
		case 3:
			hapusBarang()
		case 4:
			lihatBarang()
		case 5:
			cariBarang()
		case 6:
			tambahTransaksi()
		case 7:
			editTransaksi()
		case 8:
			hapusTransaksi()
		case 9:
			lihattransaksi()
		case 10:
			lihatLaporan()
		case 11:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahBarang() {
	var nama string
	var harga float64
	var stock int

	if jumlahBarang >= maxBarang {
		fmt.Println("Data barang penuh.")
		return
	}

	fmt.Print("Nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("Harga barang: ")
	fmt.Scan(&harga)
	fmt.Print("Stock barang: ")
	fmt.Scan(&stock)

	barang := Barang{
		ID:    nextBarangID,
		Nama:  nama,
		Harga: harga,
		Stock: stock,
	}

	barangList[jumlahBarang] = barang
	jumlahBarang++
	nextBarangID++

	fmt.Println("Barang berhasil ditambahkan.")
}

func editBarang() {
	var id, stock int
	var nama string
	var harga float64

	fmt.Print("ID barang yang akan diedit: ")
	fmt.Scan(&id)

	index := sequentialSearchBarang(id)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru: ")
	fmt.Scan(&nama)
	fmt.Print("Harga baru: ")
	fmt.Scan(&harga)
	fmt.Print("Stock baru: ")
	fmt.Scan(&stock)

	barangList[index].Nama = nama
	barangList[index].Harga = harga
	barangList[index].Stock = stock

	fmt.Println("Barang berhasil diedit.")
}

func hapusBarang() {
	var id int
	fmt.Print("ID barang yang akan dihapus: ")
	fmt.Scan(&id)

	index := sequentialSearchBarang(id)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	for i := index; i < jumlahBarang-1; i++ {
		barangList[i] = barangList[i+1]
	}
	jumlahBarang--

	fmt.Println("Barang berhasil dihapus.")
}

func lihatBarang() {
	var kategori, urutan int
	fmt.Println("Urutkan berdasarkan: 1. Nama 2. Harga 3. Stock")
	fmt.Scan(&kategori)
	fmt.Println("Urutan: 1. Ascending 2. Descending")
	fmt.Scan(&urutan)

	switch kategori {
	case 1:
		insertionSortNama(urutan == 1)
	case 2:
		selectionSortHarga(urutan == 1)
	case 3:
		selectionSortStock(urutan == 1)
	default:
		fmt.Println("Kategori tidak valid.")
		return
	}

	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("ID: %d, Nama: %s, Harga: %.2f, Stock: %d\n", barangList[i].ID, barangList[i].Nama, barangList[i].Harga, barangList[i].Stock)
	}
}

func cariBarang() {
	var keyword string
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scan(&keyword)

	for i := 0; i < jumlahBarang; i++ {
		if strings.Contains(strings.ToLower(barangList[i].Nama), strings.ToLower(keyword)) {
			fmt.Printf("ID: %d, Nama: %s, Harga: %.2f, Stock: %d\n", barangList[i].ID, barangList[i].Nama, barangList[i].Harga, barangList[i].Stock)
		}
	}
}

func tambahTransaksi() {
	var barangID, jumlah int
	fmt.Print("ID barang: ")
	fmt.Scan(&barangID)
	fmt.Print("Jumlah: ")
	fmt.Scan(&jumlah)

	index := sequentialSearchBarang(barangID)
	if index == -1 || barangList[index].Stock < jumlah {
		fmt.Println("Barang tidak ditemukan atau stok tidak cukup.")
		return
	}

	if jumlahTransaksi >= maxTransaksi {
		fmt.Println("Data transaksi penuh.")
		return
	}

	totalHarga := float64(jumlah) * barangList[index].Harga
	transaksi := Transaksi{
		ID:         nextTransaksiID,
		BarangID:   barangID,
		Jumlah:     jumlah,
		HargaTotal: totalHarga,
	}

	transaksiList[jumlahTransaksi] = transaksi
	jumlahTransaksi++
	barangList[index].Stock -= jumlah
	nextTransaksiID++

	fmt.Println("Transaksi berhasil ditambahkan.")
}

func editTransaksi() {
	var id, barangID, jumlah int
	fmt.Print("ID transaksi yang akan diedit: ")
	fmt.Scan(&id)

	index := binarySearchTransaksi(id)
	if index == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}

	fmt.Print("ID barang baru: ")
	fmt.Scan(&barangID)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&jumlah)

	indexBarang := sequentialSearchBarang(barangID)
	if indexBarang == -1 || barangList[indexBarang].Stock < jumlah {
		fmt.Println("Barang tidak ditemukan atau stok tidak cukup.")
		return
	}

	totalHarga := float64(jumlah) * barangList[indexBarang].Harga
	transaksiList[index].BarangID = barangID
	transaksiList[index].Jumlah = jumlah
	transaksiList[index].HargaTotal = totalHarga
	barangList[indexBarang].Stock -= jumlah

	fmt.Println("Transaksi berhasil diedit.")
}

func hapusTransaksi() {
	var id int
	fmt.Print("ID transaksi yang akan dihapus: ")
	fmt.Scan(&id)

	index := binarySearchTransaksi(id)
	if index == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}

	for i := index; i < jumlahTransaksi-1; i++ {
		transaksiList[i] = transaksiList[i+1]
	}
	jumlahTransaksi--

	fmt.Println("Transaksi berhasil dihapus.")
}

func lihattransaksi() {
	if jumlahTransaksi == 0 {
		fmt.Println("Tidak ada transaksi yang tercatat.")
		return
	}

	for i := 0; i < jumlahTransaksi; i++ {
		transaksi := transaksiList[i]
		indexBarang := sequentialSearchBarang(transaksi.BarangID)
		if indexBarang == -1 {
			fmt.Printf("ID Transaksi: %d - Barang tidak ditemukan\n", transaksi.ID)
			continue
		}

		barang := barangList[indexBarang]
		fmt.Printf("ID Transaksi: %d\n", transaksi.ID)
		fmt.Printf("ID Barang: %d\n", transaksi.BarangID)
		fmt.Printf("Nama Barang: %s\n", barang.Nama)
		fmt.Printf("Jumlah: %d\n", transaksi.Jumlah)
		fmt.Printf("Harga Total: %.2f\n", transaksi.HargaTotal)
	}
}

func lihatLaporan() {
	totalModal := 0.0
	totalPendapatan := 0.0

	for i := 0; i < jumlahBarang; i++ {
		totalModal += float64(barangList[i].Stock) * barangList[i].Harga
	}

	for i := 0; i < jumlahTransaksi; i++ {
		totalPendapatan += transaksiList[i].HargaTotal
	}

	fmt.Printf("Total Modal: %.2f\n", totalModal)
	fmt.Printf("Total Pendapatan: %.2f\n", totalPendapatan)
}

func sequentialSearchBarang(id int) int {
	for i := 0; i < jumlahBarang; i++ {
		if barangList[i].ID == id {
			return i
		}
	}
	return -1
}

func binarySearchTransaksi(id int) int {
	low := 0
	high := jumlahTransaksi - 1

	for low <= high {
		mid := (low + high) / 2
		if transaksiList[mid].ID == id {
			return mid
		} else if transaksiList[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}


func insertionSortNama(ascending bool) {
    for i := 1; i < jumlahBarang; i++ {
        key := barangList[i]
        j := i - 1

        for j >= 0 && ((ascending && barangList[j].Nama > key.Nama) || (!ascending && barangList[j].Nama < key.Nama)) {
            barangList[j+1] = barangList[j]
            j = j - 1
        }
        barangList[j+1]=key
	}
}
func selectionSortHarga(ascending bool) {
	for i := 0; i < jumlahBarang-1; i++ {
		idx := i
		for j := i + 1; j < jumlahBarang; j++ {
			if (ascending && barangList[j].Harga < barangList[idx].Harga) ||
				(!ascending && barangList[j].Harga > barangList[idx].Harga) {
				idx = j
			}
		}
		barangList[i], barangList[idx] = barangList[idx], barangList[i]
	}
}

func selectionSortStock(ascending bool) {
	for i := 0; i < jumlahBarang-1; i++ {
		idx := i
		for j := i + 1; j < jumlahBarang; j++ {
			if (ascending && barangList[j].Stock < barangList[idx].Stock) ||
				(!ascending && barangList[j].Stock > barangList[idx].Stock) {
				idx = j
			}
		}
		barangList[i], barangList[idx] = barangList[idx], barangList[i]
	}
}