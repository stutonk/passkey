package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/pflag"
	flag "github.com/spf13/pflag"
	"github.com/stutonk/boxutil"
)

const (
	errFmt   = "%v: fatal; %v\n"
	saltLen  = 128
	usageFmt = `usage: %v [-h, -v] [-s salt] [passphrase]
If no passphrase given, read from STDIN
Options are:
`
	verFmt  = "%v version %v\n"
	version = "1.2.0"
)

var (
	appName  = os.Args[0]
	helpFlag bool
	saltFlag setString
	verFlag  bool
)

type setString struct {
	set   bool
	value string
}

func (sf *setString) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *setString) String() string {
	return sf.value
}

func (sf *setString) Type() string {
	return "string"
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageFmt, appName)
		flag.PrintDefaults()
		fmt.Println()
	}

	flag.BoolVarP(
		&helpFlag,
		"help",
		"h",
		false,
		"display this help and exit",
	)
	flag.VarP(
		&saltFlag,
		"salt",
		"s",
		"provide salt as a hexidecimal string",
	)
	flag.BoolVarP(
		&verFlag,
		"version",
		"v",
		false,
		"output version information and exit",
	)
	flag.Parse()
}

func main() {
	switch {
	case verFlag:
		fmt.Printf(verFmt, appName, version)
		return
	case helpFlag:
		flag.Usage()
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, errFmt, appName, r)
		}
	}()

	var (
		err   error
		input []byte
		salt  []byte
	)
	if len(pflag.Args()) > 0 {
		input = []byte(strings.Join(pflag.Args(), " "))
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}
	if saltFlag.set {
		salt, err = hex.DecodeString(saltFlag.value)
		if err != nil {
			panic(err)
		}
	} else {
		salt = make([]byte, saltLen)
		if _, err := rand.Read(salt); err != nil {
			panic(err)
		}
	}

	fmt.Println(hex.EncodeToString((*boxutil.Passkey(input, salt))[:]))
	fmt.Println(hex.EncodeToString(salt))
}
