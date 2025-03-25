module github.com/meesooqa/cheque/server

go 1.24.1

replace github.com/meesooqa/cheque/api => ../api

replace github.com/meesooqa/cheque/common => ../common

replace github.com/meesooqa/cheque/db => ../db

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	github.com/joho/godotenv v1.5.1
	github.com/meesooqa/cheque/api v0.0.0-00010101000000-000000000000
	github.com/meesooqa/cheque/common v0.0.0-00010101000000-000000000000
	github.com/meesooqa/cheque/db v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.71.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250313205543-e70fdf4c4cb4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250313205543-e70fdf4c4cb4 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.5.11 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
