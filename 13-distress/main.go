package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	div1, div2 := mustParse(parse("[[2]]")), mustParse(parse("[[6]]"))
	var lists []listObject

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if s.Text() == "" {
			continue
		}

		listObj, err := parse(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		lists = append(lists, listObj)
	}

	lists = append(lists, div1, div2)

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].LessThan(lists[j]) == -1
	})

	div1Index, div2Index := indexOf(lists, div1)+1, indexOf(lists, div2)+1

	fmt.Println(div1Index * div2Index)
}

func indexOf(list []listObject, obj listObject) int {
	for i, e := range list {
		if e.LessThan(obj) == 0 {
			return i
		}
	}
	return -1
}

func mustParse(obj listObject, err error) listObject {
	if err != nil {
		panic(err.Error())
	}

	return obj
}

func parse(s string) (listObject, error) {
	if isNum(s) {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return listObject{}, fmt.Errorf("parse error: cannot decode %q to integer: %s", s, err.Error())
		}
		return listObject{
			value: int(num),
		}, nil
	}

	if s[0] != '[' {
		return listObject{}, fmt.Errorf("parse error: expected '['")
	}

	if s[len(s)-1] != ']' {
		return listObject{}, fmt.Errorf("parse error: expected ']'")
	}

	s = s[1 : len(s)-1]

	obj := listObject{isList: true}

	for len(s) > 0 {
		if s[0] == ',' {
			s = s[1:]
			continue
		}

		if s[0] != '[' {
			word := untilNextComma(s)
			listElem, err := parse(word)
			if err != nil {
				return obj, err
			}
			obj.elems = append(obj.elems, listElem)
			s = s[len(word):]
			continue
		}

		bracks := 0
		end := 0
		for i, r := range s {
			switch r {
			case '[':
				bracks++
			case ']':
				bracks--
			}
			if bracks == 0 {
				end = i + 1
				break
			}
		}

		word := s[:end]
		listElem, err := parse(word)
		if err != nil {
			return obj, err
		}

		obj.elems = append(obj.elems, listElem)

		s = s[len(word):]
	}

	return obj, nil
}

func untilNextComma(s string) string {
	return strings.Split(s, ",")[0]
}

func isNum(s string) bool {
	return !strings.ContainsAny(s, "[],")
}

type listObject struct {
	value  int
	isList bool
	elems  []listObject
}

func (obj listObject) LessThan(other listObject) int {
	if !obj.isList && !other.isList {
		if obj.value < other.value {
			return -1
		} else if obj.value == other.value {
			return 0
		}
		return 1
	}

	if !obj.isList && other.isList {
		return listObject{isList: true, elems: []listObject{obj}}.LessThan(other)
	}

	if obj.isList && !other.isList {
		return obj.LessThan(listObject{isList: true, elems: []listObject{other}})
	}

	for i := 0; i < min(len(obj.elems), len(other.elems)); i++ {
		less := obj.elems[i].LessThan(other.elems[i])
		if less == -1 {
			return -1
		} else if less == 1 {
			return 1
		}
	}

	if len(obj.elems) < len(other.elems) {
		return -1
	} else if len(obj.elems) > len(other.elems) {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (obj listObject) String() string {
	if !obj.isList {
		return strconv.FormatInt(int64(obj.value), 10)
	}

	var elems []string
	for _, e := range obj.elems {
		elems = append(elems, e.String())
	}

	return "[" + strings.Join(elems, ", ") + "]"
}
