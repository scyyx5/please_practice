#!/bin/sh
./plz-out/bin/common/docker/base.sh && docker build  -t hub.docker.com/image:2d6210db1a70cedc68c659925513f535130f4108bc0144a7efe6f4c8eb4bd349 -f Dockerfile - < plz-out/gen/hello_service/k8s/_image#docker_context.tar.gz
