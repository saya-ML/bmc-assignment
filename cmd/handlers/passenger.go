package handlers

import (
	"bmc/bmc-assignment/internal/passenger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IPassengerHandler interface {
	Get(c *gin.Context)
	All(c *gin.Context)
}

type passengerRequest struct {
	Fields []string `form:"fields"`
}

type passengerHandler struct {
	passengerService passenger.IService
}

// Get Passenger godoc
// @Summary      Get a passenger by key
// @Description  Get a passenger by key and filter fields (optional)
// @Param        id path string true "Key of the passenger"
// @Produce      application/json
// @Success      200 {object} database.Passenger{}
// @Failure      400
// @Router       /passenger/{id} [get]
func (p *passengerHandler) Get(c *gin.Context) {
	var pr passengerRequest
	if errBind := c.ShouldBindQuery(&pr); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}
	key := c.Param("key")
	passenger, errGet := p.passengerService.Get(key)
	if errGet != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errGet.Error()})
		return
	}
	if len(pr.Fields) > 0 {
		c.JSON(http.StatusOK, passenger.SelectFields(pr.Fields))
		return
	}
	c.JSON(http.StatusOK, passenger)
}

// All godoc
// @Summary     Get all passengers
// @Description Get all passengers of the database
// @Produce     application/json
// @Success     200 {object} []database.Passenger{}
// @Failure     400
// @Router      /passenger [get]
func (p *passengerHandler) All(c *gin.Context) {
	passenger, errGet := p.passengerService.All()
	if errGet != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errGet.Error()})
		return
	}
	c.JSON(http.StatusOK, passenger)
}

func NewPassenger(ps passenger.IService) IPassengerHandler {
	return &passengerHandler{
		passengerService: ps,
	}
}
