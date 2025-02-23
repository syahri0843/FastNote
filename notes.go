package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var notes []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addNote(scanner)
		case "2":
			viewNotes()
		case "3":
			fmt.Println("Keluar... Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func showMenu() {
	fmt.Println("\n=== APLIKASI CATATAN ===")
	fmt.Println("1. Tambah Catatan")
	fmt.Println("2. Lihat Catatan")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
}

func addNote(scanner *bufio.Scanner) {
	fmt.Print("Masukkan catatan: ")
	scanner.Scan()
	note := strings.TrimSpace(scanner.Text())

	if note == "" {
		fmt.Println("Catatan tidak boleh kosong!")
		return
	}

	notes = append(notes, note)
	fmt.Println("Catatan berhasil ditambahkan!")
}

func viewNotes() {
	fmt.Println("\n=== DAFTAR CATATAN ===")
	if len(notes) == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}

	for i, note := range notes {
		fmt.Printf("%d. %s\n", i+1, note)
	}
}