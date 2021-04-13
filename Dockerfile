FROM alpine:3.5 

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY ./genshin-daily-login .

CMD [ "./genshin-daily-login" ]