FROM golang:1.14-alpine AS build

WORKDIR /src/ms-dynamic-qrcode/

COPY ./ /src/ms-dynamic-qrcode/

EXPOSE 8081
 
RUN CGO_ENABLED=0 go build -o /bin/ms-dynamic-qrcode

FROM scratch
COPY --from=build /bin/ms-dynamic-qrcode /bin/ms-dynamic-qrcode
ENTRYPOINT ["/bin/ms-dynamic-qrcode"]