

# remove stopped containers
docker rm $(docker ps -a -q)
# remove untagged images
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
docker build -t pinnaman/go_1.12_hello .
docker run -d -p 8080:8080 pinnaman/go_1.12_hello
