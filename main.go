package main

import (
	"fmt"
	"math"
)

func main() {
	var angka1, angka2, tinggi, jari float64
	var angka int

	fmt.Println("====== Pilihan Calculator : ======")
	fmt.Println("1. Kali")
	fmt.Println("2. Bagi")
	fmt.Println("3. Tambah ")
	fmt.Println("4. Kurang ")
	fmt.Println("5. Akar")
	fmt.Println("6. Pangkat")
	fmt.Println("7. Luas Persegi")
	fmt.Println("8. Luas Lingkaran")
	fmt.Println("9. Volume Tabung")
	fmt.Println("10. Volume Balok ")
	fmt.Println("11. Volume Prisma ")
	fmt.Print("Masukkan metode perhitungan yang diinginkan : 	")
	fmt.Scan(&angka)

	switch angka {
	case 1:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		kali := angka1 * angka2
		fmt.Println(angka1, " x ", angka2, " = ", kali)
	case 2:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		bagi := angka1 / angka2
		fmt.Println(angka1, " : ", angka2, " = ", bagi)
	case 3:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		tambah := angka1 + angka2
		fmt.Println(angka1, " + ", angka2, " = ", tambah)
	case 4:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		kurang := angka1 - angka2
		fmt.Println(angka1, " / ", angka2, " = ", kurang)
	case 5:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		akar := math.Sqrt(angka1)
		fmt.Println(" akar ", angka1, " = ", akar)
	case 6:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		pangkat := math.Pow(angka1, angka2)
		fmt.Println(angka1, " pangkat ", angka2, " = ", pangkat)
	case 7:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		persegi := angka1 * angka1
		fmt.Println("Luas persegi =  ", persegi)
	case 8:
		fmt.Print("Masukkan jari-jari = ")
		fmt.Scan(&jari)
		lingkaran := math.Pi * math.Pow(jari, 2)
		fmt.Println("Luas lingkaran dengan jari-jari ", jari, " = ", lingkaran)
	case 9:
		fmt.Print("Masukkan tinggi = ")
		fmt.Scan(&tinggi)
		tabung := math.Pi * math.Pow(angka1, 2) * tinggi
		fmt.Println("Volume tabung = ", tabung)
	case 10:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		fmt.Print("Masukkan tinggi = ")
		fmt.Scan(&tinggi)
		balok := angka1 * angka2 * tinggi
		fmt.Println("Volum balok = ", balok)
	case 11:
		fmt.Print("Masukkan angka1 = ")
		fmt.Scan(&angka1)
		fmt.Print("Masukkan angka2 = ")
		fmt.Scan(&angka2)
		fmt.Print("Masukkan tinggi = ")
		fmt.Scan(&tinggi)
		prisma := 0.5 * angka1 * angka2 * tinggi
		fmt.Println("Volume prisma = ", prisma)

	}

}
