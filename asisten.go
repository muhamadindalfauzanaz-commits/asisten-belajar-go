package main

import "fmt"

const NMAX int = 100

// untuk Manajemen Catatan, menu 1-3
type Note struct {
	Id        int
	Topik     string
	Konten    string
	Tanggal   string
	Kesulitan int
}

var notes [NMAX]Note
var noteCount int
var nextId int = 1

// untuk Smart Planner, menu 4
type Slot struct {
	Hari          string
	Mulai         string
	Selesai       string
	Menit         int
	Terisi        bool
	Topik         string
	MenitTerpakai int
}

type Aktivitas struct {
	Nama   string
	Durasi int
	Sisa   int
}

var slots [NMAX]Slot
var slotCount int

var aktivitas [NMAX]Aktivitas
var aktivitasCount int

func main() {
	var pilih int
	for {
		menuUtama()
		fmt.Scanln(&pilih)
		switch pilih {
		case 1:
			menuCatatan()
		case 2:
			menuPencarian()
		case 3:
			menuPengurutan()
		case 4:
			menuPlanner()
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUtama() {
	fmt.Println("------------------------------")
	fmt.Println(" Aplikasi AI Asisten Belajar ")
	fmt.Println("------------------------------")
	fmt.Println("1. Manajemen Catatan")
	fmt.Println("2. Pencarian Materi")
	fmt.Println("3. Pengurutan Catatan")
	fmt.Println("4. Smart Planner")
	fmt.Println("0. Keluar")
	fmt.Println("------------------------------")
	fmt.Print("Pilih menu: ")
}

func menuCatatan() {
	var pilih int
	for {
		fmt.Println("\n=== MENU CATATAN ===")
		fmt.Println("1. Tambah Catatan")
		fmt.Println("2. Tampilkan Catatan")
		fmt.Println("3. Ubah Catatan")
		fmt.Println("4. Hapus Catatan")
		fmt.Println("5. Tampilkan Ringkasan Catatan")
		fmt.Println("0. Kembali")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahCatatan()
		case 2:
			tampilkanCatatan()
		case 3:
			ubahCatatan()
		case 4:
			hapusCatatan()
		case 5:
			ringkasanCatatan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPencarian() {
	var pilih int
	fmt.Println("\n=== PENCARIAN CATATAN ===")
	fmt.Println("1. Topik (Sequiential Search)")
	fmt.Println("2. Tanggal (Binary Search)")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih berdasarkan: ")
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		pencarianSequential()
	case 2:
		insertionSortCatatanByTanggal()
		pencarianBinaryTanggal()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func menuPengurutan() {
	var pilih int
	for {
		fmt.Println("\n=== PENGURUTAN CATATAN ===")
		fmt.Println("1. Urutkan berdasarkan Tingkat Kesulitan (Selection Sort)")
		fmt.Println("2. Urutkan berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih metode: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			selectionSortCatatanByKesulitan()
			tampilkanCatatan()
		case 2:
			insertionSortCatatanByTanggal()
			tampilkanCatatan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahCatatan() {
	if noteCount >= NMAX {
		fmt.Println("Maaf, jumlah catatan sudah maksimal.")
		return
	}

	var topik, konten, tanggal string
	var kesulitan int

	fmt.Print("Masukkan topik: ")
	fmt.Scanln(&topik)
	fmt.Print("Masukkan isi catatan (Ubah spasi menjadi '_'): ")
	fmt.Scanln(&konten)
	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&tanggal)
	fmt.Print("Masukkan tingkat kesulitan (1: Mudah, 2: Sedang, 3: Sulit): ")
	fmt.Scanln(&kesulitan)
	if kesulitan < 1 || kesulitan > 3 {
		fmt.Println("Tingkat kesulitan tidak valid. Masukkan 1, 2, atau 3.")
		return
	}

	notes[noteCount] = Note{
		Id:        nextId,
		Topik:     topik,
		Konten:    konten,
		Tanggal:   tanggal,
		Kesulitan: kesulitan,
	}

	noteCount++
	nextId++
	fmt.Println("Catatan berhasil ditambahkan.")
}

func tampilkanCatatan() {
	if noteCount == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}
	var i int
	for i < noteCount {
		Catatan(notes[i])
		i++
	}
}

func Catatan(A Note) {
	fmt.Println("------------------------------")
	fmt.Println("ID:", A.Id)
	fmt.Println("Topik:", A.Topik)
	fmt.Println("Isi:", A.Konten)
	fmt.Println("Tanggal:", A.Tanggal)
	fmt.Println("Tingkat Kesulitan:", A.Kesulitan)
	fmt.Println("------------------------------")
}

func ubahCatatan() {
	var cariTopik string
	var found bool
	var i int

	fmt.Print("Masukkan topik catatan yang ingin diubah: ")
	fmt.Scanln(&cariTopik)

	for i < noteCount && !found {
		if notes[i].Topik == cariTopik {
			fmt.Print("Masukkan isi baru: ")
			fmt.Scanln(&notes[i].Konten)
			fmt.Println("Catatan berhasil diubah.")
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

func hapusCatatan() {
	var id, i int
	var found bool

	fmt.Print("Masukkan ID catatan yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i < noteCount && !found {
		if notes[i].Id == id {
			var j int = i
			for j < noteCount-1 {
				notes[j] = notes[j+1]
				j++
			}
			noteCount--
			fmt.Println("Catatan berhasil dihapus.")
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

func ringkasanCatatan() {
	if noteCount == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}
	fmt.Println("\nRingkasan Catatan:")
	var i int
	for i < noteCount {
		fmt.Print("Catatan ID", notes[i].Id, " (", notes[i].Topik, "): ")
		fmt.Println(ambilKalimatPertama(notes[i].Konten))
		i++
	}
}

func ambilKalimatPertama(s string) string {
	var i int
	for i < len(s) {
		if s[i] == '.' {
			return s[:i+1]
		}
		i++
	}
	return s
}

func pencarianSequential() {
	var keyword string
	var found bool
	var i int

	fmt.Print("Masukkan topik yang ingin dicari: ")
	fmt.Scanln(&keyword)

	for i < noteCount {
		if notes[i].Topik == keyword {
			Catatan(notes[i])
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

func pencarianBinaryTanggal() {
	var keyword string
	var found bool
	var mid, left, right int

	fmt.Print("Masukkan tanggal yang ingin dicari (YYYY-MM-DD): ")
	fmt.Scanln(&keyword)

	left = 0
	right = noteCount - 1

	for left <= right && !found {
		mid = (left + right) / 2
		if notes[mid].Tanggal == keyword {
			found = true
		} else if notes[mid].Tanggal < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if found {
		Catatan(notes[mid])
	} else {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

func selectionSortCatatanByKesulitan() {
	var pass, idx, i int
	var temp Note

	pass = 0
	for pass < noteCount-1 {
		idx = pass
		i = pass + 1
		for i < noteCount {
			if notes[i].Kesulitan < notes[idx].Kesulitan {
				idx = i
			}
			i++
		}
		if idx != pass {
			temp = notes[pass]
			notes[pass] = notes[idx]
			notes[idx] = temp
		}
		pass++
	}
}

func insertionSortCatatanByTanggal() {
	var i, j int

	i = 1
	for i < noteCount {
		temp := notes[i]
		j = i - 1
		for j >= 0 && notes[j].Tanggal > temp.Tanggal {
			notes[j+1] = notes[j]
			j--
		}
		notes[j+1] = temp
		i++
	}
}

func menuPlanner() {
	var pilih int
	for {
		fmt.Println("\n=== SMART PLANNER ===")
		fmt.Println("1. Tambah Waktu Luang")
		fmt.Println("2. Tambah Aktivitas")
		fmt.Println("3. Generate & Tampilkan Jadwal")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih metode: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahWaktuLuang()
		case 2:
			tambahAktivitas()
		case 3:
			generateJadwal()
			tampilkanJadwal()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahWaktuLuang() {
	var jamMulai, menitMulai, jamSelesai, menitSelesai, totalMulai, totalSelesai int

	if slotCount >= NMAX {
		fmt.Println("Slot penuh.")
		return
	}

	fmt.Print("Masukkan hari (misal: Senin): ")
	fmt.Scanln(&slots[slotCount].Hari)
	fmt.Print("Jam mulai (HH:MM): ")
	fmt.Scanln(&slots[slotCount].Mulai)
	fmt.Print("Jam selesai (HH:MM): ")
	fmt.Scanln(&slots[slotCount].Selesai)

	fmt.Sscanf(slots[slotCount].Mulai, "%d:%d", &jamMulai, &menitMulai)
	fmt.Sscanf(slots[slotCount].Selesai, "%d:%d", &jamSelesai, &menitSelesai)

	totalMulai = jamMulai*60 + menitMulai
	totalSelesai = jamSelesai*60 + menitSelesai

	slots[slotCount].Menit = totalSelesai - totalMulai
	slots[slotCount].Terisi = false
	slots[slotCount].Topik = ""
	slots[slotCount].MenitTerpakai = 0
	slotCount++

	fmt.Println("Waktu luang berhasil ditambahkan.")
}

func tambahAktivitas() {
	if aktivitasCount >= NMAX {
		fmt.Println("Aktivitas penuh.")
		return
	}

	fmt.Print("Masukkan nama aktivitas: ")
	fmt.Scanln(&aktivitas[aktivitasCount].Nama)
	fmt.Print("Durasi aktivitas (dalam menit): ")
	fmt.Scanln(&aktivitas[aktivitasCount].Durasi)
	aktivitas[aktivitasCount].Sisa = aktivitas[aktivitasCount].Durasi

	aktivitasCount++
	fmt.Println("Aktivitas berhasil ditambahkan.")
}

func generateJadwal() {
	if slotCount == 0 || aktivitasCount == 0 {
		fmt.Println("Tambahkan slot dan aktivitas terlebih dahulu.")
		return
	}

	for i := 0; i < aktivitasCount; i++ {
		j := 0
		for j < slotCount && aktivitas[i].Sisa > 0 {
			sisaLuang := slots[j].Menit - slots[j].MenitTerpakai
			if sisaLuang > 0 {
				pakai := aktivitas[i].Sisa
				if sisaLuang < aktivitas[i].Sisa {
					pakai = sisaLuang
				}

				if !slots[j].Terisi || slots[j].Topik == aktivitas[i].Nama {
					slots[j].Topik = aktivitas[i].Nama
					slots[j].Terisi = true
					slots[j].MenitTerpakai += pakai
					aktivitas[i].Sisa -= pakai
				}
			}
			j++
		}
	}
	fmt.Println("Jadwal berhasil dibuat.\n")
}

func tampilkanJadwal() {
	if slotCount == 0 {
		fmt.Println("Tidak ada slot.")
		return
	}

	for i := 0; i < aktivitasCount; i++ {
		fmt.Printf("Aktivitas : %s\n", aktivitas[i].Nama)
		ada := false
		for j := 0; j < slotCount; j++ {
			if slots[j].Topik == aktivitas[i].Nama && slots[j].MenitTerpakai > 0 {
				fmt.Printf("Hari: %s\nWaktu: %s (berlangsung %d menit)\n", slots[j].Hari, slots[j].Mulai, slots[j].MenitTerpakai)
				ada = true
			}
		}
		if !ada {
			fmt.Println("Belum dijadwalkan (tidak cukup waktu)")
		}
	}

	fmt.Println("\nSlot kosong yang belum digunakan:")
	tampilkanWaktuLuang()
}

func tampilkanWaktuLuang() {
	for i := 0; i < slotCount; i++ {
		sisa := slots[i].Menit - slots[i].MenitTerpakai
		if sisa > 0 {
			fmt.Printf("Hari: %s, Sisa waktu: %d menit\n", slots[i].Hari, sisa)
			fmt.Println("Aktivitas: (kosong)")
		}
	}
}
