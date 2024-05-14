# golang base image
FROM docker.io/golang:1.22.3

# set the working directory
WORKDIR /usr/src/app

# copy all the project files onto the working directory
COPY . .

# install needed dependencies
RUN go mod download && \
 	go install github.com/cosmtrek/air@latest && \
 	go install github.com/go-task/task/v3/cmd/task@latest

# run the tail command to keep the container running
CMD ["tail", "-f", "/dev/null"]
