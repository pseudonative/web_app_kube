FROM golang:1.18

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download 

COPY backend/ . 

RUN go build -o main .

EXPOSE 8080

CMD [ "/backend/main" ]