package main

func main()  {
	server := NewServer(":3000")

	//uso de la funcion handle con el path "/" y el handler
	server.Handle("GET","/", HandleRoot)
	server.Handle("POST","/create", PostRequest)
    server.Handle("POST","/user", UserPostRequest)
	server.Handle("POST","/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}