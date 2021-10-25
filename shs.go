//
// Static HTTP server (c) 2021 Kostiantyn Cherednichenko
//

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	log.Println("Static HTTP Server")

	var (
		listen string
		static_folder string
		path_prefix string
		usage bool
	)

	flag.BoolVar(&usage, "h", false, "Show help and exit")
	flag.StringVar(&listen, "l", ":3000", "Listening on interface:port")
	flag.StringVar(&path_prefix, "p", "/", "Path prefix")
	flag.StringVar(&static_folder, "s", "/var/www", "Static folder")

	flag.Parse()
	
	if usage {
		flag.Usage()
		log.Fatal("Exiting...")
	}

	log.Println("Listen: ", listen)
	log.Println("Prefix: ", path_prefix)
	log.Println("Folder: ", static_folder)

	// Validation
	if listen == "" {
		log.Fatal("Interface:port is invalid!")
	}

	if static_folder == "" || static_folder[0:1] != "/" {
        log.Fatal("Static folder is invalid!")
    }

	if path_prefix == "" || path_prefix[0:1] != "/" {
		log.Fatal("Path prefix is invalid!")
	}

	// Prepare server
	fs := http.FileServer(http.Dir(static_folder))
	http.Handle(path_prefix, fs)

	// Start server
	log.Println("Listening for connection...")
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
