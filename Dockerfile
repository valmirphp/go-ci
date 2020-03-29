FROM scratch

# COPY gopath/bin/calc /calc
# COPY gopath/src/calc /calc

ADD . .

RUN ls ./

ENTRYPOINT ["/calc"]