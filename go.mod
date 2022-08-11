module github.com/NpoolPlatform/authing-gateway

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/NpoolPlatform/api-manager v0.0.0-20220121051827-18c807c114dc
	github.com/NpoolPlatform/appuser-manager v0.0.0-20220206145102-74eb783ed3b5
	github.com/NpoolPlatform/go-service-framework v0.0.0-20220127101433-3def4e496433
	github.com/NpoolPlatform/libent-cruder v0.0.0-20220801075201-cab5db8b6290
	github.com/NpoolPlatform/login-gateway v0.0.0-20220328094651-99c681b06955
	github.com/NpoolPlatform/message v0.0.0-20220802142552-733f92b9ee5c
	github.com/NpoolPlatform/third-gateway v0.0.0-20220808033714-72ad73b09ef3
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.47.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
