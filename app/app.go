package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-employee/config"
	"web-employee/internal/handler"
	"web-employee/internal/logger"
	"web-employee/internal/repository"
	repo "web-employee/internal/repository/employee"
	"web-employee/internal/service/employee"
)

type App struct {
	cfg                *config.Config
	conn               *pgx.Conn
	httpServer         *http.Server
	employeeRepository repository.EmployeeRepository
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	a.cfg = config.GetConfig()

	err := logger.SetupLogger(a.cfg.Env)
	if err != nil {
		return fmt.Errorf("failed to setup logger: %w", err)
	}
	log.Printf("Environment: %s", a.cfg.Env)

	if a.cfg.Db.Dsn == "" {
		return fmt.Errorf("database DSN is empty")
	}
	conn, err := pgx.Connect(ctx, a.cfg.Db.Dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	logger.Log.Print("Database connected")
	a.conn = conn
	a.employeeRepository = repo.NewEmployeeRepository(a.conn)
	service := employee.NewEmployeeService(a.employeeRepository)
	handlers := handler.NewHandler(service)
	a.httpServer = &http.Server{
		Addr:           a.cfg.Server.Address,
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return nil
}

func (a *App) Run() error {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Printf("Server error: %v", err)
		}
	}()
	logger.Log.Print("Server started on: " + a.cfg.Server.Address)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Log.Print("Server shutting down")
	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		logger.Log.Printf("Server Shutdown Error: %v", err)
		return err
	}
	if err := a.conn.Close(context.Background()); err != nil {
		logger.Log.Printf("Connection Close Error: %v", err)
		return err
	}
	return nil
}
