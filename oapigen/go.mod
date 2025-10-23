module github.com/HewlettPackard/hpe-morpheus-go-sdk

go 1.24

require gopkg.in/validator.v2 v2.0.1

require github.com/go-viper/mapstructure/v2 v2.4.0

require (
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
)

replace github.com/go-viper/mapstructure/v2 v2.4.0 => github.com/ZoeySimone/mapstructure/v2 v2.4.1-0.20250902104126-099f6e5a61b8
