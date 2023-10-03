package main

func main() {
	server := NewAPIServer(":3000")
	server.Run()
	//fmt.Println("Yeah")
}
