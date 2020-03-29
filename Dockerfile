FROM scratch

COPY gopath/bin/calc /calc

ENTRYPOINT ["/calc"]