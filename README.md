# asisten-belajar-go
Aplikasi AI Asisten Belajar berbasis Go (Golang) yang memiliki fitur manajemen catatan, pencarian, pengurutan, dan planner waktu

# 🧠 AI Asisten Belajar (Golang)

![Go](https://img.shields.io/badge/Language-Go-blue)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen)
![Made by](https://img.shields.io/badge/Made%20by-Muhamad%20Indal%20Fauzan%20Az-red)

Aplikasi terminal berbasis **Go (Golang)** untuk membantu mahasiswa mencatat, mengelola, mencari, dan menjadwalkan waktu belajar secara efisien.  
Dikembangkan dengan prinsip *structured programming* untuk menerapkan dasar-dasar algoritma dan manajemen data dalam konteks nyata.

---

## ✨ Fitur Utama

### 📘 Manajemen Catatan
- Tambah, ubah, tampilkan, dan hapus catatan belajar.
- Menampilkan ringkasan otomatis dari setiap catatan.

### 🔍 Pencarian Cerdas
- **Sequential Search** untuk mencari berdasarkan *topik*.
- **Binary Search** untuk mencari berdasarkan *tanggal catatan*.

### 📊 Pengurutan Data
- **Selection Sort:** Mengurutkan catatan berdasarkan tingkat kesulitan.
- **Insertion Sort:** Mengurutkan catatan berdasarkan tanggal.

### 🗓️ Smart Planner
- Menambahkan waktu luang dan aktivitas belajar.
- Menyusun jadwal otomatis berdasarkan durasi dan ketersediaan waktu.
- Menampilkan sisa waktu kosong dan aktivitas yang terjadwal.

---

## ⚙️ Teknologi
- **Bahasa:** Go (Golang)  
- **Paradigma:** Structured Programming  
- **Antarmuka:** Command Line Interface (CLI)  
- **Struktur Data:** Array dan Record (Struct)  

---

## 🧩 Cara Menjalankan
Pastikan Go sudah terinstal di komputer kamu.

```bash
# Clone repository
git clone https://github.com/muhamadindalfauzanaz/asisten-belajar-go.git

# Masuk ke folder proyek
cd asisten-belajar-go

# Jalankan program
go run asisten.go
