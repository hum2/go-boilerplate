# base
FROM golang:1.20.5-alpine as base

RUN apk update
RUN apk add build-base make
RUN apk add tzdata

ARG HOME_PATH="/go/src/app"
RUN mkdir -p $HOME_PATH
WORKDIR $HOME_PATH
COPY . .
RUN make setup

# local
FROM base as local

ENV ENV=local
ENV PROJECT_DIR=$HOME_PATH

RUN apk add bash git vim
RUN apk add nodejs npm

RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/golang/mock/mockgen@latest
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.53.3
RUN go get -d entgo.io/ent/cmd/ent
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT ["make", "start_local"]

# development
FROM base as dev_builder
RUN make build

FROM alpine as dev
ENV ENV=dev
COPY --from=dev_builder $HOME_PATH/bin/app /app
ENTRYPOINT ["/app"]
CMD ["server"]

# staging
FROM base as stg_builder
RUN make build

FROM alpine as stg
ENV ENV=stg
COPY --from=stg_builder $HOME_PATH/bin/app /app
ENTRYPOINT ["/app"]
CMD ["server"]

# production
FROM base as prd_builder
RUN make build_release

FROM alpine as prd
ENV ENV=prd
COPY --from=prd_builder $HOME_PATH/bin/app /app
ENTRYPOINT ["/app"]
CMD ["server"]
