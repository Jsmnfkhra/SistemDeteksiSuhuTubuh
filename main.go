package main

import (
	"fmt"
)

func main() {
	for {
		tampilkanMenu()
		pilihan := inputMenu()

		switch pilihan {
		case 1:
			cekSuhu()
		case 2:
			infoKategori()
		case 3:
			fmt.Println("\nTerima kasih! Program selesai.")
			return
		}
	}
}

// =====================================================
// MENU
// =====================================================
func tampilkanMenu() {
	fmt.Println("=========================================")
	fmt.Println("         CEK STATUS SUHU TUBUH           ")
	fmt.Println("=========================================")
	fmt.Println("1. Cek Suhu Tubuh")
	fmt.Println("2. Info Kategori Suhu")
	fmt.Println("3. Keluar")
	fmt.Println("=========================================")
}

// =====================================================
// VALIDASI INPUT MENU
// =====================================================
func inputMenu() int {
	for {
		var pilihan int
		fmt.Print("Masukkan pilihan (1-3): ")
		_, err := fmt.Scan(&pilihan)

		if err != nil {
			fmt.Println("❌ Input harus berupa angka 1-3!\n")
			clearBuffer()
			continue
		}

		if pilihan < 1 || pilihan > 3 {
			fmt.Println("❌ Pilihan tidak tersedia! Pilih 1, 2, atau 3.\n")
			continue
		}

		return pilihan
	}
}

// =====================================================
// FITUR CEK SUHU
// =====================================================
func cekSuhu() {
	for {
		var suhu float64

		fmt.Println("\n=========================================")
		fmt.Println("              INPUT SUHU TUBUH           ")
		fmt.Println("=========================================")

		fmt.Print("Masukkan suhu (°C): ")
		_, err := fmt.Scan(&suhu)

		if err != nil {
			fmt.Println("❌ Input tidak valid! Harus berupa angka.")
			clearBuffer()
			continue
		}

		status, saran := tentukanStatus(suhu)

		fmt.Println("\n=========================================")
		fmt.Printf("Suhu   : %.2f °C\n", suhu)
		fmt.Println("Status :", status)
		fmt.Println("Saran  :", saran)
		fmt.Println("=========================================")

		if !tanyaUlang() {
			fmt.Println()
			return
		}

		fmt.Println()
	}
}

// =====================================================
// LOGIKA STATUS SUHU
// =====================================================
func tentukanStatus(suhu float64) (string, string) {
	if suhu < 35 {
		return "Hipotermia (Bahaya)", "Segera cari bantuan medis."
	} else if suhu >= 35 && suhu <= 37 {
		return "Normal", "Tetap jaga pola makan dan tidur."
	} else if suhu > 37 && suhu <= 38 {
		return "Demam Ringan", "Perbanyak minum air putih."
	} else if suhu > 38 && suhu <= 39 {
		return "Demam Sedang", "Konsumsi obat demam dan istirahat."
	} else {
		return "Demam Tinggi (Bahaya)", "Segera ke dokter."
	}
}

// =====================================================
// VALIDASI CEK LAGI (y/n)
// =====================================================
func tanyaUlang() bool {
	for {
		var ulang string
		fmt.Print("Cek lagi? (y/n): ")
		fmt.Scan(&ulang)

		if ulang == "y" || ulang == "Y" {
			return true
		}

		if ulang == "n" || ulang == "N" {
			return false
		}

		fmt.Println("❌ Input tidak valid! Masukkan y atau n.")
	}
}

// =====================================================
// INFO KATEGORI
// =====================================================
func infoKategori() {
	fmt.Println("\n=========================================")
	fmt.Println("          INFORMASI KATEGORI SUHU        ")
	fmt.Println("=========================================")
	fmt.Println("< 35°C     : Hipotermia (Bahaya)")
	fmt.Println("35 – 37°C  : Normal")
	fmt.Println("37.1 – 38°C: Demam Ringan")
	fmt.Println("38.1 – 39°C: Demam Sedang")
	fmt.Println("> 39°C     : Demam Tinggi (Bahaya)")
	fmt.Println("=========================================\n")
}

// =====================================================
// CLEAR BUFFER UNTUK HANDLE INPUT ERROR
// =====================================================
func clearBuffer() {
	var buang string
	fmt.Scanln(&buang)
}
