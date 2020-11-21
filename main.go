package main

func main(){
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/Home", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
