package main

import "fmt"

// interface itu secara konsep bisa kita padankan dengan contract, ada aturan mainnya
type BangungDatar interface {
	// biasanya interface contractnya adalah method
	HitungLuas() int
}

// cara untuk memenuhi interface, kita buat sebuah struct
type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

type RandomNumber struct {
	Angka int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func (random RandomNumber) HitungLuas() int {
	return 2016 + random.Angka
}

func main() {
	// hitung persegi
	persegi := Persegi{Sisi: 4}
	luasPersegi := LuasPersegi(persegi)
	fmt.Println("Luas Persegi: ", luasPersegi)

	// hitung persegi panjang
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luasPersegiPanjang := LuasPersegiPanjang(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luasPersegiPanjang)

	// menggunakan interface
	fmt.Println("==============Menggunakan Interface Paramnya==============")
	resultPersegi := PersegiDanPersegiPanjang(persegi)
	fmt.Println("Luas Persegi Dengan Param Interface: ", resultPersegi)
	resultpersegiPanjang := PersegiDanPersegiPanjang(persegiPanjang)
	fmt.Println("Luas Persegi Panjang Dengan Param Interface: ", resultpersegiPanjang)

	newRandom := RandomNumber{Angka: 5}
	resultRandom := PersegiDanPersegiPanjang(newRandom)
	fmt.Println("New Random: ", resultRandom)
}

// masalah tanpa interface, akan membuat banyak function jika ingin menghitung sebuah param bangunDatar (kasusnya bangun data dalam hal ini)
// jadi satu function hanya bisa memiliki satu type data yg paten untuk parameternya
func LuasPersegi(bangunDatar Persegi) int {
	return bangunDatar.HitungLuas()
}

func LuasPersegiPanjang(bangunDatar PersegiPanjang) int {
	return bangunDatar.HitungLuas()
}

// solusi param menggunakan interface
func PersegiDanPersegiPanjang(bangunDatar BangungDatar) int {
	return bangunDatar.HitungLuas()
}
