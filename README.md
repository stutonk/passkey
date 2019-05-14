[![Go Report Card](https://goreportcard.com/badge/github.com/stutonk/passkey)](https://goreportcard.com/report/github.com/stutonk/passkey)  
passkey
===
passkey is a simple utility for generating cryptographic keys from a
passphrase using the Argon2 algorithm. Passphrases may be any string when
passed as an argument or any sequence of bytes when read from STDIN. The
program will automatically generate a fresh 128-byte salt if one is not
provided as an option. Keys are output as 64 hexadecimal characters (32
bytes when using binary output) prepended to the salt that was used to
generate them.

### release binaries
are available [here](https://github.com/stutonk/passkey/releases) for amd64/all major OSes

### for unixes
`make && make install`

### everybody else
`go build` and drop it somewhere in your tree
