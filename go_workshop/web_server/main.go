package main


func main() {
	server := newServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()
}
