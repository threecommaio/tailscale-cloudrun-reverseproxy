#!/bin/sh

TARGET_URL=$1
/app/tailscaled --tun=userspace-networking --socks5-server=localhost:1055 &
until /app/tailscale up --authkey=${TAILSCALE_AUTHKEY} --hostname=cloudrun-app
do
    sleep 0.1
done
echo Tailscale started
echo "TargetURL: $TARGET_URL"
ALL_PROXY=socks5://localhost:1055/ /app/app -targetURL $TARGET_URL