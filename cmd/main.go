package main

import (
	"database/sql"
	_"github.com/lib/pq"
	"log/slog"
	"os"
)


func main(){
	conn:="postgres://postgres:2005@localhost/fitness?sslmode=disable"
	logger:=slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{
		AddSource: true,
	}))
	db,err:=sql.Open("postgres",conn)
	if err!=nil{
		logger.Error("fialed to connect")
		os.Exit(1)
	}
	defer db.Close()

	if err:=db.Ping();err!=nil{
		logger.Error("fialed to ping")
		os.Exit(1)
	}

}