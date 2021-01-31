package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wells7778/buDai/model"
)

func Index(c *gin.Context) {
	roles := model.All()
	c.JSON(http.StatusOK, roles)
}

func Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	role := model.Find(uint(id))

	c.JSON(http.StatusOK, role)
}

func Create(c *gin.Context) {
	var r model.Role
	model.Create(r)
	c.Status(http.StatusOK)
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	role, err := model.Find(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	role.Name = c.Param("name")
	c.Status(http.StatusNoContent)
}

func Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = model.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
