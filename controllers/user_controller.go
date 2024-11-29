package controllers

import (
	"gin_aws/models"
	"gin_aws/services"
	"gin_aws/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Error al obtener los usuarios")
		return
	}
	utils.RespondJSON(c, http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Datos inválidos")
		return
	}
	if err := services.CreateUser(&user); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Error al crear el usuario")
		return
	}
	utils.RespondJSON(c, http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUser(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Usuario no encontrado")
		return
	}
	utils.RespondJSON(c, http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteUser(uint(id)); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Error al eliminar el usuario")
		return
	}
	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
