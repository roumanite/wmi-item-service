# Test and build
FROM golang as builder

ARG USER=wmi_user
ARG WORK_DIR=/app

ENV CGO_ENABLED 0

RUN apt-get update -qq && apt-get upgrade -y

RUN groupadd ${USER} && useradd -m -g ${USER} -l ${USER}
RUN mkdir -p ${WORK_DIR} && chown -R ${USER}:${USER} ${WORK_DIR}

WORKDIR ${WORK_DIR}
USER ${USER}

COPY . ${WORK_DIR}

RUN go mod verify

RUN go test -cover ./...

RUN go build -o app ./cmd/httpd/main.go

FROM alpine

ARG USER=wmi_user
ARG WORK_DIR=/app
ARG PORT=8080

ENV GIN_MODE=release
ENV PORT=${PORT}

RUN addgroup -S ${USER} && adduser -S ${USER} -G ${USER}

WORKDIR ${WORK_DIR}

COPY --chown=0:0 --from=builder ${WORK_DIR}/app ${WORK_DIR}

USER ${USER}
EXPOSE ${PORT}

CMD ["./app"]
