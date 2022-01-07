package main

import "fmt"

func main()  {
	var cities = make(map[string]string)
	cities["Shanghai"] = "上海"
	cities["Beijing"] = "北京"

	for _, city := range cities{
		fmt.Println(city)
	}
	fmt.Println(cities["Shanghai"])

	delete(cities, "Shanghai")

	fmt.Println(cities)
}