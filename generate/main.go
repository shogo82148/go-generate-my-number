package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	w := bufio.NewWriterSize(os.Stdout, 1<<20)
	for i := range 100_000_000_000 {
		if err := checkDigit(w, i); err != nil {
			panic(err)
		}
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
}

func checkDigit(w io.Writer, num int) error {
	var buf [13]byte
	a := num % 10
	num /= 10
	b := num % 10
	num /= 10
	c := num % 10
	num /= 10
	d := num % 10
	num /= 10
	e := num % 10
	num /= 10
	f := num % 10
	num /= 10
	g := num % 10
	num /= 10
	h := num % 10
	num /= 10
	i := num % 10
	num /= 10
	j := num % 10
	num /= 10
	k := num % 10
	num /= 10

	if num != 0 {
		panic("num is larger than 11 digits")
	}
	sum := a*2 + b*3 + c*4 + d*5 + e*6 + f*7 + g*2 + h*3 + i*4 + j*5 + k*6
	sum %= 11
	if sum <= 1 {
		sum = 0
	} else {
		sum = 11 - sum
	}
	buf[0] = byte('0' + k)
	buf[1] = byte('0' + j)
	buf[2] = byte('0' + i)
	buf[3] = byte('0' + h)
	buf[4] = byte('0' + g)
	buf[5] = byte('0' + f)
	buf[6] = byte('0' + e)
	buf[7] = byte('0' + d)
	buf[8] = byte('0' + c)
	buf[9] = byte('0' + b)
	buf[10] = byte('0' + a)
	buf[11] = byte('0' + sum)
	buf[12] = '\n'
	_, err := w.Write(buf[:])
	return err
}
