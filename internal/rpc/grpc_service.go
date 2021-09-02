package rpc

import (
	"context"
	"net"

	pb "github.com/ffo32167/weather/api"
	p "github.com/ffo32167/weather/internal/processer"
	ch "github.com/ffo32167/weather/internal/storage"
	tp "github.com/ffo32167/weather/internal/types"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// GrpcServer реализует сервер grpc
type GrpcServer struct {
	pb.UnimplementedWeatherParserServer
	config *tp.Config
	cache  ch.Cacher
}

// ServerStart запускает сервер GRPC
func ServerStart(cfg *tp.Config, cache ch.Cacher) {
	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("can't listen port")
	}
	serverGRPC := grpc.NewServer()
	serv := NewGrpcServer(cfg, cache)
	pb.RegisterWeatherParserServer(serverGRPC, serv)
	if err := serverGRPC.Serve(lis); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("can't start grpc Server")
	}
	logrus.Info("server started")
}

// NewGrpcServer создаёт сервер из конфига и кэша
func NewGrpcServer(cfg *tp.Config, cache ch.Cacher) (s *GrpcServer) {
	return &GrpcServer{config: cfg, cache: cache}
}

// ProcessGRPCRequest обрабатывает grpc запрос
func (server *GrpcServer) ProcessGRPCRequest(ctx context.Context, params *pb.WeatherParams) (*pb.DayWeather, error) {
	logrus.WithFields(logrus.Fields{"params": params, "config": server.config}).Debug("grpc params")
	buffer, format := p.ProcessRequest(
		tp.WeatherParams{
			MonthsNumbers: params.MonthsNumbers,
			Cities:        params.Cities,
			Site:          params.Site,
			Months:        params.Months,
			Year:          params.Year,
			ReplyFormat:   params.ReplyFormat,
		},
		server.config,
		server.cache)
	return &pb.DayWeather{ComparisonCSV: buffer.Bytes(), Format: format}, nil
}
