### Altera Golang Mini Course

#### A. Docker
- How to run project

      $ go run main.go -migrate=migrate

#### B. Dockerize
- Docker hub repository: https://hub.docker.com/r/faizalnurrozi/amgc
- How to build an image

      $ docker build -t faizalnurrozi/agmc:latest . -f docker/go/Dockerfile
  
- How to push image to repository

      $ docker login
      $ docker push faizalnurrozi/agmc