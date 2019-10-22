package models

import (
	"crypto/md5"
	"fmt"
	"io"
)

var GlobalNumber int

func init() {
	GlobalNumber = 10000000
}

func GetMD5(longurl string) string {
	h := md5.New()
	salt := "salty"

	io.WriteString(h, longurl+salt)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))

	return urlmd5
}

func Generate() (tiny string) {
	GlobalNumber++
	num := GlobalNumber
	fmt.Println(num)

	a := merge(getRange(48, 57), getRange(65, 90))
	a = merge(a, getRange(97, 122))

	if num < 62 {
		tiny = string(a[num])
	} else {
		var runes []rune
		runes = append(runes, a[num%62])
		num = num / 62
		for num >= 1 {
			if num < 62 {
				runes = append(runes, a[num-1])
			} else {
				runes = append(runes, a[num%62])
			}
			num = num / 62
		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		tiny = string(runes)
	}

	return tiny
}

func getRange(start, end rune) (ran []rune) {
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
