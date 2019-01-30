FROM busybox

ENV HELLO world

RUN mkdir /www && touch /www/index.html

EXPOSE 8000

CMD httpd -p 8000 -h /www -f & wait

# docker build -t hello .
# docker run hello
# docker exec $(docker ps -q -l) cat /proc/1/environ
