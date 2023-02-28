docker build -t rzkynovan/appgolang . --no-cache
docker container create --name appgolang --network kanonetwork rzkynovan/appgolang

docker container start appgolang

docker container logs golangapp

docker container rm golangapp
docker image rm rzkynovan/golangapp


docker container create --name appgolang -p 8080:8080 --network kanonetwork rzkynovan/appgolang
docker container start appgolang 

"mongodb://rizky:password@dbkano:27017/"