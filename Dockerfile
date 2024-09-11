# Stage 1: Build the third-party C programs and Go app
FROM ghcr.io/z3prover/z3:ubuntu-20.04-bare-z3-sha-c1454dc
#Remopve z3 default entrypoint
ENTRYPOINT []
# Install dependencies (C compiler, Python, Git)
RUN apt-get update && apt-get install -y \
    build-essential \
    python3 \
    python3-pip \
    git \
    wget \
    tar

# Install Go
RUN wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz -O /tmp/go1.22.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf /tmp/go1.22.0.linux-amd64.tar.gz \
    && rm /tmp/go1.22.0.linux-amd64.tar.gz


# Set up Go environment
ENV PATH="/usr/local/go/bin:$PATH"

# Clone and build Prover9
WORKDIR /lib/prover9
RUN git clone https://github.com/ai4reason/Prover9.git .
RUN make all

# Ensure p9 and mace4 are executable
RUN chmod +x /lib/prover9/bin/prover9
RUN chmod +x /lib/prover9/bin/mace4

# (Optional) Run tests for Prover9, could be skipped in production
RUN make test1 && make test2 && make test3

# Clone and build Z3
#WORKDIR /lib/z3
#RUN git clone https://github.com/Z3Prover/z3.git .
#RUN python3 scripts/mk_make.py --prefix=/usr/local
#WORKDIR /lib/z3/build
#RUN make && make install

# Build Go prover
WORKDIR /src
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ ./
RUN go build -o /usr/local/bin/ddlogic

# Copy only the necessary binaries from the build stage
COPY /src/config/config.json /usr/local/bin/config/config.json
COPY /lib/prover9/bin/prover9 /usr/local/bin/prover9
COPY /lib/prover9/bin/mace4 /usr/local/bin/mace4
# Expose the necessary port (if your Go app serves HTTP)
EXPOSE 8080 1433

# Set DB secrets
ENV DB_PORT=""
ENV DB_SERVER=""
ENV DB_NAME=""
ENV DB_USER=""
ENV DB_PASSWORD=""


# Command to run your Go application
CMD ["/usr/local/bin/ddlogic"]