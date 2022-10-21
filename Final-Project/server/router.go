package server

import (
	"dalas98/golang-lesson/Final-Project/controller/controllercomment"
	"dalas98/golang-lesson/Final-Project/controller/controllerphoto"
	"dalas98/golang-lesson/Final-Project/controller/controllersocialmedia"
	"dalas98/golang-lesson/Final-Project/controller/controlleruser"
	_ "dalas98/golang-lesson/Final-Project/docs"
	"dalas98/golang-lesson/Final-Project/middleware"
	"dalas98/golang-lesson/Final-Project/repository/repositorycomment"
	"dalas98/golang-lesson/Final-Project/repository/repositoryphoto"
	"dalas98/golang-lesson/Final-Project/repository/repositorysocialmedia"
	"dalas98/golang-lesson/Final-Project/repository/repositoryuser"
	"dalas98/golang-lesson/Final-Project/service/servicecomment"
	"dalas98/golang-lesson/Final-Project/service/servicephoto"
	"dalas98/golang-lesson/Final-Project/service/servicesocialmedia"
	"dalas98/golang-lesson/Final-Project/service/serviceuser"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("", middleware.Authorization, ctrlUser.DeleteByID)

	// route photos
	repoPhoto := repositoryphoto.New(db)
	srvPhoto := servicephoto.New(repoPhoto)
	ctrlPhoto := controllerphoto.New(srvPhoto)

	r.GET("photos", middleware.Authorization, ctrlPhoto.GetPhotos)
	r.POST("photos", middleware.Authorization, ctrlPhoto.Create)
	r.PUT("photos/:photoID", middleware.Authorization, ctrlPhoto.Update)
	r.DELETE("photos/:photoID", middleware.Authorization, ctrlPhoto.Delete)

	// route comment
	repoComment := repositorycomment.New(db)
	srvComment := servicecomment.New(repoComment)
	ctrlComment := controllercomment.New(srvComment)
	r.GET("comments", middleware.Authorization, ctrlComment.Get)
	r.POST("comments", middleware.Authorization, ctrlComment.Create)
	r.PUT("comments/:commentID", middleware.Authorization, ctrlComment.Update)
	r.DELETE("comments/:commentID", middleware.Authorization, ctrlComment.Delete)

	// route social media
	repoSocialmedia := repositorysocialmedia.New(db)
	srvSocialmedia := servicesocialmedia.New(repoSocialmedia, repoPhoto)
	ctrlSocialmedia := controllersocialmedia.New(srvSocialmedia)
	routerSocialmedia := r.Group("/socialmedias")
	routerSocialmedia.POST("", middleware.Authorization, ctrlSocialmedia.Create)
	routerSocialmedia.GET("", middleware.Authorization, ctrlSocialmedia.GetList)
	routerSocialmedia.PUT("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.UpdateByID)
	routerSocialmedia.DELETE("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.DeleteByID)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
