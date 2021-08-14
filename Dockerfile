FROM golang as builder
RUN mkdir -p /source/app
WORKDIR /source/app
COPY . .
RUN go mod download
RUN go build -o hello

FROM scratch
COPY --from=builder /source/app/hello /bin/hello
ENTRYPOINT [ "/bin/hello" ]