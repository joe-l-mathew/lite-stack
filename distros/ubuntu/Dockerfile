FROM ubuntu:latest

# Update and install necessary packages
RUN apt-get update && apt-get install -y \
    openssh-server \
    sudo \
    net-tools \
    && python3 \
    && rm -rf /var/lib/apt/lists/*

# Ensure the 'ubuntu' user exists, add it to the 'sudo' group, and change its password
RUN useradd -m -s /bin/bash ubuntu || true && \
    echo 'ubuntu:root' | chpasswd && \
    usermod -aG sudo ubuntu

# Create SSH run directory
RUN mkdir /var/run/sshd

# Expose SSH port
EXPOSE 22

# Run SSH in daemon mode
CMD ["/usr/sbin/sshd", "-D"]
