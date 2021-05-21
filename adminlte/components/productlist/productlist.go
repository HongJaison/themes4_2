package productlist

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/HongJaison/go-admin4/modules/logger"
	adminTemplate "github.com/HongJaison/go-admin4/template"
	"github.com/HongJaison/themes4_2/adminlte/components"
)

type ProductList struct {
	components.Base
	Data []map[string]string
}

func New() ProductList {
	return ProductList{}
}

func (p ProductList) SetData(value []map[string]string) ProductList {
	p.Data = value
	return p
}

func (p ProductList) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("productlist").
		Funcs(adminTemplate.DefaultFuncMap).
		Parse(List["productlist"])

	if err != nil {
		logger.Error("ProductList GetTemplate Error: ", err)
	}

	return tmpl, "productlist"
}

func (p ProductList) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := p.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, p)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
