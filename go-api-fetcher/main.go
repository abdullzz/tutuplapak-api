package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    // "strings"
)

type Data []struct {
	ID           int      `json:"id"`
	NamaLapak    string   `json:"nama lapak"`
	NamaPemilik  string   `json:"nama pemilik"`
	NomorTelefon string   `json:"nomor telefon"`
	Membership   string   `json:"membership"`
	Alamat       []Alamat `json:"alamat"`
}
type Alamat struct {
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	KodePos   string `json:"kode pos"`
}

func main() {

    resp, err := http.Get("https://tutuplapak-api.herokuapp.com/lapak")
    if err != nil {
        log.Fatal(err)
    }

    body, _ := ioutil.ReadAll(resp.Body)

    // var f interface{}
    var f Data
    
    
    err = json.Unmarshal(body, &f)
    // fmt.Printf("%+v", f)


    fmt.Printf("%3s | %20s | %20s | %15s | %10s | %15s |\n", "ID", "NAMA LAPAK", "NAMA PEMILIK", "NOMOR TELEFON", "MEMBERSHIP","KOTA")
    for _, i := range f{
        fmt.Printf("%3d | %20s | %20s | %15s | %10s | %15s |\n", i.ID, i.NamaLapak, i.NamaPemilik, i.NomorTelefon, i.Membership, i.Alamat[0].Kota)
    }
}