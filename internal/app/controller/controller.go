package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hhly234/passport-tracking/internal/app/home"
	"github.com/hhly234/passport-tracking/internal/pkg/db"
	"github.com/hhly234/passport-tracking/internal/pkg/db/mongodb"
	"github.com/hhly234/passport-tracking/internal/pkg/middleware"
	"github.com/hhly234/passport-tracking/internal/pkg/types"
)

type static struct {
	prefix string
	dir    string
}

// DBConns connections to database
type DBConns struct {
	Database db.Connections
}

// InitRoute all routes
func InitRoute(conns *DBConns) (http.Handler, error) {
	routes := []types.Route{}
	home.NewRouter(&routes)

	r := mux.NewRouter()
	r.Use(middleware.Log)
	for _, rt := range routes {
		r.HandleFunc(rt.Path, rt.Handler).Methods(rt.Method)
	}

	s := static{"/static/", "web/static"}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}

// InitDB connections to database
func InitDB(conf *types.Server) *DBConns {
	conns := &DBConns{}
	conns.Database.Type = conf.DB.Type

	switch conf.DB.Type {
	case db.TypeMongoDB:
		s, err := mongodb.Dial(&conf.DB.ConfigDB)
		if err != nil {
			log.Panicln("failed to dial to target server, err:", err)
		}
		conns.Database.MongoDB = s
	}
	return conns
}

// Close all underlying connections
func (c *DBConns) Close() {
	c.Database.Close()
}
