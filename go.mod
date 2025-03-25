module worktools

go 1.23.0

require github.com/mniak/hsmlib v0.0.0

replace github.com/mniak/hsmlib => ../hsmlib

require (
	github.com/mniak/krypton v0.0.4
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/codefresh-io/go-sdk v1.4.7 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)

replace github.com/mniak/krypton => ../krypton
