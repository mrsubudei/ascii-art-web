name = ascii-art-web
port = 8087

start: build run

build:
		@docker build -t $(name) .
		@docker image prune --filter label=stage=build -f

run:
		@docker run -p $(port):8087 --name $(name) -d $(name)
		@echo  
		@echo Starting server.. http://localhost:$(port)

exec:
		@docker exec -ti $(name) sh

stop:
		@docker stop $(name)
		@docker rm $(name)

remove:
		@docker rmi $(name)

kill: stop remove

meta:
		@echo Authors:
		@docker image inspect --format='{{.ContainerConfig.Labels.authors}}' $(name)
		@echo DockerVersion:
		@docker image inspect --format='{{.DockerVersion}}' $(name)