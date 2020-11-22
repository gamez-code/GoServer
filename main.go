package main

func main(){
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/Home", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
