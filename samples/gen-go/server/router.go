package server

// Code auto-generated. Do not edit.

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/Clever/go-process-metrics/metrics"
	"github.com/Clever/wag/middleware"
	"github.com/gorilla/mux"
	"gopkg.in/Clever/kayvee-go.v5/logger"
	kvMiddleware "gopkg.in/Clever/kayvee-go.v5/middleware"
	"gopkg.in/tylerb/graceful.v1"
)

type contextKey struct{}

// Server defines a HTTP server that implements the Controller interface.
type Server struct {
	// Handler should generally not be changed. It exposed to make testing easier.
	Handler http.Handler
	addr    string
	l       *logger.Logger
}

// Serve starts the server. It will return if an error occurs.
func (s *Server) Serve() error {

	go func() {
		metrics.Log("swagger-test", 1*time.Minute)
	}()

	go func() {
		// This should never return. Listen on the pprof port
		log.Printf("PProf server crashed: %s", http.ListenAndServe(":6060", nil))
	}()

	s.l.Counter("server-started")

	// Give the sever 30 seconds to shut down
	return graceful.RunWithErr(s.addr, 30*time.Second, s.Handler)
}

type handler struct {
	Controller
}

func withMiddleware(serviceName string, router http.Handler, m []func(http.Handler) http.Handler) http.Handler {
	handler := router

	// Wrap the middleware in the opposite order specified so that when called then run
	// in the order specified
	for i := len(m) - 1; i >= 0; i-- {
		handler = m[i](handler)
	}
	handler = middleware.Tracing(handler)
	handler = middleware.Panic(handler)
	// Logging middleware comes last, i.e. will be run first.
	// This makes it so that other middleware has access to the logger
	// that kvMiddleware injects into the request context.
	handler = kvMiddleware.New(handler, serviceName)
	return handler
}

// New returns a Server that implements the Controller interface. It will start when "Serve" is called.
func New(c Controller, addr string) *Server {
	return NewWithMiddleware(c, addr, []func(http.Handler) http.Handler{})
}

// NewWithMiddleware returns a Server that implemenets the Controller interface. It runs the
// middleware after the built-in middleware (e.g. logging), but before the controller methods.
// The middleware is executed in the order specified. The server will start when "Serve" is called.
func NewWithMiddleware(c Controller, addr string, m []func(http.Handler) http.Handler) *Server {
	router := mux.NewRouter()
	h := handler{Controller: c}

	l := logger.New("swagger-test")

	router.Methods("GET").Path("/v1/books").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "getBooks")
		h.GetBooksHandler(r.Context(), w, r)
		ctx := middleware.WithTracingOpName(r.Context(), "getBooks")
		r = r.WithContext(ctx)
	})

	router.Methods("POST").Path("/v1/books").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "createBook")
		h.CreateBookHandler(r.Context(), w, r)
		ctx := middleware.WithTracingOpName(r.Context(), "createBook")
		r = r.WithContext(ctx)
	})

	router.Methods("GET").Path("/v1/books/{book_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "getBookByID")
		h.GetBookByIDHandler(r.Context(), w, r)
		ctx := middleware.WithTracingOpName(r.Context(), "getBookByID")
		r = r.WithContext(ctx)
	})

	router.Methods("GET").Path("/v1/books2/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "getBookByID2")
		h.GetBookByID2Handler(r.Context(), w, r)
		ctx := middleware.WithTracingOpName(r.Context(), "getBookByID2")
		r = r.WithContext(ctx)
	})

	router.Methods("GET").Path("/v1/health/check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "healthCheck")
		h.HealthCheckHandler(r.Context(), w, r)
		ctx := middleware.WithTracingOpName(r.Context(), "healthCheck")
		r = r.WithContext(ctx)
	})

	handler := withMiddleware("swagger-test", router, m)
	return &Server{Handler: handler, addr: addr, l: l}
}
