package main

const version = 1

func main() {
	initLogger()
	initConfig()
	initDB()
	saveConfig()
	initHttpServer()
}



