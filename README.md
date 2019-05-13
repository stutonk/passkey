passkey
===
passkey is a simple utility for generating cryptographic keys from passwords
using the Argon2 algorithm. Keys are 64 hexadecimal characters (32 bytes 
when using binary output) long and are appended with the 256 character 
(128-byte) salt used to generate them.

### release binaries
are available [here](https://github.com/stutonk/passkey/releases) for amd64/all major OSes

### for unixes
`make && make install`

### everybody else
`go build` and drop it somewhere in your tree
