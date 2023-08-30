package handlers

import (
	"bmc/bmc-assignment/internal/histogram"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IHistogramHandler interface {
	HistPlot(c *gin.Context)
}

type histogramHandler struct {
	histogramService histogram.IService
}

// HistPlot     godoc
// @Summary     Get histogram
// @Description Generate a histogram with the fares
// @Produce     image/svg+xml
// @Success     200
// @Failure     400 {object} error "error message"
// @Router      /histogram [get]
func (h *histogramHandler) HistPlot(c *gin.Context) {
	histogram, errGet := h.histogramService.HistPlot()
	if errGet != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errGet.Error()})
		return
	}
	c.Data(http.StatusOK, "image/svg+xml", histogram)
}

func NewHistogram(hs histogram.IService) IHistogramHandler {
	return &histogramHandler{
		histogramService: hs,
	}
}
