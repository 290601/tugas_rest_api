package gamescontroller

import (
	"net/http"
	"tugas-rest-api/models"
	"tugas-rest-api/models/games"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	var games []games.Games

	if err := models.DB.Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"games": games})
}

func Show(c *gin.Context) {
	var game games.Games

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game": game})
}

func Create(c *gin.Context) {
	var input games.Games

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game": input})
}

func Update(c *gin.Context) {
	var game games.Games

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game tidak ditemukan!"})
		return
	}

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// models.DB.Save(&game)

	c.JSON(http.StatusOK, gin.H{"game": game})
}

func Delete(c *gin.Context) {
	var game games.Games

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game tidak ditemukan!"})
		return
	}

	if err := models.DB.Delete(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// models.DB.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"message": "Game berhasil dihapus!"})
}
