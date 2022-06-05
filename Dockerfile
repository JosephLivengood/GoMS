# syntax=docker/dockerfile:1

##
## Build Go Source
##
FROM golang:1.18 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /build


##
## Build Deployment Image
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build /build

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/build"]
