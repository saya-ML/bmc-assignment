package histogram

import (
	"bmc/bmc-assignment/internal/passenger"
	"bytes"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type IService interface {
	HistPlot() ([]byte, error)
}
type histogramService struct {
	passengerService passenger.IService
}

func New(ps passenger.IService) IService {
	return &histogramService{
		passengerService: ps,
	}
}

func (h *histogramService) HistPlot() ([]byte, error) {
	passengers, errPassengerService := h.passengerService.All()
	if errPassengerService != nil {
		return nil, errPassengerService
	}
	values := make(plotter.Values, len(passengers))
	var sum float64
	for i := range passengers {
		values[i] = passengers[i].Fare
		sum += passengers[i].Fare
	}
	p := plot.New()
	p.Title.Text = "histogram plot"
	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	//if errSave := p.Save(10*vg.Inch, 10*vg.Inch, "test.png"); errSave != nil {
	//	return nil, errSave
	//}
	//return nil, nil

	w, err := p.WriterTo(3*vg.Inch, 3*vg.Inch, "svg")
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	_, errHist := w.WriteTo(&buf)
	if errHist != nil {
		return nil, errHist
	}
	return buf.Bytes(), nil

}
