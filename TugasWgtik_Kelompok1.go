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

// CheckViolation memeriksa pelanggaran berdasarkan tanggal dan daftar kendaraan
func CheckViolation(tanggal int, kendaraan []Vehicle) []Violation {
	var pelanggaran []Violation

	for _, kendaraan := range kendaraan {
		// Periksa apakah kendaraan adalah mobil
		if strings.ToLower(kendaraan.Tipe) != "mobil" {
			continue
		}

		// Dapatkan digit terakhir dari nomor plat
		bagianPlat := strings.Fields(kendaraan.NomorPlat)
		if len(bagianPlat) < 2 {
			continue
		}

		bagianAngka := bagianPlat[1]
		digitTerakhirStr := string(bagianAngka[len(bagianAngka)-1])
		digitTerakhir, err := strconv.Atoi(digitTerakhirStr)
		if err != nil {
			continue
		}

		jumlahPelanggaran := 0
		for _, rute := range kendaraan.Rute {
			// Periksa apakah rute tersebut tunduk pada aturan ganjil-genap
			if !isOddEvenRoute(rute) {
				continue
			}

			// Periksa apakah tanggal dan digit terakhir dari nomor plat sesuai dengan aturan ganjil-genap
			if (tanggal%2 == 0 && digitTerakhir%2 != 0) || (tanggal%2 != 0 && digitTerakhir%2 == 0) {
				jumlahPelanggaran++
			}
		}

		if jumlahPelanggaran > 0 {
			// Periksa apakah pengemudi sudah memiliki entri dalam slice pelanggaran
			ditemukan := false
			for i := range pelanggaran {
				if pelanggaran[i].NamaPengemudi == kendaraan.NamaPengemudi {
					pelanggaran[i].Jumlah += jumlahPelanggaran
					ditemukan = true
					break
				}
			}

			// Jika pengemudi tidak ditemukan dalam slice pelanggaran, tambahkan entri baru
			if !ditemukan {
				pelanggaran = append(pelanggaran, Violation{NamaPengemudi: kendaraan.NamaPengemudi, Jumlah: jumlahPelanggaran})
			}
		}
	}

	return pelanggaran
}



