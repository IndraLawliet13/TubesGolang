package main

// Notes:
// 1. tambah id barang (di tambahData() id = jumlah data sekarang)
// 2.

import (
	"fmt"
	"strings"
)

const maxarray = 10000

type produk struct {
	nama, merk, jenis string
	harga, stok, id   int
}
type transaksi struct {
	idBarang, jumlahBarang, totalHarga int
}
type tabProduk [maxarray]produk
type tabTransaksi struct {
	riwayat [maxarray]transaksi
	n       int
}

func menu() {
	var pilihan, pilihan2, n, i int
	var m string
	var p tabProduk
	var t tabTransaksi
	m = "=============== Menu Inventaris Toko ===============\n\n1. Tambah Data\n2. Edit Data\n3. Hapus Data\n4. Tampilkan Data\n5. Cari Data\n6. Catat Transaksi\n7. Exit\nJumlah Data Sekarang (%d)\nMasukkan Pilihan >> "
	for pilihan != 7 {
		fmt.Printf(m, n)
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			var tData int
			fmt.Print("Mau Tambah Berapa Data ? ")
			fmt.Scan(&tData)
			if tData > 0 {
				for i = 0; i < tData; i++ {
					tambahData(&p, &n)
				}
			}
		} else if pilihan == 2 {
			if n > 0 {
				editData(&p, &n)
			}
		} else if pilihan == 3 {
			if n > 0 {
				hapusData(&p, &n)
			}
		} else if pilihan == 4 {
			if n != 0 {
				pilihan2 = 0
				for pilihan2 != 8 {
					fmt.Print("1. Urutkan berdasarkan nama (asc)\n2. Urutkan berdasarkan nama (desc)\n3. Urutkan berdasarkan harga (asc)\n4. Urutkan berdasarkan harga (desc)\n5. Urutkan berdasarkan stok (asc)\n6. Urutkan berdasarkan stok (desc)\n7. Tampilkan Data Transaksi\n8. Exit\nMasukkan Pilihan >> ")
					fmt.Scan(&pilihan2)
					if pilihan2 == 1 {
						tampilData(p, n, "nama", "asc")
					} else if pilihan2 == 2 {
						tampilData(p, n, "nama", "desc")
					} else if pilihan2 == 3 {
						tampilData(p, n, "harga", "asc")
					} else if pilihan2 == 4 {
						tampilData(p, n, "harga", "desc")
					} else if pilihan2 == 5 {
						tampilData(p, n, "stok", "asc")
					} else if pilihan2 == 6 {
						tampilData(p, n, "stok", "desc")
					} else if pilihan2 == 7 {
						if t.n > 0 {
							tampilDataTransaksi(t, p, n)
						}
					}
				}
			}
		} else if pilihan == 5 {
			if n > 0 {
				cariData(p, n)
			}
		} else if pilihan == 6 {
			if n > 0 {
				catatTransaksi(&t, &p, &n)
			}
		}
	}
}

// Menu Pertama
func tambahData(p *tabProduk, n *int) {
	if *n < maxarray {
		fmt.Println("=============== Tambah Data ===============")
		fmt.Print("Masukkan Nama Barang : ")
		fmt.Scan(&p[*n].nama)
		fmt.Print("Masukkan Merk Barang : ")
		fmt.Scan(&p[*n].merk)
		fmt.Print("Masukkan Jenis Barang : ")
		fmt.Scan(&p[*n].jenis)
		fmt.Print("Masukkan Harga : ")
		fmt.Scan(&p[*n].harga)
		fmt.Print("Masukkan Stok : ")
		fmt.Scan(&p[*n].stok)
		p[*n].id = *n
		*n += 1
		fmt.Println("Data Berhasil Ditambahkan!\n==============================")
	}
}

// Menu kedua
func editData(p *tabProduk, n *int) {
	var found int
	var nama, merk, jenis string
	var harga, stok int
	found = -1
	fmt.Println("=============== Edit Data ===============")
	for found == -1 && nama != "x" {
		fmt.Print("Masukkan Nama Barang (x = batal) >> ")
		fmt.Scan(&nama)
		if nama == "x" {
			fmt.Println("Data Batal Diedit\n==============================")
			continue
		}
		found = searchBarangByName(*p, nama, *n)
		if found == -1 {
			fmt.Println("Barang dengan nama tersebut tidak ditemukan, mohon input lagi nama barang yang benar!")
		}
	}
	// Jika ketemu dan nama != x:
	if nama != "x" {
		fmt.Println("Masukkan data baru, ketik x/-1 untuk tidak mengubahnya")
		fmt.Print("Masukkan Nama Barang (x = skip): ")
		fmt.Scan(&nama)
		if nama != "x" {
			p[found].nama = nama
		}
		fmt.Print("Masukkan Merk Barang (x = skip): ")
		fmt.Scan(&merk)
		if merk != "x" {
			p[found].merk = merk
		}
		fmt.Print("Masukkan Jenis Barang (x = skip): ")
		fmt.Scan(&jenis)
		if jenis != "x" {
			p[found].jenis = jenis
		}
		fmt.Print("Masukkan Harga (-1 = skip): ")
		fmt.Scan(&harga)
		if harga != -1 {
			p[found].harga = int(harga)
		}
		fmt.Print("Masukkan Stok (-1 = skip): ")
		fmt.Scan(&stok)
		if stok != -1 {
			p[found].stok = int(stok)
		}
		fmt.Println("Data Berhasil Diedit\n==============================")
	}
}

// Menu ketiga
func hapusData(p *tabProduk, n *int) {
	var found, no, i int
	var nama string
	var new tabProduk
	found = -1
	fmt.Println("=============== Hapus Data ===============")
	for found == -1 {
		fmt.Print("Masukkan Nama Barang >> ")
		fmt.Scan(&nama)
		found = searchBarangByName(*p, nama, *n)
		if found == -1 {
			fmt.Println("Barang dengan nama tersebut tidak ditemukan, mohon input lagi nama barang yang benar!")
		}
	}
	no = 0
	for i = 0; i < *n; i++ {
		if i != found {
			new[no] = p[i]
			no++
		}
	}
	*p = new
	*n = no
	fmt.Println("Data Berhasil Dihapus\n==============================")
}

// Menu keempat
func tampilData(p tabProduk, n int, by, method string) {
	if n != 0 {
		var i int
		if by == "nama" {
			insertionSortNama(&p, n, method)
		} else if by == "harga" {
			insertionSortHarga(&p, n, method)
		} else if by == "stok" {
			selectionSortStok(&p, n, method)
		} else {
			insertionSortHarga(&p, n, "asc")
		}
		fmt.Print("|\tNo\t|\tNama\t|\tMerk\t|\tJenis\t|\tHarga\t|\tStok\t|\n")
		for i = 0; i < n; i++ {
			fmt.Printf("|\t%d\t|\t%s\t|\t%s\t|\t%s\t|\t%d\t|\t%d\t|\n", i+1, p[i].nama, p[i].merk, p[i].jenis, p[i].harga, p[i].stok)
		}
	}
}
func tampilDataTransaksi(t tabTransaksi, P tabProduk, n int) {
	var i int
	var p produk
	fmt.Print("|\tNo\t|\tNama\t|\tMerk\t|\tJenis\t|\tJumlah\t|\tTotal\t|\n")
	for i = 0; i < t.n; i++ {
		p = searchBarangById(P, n, t.riwayat[i].idBarang)
		fmt.Printf("|\t%d\t|\t%s\t|\t%s\t|\t%s\t|\t%d\t|\t%d\t|\n", i+1, p.nama, p.merk, p.jenis, t.riwayat[i].jumlahBarang, t.riwayat[i].totalHarga)
	}
}

// Menu kelima
func cariData(p tabProduk, n int) {
	var pil, found, harga, stok int
	var nama string
	pil = 0
	for pil != 4 && (nama != "x" || harga != -1 || stok != -1) {
		fmt.Print("=============== Cari Data ===============\n1. Cari berdasarkan nama\n2. Cari berdasarkan harga\n3. Cari berdasarkan stok\n4. Exit\nMasukkan Pilihan >> ")
		fmt.Scan(&pil)
		if pil == 1 {
			found = -1
			for found == -1 && nama != "x" {
				fmt.Print("Masukkan Nama Barang (x = batal) >> ")
				fmt.Scan(&nama)
				if nama == "x" {
					fmt.Println("Data Batal Dicari\n==============================")
					continue
				}
				found = searchBarangByName(p, nama, n)
				if found == -1 {
					fmt.Println("Barang dengan harga tersebut tidak ditemukan, mohon input lagi harga barang yang benar!")
				} else {
					fmt.Print("|\tNama\t|\tMerk\t|\tJenis\t|\tHarga\t|\tStok\t|\n")
					for i := 0; i < n; i++ {
						if p[i].nama == nama {
							fmt.Printf("|\t%s\t|\t%s\t|\t%s\t|\t%d\t|\t%d\t|\n==============================\n", p[i].nama, p[i].merk, p[i].jenis, p[i].harga, p[i].stok)
						}
					}
				}
			}
		} else if pil == 2 {
			found = -1
			for found == -1 && harga != -1 {
				fmt.Print("Masukkan Harga Barang (-1 = batal) >> ")
				fmt.Scan(&harga)
				if harga == -1 {
					fmt.Println("Data Batal Dicari\n==============================")
					continue
				}
				found = searchBarangByHarga(&p, harga, n)
				if found == -1 {
					fmt.Println("Barang dengan harga tersebut tidak ditemukan, mohon input lagi harga barang yang benar!")
				} else {
					fmt.Print("|\tNama\t|\tMerk\t|\tJenis\t|\tHarga\t|\tStok\t|\n")
					for i := 0; i < n; i++ {
						if p[i].harga == harga {
							fmt.Printf("|\t%s\t|\t%s\t|\t%s\t|\t%d\t|\t%d\t|\n==============================\n", p[i].nama, p[i].merk, p[i].jenis, p[i].harga, p[i].stok)
						}
					}
				}
			}
		} else if pil == 3 {
			found = -1
			for found == -1 && stok != -1 {
				fmt.Print("Masukkan Stok Barang (-1 = batal) >> ")
				fmt.Scan(&stok)
				if stok == -1 {
					fmt.Println("Data Batal Dicari\n==============================")
					continue
				}
				found = searchBarangByStok(p, stok, n)
				if found == -1 {
					fmt.Println("Barang dengan stok tersebut tidak ditemukan, mohon input lagi stok barang yang benar!")
				} else {
					fmt.Print("|\tNama\t|\tMerk\t|\tJenis\t|\tHarga\t|\tStok\t|\n")
					for i := 0; i < n; i++ {
						if p[i].stok == stok {
							fmt.Printf("|\t%s\t|\t%s\t|\t%s\t|\t%d\t|\t%d\t|\n==============================\n", p[i].nama, p[i].merk, p[i].jenis, p[i].harga, p[i].stok)
						}
					}
				}
			}
		}
	}
}

// Menu keenam
func catatTransaksi(t *tabTransaksi, p *tabProduk, n *int) {
	var found, jumlah int
	var nama string
	found = -1
	for found == -1 && nama != "x" {
		fmt.Print("Masukkan Nama Barang (x = batal) >> ")
		fmt.Scan(&nama)
		if nama == "x" {
			fmt.Println("Data Batal Dicari\n==============================")
			continue
		}
		found = searchBarangByName(*p, nama, *n)
		if found == -1 {
			fmt.Println("Barang dengan harga tersebut tidak ditemukan, mohon input lagi harga barang yang benar!")
		}
	}
	// Jika ketemu dan nama != x:
	if nama != "x" && p[found].stok > 0 {
		fmt.Print("Masukkan jumlah pembelian (stok saat ini = ", p[found].stok, ") >> ")
		fmt.Scan(&jumlah)
		for jumlah > p[found].stok {
			fmt.Println("Jumlah melebihi stok yang ada, mohon masukkan jumlah ulang!")
			fmt.Print("Masukkan jumlah pembelian (stok saat ini = ", p[found].stok, ") >> ")
			fmt.Scan(&jumlah)
		}
		// Jika jumlah <= stok maka lakukan penjumlahan total biaya/harga dan kurangi stok yang ada
		p[found].stok = p[found].stok - jumlah
		t.riwayat[t.n] = transaksi{p[found].id, jumlah, jumlah * p[found].harga}
		t.n += 1
		fmt.Println("Data Transaksi Berhasil Dicatat\n==============================")
	}
}

// function pendukung
// Searching
func searchBarangByName(p tabProduk, nama string, n int) int {
	var i, found int
	found = -1
	// sequential search
	for i = 0; i < n && found == -1; i++ {
		if p[i].nama == nama {
			found = i
		}
	}
	return found
}
func searchBarangByHarga(p *tabProduk, harga int, n int) int {
	insertionSortHarga(*&p, n, "asc")
	var min, max, mid, index int
	min = 0
	max = n - 1
	index = -1
	for min <= max && index == -1 {
		mid = (min + max) / 2
		if p[mid].harga == harga {
			index = mid
		} else if p[mid].harga < harga {
			min = mid + 1
		} else if p[mid].harga > harga {
			max = mid - 1
		}
	}
	return index
}
func searchBarangByStok(p tabProduk, stok, n int) int {
	var i, found int
	found = -1
	// sequential search
	for i = 0; i < n && found == -1; i++ {
		if p[i].stok == stok {
			found = i
		}
	}
	return found
}
func searchBarangById(P tabProduk, n, id int) produk {
	var i, found int
	var p produk
	found = -1
	// sequential search
	for i = 0; i < n && found == -1; i++ {
		if P[i].id == id {
			found = i
			p = P[i]
		}
	}
	return p
}

// Sorting
func insertionSortHarga(p *tabProduk, n int, method string) {
	var i, j int
	if method == "asc" {
		for i = 0; i < n; i++ {
			for j = i; j > 0 && p[j].harga < p[j-1].harga; j-- {
				p[j-1], p[j] = p[j], p[j-1]
			}
		}
	} else if method == "desc" {
		for i = 0; i < n; i++ {
			for j = i; j > 0 && p[j].harga > p[j-1].harga; j-- {
				p[j-1], p[j] = p[j], p[j-1]
			}
		}
	}
}

func insertionSortNama(p *tabProduk, n int, method string) {
	var i, j int
	var temp produk
	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1
		// Jika method adalah "asc"
		if method == "asc" {
			for j >= 0 && strings.ToLower(p[j].nama) > strings.ToLower(temp.nama) {
				p[j+1] = p[j]
				j--
			}
		} else if method == "desc" {
			for j >= 0 && strings.ToLower(p[j].nama) < strings.ToLower(temp.nama) {
				p[j+1] = p[j]
				j--
			}
		}
		p[j+1] = temp
	}
}

func selectionSortStok(p *tabProduk, n int, method string) {
	var m, i, j int
	if method == "asc" {
		for i = 0; i < n-1; i++ {
			m = i
			for j = i; j < n; j++ {
				if p[m].stok > p[j].stok {
					m = j
				}
			}
			p[i], p[m] = p[m], p[i]
		}
	} else if method == "desc" {
		for i = 0; i < n-1; i++ {
			m = i
			for j = i; j < n; j++ {
				if p[m].stok < p[j].stok {
					m = j
				}
			}
			p[i], p[m] = p[m], p[i]
		}
	}
}

// Main funcnya
func main() {
	menu()
}
