#!/bin/sh
docker build  -t hub.docker.com/base:54c9fc09fa64a4e1039f6f46024c261e5743d916e69fdc96831bb044908fc099 -f Dockerfile-base - < plz-out/gen/common/docker/_base#docker_context.tar.gz
