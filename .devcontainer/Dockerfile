# See here for image contents: https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/Dockerfile

# This is pinned to a particular version of go:
FROM mcr.microsoft.com/devcontainers/go:1.23

# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETARCH

# APT dependencies
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends bash-completion software-properties-common lsb-release graphviz zip nodejs npm python3-pip \
    # install az-cli
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash -

# install docker
# - not yet needed?
# RUN curl -fsSL https://get.docker.com | sh -
# RUN usermod -aG docker vscode

COPY install-dependencies.sh .
RUN ./install-dependencies.sh devcontainer && rm install-dependencies.sh

# Setup envtest
# NB: if you change this, dev.sh also likely needs updating, also need to update the env below
RUN setup-envtest use 1.27.1 --bin-dir /usr/local/envtest/bin
# Set the KUBEBUILDER_ASSETS variable. Ideally we could do source <(setup-envtest use --bin-dir /usr/local/envtest/bin -i -p env)
# but there's no way to dynamically set an env variable for all container users.
# Many usages of this container are from docker exec task <x>, where no shell is invoked and the entrypoint is not run
# (entrypoint is only run on start, not on exec). Due to that, the following approaches do not work:
# - ~/.bashrc - only works for one user in a shell but we must support -u $(id -u ${USER}):$(id -g ${USER}) which means the container could run as more than 1 user
# - /etc/profile or /etc/profile.d - only works for one user in a login shell
ENV KUBEBUILDER_ASSETS=/usr/local/envtest/bin/k8s/1.27.1-linux-${TARGETARCH}
ENV PATH=$KUBEBUILDER_ASSETS:$PATH

# Make kubectl completions work with 'k' alias
RUN echo 'alias k=kubectl' >> "/etc/bash.bashrc"
RUN echo 'complete -F __start_kubectl k' >> "/etc/bash.bashrc"
RUN echo 'source <(kubectl completion bash)' >> "/etc/bash.bashrc"

# Setup go-task completions
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/v3.2.0/completion/bash/task.bash" > "/etc/.task.completion.sh" \
    && echo 'source /etc/.task.completion.sh' >> "/etc/bash.bashrc"

ENV KIND_CLUSTER_NAME=aso

# install docker, from: https://github.com/microsoft/vscode-dev-containers/blob/main/script-library/docs/docker.md
COPY library-scripts/docker-debian.sh /tmp/library-scripts/

# these are all the default values except for the last one (USE_MOBY) which is what we want to set.
# currently packages.microsoft.com does not have Moby packages for Debian Bullseye, which the
# devcontainer image is based on: https://github.com/microsoft/vscode-dev-containers/issues/1008
RUN bash /tmp/library-scripts/docker-debian.sh true /var/run/docker-host.sock /var/run/docker.sock automatic false

ENTRYPOINT ["/usr/local/share/docker-init.sh"]
CMD ["sleep", "infinity"]
