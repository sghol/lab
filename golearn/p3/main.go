package main

import (
	"fmt"
	"strings"
)

func toLatinDigits(s string) string {
	persian := []rune("۰۱۲۳۴۵۶۷۸۹")
	arabic := []rune("٠١٢٣٤٥٦٧٨٩")

	for i := range 10 {
		s = strings.ReplaceAll(s, string(persian[i]), string(rune('0'+i)))
		s = strings.ReplaceAll(s, string(arabic[i]), string(rune('0'+i)))
	}
	return s
}

// Normalize phone to store magfa income format : 989xxxxxxxxx
func normalizePhone(phone string) string {
	// convert phone number to latin digits
	p := toLatinDigits(phone)
	normalized := ""

	// 09123456789 -> 989123456789
	if p[0] == '0' {
		normalized = "98" + p[1:]
	}
	// 9123456789 -> 989123456789
	if p[0] == '9' && len(p) == 10 {
		normalized = "98" + p
	}
	// +989123456789 -> 989123456789
	if p[0] == '+' {
		normalized = p[1:]
	}
	return normalized
}
func main() {
	res := toLatinDigits("۹۳۷۱۲۳۴۵۶۷")
	res1 := toLatinDigits("۰۹۳۷۱۲۳۴۵۶۷")
	res2 := toLatinDigits("۹۸۹۳۷۱۲۳۴۵۶۷")
	res3 := toLatinDigits("+۹۸۹۳۷۱۲۳۴۵۶۷")
	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	fmt.Println("--------------Normalize -------------")
	n1 := normalizePhone("09123456789")
	n2 := normalizePhone("9123456789")
	n3 := normalizePhone("+989123456789")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
