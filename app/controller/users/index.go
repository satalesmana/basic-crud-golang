package users

import (
	usermodel "app-basic-crud/app/database/model/user"
	"app-basic-crud/app/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
	AddNew(c *gin.Context)
	FindBy(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type uscase struct {
	userModel usermodel.Handler
}

func NewUserHandler() Handler {
	return &uscase{
		userModel: usermodel.NewUserHandler(),
	}
}

func (m *uscase) GetAll(c *gin.Context) {
	data, err := m.userModel.GetAll()
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

func (m *uscase) AddNew(c *gin.Context) {
	var (
		data usermodel.User
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	lastID, err := m.userModel.Insert(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}
	data.ID = lastID

	c.JSON(response.Format(http.StatusOK, nil, data))

}

func (m *uscase) FindBy(c *gin.Context) {
	id := c.Param("id")

	data, err := m.userModel.GetByid(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))

}

func (m *uscase) Edit(c *gin.Context) {
	var (
		data usermodel.User
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}
	strID := c.Param("id")
	intID, _ := strconv.ParseInt(strID, 10, 64)
	data.ID = intID

	res, err := m.userModel.Update(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, res))
}

func (m *uscase) Delete(c *gin.Context) {
	id := c.Param("id")

	data, err := m.userModel.Delete(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))

}
