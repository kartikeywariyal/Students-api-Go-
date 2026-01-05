package main

import "github.com/kartikeywariyal/students-api-Go-/internal/config"

func main() {
	config.MustLoad()
	println("Config loaded with enwvironment:")
}
