package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {

	ParseCSV()
	brazilTotal := 45
	argentinaTotal := 15

	resBra, err1 := GetTotalTickets("Brazil")
	resArg, err2 := GetTotalTickets("Argentina")

	assert.NoError(t, err1)
	assert.Equal(t, brazilTotal, resBra)

	assert.NoError(t, err2)
	assert.Equal(t, argentinaTotal, resArg)

}

func TestCountByPeriod(t *testing.T) {

	ParseCSV()
	madrugadaP := 304
	tardeP := 289

	resMadrugada, err1 := GetCountByPeriod("madrugada")
	resTarde, err2 := GetCountByPeriod("tarde")

	assert.NoError(t, err1)
	assert.Equal(t, madrugadaP, resMadrugada)

	assert.NoError(t, err2)
	assert.Equal(t, tardeP, resTarde)
}
