# Static HTTP Server

[![Docker Size](https://img.shields.io/docker/image-size/dshadow/static-http-server/latest)](#) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/dshadow/static-http-server)

Pico http server to serve static content only.

## Features
- Fast and small! It's only 4.39MB !!!
- Based on github.com/valyala/fasthttp + scratch container
- Designed for Docker Compose and Kubernetes (but can be used in other env.)

### Simple usage with an example index.html

```
$ docker run -p 127.0.0.1:8080:8080 --name test-static-http-server -d dshadow/static-http-server
```

### Simple usage with a custom www folder

```
$ docker run -p 127.0.0.1:8080:8080 --name test-static-http-server -v /my/custom/www/folder:/www -d dshadow/static-http-server
```

### Usage with docker-composer

1. Create custom docker-compose.yml
```
version: '3.8'
services:
  web:
    image: dshadow/static-http-server
    expose:
      - 8080:8080
    volumes:
	  - /my/custom/www/folder:/www
```
2. Replace /my/custom/www/folder with your own static folder
3. Build and run
```
$ docker-compose up --build -d
```
4. Stop and remove
```
$ docker-compose down
```

### Compile and use without containers
```
$ git clone https://github.com/dshadow/static-http-server.git
$ go build -o shs shs.go
$ ./shs -l :8080 -p /example/images -s /var/share/www
```

### Command line arguments
- -h Show help and exit
- -l Listening on all interfaces with a specified tcp port (default value: ":3000")
- -c Enable compression
- -s Static folder with index.html and other files (default value: "/www")

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

*  [Kostiantyn Cherednichenko](https://github.com/dshadow)

## License

MIT

## Copyright

```console
Copyright (c) 2021 Kostiantyn Cherednichenko
```
