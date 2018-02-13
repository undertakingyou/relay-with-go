# Relay with Go
This relay example is written in Go, and uses Go's built in http server to both
listen for and send requests.

For specific instructions on how to set up and use Go, please see
https://golang.org.

For a specific description of what this code is trying to solve, please see https://github.com/enderlabs/teem-code-examples/tree/master/teem-google-service-account-relay

## Build
Because Go is a compiled language, you will have to compile this in order to
deploy and run it. Here are the steps to do so.

### Make sure you have golang installed.
For whatever system you are planning to run the binary, you will want to build
this either on the same system, or a similar target system. For details on
installing Go please see https://golang.org/doc/install.

### Compile the binary
Navigate to the folder that contains relay.go. After any modifications that you
may need (see below for an example) run go build relay.go. The result will be a
new file in the directory, relay. This can then be moved wherever needed on
this system.

### Get publically signed certificates
Google requires that these pushes go to https domain with valid publically
accessible certs. This example code will look for a server.crt, and server.key
in the same file that the binary runs in. A good way to do this could be to use
letsencrypt.org to create a cert for you, and then either copy or link the
fullchain.pem and privkey.pem to server.crt and server.key respectively. The Go
application will start listening for the Google push endpoint on port 443 using
those certificate files.

### Decide on deployment method
The Go binary could be run in many ways. The code provided here as an example
will listen on port 443 and serve for the Teem google push endpoint. Modifying
the code could allow you to change this deployment method. No matter what you
choose to deploy this, please ensure that it has valid HTTPS. Examples could be

 * Using a proxy to pass information to the binary, such as nginx, haproxy, or
caddy. Please note, that with this method you will want to modify this code to
listen on a different port and not use TLS BEFORE you compile it.
 * Using a tool like supervisord to ensure that the binary continues to run
 * A docker container deployment. See https://blog.golang.org/docker for more
   details.

## Reference Links
The following links helped me in creating this example, and can help you as you
consider how to deploy this or similar Go code.
https://golang.org/
https://github.com/denji/golang-tls
https://www.kaihag.com/https-and-go/
https://www.reddit.com/r/golang/comments/424ple/deploying_a_go_https_server_in_real_life/
https://letsencrypt.org
https://certbot.eff.org


## Licenses
Go is licensed under a BSD license https://golang.org/LICENSE
This software is provided under an MIT license. Please see LICENSE.md
