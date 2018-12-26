package main

// https://www.codewars.com/kata/meeting/train/go
import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"
	// (BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)
	s = "Alex:STAN"
	//fmt.Println(Meeting(s))
	fmt.Println(MeetingC(s))
}

type member struct {
	firstname string
	lastname  string
}

type members []member

func (m members) Len() int {
	return len(m)
}
func (m members) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m members) Less(i, j int) bool {
	if m[i].lastname == m[j].lastname {
		return m[i].firstname < m[j].firstname
	}
	return m[i].lastname < m[j].lastname
}

// Meeting :
func Meeting(s string) string {
	s = strings.ToUpper(s)
	var buffer bytes.Buffer
	arr := strings.Split(s, ";")
	lst := make(members, len(arr))

	for i := 0; i < len(arr); i++ {
		lst[i].firstname = strings.Split(arr[i], ":")[0]
		lst[i].lastname = strings.Split(arr[i], ":")[1]
	}

	sort.Sort(lst)
	for _, v := range lst {
		buffer.WriteString("(")
		buffer.WriteString(v.lastname)
		buffer.WriteString(", ")
		buffer.WriteString(v.firstname)
		buffer.WriteString(")")
	}
	return buffer.String()
}

// MeetingB :
func MeetingB(s string) string {
	var buffer bytes.Buffer
	lst := members{}
	mem := member{}
	for k, v := range s {
		if v == ':' {
			mem.firstname = buffer.String()
			buffer.Reset()
			continue
		}
		if v == ';' || k == len(s)-1 {
			if k == len(s)-1 {
				v -= 'a' - 'A'
				buffer.WriteRune(v)
			}
			mem.lastname = buffer.String()
			buffer.Reset()
			lst = append(lst, mem)
			mem = member{}
			continue

		}
		if v >= 'a' && v <= 'z' {
			v -= 'a' - 'A'
		}
		buffer.WriteRune(v)
	}
	sort.Sort(lst)
	for _, v := range lst {
		buffer.WriteString("(")
		buffer.WriteString(v.lastname)
		buffer.WriteString(", ")
		buffer.WriteString(v.firstname)
		buffer.WriteString(")")
	}
	return buffer.String()
}

// MeetingC :
func MeetingC(s string) string {
	var finalname bytes.Buffer
	var firstname bytes.Buffer
	var lastname bytes.Buffer
	var isfirstname = true
	lst := []string{}
	for k, v := range s {
		if v == ':' {
			isfirstname = false
		}
		if v == ';' || k == len(s)-1 {
			if k == len(s)-1 {
				if v >= 'a' && v <= 'z' {
					v -= 'a' - 'A'
				}
				lastname.WriteRune(v)
			}
			finalname.WriteString("(")
			finalname.WriteString(lastname.String())
			finalname.WriteString(", ")
			finalname.WriteString(firstname.String())
			finalname.WriteString(")")
			lst = append(lst, finalname.String())
			lastname.Reset()
			firstname.Reset()
			finalname.Reset()
			isfirstname = true
		}

		if v >= 'a' && v <= 'z' {
			v -= 'a' - 'A'
		}
		if v >= 'A' && v <= 'Z' {
			if isfirstname {
				firstname.WriteRune(v)
			} else {
				lastname.WriteRune(v)
			}
		}
	}

	sort.Strings(lst)
	return strings.Join(lst, "")
	// var res bytes.Buffer
	// for _, v := range lst {
	// 	res.WriteString(v)
	// }
	// return res.String()
}

// MeetingD :
func MeetingD(s string) string {
	sw := strings.Split(strings.ToUpper(s), ";")
	res := []string{}
	for _, pn := range sw {
		a := strings.Split(pn, ":")
		s := "(" + a[1] + ", " + a[0] + ")"
		res = append(res, s)
	}
	sort.Strings(res)
	return strings.Join(res, "")
}
