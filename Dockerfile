#Start from base image 1.19.4
FROM golang:1.19.4

#configure the repo url so we can configure our work directory
ENV REPO_URL = github.com/BFDavidGamboa/bookstore_items-api

#Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH =${GOPATH}/SRC/${REPO_URL}

# /app/src/github.com/BFDavidGamboa/bookstore_items-api/src

#Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH = ${APP_PATH}/src
COPY src ${WORKPATH}
WORKDIR ${WORKPATH}

RUN go build -o items-api . 
#EXPOSE PORT 8081
EXPOSE 8081
CMD [ "./items-api"]