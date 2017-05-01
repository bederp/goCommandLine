package main

import(
	"flag"
	"os"
	"net/http"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve http on")
	path := flag.String("path", "", "path to serve")

	if *path == ""{
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
