# This docker file just written for test of MultiProcessController
FROM yindaheng98/go-git
ADD . .

RUN go get -d -v ./... && \
    cd MultiProcessController/main && \
    go build -v -o /mpc

FROM egaillardon/jmeter
STOPSIGNAL SIGINT

COPY --from=0 /mpc /jmeter
COPY entrypoint.sh /usr/local/bin/
WORKDIR /jmeter

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/jmeter/mpc"]