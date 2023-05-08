FROM node:16-alpine3.11

RUN apk add --no-cache bc
RUN apk add --update coreutils && rm -rf /var/cache/apk/*

WORKDIR /app

ADD script.sh /bin/scripts.sh
RUN chmod +x /bin/scripts.sh

ENTRYPOINT ["sh","/bin/scripts.sh"]
