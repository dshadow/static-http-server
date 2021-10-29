//
// Static HTTP server
// (c) 2021 Kostiantyn Cherednichenko
//

package main

import (
	"flag"
	"log"
//	"net/http"
	"github.com/valyala/fasthttp"
)

const version string = "2.0.0"

func main() {
	log.Println("Static HTTP Server v" + version)

	var (
		listen string
		static_folder string
		usage bool
		compress bool
	)

	flag.BoolVar(&usage, "h", false, "Show help and exit")
	flag.StringVar(&listen, "l", ":3000", "Listening on interface:port")
	flag.StringVar(&static_folder, "s", "/var/www", "Static folder")
	flag.BoolVar(&compress, "c", false, "Use compression")

	flag.Parse()
	
	if usage {
		flag.Usage()
		log.Fatal("Exiting...")
	}

	log.Println("Listen: ", listen)
	log.Println("Folder: ", static_folder)
	log.Println("Compresion: ", compress)

	// Validation
	if listen == "" {
		log.Fatal("Interface:port is invalid!")
	}

	if static_folder == "" || static_folder[0:1] != "/" {
        log.Fatal("Static folder is invalid!")
    }

	// Prepare server
	//http.Handle(path_prefix, fs)
	fs := &fasthttp.FS{
		Root: static_folder,
		IndexNames: []string{"index.html"},
		GenerateIndexPages: false,
		Compress: compress,
	}

	fsHandler := fs.NewRequestHandler()

	// Start server
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fsHandler(ctx)
	}

	log.Println("Listening for connection...")
	if err := fasthttp.ListenAndServe(listen, requestHandler); err != nil {
		log.Fatal(err)
	}

	select {}
}
