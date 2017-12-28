package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	v, err := readIn(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(v); i++ {
		reverse(v[i])
	}
	for k, value := range v {
		if k == 0 {
			fmt.Printf("%s", printArr(value))
		} else {
			fmt.Printf("\n%s", printArr(value))
		}
	}
}

func printArr(a []int) string {
	o := ""
	for k, v := range a {
		if k == 0 {
			o += strconv.FormatInt(int64(v), 10)
		} else {
			o += " " + strconv.FormatInt(int64(v), 10)
		}
	}
	return o
}

func readIn(r io.Reader) ([][]int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var src [][]int
	n := true
	for s.Scan() {
		if !n {
			txt := strings.Split(
				strings.TrimSpace(s.Text()), " ")
			var o []int
			for _, v := range txt {
				value, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				o = append(o, value)
			}
			src = append(src, o)
			n = true
			continue
		}
		n = false
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return src, nil
}

func reverse(a []int) {
	left := 0
	right := len(a) - 1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}
