package main

import "fmt"

type Hero struct {
	Name string
	Ad int
	level int
}

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("level = ", this.level)
}

