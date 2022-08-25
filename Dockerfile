FROM alpine
ADD grocer /grocer
ENTRYPOINT [ "/grocer" ]
