# syntax=docker/dockerfile:1
FROM golang:1.18beta1-buster
WORKDIR /simpleenvironmentvariables
ENV TestName="Alexander Miles"
ENV TestPort=7002
COPY . .
CMD ["go", "run", "."]