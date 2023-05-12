package userscontroller

import (
	"net/http"
	"tugas-rest-api/models"
	"tugas-rest-api/models/users"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	var users []users.Users

	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context) {
	var user users.Users

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
