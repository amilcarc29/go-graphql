package dependencies

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	routerHandler *mux.Router
}

// NewContainer returns a container with the dependencies
func NewContainer() (*Container, error) {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", os.Getenv("DBUsername"),
			os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBName")),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("DB connected successfully.")

	routerHandler := mux.NewRouter()

	return &Container{
		db:            db,
		routerHandler: routerHandler,
	}, nil
}

// Database returns database
func (container *Container) Database() *gorm.DB {
	return container.db
}

// RouterHandler returns router handler
func (container *Container) RouterHandler() *mux.Router {
	return container.routerHandler
}
