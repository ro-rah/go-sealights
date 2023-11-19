# Run like this:
#docker run -p 8080:8080 -e SEALIGHTS_TOKEN=your_token -e APP_NAME=your_app_name -e BRANCH_NAME=your_branch -e BUILD_NAME=your_build my-go-app



# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Download and extract SeaLights Go Agent and CLI
RUN echo "[Sealights] Downloading Sealights Golang & CLI Agents..." \
    && wget -nv -O sealights-go-agent.tar.gz https://agents.sealights.co/slgoagent/latest/slgoagent-linux-amd64.tar.gz \
    && wget -nv -O sealights-slcli.tar.gz https://agents.sealights.co/slcli/latest/slcli-linux-amd64.tar.gz \
    && tar -xzf ./sealights-go-agent.tar.gz && tar -xzf ./sealights-slcli.tar.gz \
    && rm -f ./sealights-go-agent.tar.gz ./sealights-slcli.tar.gz \
    && ./slgoagent -v 2> /dev/null | grep version && ./slcli -v 2> /dev/null | grep version

# Copy the local code to the container
COPY . .

# ARG directive for build-time variables
ARG SEALIGHTS_TOKEN
ARG APP_NAME
ARG BRANCH_NAME
ARG BUILD_NAME

# Set environment variables for SeaLights configuration
ENV SEALIGHTS_TOKEN=${SEALIGHTS_TOKEN}
ENV APP_NAME=${APP_NAME}
ENV BRANCH_NAME=${BRANCH_NAME}
ENV BUILD_NAME=${BUILD_NAME}

# Build the Go application
RUN echo "$SEALIGHTS_TOKEN" > ./sltoken.txt
RUN ./slcli config init --lang go --token ./sltoken.txt
RUN ./slcli config create-bsid --app $APP_NAME --branch $BRANCH_NAME --build $BUILD_NAME
RUN ./slcli scan --bsid buildSessionId.txt --path-to-scanner ./slgoagent --workspacepath ./ --scm git --scmVersion "0" --scmProvider github --disable-on-init true
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the executable with SeaLights agent
CMD ["sealights-go", "run", "./main"]
