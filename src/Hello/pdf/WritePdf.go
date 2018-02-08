package main
import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)
//go get github.com/jung-kurt/gofpdf
func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 26)
	pdf.Cell(40, 10, "Hello PDF World")
	err := pdf.OutputFileAndClose("write_pdf.pdf")
	if err != nil {
		fmt.Println(err)
	}
}