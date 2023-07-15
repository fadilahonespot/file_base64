package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// Baca file ke dalam []byte
	filePath := "golang-fly.jpg"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	// Encode file menjadi Base64
	encodedData := base64.StdEncoding.EncodeToString(fileData)

	// Tulis data Base64 ke dalam file baru
	encodedFilePath := "encoded_file.txt"
	err = os.WriteFile(encodedFilePath, []byte(encodedData), 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}

	fmt.Println("File berhasil diencode dan disimpan sebagai", encodedFilePath)
}
