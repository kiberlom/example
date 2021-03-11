package main

import (
	"fmt"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

func pd() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		fmt.Println(err)
		return
	}
	htmlStr := `<html><body><h1 style="color:red;">This is an html
 from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("Your_pdfname.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func main() {

	//const exe = "wkhtmltopdf"
	//
	//exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(exeDir)
	//
	//path, err := exec.LookPath(filepath.Join(exeDir, exe))
	//if err == nil && path != "" {
	//	fmt.Println(path)
	//}
	//fmt.Println(path)
	//
	//
	//path, err = exec.LookPath(exe)
	//if err == nil && path != "" {
	//	fmt.Println(path)
	//}
	//
	//fmt.Println(path)
	//

	//dir := os.Getenv("WKHTMLTOPDF_PATH")
	////dir := os.Getenv("GOMOD")
	//if dir == "" {
	//	fmt.Println(12)
	//}
	//
	//fmt.Println(dir)

	//path, err = exec.LookPath(filepath.Join(dir, exe))
	//if err == nil && path != "" {
	//	fmt.Println(path)
	//}

	pd()

}
