package file

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

//Page to read file content
type Page struct {
	title   string
	ext     string
	content []byte
}

//Read get file of filename
func (page *Page) Read(filename string) {
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("open error")
		return
	}
	defer inputFile.Close()
	page.title = filename
	page.ext = filename
	inputReader := bufio.NewReader(inputFile)
	for {
		content, err := inputReader.ReadBytes('\n')
		if err != nil {
			return
		}
		if err == io.EOF {
			return
		}
		for ix := 0; ix < cap(content); ix++ {
			page.content = append(page.content, content[ix])
		}
	}
}

//PrintTitle print title of page
func (page *Page) PrintTitle() {
	fmt.Println(page.title)
}

//PrintContent print content of page
func (page *Page) PrintContent(end int) {
	fmt.Println(string(page.content[:end]))
}

//Product store info of product
type Product struct {
	name     string
	price    float64
	quentity int
}

//ReadProduct read info from file
func ReadProduct(filename string) []*Product {

	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("open error")
		panic(inputError)
	}
	defer inputFile.Close()

	products := make([]*Product, 0)
	inputReader := csv.NewReader(inputFile)
	inputReader.Comma = ';'

	var v1, v2, v3 string
	for {
		content, err := inputReader.Read()
		if err != nil {
			break
		}

		v1, v2, v3 = content[0], content[1], content[2]
		a, _ := strconv.ParseFloat(v2, 10)
		b, _ := strconv.ParseInt(v3, 10, 32)
		product := &Product{v1, a, int(b)}
		products = append(products, product)

	}
	return products
}

//String return value
func (p *Product) String() (string, float64, int) {
	return p.name, p.price, p.quentity
}

//PrintProducts print all
func PrintProducts(products []*Product) {
	for ix := 0; ix < len(products); ix++ {
		fmt.Println(products[ix])
	}
}

//UnCompress to read gzip
//
func UnCompress(filename string) {
	var r *bufio.Reader
	input, err := os.OpenFile(filename, os.O_RDONLY, 744)
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer input.Close()
	file, err := gzip.NewReader(input)
	if err != nil {
		r = bufio.NewReader(input)
	} else {
		r = bufio.NewReader(file)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}

}

//WriteFile to put string to filename
func WriteFile(filename string) {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1:])
	}
	input, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer input.Close()
	file := bufio.NewWriter(input)
	for ix := 0; ix < 10; ix++ {
		a := strconv.FormatInt(int64(ix), 10)
		str := a + ":lines to writed\n"
		file.WriteString(str)
		fmt.Println(ix, "writed")
	}
	file.Flush()
}

//CopyFile copy file from a to b
func CopyFile(desName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(desName)
	if err != nil {
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}
