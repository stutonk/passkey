package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/stutonk/boxutil"
)

const (
	errFmt   = "%v: fatal; %v\n"
	saltLen  = 128
	usageFmt = `usage: %v [-h, -v] [-b] [-p password]
If -p is not given, read from STDIN
Options are:
`
	verFmt  = "%v version %v\n"
	version = "1.0.0"
)

var (
	appName  = os.Args[0]
	binFlag  bool
	helpFlag bool
	passFlag setString
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
		&binFlag,
		"binary",
		"b",
		false,
		"output in binary mode",
	)
	flag.BoolVarP(
		&helpFlag,
		"help",
		"h",
		false,
		"display this help and exit",
	)
	flag.VarP(
		&passFlag,
		"password",
		"p",
		"specify the password as an argument",
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
	)
	if passFlag.set {
		input = []byte(passFlag.value)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	output := append((*boxutil.Passkey(input, salt))[:], salt...)
	if binFlag {
		os.Stdout.Write(output)
	} else {
		fmt.Println(hex.EncodeToString(output))
	}
}
