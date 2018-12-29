package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(ByState(`John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`))
	// "California\n..... Amy Wilde 334 Bayshore Pkwy Mountain View California\n Massachusetts\n..... Eric Adams 20 Post Road Sudbury Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Oklahoma\n..... Orville Thomas 11345 Oak Bridge Road Tulsa Oklahoma\n Pennsylvania\n..... Terry Kalkas 402 Lans Road Beaver Falls Pennsylvania\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia\n..... Hubert Sims 328A Brook Road Roanoke Virginia"
}

type state struct {
	name    string
	address bytes.Buffer
}

// ByState :
func ByState(str string) string {
	states := map[string]*state{}
	states["AZ"] = &state{name: "Arizona"}
	states["CA"] = &state{name: "California"}
	states["ID"] = &state{name: "Idaho"}
	states["IN"] = &state{name: "Indiana"}
	states["MA"] = &state{name: "Massachusetts"}
	states["OK"] = &state{name: "Oklahoma"}
	states["PA"] = &state{name: "Pennsylvania"}
	states["VA"] = &state{name: "Virginia"}

	lines := strings.Split(str, "\n")
	sort.Strings(lines)
	for _, l := range lines {
		k := l[len(l)-2:]
		s := states[k]
		l = strings.TrimSpace(l)
		l = strings.Replace(l, ",", "", -1)
		l = strings.Replace(l, k, s.name, -1)
		s.address.WriteString("\n..... ")
		s.address.WriteString(l)
	}
	resarr := []string{}
	for _, v := range states {
		if v.address.Len() > 0 {
			resarr = append(resarr, (v.name + v.address.String()))
		}
	}
	sort.Strings(resarr)
	return strings.Join(resarr, "\n ")
}

// ByStateB :
func ByStateB(str string) string {
	res := ""
	m := map[string]string{}
	m["AZ"] = "Arizona"
	m["CA"] = "California"
	m["ID"] = "Idaho"
	m["IN"] = "Indiana"
	m["MA"] = "Massachusetts"
	m["OK"] = "Oklahoma"
	m["PA"] = "Pennsylvania"
	m["VA"] = "Virginia"

	states := map[string]string{}
	lines := strings.Split(str, "\n")
	sort.Strings(lines)
	for _, v := range lines {
		v = strings.TrimSpace(v)
		v = strings.Replace(v, ",", "", -1)
		s := v[len(v)-2:]
		f := strings.Replace(v, s, m[s], -1)
		if _, ok := states[m[s]]; !ok {
			states[m[s]] = m[s]
		}
		states[m[s]] += fmt.Sprintf("\n..... %v", f)
	}

	myarr := []string{"Arizona", "California", "Idaho", "Indiana", "Massachusetts", "Oklahoma", "Pennsylvania", "Virginia"}
	for _, v := range myarr {
		if _, ok := states[v]; ok && res != "" {
			res += fmt.Sprintf("\n %v", states[v])
		} else if _, ok := states[v]; ok {
			res += states[v]
		}
	}
	return res
}

// ByStateC :
func ByStateC(str string) string {
	type statefriend struct {
		state   string
		name    string
		address string
		town    string
	}
	var s = []string{" MA", " VA", " OK", " PA", " CA", " AZ", " ID", " IN"}
	var st = []string{", Massachusetts", ", Virginia", ", Oklahoma", ", Pennsylvania", ", California", ", Arizona", ", Idaho", ", Indiana"}
	var u = str
	for i := range s {
		u = strings.Replace(u, s[i], st[i], -1)
	}
	re := regexp.MustCompile(`\n+`)
	arr := re.Split(u, -1)
	narr := make([]statefriend, len(arr))
	for i, x := range arr {
		y := strings.Split(x, ", ")
		narr[i].state = y[3]
		narr[i].name = y[0]
		narr[i].address = y[1]
		narr[i].town = y[2]
	}
	sort.Slice(narr, func(i, j int) bool {
		if narr[i].state == narr[j].state {
			return narr[i].name < narr[j].name
		}
		return narr[i].state < narr[j].state
	})
	output := ""
	last := ""

	for _, v := range narr {
		e := v.state
		if e != last {
			last = e
			output += "\n" + " " + e + "\n..... " + v.name + " " + v.address + " " + v.town + " " + v.state
		} else {
			output += "\n..... " + v.name + " " + v.address + " " + v.town + " " + v.state
		}
	}

	output = output[2:len(output)]
	return output
}
