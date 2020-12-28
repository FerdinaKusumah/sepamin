package main

import "github.com/FerdinaKusumah/sepamin/internal"

func main(){
	opt := internal.ParseOption()
	internal.Run(opt)
}