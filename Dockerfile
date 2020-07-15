FROM golang:1.13.7-alpine3.11 as binaryBuilder

RUN \
  echo -e "\e[32madd build dependency packages\e[0m" \
  && apk --no-cache add ca-certificates git

WORKDIR /go/src/service_owner_api
COPY golang_binary/. .

RUN \
  echo -e "\e[32m'go get' all build dependencies\e[0m" \
  && go get -v -d ./... \
  \
  && echo -e "\e[32mBuild the binary\e[0m" \
  && env GOOS=linux GOARCH=386 go build -v -o main main.go

FROM scratch

LABEL \
  build-date="2020-02-26" \
  description="Service Owner API" \
  maintainer="application-support@polarisalpha.com" \
  name="service_owner_api" \
  vendor="Parsons" \
  version=""

WORKDIR /bin
COPY --from=binaryBuilder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=binaryBuilder /go/src/service_owner_api/main .
COPY --from=binaryBuilder /go/src/service_owner_api/service_owner.yml .
EXPOSE 9858
CMD ["/bin/main"]
