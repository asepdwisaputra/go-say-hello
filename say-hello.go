package gosayhello

import "fmt"

type Mahasiswa struct {
	Nama, Alamat string
	NoHP         int
}

func (mahasiswa Mahasiswa) SayHello(namaOrang string) {
	fmt.Println("Haii"+namaOrang, "aku"+mahasiswa.Nama, "Salam kenal!")
}
