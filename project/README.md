# Build Docker image:
docker build --tag docker-gs-ping .

# Run docker cmd:
docker run -d --rm -p 8080:8080 docker-gs-ping

# Run docker in jenkins from dockerhub:
docker run -d --rm --name docker-gs-ping -p 8081:8080 thehaohcm/jenkinslab:latest
