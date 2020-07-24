FROM golang:1.14.6-stretch AS build_stage

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download

RUN go build -o main

FROM oracle/instantclient:19 AS production_state
# COPY network/dev/ /usr/lib/oracle/19.6/client64/lib/network/admin/
COPY --from=build_stage /app .
CMD ["./main"]
