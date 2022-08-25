package main

import (
	"fmt"
	"os"
	"strconv"
)

type Human struct {
	Name       string
	Address    string
	Occupation string
	Reason     string
}

func main() {
	persons := listPersons()
	input := os.Args
	id, _ := strconv.Atoi(input[1:][0])
	if id > 0 {
		id--
		fmt.Println("Nama:", persons[id].Name)
		fmt.Println("Alamat:", persons[id].Address)
		fmt.Println("Pekerjaan:", persons[id].Occupation)
		fmt.Println("Alasan:", persons[id].Reason)
	}

}

func listPersons() []Human {
	persons := []Human{
		{Name: "Edi", Address: "Ciledug", Occupation: "IT", Reason: "Ingin Menambah Pengalaman"},
		{Name: "Fahmi", Address: "Depok", Occupation: "IT", Reason: "ingin bisa bahasa go"},
		{Name: "Giva", Address: "Depok", Occupation: "IT", Reason: "Agar semakin jago menggunakan golang"},
		{Name: "Clara", Address: "Ciledug", Occupation: "IT", Reason: "Ingin mengetahui golang dan menggunakannya"},
		{Name: "Bayu", Address: "Parung", Occupation: "IT", Reason: "ingin bisa golang"},
		{Name: "Eka", Address: "Cilengsi", Occupation: "IT", Reason: "Agar bisa golang"},
		{Name: "Talia", Address: "Bekasi", Occupation: "IT", Reason: "supaya memperbayak pengetahuan tentang programming"},
		{Name: "Irfan", Address: "Mampang", Occupation: "IT", Reason: "ingin memperbayak penguasaan bahasa pemrograman"},
		{Name: "Fiqri", Address: "Mencong", Occupation: "IT", Reason: "Ingin mengetahui banyak bahasa, salah satunya golang"},
		{Name: "Ricky", Address: "Cipindoh", Occupation: "IT", Reason: "Ingin lebih tau tentang golang"},
	}
	return persons
}
