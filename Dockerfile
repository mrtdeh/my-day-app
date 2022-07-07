FROM golang:1.18-alpine AS build

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./
RUN --mount=type=cache,target=/root/.cache/go-build \ 
go build -o /out/app .

FROM scratch AS bin-unix
COPY --from=build /out/app /

EXPOSE 8080

CMD ["/app"]
