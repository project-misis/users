// Package application
package application

import (
	"net"

	"echo-template/internal/delivery/handlers"
	"echo-template/internal/infrastructure"
	"echo-template/internal/infrastructure/database"
	"echo-template/internal/infrastructure/logger"
	"echo-template/internal/infrastructure/repository"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"

	"buf.build/go/protovalidate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/project-misis/users_proto/pb"

	"google.golang.org/grpc"
)

type Application struct {
	Server  *grpc.Server
	Address string
	DB      *pgxpool.Pool
	Logger  *logger.Logger
}

func NewApplication(config *infrastructure.Config) *Application {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	interceptor := protovalidate_middleware.UnaryServerInterceptor(validator)
	gs := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	l := logger.NewLogger()

	db, err := database.NewPostgresDB(config, l)
	if err != nil {
		l.Errorf("failed to connect to database: %s", err.Error())
		return nil
	}

	repo := repository.NewUserRepository(db)
	s := handlers.NewUserService(repo)
	pb.RegisterCrudServer(gs, s)
	return &Application{
		Server:  gs,
		Address: config.Port,
		DB:      db,
		Logger:  l,
	}
}

func (a *Application) RunServer() error {
	a.Logger.Info("Starting server on " + a.Address)

	lis, err := net.Listen("tcp", a.Address)
	if err != nil {
		a.Logger.Fatalf("failed to listen: %v", err)
	}
	if err := a.Server.Serve(lis); err != nil {
		a.Logger.Fatalf("Failed to start server: %s", err.Error())
	}
	return nil
}
