module github.com/NpoolPlatform/authing-gateway

go 1.16

require (
	entgo.io/ent v0.11.2
	github.com/NpoolPlatform/api-manager v0.0.0-20220328101926-8907b2f76c6d
	github.com/NpoolPlatform/appuser-manager v0.0.0-20220820125511-0b87b81576ef
	github.com/NpoolPlatform/go-service-framework v0.0.0-20220812032117-44ecffa2bb95
	github.com/NpoolPlatform/libent-cruder v0.0.0-20220801075201-cab5db8b6290
	github.com/NpoolPlatform/login-gateway v0.0.0-20220328094651-99c681b06955
	github.com/NpoolPlatform/message v0.0.0-20220820095345-7f2da0b48358
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/stretchr/testify v1.7.1
	github.com/urfave/cli/v2 v2.4.0
	google.golang.org/grpc v1.48.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
