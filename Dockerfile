from bufbuild/buf as proto

COPY proto /proto

COPY buf.gen.yaml /pb/
RUN buf generate /proto -o /pb --template /pb/buf.gen.yaml

COPY ui/buf.gen.yaml /uipb/
RUN buf generate /proto -o /uipb --template /uipb/buf.gen.yaml

from node:lts-alpine as frontend

run corepack enable

WORKDIR /app

COPY ui/package.json ui/pnpm-lock.yaml ./

RUN pnpm i

COPY ui .
COPY --from=proto /uipb/ src/pb

RUN pnpm build

FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=frontend /app/build ui/build
COPY --from=proto /pb .

RUN go build -o dburst . 

FROM alpine

COPY --from=builder /app/dburst /usr/bin

ENTRYPOINT ["dburst"]