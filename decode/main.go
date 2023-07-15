package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// Baca file encoded Base64 ke dalam []byte
	encodedFilePath := "encoded_file.txt"
	encodedData, err := os.ReadFile(encodedFilePath)
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	// Decode Base64 menjadi []byte
	decodedData, err := base64.StdEncoding.DecodeString(string(encodedData))
	if err != nil {
		fmt.Println("Gagal melakukan decoding:", err)
		return
	}

	// Tulis data decoded ke dalam file baru
	decodedFilePath := "golang.jpg"
	err = os.WriteFile(decodedFilePath, decodedData, 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}

	fmt.Println("File berhasil didecode dan disimpan sebagai", decodedFilePath)
}
