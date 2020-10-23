package main

import (
	"context"
	"log"

	"github.com/Admin/app/controllers"
	_ "github.com/Admin/app/docs"
	"github.com/Admin/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Name     string
	Email    string
	Password string
}

type Drugs struct {
	Drug []Drug
}
type Drug struct {
	Name string
}

type RegisterStores struct {
	RegisterStore []RegisterStore
}
type RegisterStore struct {
	Name string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Set Drug Data
	drugs := Drugs{
		Drug: []Drug{
			Drug{"Aspirin"},
			Drug{"ยาแก้ไอ"},
			Drug{"Paracetamol"},
		},
	}
	for _, d := range drugs.Drug {
		client.Drug.
			Create().
			SetName(d.Name).
			Save(context.Background())
	}

	// Set RegisterStore Data
	registerstores := RegisterStores{
		RegisterStore: []RegisterStore{
			RegisterStore{"คลังนอก"},
			RegisterStore{"คลังใน"},
		},
	}

	for _, rs := range registerstores.RegisterStore {
		client.RegisterStore.
			Create().
			SetName(rs.Name).
			Save(context.Background())
	}

	// Set Users Data
	users := Users{
		User: []User{
			User{"สมชาย หมั่นเพียร ", "somchai@gmail.com", "12345"},
			User{"วรรณี ใจชื่น", "่wanee@gmail.com", "123456"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetName(u.Name).
			SetEmail(u.Email).
			SetPassword(u.Password).
			Save(context.Background())
	}

	v1 := router.Group("/api/v1")
	controllers.NewRequisitionController(v1, client)
	controllers.NewDrugController(v1, client)
	controllers.NewRegisterStoreController(v1, client)
	controllers.NewUserController(v1, client)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
