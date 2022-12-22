package server

import (
	"erik-craigo-code-challenge/controllers"
	"erik-craigo-code-challenge/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func SetupRouter(studentController controllers.StudentController, studentExamController controllers.StudentExamController, examController controllers.ExamController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Routes for healthcheck of api server
	healthcheck := router.Group("health")
	{
		health := new(controllers.HealthController)
		healthcheck.GET("/health", health.Status)
	}

	//Routes for swagger
	swagger := router.Group("swagger")
	{
		// programatically set swagger info
		docs.SwaggerInfo.Title = "Golang REST API for Code Challenge"
		docs.SwaggerInfo.Description = "This is a sample backend written in Go."
		docs.SwaggerInfo.Version = "1.0"

		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	studentApi := router.Group("students")
	{
		//  1. A REST API `/students` that lists all users that have received at least one test score
		studentApi.GET("/", studentController.Get)
		////  2. A REST API `/students/{id}` that lists the test results for the specified student, and provides the student's average score across all exams
		studentApi.GET("/:id", studentController.GetById)

	}

	examApi := router.Group("exams")
  {
		////  3. A REST API `/exams` that lists all the exams that have been recorded
		examApi.GET("/", examController.List)
		////  4. A REST API `/exams/{number}` that lists all the results for the specified exam, and provides the average score across all students
    examApi.GET("/:number", examController.GetByNumber)
  }

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	return router

}
