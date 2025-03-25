module github.com/meesooqa/cheque/api

go 1.24.1

replace github.com/meesooqa/cheque/common => ../common

replace github.com/meesooqa/cheque/db => ../db

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	github.com/meesooqa/cheque/common v0.0.0-00010101000000-000000000000
	github.com/meesooqa/cheque/db v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250313205543-e70fdf4c4cb4
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.6
	gorm.io/gorm v1.25.12
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250313205543-e70fdf4c4cb4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
