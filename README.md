[![Go Report Card](https://goreportcard.com/badge/github.com/stutonk/passkey)](https://goreportcard.com/report/github.com/stutonk/passkey)  
passkey is a simple utility for generating cryptographic keys from a
passphrase using the Argon2 algorithm. Passphrases may be any string when
passed as an argument or any sequence of bytes when read from STDIN. The
program will automatically generate a fresh 128-byte salt if one is not
provided as an option. Keys are output as 64 hexadecimal characters. The salt
that was used to generate the key is output as 256 hexidecimal characters on
the subsequent line.

```
usage: passkey [-h, -v] [-s salt] [passphrase]
If no passphrase given, read from STDIN
Options are:
  -h, --help          display this help and exit
  -s, --salt string   provide salt as a hexidecimal string
  -v, --version       output version information and exit
```

### release binaries
are available [here](https://github.com/stutonk/passkey/releases) for amd64/all major OSes

### for unixes
`make && make install`

### everybody else
`go build`

### license
```
##########
###    ###
##  ##  ##   To the extent possible under law,
#  # ##  #   the authors have waived all copyright
#  # ##  #   and related and neighboring rights to
#  ## #  #   this work.  For more information, please see
#  ## #  #   <https://creativecommons.org/publicdomain/zero/1.0/>.
##  ##  ##
###    ### 
##########
```