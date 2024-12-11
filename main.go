package main

func main() {
	db := NewDB("FIFO")
	StartServer(db)
}
