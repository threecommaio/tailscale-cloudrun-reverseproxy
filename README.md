# Cloudrun Tailscale Reverse Proxy

## Setup
1. Create a ephemeral key in Tailscale
2. Set `TAILSCALE_AUTHKEY` in your Cloud Run environment variables
3. Set `TARGET_URL` in your Cloud Run environment variables to the your device on your network.

Due to how ephemeral keys work, they do not allocate an IPv4, so you must hit the IPv6 address of the device. Thankfully Magic DNS makes this easy with the AAAA record. So you can set the `TARGET_URL` to something like: `http://[device].[namespace].io.beta.tailscale.net:8080`

## Usage
```bash
$ mage -l
Targets:
  build*    handles building the app
  clean     handles cleaning the project up
  deploy    handles deploying to cloud run
  push      handles pushing the image

* default target
```