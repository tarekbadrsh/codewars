package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func main() {
	s := `<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>

	<prod><name>hammer</name><prx>10</prx><qty>50</qty></prod>

	<prod><name>screwdriver</name><prx>5</prx><qty>51</qty></prod>

	<prod><name>table saw</name><prx>1099.99</prx><qty>5</qty></prod>

	<prod><name>saw</name><prx>9</prx><qty>10</qty></prod>`
	fmt.Println(CatalogB(s, "saw"))
	// "table saw > prx: $1099.99 qty: 5\nsaw > prx: $9 qty: 10\n..."
}

// Prod :
type Prod struct {
	Name string `xml:"name"`
	Prx  string `xml:"prx"`
	Qty  string `xml:"qty"`
}

// Catalog :
func Catalog(s string, article string) string {
	d := xml.NewDecoder(bytes.NewBufferString(s))
	res := bytes.Buffer{}
	for {
		var p Prod
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if strings.Contains(p.Name, article) {
			if res.Len() > 0 {
				res.WriteString("\n")
			}
			res.WriteString(p.Name)
			res.WriteString(" > prx: $")
			res.WriteString(p.Prx)
			res.WriteString(" qty: ")
			res.WriteString(p.Qty)
		}
	}
	if res.Len() == 0 {
		return "Nothing"
	}

	return res.String()
}

// CatalogB :
func CatalogB(s string, article string) string {
	var r = []string{}
	re := regexp.MustCompile(`\n+`)
	arr := re.Split(s, -1)
	for _, line := range arr {
		reg := regexp.MustCompile(`<.*?><.*?>`)
		l := reg.Split(line, -1)
		if strings.Contains(l[1], article) {
			st := fmt.Sprintf("%s > prx: $%s qty: %s", l[1], l[2], l[3])
			r = append(r, st)
		}
	}
	if len(r) == 0 {
		return "Nothing"
	}
	return strings.Join(r, "\n")
}
