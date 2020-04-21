package main

import "fmt"

type Course struct {
  Name    string
  Price   float64
  IsFree  bool
  UserIDs []uint
  Classes map[uint]string
}

func (c Course) PrintClasses() {
  text := "Las clases son: "
  for _, class := range c.Classes {
    text += class + ", "
  }

  fmt.Println(text[:len(text) - 2])
}

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
