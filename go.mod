module worktools

go 1.23.0

require github.com/mniak/hsmlib v0.0.0

replace github.com/mniak/hsmlib => ../hsmlib

require (
	github.com/codefresh-io/go-sdk v1.4.7
	github.com/mniak/krypton v0.0.4
	github.com/samber/lo v1.49.1
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/mniak/krypton => ../krypton
