FROM busybox

MAINTAINER Bryan Lee, <dragonme@gmail.com> 

COPY ./tiny-goweb /home/
COPY ./entrypoint.sh /

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
