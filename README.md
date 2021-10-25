# Static HTTP Server
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
