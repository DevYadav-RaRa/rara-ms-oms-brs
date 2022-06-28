FROM --platform=linux/amd64 ubuntu:20.04

# Update aptitude
RUN apt-get update

# install core deps
RUN apt-get install -y autoconf automake gcc cpp git zip unzip openjdk-8-jdk-headless

#Install golang
RUN mkdir -p /opt
WORKDIR /opt
RUN apt-get install -y wget
ADD https://go.dev/dl/go1.18.3.linux-amd64.tar.gz /opt/
RUN ls -lart
RUN tar -xvf go1.18.3.linux-amd64.tar.gz
RUN ln -sf /opt/go1.18.3.linux-amd64 go
RUN ls -lart
# replace shell with bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
ENV GO_INST_PATH /opt/go
ENV PATH $GO_INST_PATH/bin:$PATH

# confirm installation
RUN go version

#Mount project
WORKDIR /root
RUN mkdir app
VOLUME /root/app

#RUN prod
ENV GO_ENV development
ENV GO_PORT 8181
WORKDIR /root/app
EXPOSE 8181

# Start
RUN mkdir /logs
VOLUME /logs

#Skynet mapping with host machine
RUN mkdir /root/.skynet
VOLUME /root/.skynet

#CMD nohup node start parser_app >> /logs/node.out 2>&1&
CMD ./prun.sh >> /logs/node.out 2>&1
