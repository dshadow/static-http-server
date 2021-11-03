//
// Static HTTP server
// (c) 2021 Kostiantyn Cherednichenko
//

package main

import (
	"flag"
	"log"
	"strings"
	"github.com/valyala/fasthttp"
)

const version string = "2.0.2"

func main() {
	log.Println("Static HTTP Server v" + version)

	var (
		listen string
		static_folder string
		usage bool
		compress bool
		redirect_path string
		headers []string
	)

	flag.BoolVar(&usage, "h", false, "Show help and exit")
	flag.StringVar(&listen, "l", ":3000", "Listening on interface:port")
	flag.StringVar(&static_folder, "s", "/var/www", "Static folder")
	flag.BoolVar(&compress, "c", false, "Use compression")
	flag.StringVar(&redirect_path, "r", "", "Redirect path if not found (example: /index.html), except: favicon.ico, robots.txt")

	// Parse custom headers
	flag.Func("H", "Add additional HTTP headers (example: -H \"Access-Control-Allow-Origin: *\" -H \"X-Content-Type-Options: nosniff\")", func(s string) error {
		header := strings.Split(s, ": ")

		if len(header) != 2 {
            log.Fatal("Invalid header: " + s)
        }

		headers = append(headers, header...)
        
        return nil
    })

	flag.Parse()

	if usage {
		flag.Usage()
		log.Fatal("Exiting...")
	}

	log.Println("Listen: ", listen)
	log.Println("Folder: ", static_folder)
	log.Println("Compresion: ", compress)
	log.Println("Redirect path: ", redirect_path)
	log.Println("Headers: ", headers)

	// Validation
	if listen == "" {
		log.Fatal("Interface:port is invalid!")
	}

	if static_folder == "" || static_folder[0:1] != "/" {
        log.Fatal("Static folder is invalid!")
    }

	// Prepare server
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

		if redirect_path != "" {
			response := &ctx.Response
			path := string(ctx.Path())
			if response.StatusCode() == fasthttp.StatusNotFound && path != "/favicon.ico" && path != "/robots.txt" {
				ctx.Response.Reset()

				if compress {
					fasthttp.ServeFile(ctx, static_folder + redirect_path)
				} else {
					fasthttp.ServeFileUncompressed(ctx, static_folder + redirect_path)
				}
			}
		}

		// Add custom headers
		i := 0
		for i < len(headers) {
			ctx.Response.Header.Set(headers[i], headers[i + 1])
			i += 2
		}
	}

	log.Println("Listening for connection...")
	if err := fasthttp.ListenAndServe(listen, requestHandler); err != nil {
		log.Fatal(err)
	}

	select {}
}

