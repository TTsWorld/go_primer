package csv_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
	readcsv()

}

func readcsv() {
	//准备读取文件
	fileName := "/tmp/market02.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
			continue
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
		fmt.Println(len(row))
		if len(row) != 9 {
			fmt.Println(len(row))
			continue
		}
		uid := row[0]
		firstAfid := row[3]
		fmt.Print(uid + " - ")
		fmt.Println(firstAfid)
	}
}
