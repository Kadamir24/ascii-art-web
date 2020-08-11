# ascii-art-web

# Docker build только если запускаешь на новом компьютере
docker build -t app-ascii . 

# Docker run
docker container run --publish 8080:8080 --name bb app-ascii

# Docker stop
1) запускаешь команду <docker ps> и смотришь ID_CONTAINER
2) запускаешь команду <docker stop { ID_CONTAINER }>

