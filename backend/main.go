package main

import (
	"context"
	"log"
	//"fmt"

	"github.com/beam19857/app/controllers"
	_ "github.com/beam19857/app/docs"
	"github.com/beam19857/app/ent"
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
	DoctorName string
	DoctorEmail string
}
type Departments struct {
	Department []Department
}
type Department struct {
	DepartmentName string
}
type Expertises struct {
	Expertise []Expertise
}
type Expertise struct {
	ExpertiseName string
}
type Positions struct {
	Position []Position
}
type Position struct {
	PositionName string
}
// @title SUT SA Example API
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

	client, err := ent.Open("sqlite3", "file:doctor.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewExpertiseController(v1, client)
	controllers.NewPositionController(v1, client)

	// Set Department Data
	departments := Departments{
		Department: []Department{
			Department{"Antenatal"},
			Department{"Operating Room"},
			Department{"Deceased Receive"},
			Department{"Treatment"},
		},
	}
	for _, d := range departments.Department {
		client.Department.
			Create().
			SetDepartmentName(d.DepartmentName).
			Save(context.Background())
	}
		// Set Position Data
		positions := Positions{
			Position: []Position{
				Position{"Nurse"},
				Position{"Doctor"},
				Position{"Parcel worker"},
				Position{"Assistant"},
				Position{"Medical records officer"},
			},
		}
		for _, p := range positions.Position {
			client.Position.
				Create().
				SetPositionName(p.PositionName).
				Save(context.Background())
		}
		// Set Expertise Data
		expertises := Expertises{
			Expertise: []Expertise{
				Expertise{"Internal Medicine"},
				Expertise{"Obstetrics"},
				Expertise{"Bone surgery"},
			},
		}
		for _, e := range expertises.Expertise {
			client.Expertise.
				Create().
				SetExpertiseName(e.ExpertiseName).
				Save(context.Background())
		}


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

	
}
