FROM golang:1.18

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download 

COPY backend/ . 

RUN go build -o main .

EXPOSE 8080

CMD [ "/backend/main" ]


#docker run --name jeremy-postgres --network jeremy-net -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:14-alpine