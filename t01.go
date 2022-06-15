package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//Какой самый эффективный способ конкатенации строк?
	a := ""
	b := ""
	for range [200]int{} {
		a += string(rand.Intn(20) + 40)
		b += string(rand.Intn(20) + 40)
	}
	s := time.Now()
	for range [10000]int{} {
		a += b
	}
	fmt.Println(time.Since(s))
	a = ""
	b = ""
	for range [200]int{} {
		a += string(rand.Intn(20) + 40)
		b += string(rand.Intn(20) + 40)
	}
	var b1 bytes.Buffer
	s = time.Now()
	b1.WriteString(a)
	for range [10000]int{} {
		b1.WriteString(b)
	}
	fmt.Println(time.Since(s))
	a = ""
	b = ""
	for range [200]int{} {
		a += string(rand.Intn(20) + 40)
		b += string(rand.Intn(20) + 40)
	}
	var b2 strings.Builder
	s = time.Now()
	b2.WriteString(a)
	for range [10000]int{} {
		b2.WriteString(b)
	}
	fmt.Println(time.Since(s))
	a = ""
	b = ""
	for range [200]int{} {
		a += string(rand.Intn(20) + 40)
		b += string(rand.Intn(20) + 40)
	}
	var b3 bytes.Buffer
	s = time.Now()
	b3.WriteString(a)
	b3.WriteString(strings.Repeat(b, 10000))
	fmt.Println(time.Since(s))
	a = ""
	b = ""
	for range [200]int{} {
		a += string(rand.Intn(20) + 40)
		b += string(rand.Intn(20) + 40)
	}
	s = time.Now()
	for range [10000]int{} {
		a = fmt.Sprintf("%s%s", a, b)
	}
	fmt.Println(time.Since(s))
}
