FROM ghcr.io/z3prover/z3:ubuntu-20.04-bare-z3-sha-c1454dc

ENTRYPOINT []

RUN apt-get update && apt-get install -y \
    build-essential \
    python3 \
    python3-pip \
    git \
    wget \
    tar

RUN wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz -O /tmp/go1.22.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf /tmp/go1.22.0.linux-amd64.tar.gz \
    && rm /tmp/go1.22.0.linux-amd64.tar.gz


ENV PATH="/usr/local/go/bin:$PATH"

WORKDIR /lib/prover9
RUN git clone https://github.com/ai4reason/Prover9.git .
RUN make all

RUN chmod +x /lib/prover9/bin/prover9
RUN chmod +x /lib/prover9/bin/mace4



RUN make test1 && make test2 && make test3

RUN mkdir -p /usr/local/bin/prover9/bin/

RUN cp /lib/prover9/bin/prover9 /usr/local/bin/prover9/bin/
RUN cp /lib/prover9/bin/mace4 /usr/local/bin/prover9/bin/


WORKDIR /src
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ ./
RUN go build -o /usr/local/bin/ddlogic

COPY /config/config.json /usr/local/bin/config/config.json
COPY /prover9/outputs_out/ /usr/local/bin/prover9/outputs_out
COPY /prover9/inputs_in/ /usr/local/bin/prover9/inputs_in
COPY /prover9/conf/ /usr/local/bin/prover9/conf



EXPOSE 8080 1433

# DB secrets
ENV DB_PORT=""
ENV DB_SERVER=""
ENV DB_NAME=""
ENV DB_USER=""
ENV DB_PASSWORD=""
ENV IS_AZURE="true"


WORKDIR /usr/local/bin
CMD ["/usr/local/bin/ddlogic"]