package main
import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)
func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "To find out what's new in this tutorial, click ")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)
	// Second page: image link and basic HTML with link
	pdf.AddPage()
	pdf.SetLink(link, 0, -1)
	pdf.Image("test.png", 10, 12, 30, 0, false, "", 0, "www.baidu.com")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
		`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
		`<center>You can also center text.</center>` +
		`<right>Or align it to the right.</right>` +
		`You can also insert links on text, such as ` +
		`<a href="http://www.fpdf.org">www.baidu.com</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)
	err := pdf.OutputFileAndClose("write_html.pdf")
	if err != nil {
		fmt.Println(err)
	}
}