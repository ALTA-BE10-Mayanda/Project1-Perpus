package main

import (
	"fmt"
	"perpus/config"
)

func main() {
	fmt.Println("Hello World")
	conn := config.InitDB()
	fmt.Println(conn.DisableAutomaticPing)
	
}