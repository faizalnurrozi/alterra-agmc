# Altera Golang Mini Course (AGMC)

This project (_**alterra-agmc**_) is a RESTful API with an agnostic approach, so developers only need to think about the business process.

## Features

- Support [RESTful API](https://en.wikipedia.org/wiki/Representational_state_transfer). e.g. User & Book.
- Support Object Relational Mapping ([ORM](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping)) concept.
- Implement clean architecture ([Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))) on main project.
- Implement common design pattern.

## Tech

AGMC project uses a number of open source project to work properly:

- [Go](https://go.dev/) - Programming language
- [Docker](https://www.docker.com/) - Containerization
- [MySQL](https://www.mysql.com/) - Relational database
- [Heroku](https://www.heroku.com) - Deployment image to server

## Installation

AGMC project requires [go](https://go.dev/) version go1.19 to run.

Initiate new table and start the server.

```sh
cd alterra-agmc
go run main.go -migrate=migrate
```
## Docker

AGMC project is very easy to install and deploy in a Docker container.

By default, the Docker will expose port **8080**, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
cd alterra-agmc
docker build -t <youruser>/alterra-agmc:<tag>
```

This will create the **alterra-agmc** image and pull in the necessary dependencies.
Be sure to swap out `<tag>` with the actual
version of AGMC.

Once done, run the Docker image and map the port to whatever you wish on
your host. In this example, we simply map port **8082** of the host to
port **8080** of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -d -p 8082:8080 --restart=always --name=alterra-agmc <youruser>/alterra-agmc:<tag>
```

> Note: `--restart=always` to configure the restart policy for a container.

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
127.0.0.1:8082
```

## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

[dill]: <https://github.com/joemccann/dillinger>
[git-repo-url]: <https://github.com/joemccann/dillinger.git>
[john gruber]: <http://daringfireball.net>
[df1]: <http://daringfireball.net/projects/markdown/>
[markdown-it]: <https://github.com/markdown-it/markdown-it>
[Ace Editor]: <http://ace.ajax.org>
[node.js]: <http://nodejs.org>
[Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
[jQuery]: <http://jquery.com>
[@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
[express]: <http://expressjs.com>
[AngularJS]: <http://angularjs.org>
[Gulp]: <http://gulpjs.com>

[PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
[PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
[PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
[PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
[PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
[PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
