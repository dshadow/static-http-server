# Static HTTP Server

[![Docker Size](https://img.shields.io/docker/image-size/dshadow/static-http-server/latest)](#) [![Docker Pulls](https://img.shields.io/docker/pulls/dshadow/static-http-server/latest)](https://hub.docker.com/r/dshadow/static-http-server) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/dshadow/static-http-server)

Pico http server to serve static content only.

## Features
- Fast and small! It's only 4.39MB !!!
- Based on Go net/http 1.17 + scratch container
- Designed for Docker Compose and Kubernetes (but can be used in other env.)

## Usage

### Run with an example index.html
```
docker-compose up --build -d
```

### Run with a custom local folder
1. Open docker-compose.yml
2. Uncomment volumes
3. Replace /local/www/folder with your own static folder
4. Run
```
docker-compose up --build -d
```

### All command line arguments with example values
```
-h					Show help and exit
-l :3000			Listening on all interfaces with tcp port 3000
-p /example/images	Prefix path
-s /www				Static folder with index.html and other files
```

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
