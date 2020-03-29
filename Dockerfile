FROM scratch

COPY bin/calc /calc

ENTRYPOINT ["/calc"]