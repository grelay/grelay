module github.com/grelay/grelay/v1/benchmark

go 1.17

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/cep21/circuit v3.0.0+incompatible
	github.com/cep21/circuit/v3 v3.2.1
	github.com/grelay/grelay v0.0.1
	github.com/iand/circuit v0.0.4
	github.com/rubyist/circuitbreaker v2.2.1+incompatible
	github.com/sony/gobreaker v0.4.1
	github.com/streadway/handy v0.0.0-20200128134331-0f66f006fb2e
)

require (
	github.com/cenk/backoff v2.2.1+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/peterbourgon/g2s v0.0.0-20170223122336-d4e7ad98afea // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.6 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/grelay/grelay => ../
