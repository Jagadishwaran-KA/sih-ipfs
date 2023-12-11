package main

import (
	"fmt"
	"io"
	"log"
	"os"
	//"path/filepath"

	"github.com/ulikunitz/xz"
)

func compressLZMA(inputPath, outputPath string) error {
	// Open the input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Create an LZMA writer
	writer, err := xz.NewWriter(outputFile)
	if err != nil {
		return err
	}
	defer writer.Close()

	// Copy data from the input file to the LZMA writer
	_, err = io.Copy(writer, inputFile)
	if err != nil {
		return err
	}

	return nil
}

func decompressLZMA(inputPath, outputPath string) error {
	// Open the input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Create an LZMA reader
	reader, err := xz.NewReader(inputFile)
	if err != nil {
		return err
	}

	// Copy data from the LZMA reader to the output file
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go compress input.txt output.lzma")
		fmt.Println("       go run main.go decompress input.lzma output.txt")
		os.Exit(1)
	}

	mode := os.Args[1]
	inputPath := os.Args[2]
	outputPath := os.Args[3]

	switch mode {
	case "compress":
		err := compressLZMA(inputPath, outputPath)
		if err != nil {
			log.Fatal("Error:", err)
		}
		fmt.Println("Compression successful.")
	case "decompress":
		err := decompressLZMA(inputPath, outputPath)
		if err != nil {
			log.Fatal("Error:", err)
		}
		fmt.Println("Decompression successful.")
	default:
		fmt.Println("Invalid mode. Use 'compress' or 'decompress'.")
		os.Exit(1)
	}
}
