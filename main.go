package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world")

	//string
	var nama string = "esa dimas aditya"
	fmt.Println(nama)

	//var dan const

	var example = "variabel"
	const ex = "const"
	fmt.Println(example)
	fmt.Println(ex)

	//boolean
	var pelajar bool = true
	fmt.Println(pelajar)

	//type declaration
	type str string
	type numb int

	var sekolah str = "smkn 1 kalitengah"
	var nomor numb = 10
	fmt.Println(sekolah)
	fmt.Println(nomor)

	//operasi matematika
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	var i = 10
	i++
	i += 10
	fmt.Println(i)

	//conversion

	var merk = "nike"
	var m = merk[0]
	var mMerk = string(m)
	fmt.Println(merk)
	fmt.Println(mMerk)

	//perbandingan

	var x = "blue"
	var y = "red"
	var z = x == y
	fmt.Println(z)
	fmt.Println(x == y)

	//type data number

	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 2147483647
	var nilai64 int64 = 9223372036854775807
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	//int8	    -128 – 127
	//init16	  -32768 – 32767
	//int32	    -2147483648 – 2147483647
	//int64	    -9223372036854775808 – 9223372036854775807

	//operasi boolean

	var ujian = 75
	var absen = 80

	var nilaiUjian = ujian >= 75
	var nilaiAbsen = absen >= 80

	var hasil = nilaiUjian && nilaiAbsen
	fmt.Println(hasil)

	//array
	var nilai = [03]int {
		80,
		75,
		90,
	}
	fmt.Println(nilai)

	var pemrograman [03]string
	pemrograman[0] = "golang"
	pemrograman[1] = "java"
	pemrograman[2] = "python"
	fmt.Println(pemrograman[0])
	fmt.Println(pemrograman[1])
	fmt.Println(pemrograman[2])
}

