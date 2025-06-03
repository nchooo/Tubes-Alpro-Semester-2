package main
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)
type Pegawai struct {
	ID    int
	Nama  string
	Shift string
	Gaji  int
}

const maxPegawai = 100

var pegawaiList [maxPegawai]Pegawai
var jumlahData int = 7

func main() {
	pegawaiList[0] = Pegawai{ID: 101, Nama: "Andi", Shift: "pagi", Gaji: 5000000}
	pegawaiList[1] = Pegawai{ID: 102, Nama: "Budi", Shift: "siang", Gaji: 4500000}
	pegawaiList[2] = Pegawai{ID: 103, Nama: "Citra", Shift: "malam", Gaji: 6000000}
	pegawaiList[3] = Pegawai{ID: 104, Nama: "Dewi", Shift: "pagi", Gaji: 5500000}
	pegawaiList[4] = Pegawai{ID: 105, Nama: "Eko", Shift: "siang", Gaji: 4700000}
	pegawaiList[5] = Pegawai{ID: 106, Nama: "Fajar", Shift: "malam", Gaji: 5200000}
	pegawaiList[6] = Pegawai{ID: 107, Nama: "Gita", Shift: "pagi", Gaji: 5800000}

	var pilihan int
	padding := "                                            "
	for pilihan != 9 {
		clearScreen()
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			tambahData(&pegawaiList, &jumlahData)
		case 2:
			clearScreen()
			tampilkanData(pegawaiList, jumlahData)
		case 3:
			clearScreen()
			editData(&pegawaiList, jumlahData)
		case 4:
			clearScreen()
			hapusData(&pegawaiList, &jumlahData)
		case 5:
			clearScreen()
			urutGajiTertinggi(&pegawaiList, jumlahData)
			fmt.Println(padding + "╔══════════════════════════════════════════════════╗")
			fmt.Println(padding + "║ Data telah diurutkan berdasarkan gaji tertinggi. ║")
			fmt.Println(padding + "╚══════════════════════════════════════════════════╝")
		case 6:
			clearScreen()
			urutGajiTerendah(&pegawaiList, jumlahData)
			fmt.Println(padding + "╔══════════════════════════════════════════════════╗")
			fmt.Println(padding + "║  Data telah diurutkan berdasarkan gaji terendah. ║")
			fmt.Println(padding + "╚══════════════════════════════════════════════════╝")
		case 7:
			clearScreen()
			totalGaji(pegawaiList, jumlahData)
		case 8:
			clearScreen()
			var nama string
			fmt.Print(padding + "Masukkan Nama pegawai yang dicari: ")
			fmt.Scan(&nama)
			pos := cariPegawaiNamaBinary(&pegawaiList, jumlahData, nama)
			if pos != -1 {
				fmt.Println(padding + "Data Pegawai Ditemukan:")
				fmt.Println(padding + "╔══════╦════════════════════╦════════════╦════════════╗")
				fmt.Println(padding + "║ ID   ║ Nama               ║ Shift      ║ Gaji       ║")
				fmt.Println(padding + "╠══════╬════════════════════╬════════════╬════════════╣")
				fmt.Printf(padding + "║ %-4d ║ %-18s ║ %-10s ║ %-10d ║\n", pegawaiList[pos].ID, pegawaiList[pos].Nama, pegawaiList[pos].Shift, pegawaiList[pos].Gaji)
				fmt.Println(padding + "╚══════╩════════════════════╩════════════╩════════════╝")
				fmt.Scanln()
				
			} else {
				fmt.Println(padding + "Pegawai tidak ditemukan.")
			}
			fmt.Scanln()
		case 9:
			clearScreen()
			fmt.Println(padding + "⠀⠀   ⣿⣽⣿⣿⣿⣿⣿⠿⣭⣛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	 ⣿⣿⣿⣿⣿⣿⢏⡟⢦⡝⣾⣿⣿⡿⣟⣿⡿⣿⣿⣯⣿⣟⣿⣽⣿⣟⣿⣟⣿⣟⣿⣽⣿⣯⣿⣿⣽⣿⣻⣿⢭⢿⣿⣿⣿⣿⣿⣿⢿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⢣⣛⡜⣣⠞⣜⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣿⡿⣿⣯⣿⣿⢿⣿⡿⣿⣻⣿⣿⣿⣾⣿⡿⣿⡯⡞⣭⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "     ⣿⣿⣿⣿⣣⣛⡴⣓⣌⢻⡰⣻⣿⣿⢿⣷⣿⣿⣿⣿⣿⣟⣿⣿⣿⣯⣿⣿⣿⢿⣿⣿⣿⣿⣿⣽⣷⣿⣿⣳⣽⣦⣽⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⡿⢿⠿⠿⡿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣿⣿⣿⣿⣿⣿⣿⣾⣿⡿⣿⣾⣿⣿⣿⣿⣿⣿⣻⣿⣿⣿⣿⢿⡟⣟⢿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⢯⡛⡭⠞⣥⠳⡡⢖⠢⢭⣙⠻⣿⣿⣯⣿⣿⣿⣿⣿⣾⣿⣟⣿⣷⣿⣿⣿⢿⣿⣾⣿⣿⣿⣿⣿⣿⡟⣬⠳⡜⢎⡞⢿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	 ⢯⡞⡱⢭⡙⢤⠣⣑⠊⡕⡒⠬⡣⢝⣿⣿⣿⣿⣿⣿⣿⣿⣟⣿⣿⣿⡿⣟⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⠳⣌⠧⡹⠜⡜⣫⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⢯⡜⣣⠣⢞⣠⣓⣨⣜⣠⣑⠣⡑⢎⠞⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⡿⢓⣌⡶⢕⣻⣾⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⡳⣜⣧⣿⣿⣿⣿⣿⣿⣶⣥⣍⡛⠾⣬⡹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣿⣿⣿⣽⣿⣿⣿⣿⣿⣿⣿⣷⠟⣩⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣷⣿⣿⢿⣹⣿⣿⣿⣿⣿⣿⣻⣿⣶⣈⠍⡙⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢠⣾⡿⢯⣿⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⡿⡙⠎⣿⣿⣿⣿⣿⣿⣿⣗⢮⠽⣿⢦⣈⠂⠩⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠿⡿⡩⠛⣿⣿⣿⣿⣿⡿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣻⣷⡀⠀⢿⣿⣻⣿⣿⣟⣿⡏⠁⠛⠽⣍⠤⠘⠄⡂⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⠞⠁⠀⠘⣿⠛⢼⡟⣯⡽⣿⣿⣿")
			fmt.Println(padding + "	⠀⠆⡽⣷⣄⠘⢦⡱⢯⠷⣞⣹⠃⠀⠀⢀⡼⢐⡉⠰⣀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣠⠑⠤⣀⠀⠈⠻⢼⣽⠶⠋⣿⣿⡿")
			fmt.Println(padding + "	⠀⣏⠰⢡⠟⣧⣄⣙⣟⣛⣚⣁⣠⠤⡒⠭⡐⢡⠠⣱⣾⣿⣿⣿⢟⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠤⣉⠰⢀⠏⢓⡒⢖⡲⢚⢻⣿⡟⢹")
			fmt.Println(padding + "	⠀⣿⡌⢣⡘⣄⢫⠙⡬⣁⢃⠖⠤⠓⣈⠔⡠⠃⣵⡿⠿⠛⠫⡐⢊⡑⢎⠿⣿⣿⣿⣿⣿⣿⣿⣿⠐⡡⠂⠆⡌⠜⡰⢈⠤⡁⠇⢺⠋⡔⣫")
			fmt.Println(padding + "	⠀⣿⣯⢡⠒⣌⠢⣙⡐⢢⢉⠒⣌⠱⣀⠚⡄⢩⠁⡔⠨⡁⢃⠡⢂⠔⡈⠲⢥⢛⣿⣿⣿⣿⣿⣿⣧⠡⢃⡘⠰⣁⠲⣁⢒⠩⠜⡠⢋⠔⣿")
			fmt.Println(padding + "	⠀⣿⣿⠣⡍⠤⢃⡔⢌⠂⢎⠒⡌⢒⠠⡑⣈⠆⡑⡈⢆⡑⡈⠆⠌⡒⢨⠑⡄⢊⠔⡹⠿⣿⣿⣿⣿⣷⡡⠄⢃⠄⢣⠐⡌⠱⡈⡔⢣⢺⢛")
			fmt.Println(padding + "	⠀⣿⣿⣷⠩⡜⢡⡘⢤⡉⠦⡑⢌⠢⡑⢆⠡⡘⠤⣑⠢⡘⠤⣉⢒⡉⠦⡙⡄⢣⠘⣀⠓⣌⠹⢻⢿⣿⣷⡜⠂⢎⢡⢊⡔⢣⡐⡡⢆⢣⣾")
			fmt.Println(padding + "	⠀⣿⣿⣿⣷⡘⢥⡘⠤⡘⠱⡈⠬⢡⠑⡊⠔⡡⢓⠠⢓⠠⡓⢤⠒⠬⡑⣡⠊⡔⣽⣧⡘⢄⠣⢌⠢⡐⢍⠛⠯⡐⣌⠢⡘⢄⠆⡱⢨⣼⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⡰⢌⢣⠘⡥⢌⡑⢢⠥⢱⠨⡑⠌⢆⢣⢣⢉⠆⣩⢒⡑⠦⡱⢌⡷⣞⣻⠎⡜⢢⡑⡜⢢⡙⢢⠱⣀⢣⠘⡄⢎⣑⣾⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⣿⣎⢆⠣⡜⠢⠜⡡⢎⡑⢢⠙⡜⡌⣒⠢⣍⠚⢤⠣⡜⣡⢃⠆⣿⣽⢇⠚⡌⣑⠢⡑⢦⠘⡥⢚⠤⠣⠜⡰⢲⣾⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⣿⣿⣿⣴⢡⠣⡙⡔⠢⢍⠢⢍⠒⠬⣐⠱⣈⠹⣀⠳⣨⢑⣊⠱⣌⠒⡌⢣⡘⢄⢣⡙⡄⠓⣌⠡⢎⡱⢊⣵⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣧⡱⢌⡃⢎⡱⢌⡍⢲⠡⡓⢬⡑⢌⠳⡐⢎⡰⢣⢌⡓⣌⠣⡘⢌⢆⠲⣉⠱⢂⡍⢆⣵⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣹⡐⠲⡌⣌⢃⡓⢬⢦⣘⣌⣡⣍⣒⣥⣓⣌⣲⢬⠦⠵⣊⠬⠓⡌⢥⢣⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println(padding + "	⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣧⣱⣌⣆⣍⣒⣦⣑⣊⣵⣌⣲⣐⣆⣲⣔⣊⣱⣃⡜⣬⣙⣼⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿")
			fmt.Println()
			fmt.Println(padding + "     THANK YOU:)  by: Cheng Ho and Zeifira")
			return
		default:
			fmt.Println( "Pilihan tidak valid.")
		}
		fmt.Println(padding + "  Tekan ENTER ")
		fmt.Scanln()
	}
}
func tampilkanMenu() {									
	// Menampilkan seluruh menu.
	padding := "                                                " 
	fmt.Println()
	fmt.Println(padding + "╔══════════════════════════════════════════════════╗")
	fmt.Println(padding + "║         MANAJEMEN DATA PEGAWAI - MENU UTAMA      ║")
	fmt.Println(padding + "╠══════════════════════════════════════════════════╣")
	fmt.Println(padding + "║ 1. Tambah Data Pegawai                           ║")
	fmt.Println(padding + "║ 2. Tampilkan Data Pegawai                        ║")
	fmt.Println(padding + "║ 3. Edit Data Pegawai (Sequential Search)         ║")
	fmt.Println(padding + "║ 4. Hapus Data Pegawai (Sequential Search)        ║")
	fmt.Println(padding + "║ 5. Urutkan Gaji Tertinggi (Selection Sort)       ║")
	fmt.Println(padding + "║ 6. Urutkan Gaji Terendah (Insertion Sort)        ║")
	fmt.Println(padding + "║ 7. Total gaji semua pegawai                      ║")
	fmt.Println(padding + "║ 8. Cari Pegawai Berdasarkan Nama (Binary)        ║")
	fmt.Println(padding + "║ 9. Keluar                                        ║")
	fmt.Println(padding + "╚══════════════════════════════════════════════════╝")
	fmt.Print(padding + ">> Pilih menu [1-9]: ")
}

func clearScreen() {
	// Function untuk membersihkan layar terminal.
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func tambahData(arr *[maxPegawai]Pegawai, n *int) {
	// Fungsi ini digunakan untuk menambahkan 1 data pegawai ke dalam array arr jika jumlah data (*n) belum mencapai batas maxPegawai.
	// I.S. Data pegawai mungkin belum penuh.
	// F.S. Jika belum penuh, satu data pegawai ditambahkan ke array.
	padding := "                                              "
	if *n >= maxPegawai { 
		fmt.Println(padding + "Data pegawai penuh.")
		return
	}
	var p Pegawai
	fmt.Print(padding + "Masukkan ID (angka): ")
	fmt.Scan(&p.ID)
	fmt.Print(padding + "Masukkan Nama (tanpa spasi): ")
	fmt.Scan(&p.Nama)
	fmt.Print(padding + "Masukkan Shift (pagi/siang/malam): ")
	fmt.Scan(&p.Shift)
	fmt.Print(padding + "Masukkan Gaji (angka): ")
	fmt.Scan(&p.Gaji)

	arr[*n] = p
	*n++
	fmt.Println(padding + "Data berhasil ditambahkan.")
}

func tampilkanData(arr [maxPegawai]Pegawai, n int) {
	// Prosedur menampilkan seluruh data yang ada dalam array.
	// I.S. Data pegawai bisa kosong atau terisi.
	// F.S. Jika ada, seluruh data pegawai ditampilkan dalam bentuk tabel ASCII.
	padding := "                                              " 
	if n == 0 {
		fmt.Println(padding + "Belum ada data pegawai.")
		return
	}
	fmt.Println(padding + "╔══════╦════════════════════╦════════════╦════════════╗")
	fmt.Println(padding + "║ ID   ║ Nama               ║ Shift      ║ Gaji       ║")
	fmt.Println(padding + "╠══════╬════════════════════╬════════════╬════════════╣")
	for i := 0; i < n; i++ {
		fmt.Printf(padding + "║ %-4d ║ %-18s ║ %-10s ║ %-10d ║\n", arr[i].ID, arr[i].Nama, arr[i].Shift, arr[i].Gaji)
	}
	fmt.Println(padding + "╚══════╩════════════════════╩════════════╩════════════╝")
}

func editData(arr *[maxPegawai]Pegawai, n int) {
	// Mengedit data pegawai di dalam array berdasarkan ID pegawai, menggunakan Sequential Search.
	// I.S. Data pegawai sudah ada dan ID yang dicari diketahui.
	// F.S. Jika ditemukan, data pegawai dengan ID tersebut diperbarui.
	var targetID int
	padding := "                                              " 
	fmt.Print(padding + "Masukkan ID pegawai yang akan diedit: ")
	fmt.Scan(&targetID)

	pos := -1
	for i := 0; i < n && pos == -1; i++ {
		if arr[i].ID == targetID {
			pos = i
		}
	}

	if pos == -1 {
		fmt.Println(padding + "Pegawai tidak ditemukan.")
		return
	}
	fmt.Println(padding + "Data ditemukan:")
	fmt.Println(padding + "╔══════╦════════════════════╦════════════╦════════════╗")
	fmt.Println(padding + "║ ID   ║ Nama               ║ Shift      ║ Gaji       ║")
	fmt.Println(padding + "╠══════╬════════════════════╬════════════╬════════════╣")
	fmt.Printf(padding + "║ %-4d ║ %-18s ║ %-10s ║ %-10d ║\n", arr[pos].ID, arr[pos].Nama, arr[pos].Shift, arr[pos].Gaji)
	fmt.Println(padding + "╚══════╩════════════════════╩════════════╩════════════╝")

	fmt.Print(padding + "Masukkan Nama baru: ")
	fmt.Scan(&arr[pos].Nama)
	fmt.Print(padding + "Masukkan Shift baru: ")
	fmt.Scan(&arr[pos].Shift)
	fmt.Print(padding + "Masukkan Gaji baru: ")
	fmt.Scan(&arr[pos].Gaji)

	fmt.Println(padding + "Data berhasil diperbarui.")
}

func hapusData(arr *[maxPegawai]Pegawai, n *int) {
	// Menghapus data pegawai di dalam array berdasarkan ID pegawai, menggunkaan Sequential Search.
	// I.S. Data pegawai sudah ada dan ID yang dicari diketahui.
	// F.S. Jika ditemukan, data pegawai dengan ID tersebut dihapus dan array digeser.
	var targetID int
	padding := "                                              " 
	fmt.Print(padding + "Masukkan ID pegawai yang akan dihapus: ")
	fmt.Scan(&targetID)

	pos := -1
	for i := 0; i < *n && pos == -1; i++ {
		if arr[i].ID == targetID {
			pos = i
		}
	}
	if pos == -1 {
		fmt.Println(padding + "Pegawai tidak ditemukan.")
		return
	}
	fmt.Println(padding + "Data yang dihapus:")
	fmt.Println(padding + "╔══════╦════════════════════╦════════════╦════════════╗")
	fmt.Println(padding + "║ ID   ║ Nama               ║ Shift      ║ Gaji       ║")
	fmt.Println(padding + "╠══════╬════════════════════╬════════════╬════════════╣")
	fmt.Printf(padding + "║ %-4d ║ %-18s ║ %-10s ║ %-10d ║\n", arr[pos].ID, arr[pos].Nama, arr[pos].Shift, arr[pos].Gaji)
	fmt.Println(padding + "╚══════╩════════════════════╩════════════╩════════════╝")

	for i := pos; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
	fmt.Println(padding + "Data berhasil dihapus.")
}

func urutGajiTertinggi(arr *[maxPegawai]Pegawai, n int) {
	// I.S. Data pegawai tidak terurut berdasarkan gaji.
	// F.S. Data pegawai diurutkan descending berdasarkan gaji menggunakan selection sort.
	// Mengurutkan data di dalam array dengan algoritma selection sort secara descending.
	var pass, idx, i int
	var temp Pegawai

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if arr[idx].Gaji < arr[i].Gaji {
				idx = i
			}
			i++
		}
		temp = arr[pass-1]
		arr[pass-1] = arr[idx]
		arr[idx] = temp
		pass++
	}
}


func urutGajiTerendah(arr *[maxPegawai]Pegawai, n int) {
	// I.S. Data pegawai tidak terurut berdasarkan gaji.
	// F.S. Data pegawai diurutkan ascending berdasarkan gaji menggunakan insertion sort.
	// Mengurutkan data di dalam array dengan algoritma insertion sort secara ascending.
	var i, j int
	var pass Pegawai

	for i = 1; i < n; i++ {
		pass = arr[i]
		j = i - 1

		for j >= 0 && arr[j].Gaji > pass.Gaji {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = pass
	}
}

func cariPegawaiNamaBinary(arr *[maxPegawai]Pegawai, n int, nama string) int {
	// I.S. Data pegawai belum terurut berdasarkan nama.
	// F.S. Data pegawai diurutkan berdasarkan nama lalu dicari menggunakan binary search.
	
	// Sorting dengan selection sort (ascending berdasarkan Nama)
    var idx int
    var temp Pegawai
    for i := 0; i < n-1; i++ {
        idx = i
        for j := i + 1; j < n; j++ {
            if arr[j].Nama < arr[idx].Nama {
                idx = j
            }
        }
        temp = arr[i]
        arr[i] = arr[idx]
        arr[idx] = temp
    }

    // Binary search setelah sorting
    low, high := 0, n-1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid].Nama == nama {
            return mid
        } else if arr[mid].Nama < nama {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1 // tidak ditemukan
}


func totalGaji(arr [maxPegawai]Pegawai, n int) {
	//function menotalkan semua gaji pegawai di dalam array.
	// I.S. Data pegawai telah terisi sebagian atau seluruhnya.
	// F.S. Total seluruh gaji pegawai dihitung dan ditampilkan dalam format tabel.

    var total int
    for i := 0; i < n; i++ {
        total += arr[i].Gaji
    }
	
	padding := "                                              " 
    fmt.Println(padding + "╔═══════════════════════════════════════════╗")
    fmt.Printf(padding + "║ Total Gaji Seluruh Pegawai: Rp %-10d ║\n", total)
    fmt.Println(padding + "╚═══════════════════════════════════════════╝")
}