package main

import (
	"log"
	"net/http"

    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
	"student-career-platform/internal/handlers"
	"student-career-platform/internal/middleware"
)

func main() {

	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}
	err = database.DB.AutoMigrate(&models.User{},&models.Profile{},&models.Project{},&models.Skill{},&models.Application{},)

    if err != nil {
	log.Fatal(err)
    }

	router := gin.Default()
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    AllowCredentials: true,
}))

	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Student Career Platform API is running",
		})
	})
	router.GET(
	"/api/v1/protected",
	middleware.AuthRequired(),
	func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		c.JSON(http.StatusOK, gin.H{
			"message": "you are authenticated",
			"user_id": userID,
		})
	},
    )
	router.POST("/api/v1/auth/register", handlers.Register)
	router.POST("/api/v1/auth/login", handlers.Login)
	router.GET(
	    "/api/v1/profile",
	    middleware.AuthRequired(),
	    handlers.GetProfile,
    )

    router.PUT(
	   "/api/v1/profile",
	   middleware.AuthRequired(),
	   handlers.UpdateProfile,
    )
	router.POST(
	"/api/v1/projects",
	middleware.AuthRequired(),
	handlers.CreateProject,
)

router.GET(
	"/api/v1/projects",
	middleware.AuthRequired(),
	handlers.GetProjects,
)

router.GET(
	"/api/v1/projects/:id",
	middleware.AuthRequired(),
	handlers.GetProject,
)

router.PUT(
	"/api/v1/projects/:id",
	middleware.AuthRequired(),
	handlers.UpdateProject,
)

router.DELETE(
	"/api/v1/projects/:id",
	middleware.AuthRequired(),
	handlers.DeleteProject,
)
router.POST(
	"/api/v1/skills",
	middleware.AuthRequired(),
	handlers.CreateSkill,
)

router.GET(
	"/api/v1/skills",
	middleware.AuthRequired(),
	handlers.GetSkills,
)

router.PUT(
	"/api/v1/skills/:id",
	middleware.AuthRequired(),
	handlers.UpdateSkill,
)

router.DELETE(
	"/api/v1/skills/:id",
	middleware.AuthRequired(),
	handlers.DeleteSkill,
)
router.POST(
	"/api/v1/applications",
	middleware.AuthRequired(),
	handlers.CreateApplication,
)

router.GET(
	"/api/v1/applications",
	middleware.AuthRequired(),
	handlers.GetApplications,
)

router.GET(
	"/api/v1/applications/:id",
	middleware.AuthRequired(),
	handlers.GetApplication,
)

router.PUT(
	"/api/v1/applications/:id",
	middleware.AuthRequired(),
	handlers.UpdateApplication,
)

router.DELETE(
	"/api/v1/applications/:id",
	middleware.AuthRequired(),
	handlers.DeleteApplication,
)
router.GET("/api/v1/dashboard", middleware.AuthRequired(), handlers.GetDashboard)

	router.Run(":8080")
}