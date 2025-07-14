module worktools

go 1.23.0

replace github.com/mniak/hsmlib => ../hsmlib

require (
	github.com/codefresh-io/go-sdk v1.4.7
	github.com/samber/lo v1.49.1
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/mniak/krypton => ../krypton
