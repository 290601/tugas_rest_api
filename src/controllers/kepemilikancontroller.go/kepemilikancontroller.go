package kepemilikancontrollergo

import (
	"net/http"
	"tugas-rest-api/models"
	"tugas-rest-api/models/kepemilikan"

	// "tugas-rest-api/models/users"
	// "tugas-rest-api/models/games"

	"time"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	// Join data dari tabel kepemilikan, users, dan games
	var kepemilikan []kepemilikan.Kepemilikan
	if err := models.DB.Preload("User").Preload("Game").Find(&kepemilikan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kepemilikan": kepemilikan})
}

func Show(c *gin.Context) {
	// Join data dari tabel kepemilikan, users, dan games
	var kepemilikan kepemilikan.Kepemilikan

	id := c.Param("id")

	if err := models.DB.Preload("User").Preload("Game").Where("id = ?", id).First(&kepemilikan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kepemilikan tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kepemilikan": kepemilikan})

}

func Create(c *gin.Context) {
	var input kepemilikan.Kepemilikan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.WaktuPembelian = time.Now().Format("2006-01-02 15:04:05")

	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kepemilikan": input})
}

func Delete(c *gin.Context) {
	var kepemilikan kepemilikan.Kepemilikan

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&kepemilikan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kepemilikan tidak ditemukan!"})
		return
	}

	if err := models.DB.Delete(&kepemilikan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kepemilikan": kepemilikan})
}
