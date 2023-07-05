module github.com/gitpod-io/gitpod/code-desktop/status

go 1.17

require google.golang.org/grpc v1.53.0

require github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect

require (
	github.com/gitpod-io/gitpod/supervisor/api v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/gitpod-io/gitpod/supervisor/api => ../../../supervisor-api/go // leeway
