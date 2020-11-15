FROM golang:1.11

# Create a workspace
RUN mkdir -p /opt/authtables/scripts
WORKDIR /opt/authtables

# install deps
RUN go get github.com/willf/bloom \
           gopkg.in/redis.v4 \
           github.com/Sirupsen/logrus

# Add our files
COPY .env *.go ./
COPY scripts/test ./scripts/

# Build app
RUN go build -o authtables authtables.go configuration.go datastore.go

# Default runs on 8080
EXPOSE 8080

# Run our binary
CMD /opt/authtables/authtables
