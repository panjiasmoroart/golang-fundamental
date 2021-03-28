package calculation

// A capitalize bersifat public agar bisa diakses di package lagin
func Add(number int, numberTwo int) int {
	// return number + numberTwo
	return penjumlahan(number, numberTwo)
}

func Multiply(dataA int, dataB int) int {
	return dataA * dataB
}

// function dengan huruf kecil hanya bisa dipanggil satu package yg sama calculation
func penjumlahan(number int, numberTwo int) int {
	return number + numberTwo
}
