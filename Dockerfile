FROM rockylinux:9
RUN dnf update -y
RUN dnf install -y git
COPY argo-image-update /usr/bin/
RUN chmod 777 /usr/bin/argo-image-update 