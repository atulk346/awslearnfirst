#####BASE IMAGE OR ENVIROMENT####
FROM golang:1.24-alpine
#####Create directory ######
WORKDIR /app

####    COPY ALL YOUR CODE INTO IMAGE ####
COPY . .

#### CREATE BUILD #####

RUN go build -o app .

### RUN THE BUILD ###
CMD ["./app"]