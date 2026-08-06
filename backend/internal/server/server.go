package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"gooffer/backend/internal/delivery/handlers"
	"gooffer/backend/internal/delivery/middleware"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

type Dependencies struct {
	Logger         *slog.Logger
	Addr           string
	ProfileHandler *handlers.ProfileHandler
	RecapHandler   *handlers.RecapHandler
}

func New(deps Dependencies) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("GET /api/profiles", deps.ProfileHandler.List)
	mux.HandleFunc("GET /api/profiles/{id}", deps.ProfileHandler.GetByID)

	mux.HandleFunc("POST /api/recap/generate", deps.RecapHandler.Generate)
	mux.HandleFunc("GET /api/recap/{user_id}/{year}", deps.RecapHandler.Get)
	mux.HandleFunc("GET /api/recap/{user_id}/{year}/share", deps.RecapHandler.Share)

	var handler http.Handler = mux
	handler = middleware.Recovery(deps.Logger)(handler)
	handler = middleware.Logger(deps.Logger)(handler)
	handler = middleware.CORS(handler)

	return &Server{
		logger: deps.Logger,
		httpServer: &http.Server{
			Addr:              deps.Addr,
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info("http server starting", slog.String("addr", s.httpServer.Addr))
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
