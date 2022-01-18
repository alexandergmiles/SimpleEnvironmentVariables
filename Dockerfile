# syntax=docker/dockerfile:1
FROM golang:1.18beta1-alpine3.15
WORKDIR /simpleenvironmentvariables
ENV TestName="Alexander Miles"
ENV TestPort=7002
COPY . .
CMD ["go", "run", "."]