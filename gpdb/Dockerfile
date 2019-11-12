# build stage
FROM golang:alpine AS dev_img
RUN apk update && apk upgrade && apk add --no-cache git
RUN go get -v \
    github.com/bmdoil/go-gpdb/gpdb \
    github.com/spf13/cobra \
    github.com/benhoyt/goawk/interp \
	github.com/benhoyt/goawk/parser \
	github.com/mholt/archiver \
	github.com/ryanuber/columnize \
    github.com/sirupsen/logrus

ENV APP_DIR=$GOPATH/src/gpdb
RUN mkdir -p ${APP_DIR}
COPY ./ ${APP_DIR}
WORKDIR ${APP_DIR}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -gcflags "all=-N -l" -o /gpdb

FROM alpine:latest AS prod_img
RUN apk update
COPY --from=dev_img /gpdb /usr/bin/
RUN adduser -h /home/gpadmin -s /bin/bash -D gpadmin gpadmin
USER gpadmin
WORKDIR /home/gpadmin
COPY ./config.yml ${WORKDIR}

CMD ["/bin/sh"]