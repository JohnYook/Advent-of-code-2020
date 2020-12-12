package helpers

import(
	"fmt"
	"log"
	"os"
)

func GetInputFile() (file *os.File) {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: go run %s [inputFile]\n", os.Args[0])
        return
	}

    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
	}
	defer file.Close()

	return file
}