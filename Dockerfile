FROM golang

WORKDIR /

COPY . /

RUN go build .

EXPOSE 8080

CMD ["./qr"]