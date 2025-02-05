package server

import (
    "database/sql"
    "dockMon/config"
    "dockMon/internal/domain/interfaces/infrastructure"
    "dockMon/internal/domain/interfaces/services"
    "dockMon/internal/handlers/machines"
    "dockMon/internal/handlers/middlewares"
    "dockMon/internal/infrastructure/postgres"
    serviceImpl "dockMon/internal/services"
    "fmt"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    _ "github.com/lib/pq"
    "go.uber.org/zap"
    "log"
    "net/http"
    "os"
    "time"
)

type Server struct {
    cfg *config.Config
    db  *sql.DB

    manager     services.Manager
    machineRepo infrastructure.MachineRepository

    server *http.Server
    router *chi.Mux
}

func (s *Server) prepare(path string) {
    cfg, err := config.Parse(path)
    if err != nil {
        log.Fatal(err)
    }
    s.cfg = cfg
}

func (s *Server) Init(path string) {
    s.prepare(path)
    s.dbConnect()

    s.machineRepo = postgres.NewMachineRepository(s.db)
    s.manager = serviceImpl.NewMachinesService(s.machineRepo)

    s.initRouter()
    s.server = &http.Server{
        Addr:         fmt.Sprintf("%s:%s", s.cfg.App.Host, s.cfg.App.Port),
        Handler:      s.router,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
}

func (s *Server) dbConnect() {
    user, password := os.Getenv(s.cfg.App.Db.UserEnvKey), os.Getenv(s.cfg.App.Db.PasswordEnvKey)
    if user == "" {
        log.Fatalf("%s env variable not set", s.cfg.App.Db.UserEnvKey)
    }
    if password == "" {
        log.Fatalf("%s env variable not set", s.cfg.App.Db.PasswordEnvKey)
    }
    db, err := sql.Open("postgres", toDbUrl(user, password, s.cfg.App.Db.Host, s.cfg.App.Db.Port, s.cfg.App.Db.Dbname, s.cfg.App.Db.Sslmode))
    if err != nil {
        log.Fatal(err)
    }
    s.db = db
    err = postgres.PerformMigration(s.db)
    if err != nil {
        log.Fatal(err)
    }
}

func (s *Server) initRouter() {
    logger := zap.Must(zap.NewProduction())
    router := chi.NewRouter()
    router.Use(middlewares.InjectLogger(logger))
    router.Use(middleware.RequestID)
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    router.Route("/api", func(r chi.Router) {
        r.Get("/machines", machines.GetMachines(s.manager))
        r.Put("/machine", machines.PutMachine(s.manager))
    })
    s.router = router
}

func toDbUrl(user, password, host, port, dbname, sslmode string) string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
}

func (s *Server) Run() {
    if err := s.server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
