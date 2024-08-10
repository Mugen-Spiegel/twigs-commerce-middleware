package main

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run()
}
