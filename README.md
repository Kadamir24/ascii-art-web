# ascii-art-web

Online tool to create and download ASCII art from text using 3 different fonts.

## Front-end
* HTML
* CSS
* JS

## Back-end
* Golang
* Docker

# Usage
```
$ go build
$ ./ascii-art-web
```
and open http://localhost:8080 in your browser

# Docker build только если запускаешь на новом компьютере
docker build -t app-ascii . 

# Docker run
docker container run --publish 8080:8080 --name bb app-ascii

# Docker stop
1) запускаешь команду <docker ps> и смотришь ID_CONTAINER
2) запускаешь команду <docker stop { ID_CONTAINER }>

