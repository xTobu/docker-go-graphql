FROM golang:1.12.1
WORKDIR /root/go
RUN pwd && ls -al
COPY . .
RUN go get github.com/oxequa/realize
CMD ["realize", "start"]
EXPOSE 8889