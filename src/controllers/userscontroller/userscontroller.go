package userscontroller

import (
	"net/http"
	"tugas-rest-api/controllers/password"
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

func Create(c *gin.Context) {
	var input users.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := users.Users{Nama: input.Nama, Email: input.Email, Password: hashedPassword}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Update(c *gin.Context) {
	var user users.Users

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan!"})
		return
	}

	var input users.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Delete(c *gin.Context) {
	var user users.Users

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func ComparePassword(c *gin.Context) {
	var user users.Users

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan!"})
		return
	}

	var input users.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := password.Verify(user.Password, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password tidak cocok!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password cocok!"})
}
