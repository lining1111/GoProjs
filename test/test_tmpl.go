package test

import (
	"fmt"
	"os"
	"text/template"
)

func TestTmpl() {
	s1, _ := template.ParseFiles("html/header.html", "html/content.html", "html/footer.html")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
