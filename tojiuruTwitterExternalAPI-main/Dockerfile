FROM golang:latest
WORKDIR /go/src/
ENV KEY 7t3VzvHG5Ge5YUmMrN2Ejt9pbthidGXj
# callback environment, default is localhost:3000
ENV CBURL default

COPY . /go/src/

RUN apt-get update
RUN apt-get install -y jq
RUN apt-get install -y sqlite3

RUN chmod +x /go/src/database/importsql.sh
RUN /go/src/database/importsql.sh