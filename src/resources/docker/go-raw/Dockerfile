FROM golang:1.19-alpine3.16

RUN apk add --no-cache bc
RUN apk add --update coreutils && rm -rf /var/cache/apk/*

WORKDIR /app
ADD script.sh /bin/scripts.sh
RUN chmod +x /bin/scripts.sh

ENTRYPOINT ["sh","/bin/scripts.sh"]
