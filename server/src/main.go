package main


import (
	"log"
	"net/http"
	"github.com/dirkarnez/mentor/service"
	"github.com/dirkarnez/mentor/middleware"
	pb "github.com/dirkarnez/mentor/proto"
	"github.com/go-chi/chi"

	"github.com/dirkarnez/mentor/proxy"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id          int
    LoginName   string
    Password	string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Id          int
    Name   		string
    Age         int16
}

/*
CREATE DATABASE mentor
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;

USE mentor
  
SELECT *
FROM user u INNER JOIN profile p ON u.profile_id = p.id
*/
var (
	o orm.Ormer
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
   	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(mariadb:3306)/mentor?charset=utf8")
	orm.RegisterModel(new(User), new(Profile))
}

func main() {
	o = orm.NewOrm()
    o.Using("default") 

	err := orm.RunSyncdb("default", true, true)
	if err != nil {
	    log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMentorServiceServer(grpcServer, service.NewMentorService())
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		middleware.NewGrpcWebMiddleware(wrappedGrpc).Handler, // Must come before general CORS handling
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}).Handler,
	)

	router.Get("/article-proxy", proxy.Article)

	if err := http.ListenAndServe(":8900", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}
