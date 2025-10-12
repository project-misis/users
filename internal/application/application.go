package application

import (
	"net"

	"echo-template/internal/infrastructure"
	"echo-template/internal/infrastructure/database"
	"echo-template/internal/infrastructure/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	// "github.com/project-misis/users_proto/pb"
	"google.golang.org/grpc"
)

type Application struct {
	Server  *grpc.Server
	Address string
	DB      *pgxpool.Pool
	Logger  *logger.Logger
}

func NewApplication(config *infrastructure.Config) *Application {
	gs := grpc.NewServer()
	l := logger.NewLogger() // Создаём логгер

	db, err := database.NewPostgresDB(config, l)
	if err != nil {
		l.Errorf("failed to connect to database: %s", err.Error())
		return nil
	}

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
