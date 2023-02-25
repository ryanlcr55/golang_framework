package configs

type Server struct {
	GrpcPort string `mapstructure:"GRPC_PORT" default:"8080"`
	HttpPort string `mapstructure:"HTTP_PORT" default:"8080"`
}
