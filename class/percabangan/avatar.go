package main

import (
	"fmt"
)

func main() {
	var totalPeople int
	fmt.Print("Masukkan jumlah orang yang akan berangkat: ")
	fmt.Scan(&totalPeople)

	// Konstanta kapasitas Appa
	kapasitasAppaDewasa := 5
	kapasitasAppaKecil := 2

	// Hitung jumlah Appa dewasa yang diperlukan
	jumlahAppaDewasa := totalPeople / kapasitasAppaDewasa
	sisaOrang := totalPeople % kapasitasAppaDewasa

	// Batasi jumlah Appa dewasa agar tidak lebih dari 3
	if jumlahAppaDewasa < 3{
		jumlahAppaDewasa++
	} else if jumlahAppaDewasa > 3 {
		jumlahAppaDewasa = 3
	}

	// Hitung jumlah Appa kecil yang diperlukan
	jumlahAppaKecil := (totalPeople - (jumlahAppaDewasa * kapasitasAppaDewasa) + kapasitasAppaKecil - 1) / kapasitasAppaKecil
	if jumlahAppaKecil > 5 {
		jumlahAppaKecil = 5
	} else if jumlahAppaKecil < 5 {
		jumlahAppaKecil += 1
	}

	// Tampilkan hasil
	fmt.Printf("Jumlah Appa dewasa yang digunakan: %d\n", jumlahAppaDewasa)
	fmt.Printf("Jumlah Appa kecil yang digunakan: %d\n", jumlahAppaKecil)
	fmt.Printf("Jumlah orang yang tidak bisa berangkat: %d\n", sisaOrang)
}
