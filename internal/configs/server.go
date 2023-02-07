package configs

type Server struct {
	GrpcProtocol string `mapstructure:"GRPC_PROTOCOL"`
	GrpcPort     string `mapstructure:"GRPC_PORT"`
}
