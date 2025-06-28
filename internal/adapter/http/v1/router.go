package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/narianapereira/logistics-go/internal/application/service"
	"net/http"
)

func NewRouter(service *service.ParserService) *gin.Engine {
	router := gin.Default()

	router.POST("/parse", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "arquivo não enviado"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao abrir o arquivo"})
			return
		}
		defer f.Close()

		content := make([]byte, file.Size)
		_, err = f.Read(content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao ler o arquivo"})
			return
		}

		result, err := service.Parse(content)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "erro ao processar arquivo"})
			return
		}

		c.Data(http.StatusOK, "application/json", result)
	})

	return router
}
