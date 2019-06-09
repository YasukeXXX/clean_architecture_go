package main

func main() {
  controller := NewUserController()
  controller.Create("name")
  controller.Create("name2")
}
