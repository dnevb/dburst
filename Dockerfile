from bufbuild/buf as proto

COPY proto /proto

COPY buf.gen.yaml /pb/
RUN buf generate /proto -o /pb --template /pb/buf.gen.yaml

COPY ui/buf.gen.yaml /uipb/
RUN buf generate /proto -o /uipb --template /uipb/buf.gen.yaml

from node:lts-alpine as frontend

run corepack enable

WORKDIR /app

COPY ui .
COPY --from=proto /uipb/ src/pb

RUN pnpm i

RUN pnpm build

FROM golang:alpine

WORKDIR /app

COPY . .

COPY --from=frontend /app/build ui/build
COPY --from=proto /pb .

RUN go mod tidy
RUN go build -o main . 

ENTRYPOINT ["./main"]