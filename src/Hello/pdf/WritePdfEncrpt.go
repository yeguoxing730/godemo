package main
import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetProtection(gofpdf.CnProtectPrint, "123", "abc")
	pdf.AddPage()
	pdf.AddLayer("layer",true)
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "You Must Enter the Password!!!")
	err := pdf.OutputFileAndClose("write_pdf_with_password.pdf")
	if err != nil {
		fmt.Println(err)
	}
}