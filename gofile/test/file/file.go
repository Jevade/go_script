package file

import (
	"encoding/gob"
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"flag"
	"log"
	"encoding/json"
	"encoding/xml"
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

type NewProduct struct{
	Product
	Year int
}
//Svjson will sava object  int o json flle by  []byte
func Svjson(value interface{},jsonname string){
	js,_ :=json.Marshal(value)
	file ,_ := os.OpenFile(jsonname,os.O_WRONLY|os.O_CREATE,0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.Encode(js)
}

//Rdjson will get info from json file
func RdJson(filename string){
	file ,_ := os.Open(filename)
	defer file.Close()
	dec := json.NewDecoder(file)
	var str []byte
    dec.Decode(&str)
	var in interface{}
    json.Unmarshal(str,&in)
    fmt.Println(in)
}
//Svxml
func Svxml(v interface{}, filename string){

}

type T struct{
	X int
	Y int
	Z int
}

type Q struct{
    X int
	Y int
	Zs int
}

//Svgob save object to gob file
func Svgob (filename string, v *T){
	file,_ := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY,0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(*v)
	if err!=nil{
		log.Println("Error in encoding gob1")
	}

}

//Rdgob test gob
func Rdgob(filename string,v *Q){
	file ,err1:= os.Open(filename)
	if err1 != nil {
		log.Fatal("encode error:", err1)
	}
	defer file.Close()
    inputReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inputReader)
	err := dec.Decode(v)
	if err != nil {
		log.Fatal("encode error:", err)
	}
}

//Rdxml 
func Rdxml(filename string){
	fmt.Println("read xml***********************8")
	file,_ := os.Open(filename) 
	defer file.Close()
	inputReader := xml.NewDecoder(file)
	for t,err:= inputReader.Token();err == nil;t,err =inputReader.Token(){
		switch token :=t.(type){
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name is %s\n",name)
			for _,attr := range token.Attr{
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("an attr is %s:%s",attrName,attrValue)
			}
		case xml.EndElement:
			fmt.Println("end of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content %v\n", content)
		default:
			continue
		}
	}
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
//Cat2 is cat things
func Cat2(r *os.File){
	fmt.Println("923")
	const NBUF = 512
	var buf [NBUF]byte
    for{
	    switch nr,err := r.Read(buf[:]);true{
		case nr < 0:
				fmt.Fprintf(os.Stderr,"errot:%s",err.Error())
			os.Exit(1)
        case nr == 0:
            return
        case nr > 0:
		    if nw ,ew := os.Stdout.Write(buf[0:nr]);nw != nr{
					fmt.Fprintf(os.Stderr,"cat:errty%s",ew.Error())
            }
        }
    }
}
//Cat is cat things
func Cat(r *bufio.Reader){
	fmt.Println("923")
    for{
	    buf ,err := r.ReadString('\n')
		if err == io.EOF{
		   fmt.Println(err)
			break
		}
	    fmt.Fprintf(os.Stdout,"%s",buf)
	}
	return
}
//MyCat is my cat
func MyCat(){
    flag.Parse()
	fmt.Println("params:",flag.Args())
	if flag.NArg() ==0{
	    Cat(bufio.NewReader(os.Stdin))
	}
	for ix :=0;ix<flag.NArg();ix++{
	    f,err := os.Open(flag.Arg(ix))
		if err != nil{
		    fmt.Fprintf(os.Stderr,"%s:errpr reading from %s:%s\n",os.Args[0],flag.Arg(ix),err.Error())
		}
	defer f.Close()
	Cat2(f)
	}

}

