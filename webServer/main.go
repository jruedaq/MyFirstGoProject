package main

func main() {
	server := NewServer(":80")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Login()))
	server.Listen()
}
