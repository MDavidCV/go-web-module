package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MDavidCV/go-web-module/internal/handler/controller"
	mw "github.com/MDavidCV/go-web-module/internal/handler/middleware"
	"github.com/MDavidCV/go-web-module/internal/repository"
	"github.com/MDavidCV/go-web-module/internal/service"
	"github.com/go-chi/chi/v5"
)

type ConfigSeverChi struct {
	// ServerAddress is the address where the server will listen and serve requests.
	ServerAddress string
	// LoaderFilePath is the path to the data that will be loaded into the server.
	LoaderFielPath string
	// Token is the token to validate the requests.
	Token string
}

type ServerChi struct {
	// ServerAddress is the address where the server will listen and serve requests.
	serverAddress string
	// LoaderFilePath is the path to the data that will be loaded into the server.
	loaderFilePath string
	// Token is the token to validate the requests.
	token string
}

func NewServerChi(cfg *ConfigSeverChi) *ServerChi {
	defaultConfig := &ConfigSeverChi{
		ServerAddress: ":8080",
		Token:         "12345",
	}

	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFielPath != "" {
			defaultConfig.LoaderFielPath = cfg.LoaderFielPath
		}
		if cfg.Token != "" {
			defaultConfig.Token = cfg.Token
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFielPath,
		token:          defaultConfig.Token,
	}
}

func (s *ServerChi) Run() error {
	storage := repository.NewStorageProduct("/Users/dcastrillonv/Documents/meli-boootcamp/go/go-web/go-web-module/docs/db/products.json")
	repository := repository.NewRepositoryProduct(nil, storage)
	service := service.NewServiceProduct(repository)
	controller := controller.NewProductController(service)

	router := chi.NewRouter()
	router.Use(mw.ResponseLoggerMid)

	router.Route("/products", func(r chi.Router) {
		// Public routes
		r.Group(func(r chi.Router) {
			r.Get("/", controller.GetProducts())
			r.Get("/{id}", controller.GetProductById())
			r.Get("/search", controller.SearchProduct())
			r.Get("/consumer_price", controller.GetConsumerPrice())
		})

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(mw.AuthValidationMid)
			r.Post("/", controller.CreateProduct())
			r.Put("/{id}", controller.UpdateProduct())
			r.Delete("/{id}", controller.DeleteProduct())
			r.Patch("/{id}", controller.UpdatePatchProduct())
		})
	})

	log.Printf("Server running on %s\n", s.serverAddress)
	if err := http.ListenAndServe(s.serverAddress, router); err != nil {
		return fmt.Errorf("error starting application: %w", err)
	}

	return nil
}
