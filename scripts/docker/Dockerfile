FROM alpine:3.2

ENV PYTHON_VERSION=2.7.12-r0
ENV PY_PIP_VERSION=6.1.1-r0
ENV SUPERVISOR_VERSION=3.2.0

RUN apk update \
	&& apk add -u python=$PYTHON_VERSION py-pip=$PY_PIP_VERSION \
	&& apk add ca-certificates \
    && update-ca-certificates

RUN mkdir -p /var/log/supervisord/
RUN pip install supervisor==$SUPERVISOR_VERSION

# get maintenance binary
RUN wget -O /tmp/pkg.zip https://github.com/BuiltLabs/maintenance-cli/releases/download/v0.1.0/linux_amd64.zip \
	&& unzip /tmp/pkg.zip -d /usr/local/bin

COPY supervisord.conf /etc/supervisord.conf

ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]