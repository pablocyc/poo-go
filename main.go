package main

func main() {
  Go := Course {
    "Go desde cero",
    12.34,
    false,
    []uint{12, 56, 32},
    map[uint]string{
      1: "Introducción",
      2: "Estructuras",
      3: "Mapas",
    },
  }

  Go.PrintClasses()
}
