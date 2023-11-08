FROM alpine

COPY gRPCMail /usr/local/bin/gRPCMail

EXPOSE 8080

ENTRYPOINT ["gRPCMail"]