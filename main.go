package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	fetchdata "weeklytask8/fetchData"
	geometristemplate "weeklytask8/geometrisTemplate"
	processnumber "weeklytask8/processNumber"
	usermanagement "weeklytask8/userManagement"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Process Number")
		fmt.Println("2. Fetch Data")
		fmt.Println("3. Management User")
		fmt.Println("4. Template Geometris")
		fmt.Println("0. Keluar Aplikasi")

		fmt.Print("pilih menu: ")
		scanner.Scan()
		choice := scanner.Text()
		choice = strings.ToLower(choice)

		switch choice {
		case "1":
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recover from panic %v\n", r)
				}
			}()
			fmt.Println("===Process Number===")
			var inputs = []int{}

			for {
				fmt.Println("ketik 'selesai' untuk lanjut")
				fmt.Print("Masukan Input(angka): ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())

				if strings.ToLower(input) == "selesai" {
					if len(inputs) == 0 {
						panic("Empty list provided")
					}
					break
				}

				if input == "" {
					fmt.Println("NO data provided")
					continue
				}

				inputConv, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Failed to process input")
					continue
				}

				if inputConv <= 0 {
					fmt.Println("Input must more than zero")
					continue
				}

				inputs = append(inputs, inputConv)
			}
			results := processnumber.ProcessNumber(inputs)
			fmt.Println(results)

		case "2":
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recover from panic: %v\n", r)
				}
			}()

			fmt.Println("== Fetch Data ==")

			fetchDataChan := make(chan string)
			var urls = []string{}

			for {
				fmt.Println("Ketik 'selesai' untuk lanjut")
				fmt.Print("Masukan Url: ")
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())

				if strings.ToLower(input) == "selesai" {
					if len(urls) == 0 {
						panic("Empty list provided")
					}
					break
				}

				if input == "" {
					fmt.Println("NO Url provided")
					continue
				}

				urls = append(urls, input)
			}

			var wg sync.WaitGroup

			for _, url := range urls {
				wg.Add(1)
				go fetchdata.WebFetcher(fetchDataChan, url, &wg)
			}

			go func() {
				wg.Wait()
				close(fetchDataChan)
				fmt.Println("\nProses fetch selesai")
			}()

			fmt.Println("\n=== HASIL FETCH ===")
			for result := range fetchDataChan {
				fmt.Println(result)
				fmt.Println("---")
			}

		case "3":
			fmt.Println("== Management User ==")

			userManager := usermanagement.NewUserManager()

			for {
				fmt.Println("1. Tambah User")
				fmt.Println("2. Cari User by ID")
				fmt.Println("0. Kembali ke Menu Utama")

				fmt.Print("Pilih menu: ")
				scanner.Scan()
				choice := strings.TrimSpace(scanner.Text())

				switch choice {
				case "1":
					fmt.Print("Masukan ID User: ")
					scanner.Scan()
					id := strings.TrimSpace(scanner.Text())

					fmt.Print("Masukan Username: ")
					scanner.Scan()
					username := strings.TrimSpace(scanner.Text())

					if id == "" || username == "" {
						fmt.Println("ID dan Username tidak boleh kosong")
						continue
					}

					user := usermanagement.User{
						Id:       id,
						Username: username,
					}

					userManager.AddUser(user)

				case "2":
					fmt.Print("Masukan ID yang dicari: ")
					scanner.Scan()
					id := strings.TrimSpace(scanner.Text())

					if user, exists := userManager.GetUserById(id); exists {
						fmt.Printf("User ditemukan: ID: %s, Username: %s\n", user.Id, user.Username)
					} else {
						fmt.Println("User tidak ditemukan")
					}

				case "0":
					fmt.Println("Keluar...")
					return

				default:
					fmt.Println("Pilihan tidak valid")
				}
			}
		case "4":
			fmt.Println("Template Geometris")
			fmt.Println("Menyiapkan lingkaran dengan radius 7cm")
			circle := geometristemplate.Circle{}
			circle.Radius = 7
			fmt.Println("Menyiapkan persegi dengan tinggi 5cm dan lebar 5cm")
			rectangle := geometristemplate.Rectangle{}
			rectangle.Tinggi = 5
			rectangle.Lebar = 5
			fmt.Println("Menghitung total area...")
			input := []geometristemplate.IGeometric{circle, rectangle}
			totalArea := geometristemplate.CalculateArea(input)
			fmt.Printf("Total Area: %fcm\n", totalArea)
		case "0":
			fmt.Println("Keluar dari aplikasi, sampai jumpa 🫡")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
