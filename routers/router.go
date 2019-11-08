package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/nguyen930411/go-gin-cms/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/nguyen930411/go-gin-cms/middleware/jwt"
	"github.com/nguyen930411/go-gin-cms/pkg/export"
	"github.com/nguyen930411/go-gin-cms/pkg/qrcode"
	"github.com/nguyen930411/go-gin-cms/pkg/upload"
	"github.com/nguyen930411/go-gin-cms/routers/api"
	"github.com/nguyen930411/go-gin-cms/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	r.GET("/login", v1.GetLogin)
	r.GET("/azure/callback", v1.GetCallback)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
