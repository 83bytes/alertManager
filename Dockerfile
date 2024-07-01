FROM gcr.io/distroless/base-debian11 

WORKDIR /

COPY ./alertmanager ./alertmanager

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./alertmanager"]