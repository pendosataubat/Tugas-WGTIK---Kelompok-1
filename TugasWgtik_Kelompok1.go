package main

// Vehicle struct untuk menyimpan informasi kendaraan
type Vehicle struct {
	NamaPengemudi string
	NomorPlat     string
	Tipe          string
	Rute          []string
}

// Violation struct untuk menyimpan informasi pelanggaran
type Violation struct {
	NamaPengemudi string
	Jumlah        int
}

// isOddEvenRoute memeriksa apakah rute tersebut tunduk pada aturan ganjil-genap
func isOddEvenRoute(rute string) bool {
	ruteGanjilGenap := []string{
		"Gajah Mada",
		"Hayam Wuruk",
		"Sisingamangaraja",
		"Panglima Polim",
		"Fatmawati",
		"Tomang Raya",
	}

	for _, ruteGanjilGenap := range ruteGanjilGenap {
		if ruteGanjilGenap == rute {
			return true
		}
	}
	return false
}

