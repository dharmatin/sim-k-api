FROM golang:alpine
MAINTAINER Deni Dharmatin Alisabri
RUN mkdir /app
EXPOSE 3000
COPY . /app
WORKDIR /app
RUN go build -o simk cmd/app/simk.go
RUN go build -o migrate cmd/migration/migration.go 
CMD ["sh", "-c", "/app/migrate;/app/simk"]