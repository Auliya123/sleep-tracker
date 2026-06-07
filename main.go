package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NMAX = 100

type DataTidur struct {
	tglTidur, jamTidur, tglBangun, jamBangun, status string
	durasi                                           int
}

var data [NMAX]DataTidur
var jumlahData int = 0

var reader = bufio.NewReader(os.Stdin)

func inputString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func validTanggal(t string) bool {
	if len(t) != 10 {
		return false
	}
	if t[2] != '-' || t[5] != '-' {
		return false
	}
	dd, err1 := strconv.Atoi(t[0:2])
	mm, err2 := strconv.Atoi(t[3:5])
	yyyy, err3 := strconv.Atoi(t[6:10])
	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}
	return dd >= 1 && dd <= 31 && mm >= 1 && mm <= 12 && yyyy >= 2000
}

func tanggalKeAngka(tgl string) int {
	var dd, mm, yyyy int
	fmt.Sscanf(tgl, "%d-%d-%d", &dd, &mm, &yyyy)

	return yyyy*10000 + mm*100 + dd
}

func validJam(j string) bool {
	if len(j) != 5 {
		return false
	}
	if j[2] != ':' {
		return false
	}
	hh, err1 := strconv.Atoi(j[0:2])
	mm, err2 := strconv.Atoi(j[3:5])
	if err1 != nil || err2 != nil {
		return false
	}
	return hh >= 0 && hh <= 23 && mm >= 0 && mm <= 59
}

func tanggalJamKeMenit(tanggal, jam string) int {
	var dd, mm, yyyy int
	var hh, menit int

	fmt.Sscanf(tanggal, "%d-%d-%d", &dd, &mm, &yyyy)
	fmt.Sscanf(jam, "%d:%d", &hh, &menit)

	// cukup untuk perbandingan
	return (((yyyy*12+mm)*31+dd)*24+hh)*60 + menit
}

func hitungDurasi(jamTidur, jamBangun string) int {
	var jt, mt, jb, mb int
	fmt.Sscanf(jamTidur, "%d:%d", &jt, &mt)
	fmt.Sscanf(jamBangun, "%d:%d", &jb, &mb)
	menitTidur := jt*60 + mt
	menitBangun := jb*60 + mb
	durasi := menitBangun - menitTidur
	if durasi < 0 {
		durasi += 24 * 60
	}
	return durasi
}

func tentukanStatus(durasi int) string {
	jam := float64(durasi) / 60

	if jam < 6 {
		return "Sangat Kurang"
	} else if jam < 7 {
		return "Kurang"
	} else if jam <= 9 {
		return "Ideal"
	} else {
		return "Berlebih"
	}
}

func tampilSemua() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Println("\nDATA TIDUR")
	fmt.Printf("%-4s| %-13s| %-11s| %-13s| %-11s| %-14s| %s\n", "No", "Tgl Tidur", "Jam Tidur", "Tgl Bangun", "Jam Bangun", "Durasi", "Status")

	for i := 0; i < jumlahData; i++ {
		d := data[i]
		jam := d.durasi / 60
		menit := d.durasi % 60
		durasi := fmt.Sprintf("%dj %dm", jam, menit)
		fmt.Printf("%-4d| %-12s| %-10s| %-12s| %-10s| %-12s| %s\n", i+1, d.tglTidur, d.jamTidur, d.tglBangun, d.jamBangun, durasi, d.status)
	}
}

func tampilByIndex(idx int) {
	if jumlahData == 0 {
		fmt.Println("Belum ada data")
		return
	}

	if idx < 0 || idx >= jumlahData {
		fmt.Println("Index tidak valid")
		return
	}

	d := data[idx]
	jam := d.durasi / 60
	menit := d.durasi % 60

	fmt.Println("\nDATA TIDUR")
	fmt.Println("Index        :", idx)
	fmt.Println("Tanggal Tidur:", d.tglTidur)
	fmt.Println("Jam Tidur    :", d.jamTidur)
	fmt.Println("Tanggal Bangun:", d.tglBangun)
	fmt.Println("Jam Bangun   :", d.jamBangun)
	fmt.Printf("Durasi       : %d jam %d menit\n", jam, menit)
	fmt.Println("Status       :", d.status)
}

// ================= MENU 1 =================

func tambahData() {
	var n int
	fmt.Print("Mau input berapa data? ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("\nData ke-", i+1)

		var tglTidur, tglBangun string
		var jamTidur, jamBangun string
		var durasi int

		// input tanggal tidur
		valid := false
		for !valid {
			tglTidur = inputString("Tanggal tidur (dd-mm-yyyy): ")

			if validTanggal(tglTidur) {
				valid = true
			} else {
				fmt.Println("Format salah!")
			}
		}

		// input jam tidur
		valid = false
		for !valid {
			jamTidur = inputString("Jam tidur (HH:MM): ")

			if validJam(jamTidur) {
				valid = true
			} else {
				fmt.Println("Format salah!")
			}
		}

		// input tanggal bangun
		valid = false

		for !valid {

			// input tanggal bangun
			tglBangun = inputString("Tanggal bangun (dd-mm-yyyy): ")

			if !validTanggal(tglBangun) {
				fmt.Println("Format salah!")
				continue
			}

			// input jam bangun
			jamBangun = inputString("Jam bangun (HH:MM): ")

			if !validJam(jamBangun) {
				fmt.Println("Format salah!")
				continue
			}

			waktuTidur := tanggalJamKeMenit(tglTidur, jamTidur)
			waktuBangun := tanggalJamKeMenit(tglBangun, jamBangun)

			if waktuBangun <= waktuTidur {
				fmt.Println("\nTanggal dan jam bangun harus setelah tanggal dan jam tidur!")
				fmt.Println("Silakan input ulang tanggal dan jam bangun.")
				continue
			}

			durasi = (waktuBangun - waktuTidur)
			valid = true
		}

		status := tentukanStatus(durasi)

		data[jumlahData] = DataTidur{
			tglTidur:  tglTidur,
			jamTidur:  jamTidur,
			tglBangun: tglBangun,
			jamBangun: jamBangun,
			durasi:    durasi,
			status:    status,
		}

		jumlahData++
	}

	fmt.Println("Data berhasil disimpan!")
	tampilSemua()
}

// ================= EDIT =================

func editData(idx int) {
	fmt.Println("Edit data:")

	var tglTidur, tglBangun string
	var jamTidur, jamBangun string
	var durasi int
	var valid bool

	// Validasi tanggal tidur
	valid = false
	for !valid {
		tglTidur = inputString("Tanggal tidur (dd-mm-yyyy): ")

		if validTanggal(tglTidur) {
			valid = true
		} else {
			fmt.Println("Format salah!")
		}
	}

	// Validasi jam tidur
	valid = false
	for !valid {
		jamTidur = inputString("Jam tidur (HH:MM): ")

		if validJam(jamTidur) {
			valid = true
		} else {
			fmt.Println("Format salah!")
		}
	}

	// Validasi tanggal bangun
	valid = false
	for !valid {
		tglBangun = inputString("Tanggal bangun (dd-mm-yyyy): ")

		if validTanggal(tglBangun) {
			valid = true
		} else {
			fmt.Println("Format salah!")
		}
	}

	// Validasi jam bangun + durasi
	valid = false
	for !valid {
		jamBangun = inputString("Jam bangun (HH:MM): ")

		if validJam(jamBangun) {

			durasi = hitungDurasi(jamTidur, jamBangun)

			if durasi > 0 {
				valid = true
			} else {
				fmt.Println("Jam bangun harus setelah jam tidur!")
			}

		} else {
			fmt.Println("Format salah!")
		}
	}

	// Hitung status
	status := tentukanStatus(durasi)

	// Update data
	data[idx].tglTidur = tglTidur
	data[idx].tglBangun = tglBangun
	data[idx].jamTidur = jamTidur
	data[idx].jamBangun = jamBangun
	data[idx].durasi = durasi
	data[idx].status = status

	fmt.Println("Data berhasil diupdate!")

	tampilSemua()
}

// ================= DELETE =================

func hapusData(idx int) {
	for i := idx; i < jumlahData-1; i++ {
		data[i] = data[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus!")
	tampilSemua()
}

// searching - okta
// sequentialSearch mencari data berdasarkan tanggal tidur
func sequentialSearch(tgl string) int {
	idx := -1
	i := 0
	ketemu := false
	for i < jumlahData && !ketemu {
		if data[i].tglTidur == tgl {
			idx = i
			ketemu = true
		}
		i++
	}
	return idx
}

// mengurutkan salinan array by tanggal, khusus untuk binary search
func sortByTanggalAscUntukBS(arr [NMAX]DataTidur, n int) [NMAX]DataTidur {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if tanggalKeAngka(arr[j].tglTidur) <
				tanggalKeAngka(arr[minIdx].tglTidur) {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// binarySearch mencari data berdasarkan tanggal tidur
func binarySearch(tgl string) int {
	sortedArr := sortByTanggalAscUntukBS(data, jumlahData)

	low := 0
	high := jumlahData - 1
	idxSorted := -1

	for low <= high && idxSorted == -1 {
		mid := (low + high) / 2
		target := tanggalKeAngka(tgl)
		midVal := tanggalKeAngka(sortedArr[mid].tglTidur)

		if midVal == target {
			idxSorted = mid
		} else if target < midVal {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	idx := -1
	if idxSorted != -1 {
		i := 0
		ketemu := false
		for i < jumlahData && !ketemu {
			if data[i].tglTidur == sortedArr[idxSorted].tglTidur {
				idx = i
				ketemu = true
			}
			i++
		}
	}
	return idx
}

// sorting - okta
// mengurutkan salinan array berdasarkan kriteria dan arah yang dipilih
func selectionSort(arr [NMAX]DataTidur, n int, byDurasi bool, asc bool) [NMAX]DataTidur {
	for i := 0; i < n-1; i++ {
		idxPilih := i
		for j := i + 1; j < n; j++ {
			lebihPrioritas := false
			if byDurasi {
				if asc {
					lebihPrioritas = arr[j].durasi < arr[idxPilih].durasi
				} else {
					lebihPrioritas = arr[j].durasi > arr[idxPilih].durasi
				}
			} else {
				if asc {
					lebihPrioritas = tanggalKeAngka(arr[j].tglTidur) < tanggalKeAngka(arr[idxPilih].tglTidur)
				} else {
					lebihPrioritas = tanggalKeAngka(arr[j].tglTidur) > tanggalKeAngka(arr[idxPilih].tglTidur)
				}
			}
			if lebihPrioritas {
				idxPilih = j
			}
		}
		arr[i], arr[idxPilih] = arr[idxPilih], arr[i]
	}
	return arr
}

// mengurutkan salinan array dengan cara menyisipkan elemen ke posisi yang tepat
func insertionSort(arr [NMAX]DataTidur, n int, byDurasi bool, asc bool) [NMAX]DataTidur {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		geser := true
		for j >= 0 && geser {
			harusGeser := false
			if byDurasi {
				if asc {
					harusGeser = arr[j].durasi > key.durasi
				} else {
					harusGeser = arr[j].durasi < key.durasi
				}
			} else {
				if asc {
					harusGeser = tanggalKeAngka(arr[j].tglTidur) > tanggalKeAngka(key.tglTidur)
				} else {
					harusGeser = tanggalKeAngka(arr[j].tglTidur) < tanggalKeAngka(key.tglTidur)
				}
			}
			if harusGeser {
				arr[j+1] = arr[j]
				j--
			} else {
				geser = false
			}
		}
		arr[j+1] = key
	}
	return arr
}

// menu 2 - okta

func menu2() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Println("\nUrutkan berdasarkan:")
	fmt.Println("1. Tanggal tidur")
	fmt.Println("2. Durasi tidur")
	var pilihKriteria int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihKriteria)
	reader.ReadString('\n')

	if pilihKriteria != 1 && pilihKriteria != 2 {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("\nUrutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	var pilihUrutan int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihUrutan)
	reader.ReadString('\n')

	if pilihUrutan != 1 && pilihUrutan != 2 {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("\nAlgoritma sorting:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	var pilihAlgo int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihAlgo)
	reader.ReadString('\n')

	if pilihAlgo != 1 && pilihAlgo != 2 {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	byDurasi := pilihKriteria == 2
	asc := pilihUrutan == 1

	// sorting dilakukan di salinan, array asli tidak akan berubah
	sortedArr := data
	if pilihAlgo == 1 {
		sortedArr = selectionSort(sortedArr, jumlahData, byDurasi, asc)
	} else {
		sortedArr = insertionSort(sortedArr, jumlahData, byDurasi, asc)
	}

	kriteria := "Tanggal"
	if byDurasi {
		kriteria = "Durasi"
	}
	urutan := "Ascending"
	if !asc {
		urutan = "Descending"
	}
	algo := "Selection Sort"
	if pilihAlgo == 2 {
		algo = "Insertion Sort"
	}

	fmt.Println()
	fmt.Printf("Kriteria: %s | Urutan: %s | Algoritma: %s\n", kriteria, urutan, algo)
	fmt.Println("No | Tgl Tidur  | Jam Tidur | Tgl Bangun | Jam Bangun | Durasi      | Status")
	fmt.Println("------------------------------------------------------------------------------")
	for i := 0; i < jumlahData; i++ {
		d := sortedArr[i]
		jam := d.durasi / 60
		menit := d.durasi % 60
		fmt.Printf("%2d | %s | %s     | %s  | %s      | %2dj %2dm    | %s\n",
			i+1, d.tglTidur, d.jamTidur, d.tglBangun, d.jamBangun, jam, menit, d.status)
	}
}

// ================= MENU 3 =================

func menu3() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data")
		return
	}
	tgl := inputString("Masukkan tanggal tidur yang dicari (dd-mm-yyyy): ")

	if !validTanggal(tgl) {
		fmt.Println("Format tanggal tidak valid!")
		return
	}

	fmt.Println("\nMetode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	var pilihCari int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihCari)
	reader.ReadString('\n')

	idx := -1
	if pilihCari == 1 {
		idx = sequentialSearch(tgl)
	} else if pilihCari == 2 {
		idx = binarySearch(tgl)
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	if idx == -1 {
		fmt.Println("Data tanggal", tgl, "tidak ditemukan.")
		return
	}

	fmt.Println("Data ditemukan:")
	tampilByIndex(idx)

	fmt.Println("\nApa yang ingin dilakukan?")
	fmt.Println("1. Edit")
	fmt.Println("2. Hapus")
	fmt.Println("3. Kembali")

	var pilih int
	fmt.Print("Mau pilih pilihan yang mana: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		editData(idx)
	} else if pilih == 2 {
		hapusData(idx)
	}
}

// menu 4 - okta
func menu4() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	batas := jumlahData - 7
	if batas < 0 {
		batas = 0
	}
	n := jumlahData - batas

	fmt.Println("\nLaporan 7 hari terakhir")
	fmt.Println()
	fmt.Println("No | Tgl Tidur  | Tidur | Bangun | Durasi      | Status")
	fmt.Println("-----------------------------------------------------------")

	totalDurasi := 0
	hariIdeal := 0
	hariKurang := 0
	hariSangatKurang := 0
	hariLebih := 0
	countLarut := 0

	for i := batas; i < jumlahData; i++ {
		d := data[i]
		jam := d.durasi / 60
		menit := d.durasi % 60
		fmt.Printf("%2d | %s | %s  | %s  | %2dj %2dm      | %s\n",
			i-batas+1, d.tglTidur, d.jamTidur, d.jamBangun, jam, menit, d.status)

		totalDurasi += d.durasi

		if d.status == "Ideal" {
			hariIdeal++
		} else if d.status == "Kurang" {
			hariKurang++
		} else if d.status == "Sangat Kurang" {
			hariSangatKurang++
		} else if d.status == "Berlebih" {
			hariLebih++
		}

		if d.jamTidur >= "23:00" || d.jamTidur < "06:00" {
			countLarut++
		}
	}

	rataRata := totalDurasi / n
	rataJam := rataRata / 60
	rataMenit := rataRata % 60
	rataJamFloat := float64(totalDurasi) / 60.0 / float64(n)

	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("Rata-rata durasi   : %d jam %d menit\n", rataJam, rataMenit)
	fmt.Printf("Hari ideal         : %d hari\n", hariIdeal)
	fmt.Printf("Hari kurang        : %d hari\n", hariKurang)
	fmt.Printf("Hari sangat kurang : %d hari\n", hariSangatKurang)
	fmt.Printf("Hari berlebih      : %d hari\n", hariLebih)

	fmt.Println("\nRekomendasi:")
	fmt.Println()

	if rataJamFloat < 6 {
		fmt.Println("- Tidurmu sangat kurang! Orang dewasa tuh minimal butuh 7 jam loh.")
		fmt.Println("  Jangan sering begadang ya, resikonya bisa nurunin konsentrasi sama daya tahan tubuh.")
	} else if rataJamFloat < 7 {
		fmt.Println("- Tidurmu dikit lagi menyentuh standar ideal nih.")
		fmt.Println("  Yuk, coba ditambah 30-60 menit lagi waktu tidurnya tiap malam.")
	} else if rataJamFloat <= 9 {
		fmt.Println("- Rata-rata jam tidurmu udah ideal banget (7-9 jam). Mantap, pertahankan!")
	} else {
		fmt.Println("- Durasi tidurmu terlalu lama nih.")
		fmt.Println("  Tidur berlebihan juga kurang baik dan bisa menyebabkan rasa lemas.")
	}

	if countLarut >= 3 {
		fmt.Println("- Minggu ini kamu sering banget tidur larut malam.")
		fmt.Println("  Usahakan mulai tidur sebelum jam 23:00 ya, biar tidurnya lebih berkualitas.")
	}

	if hariSangatKurang >= 2 {
		fmt.Println("- Ada beberapa malam yang jam tidurmu itu kurang banget.")
		fmt.Println("  Kurang-kurangin begadang yang ga perlu, utamakan istirahat dulu ya!")
	}

	if hariIdeal == n {
		fmt.Println("- Keren banget! Minggu ini semua jadwal tidurmu ideal. Luar biasa, pertahankan!")
	}
}

// ================= MENU =================

func tampilMenu() {
	fmt.Println("\nAPLIKASI PEMANTAUAN POLA TIDUR")
	fmt.Println()
	fmt.Println("1. Tambah riwayat tidur")
	fmt.Println("2. Lihat semua riwayat & sorting")
	fmt.Println("3. Cari & kelola data (edit/hapus)")
	fmt.Println("4. Laporan mingguan & rekomendasi")
	fmt.Println("0. Keluar")
	fmt.Println()
	fmt.Print("Pilih menu: ")
}

func main() {
	selesai := false
	for !selesai {
		tampilMenu()

		var pilih int
		fmt.Scan(&pilih)
		reader.ReadString('\n')

		if pilih == 0 {
			fmt.Println("\nTerima kasih! Sampai jumpa lagi.")
			selesai = true
		} else if pilih == 1 {
			tambahData()
		} else if pilih == 2 {
			menu2()
		} else if pilih == 3 {
			menu3()
		} else if pilih == 4 {
			menu4()
		} else {
			fmt.Println("Pilihan tidak valid! Masukkan angka 0-4.")
		}
	}
}
