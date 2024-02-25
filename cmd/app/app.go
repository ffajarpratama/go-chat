package app

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ffajarpratama/go-chat/config"
	"github.com/ffajarpratama/go-chat/internal/repository"
	"github.com/ffajarpratama/go-chat/internal/transporter"
	"github.com/ffajarpratama/go-chat/internal/usecase"
	"github.com/ffajarpratama/go-chat/pkg/postgres"
	"github.com/ffajarpratama/go-chat/pkg/util"
)

func StartServer() (err error) {
	cnf := config.New()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	util.SetTimeZone("UTC")

	db, err := postgres.NewDBClient(cnf)
	if err != nil {
		return err
	}

	repo := repository.New(db)
	uc := usecase.New(repo)
	handler := transporter.NewV1HttpTransporter(uc)

	log.Println("server started on:", cnf.App.URL)
	return http.ListenAndServe(fmt.Sprintf(":%d", cnf.App.Port), handler)
}
