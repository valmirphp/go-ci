FROM scratch

# COPY gopath/bin/calc /calc
# COPY gopath/src/calc /calc

ADD /workspace /tmp2

RUN ls /tmp2

# ENTRYPOINT ["/calc"]