FROM daocloud.io/ubuntu:trusty
MAINTAINER water-law <dockerwaterlaw@daocloud.io>
RUN apt-get update && \
    apt-get install -y python3 \
                        python3-dev \
                        python3-pip \
    && apt-get clean \
    && apt-get autoclean \
    && rm -rf /var/lib/apt/lists/*

RUN pip3 install --upgrade pip
RUN pip3 install schedule
RUN pip3 install oss2


WORKDIR /runner
COPY start.py /runner/
CMD ["python3 /runner/start.py"]
ENTRYPOINT ["sh","-c","python3 /runner/start.py"]

