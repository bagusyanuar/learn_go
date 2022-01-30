package routes

import (
	"net/http"
	authentication "web_go/src/auth"
	"web_go/src/controller"

	"github.com/gin-gonic/gin"
)

func Routes() {
	http.HandleFunc("/", controller.MainPage)
	http.HandleFunc("/about", controller.AboutPage)
	http.HandleFunc("/api", controller.SampleApi)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("assets"))))
}

func InitRoutes() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		// apiV1.GET("/test", controller.TestSlug)
		// apiV1.GET("/profile", authentication.IsAuth, func(c *gin.Context) {
		// 	user := c.MustGet("user")
		// 	c.JSON(http.StatusOK, map[string]interface{}{
		// 		"user": user,
		// 	})
		// })

		auth := apiV1.Group("/auth")
		{
			adminAuth := auth.Group("/admin")
			{
				adminAuth.POST("/sign-up", authentication.AdminSignUp)
				adminAuth.POST("/sign-in", authentication.AdminSignIn)
			}
			memberAuth := auth.Group("/member")
			{
				memberAuth.POST("/sign-up", authentication.MemberSignUp)
				memberAuth.POST("/sign-in", authentication.MemberSignIn)
			}

		}

		member := apiV1.Group("/member")
		{
			member.GET("/profile", authentication.IsAuth, authentication.IsMember, controller.GetMemberProfile)
		}

		groupUser := apiV1.Group("/users")
		{
			groupUser.GET("/", controller.GetUsers)
			groupUser.POST("/", controller.CreateUserMember)
			groupUser.POST("/singular", controller.CreateSingularUser)
		}

		groupMember := apiV1.Group("/members")
		{
			groupMember.GET("/", controller.GetMembers)
			groupMember.GET("/join", controller.GetMemberUsers)
		}
		teachers := apiV1.Group("/teachers")
		{
			teachers.GET("/", controller.GetTeachers)
		}

		categories := apiV1.Group("/categories")
		{
			categories.GET("/", controller.GetCategories)
			categories.POST("/", controller.CategoryCreate)
			categories.GET("/:id", controller.GetDetailCategories)
			categories.PATCH("/:id", controller.PatchCategories)
			categories.DELETE("/:id", controller.DeleteCategories)
		}
	}
	return r
}
