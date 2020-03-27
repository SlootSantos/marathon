FROM golang

WORKDIR /app/go/

COPY . .
RUN GOOS=linux GOARCH=arm GOARM=5 go build -o main .
EXPOSE 9999
CMD [ "./main" ]
