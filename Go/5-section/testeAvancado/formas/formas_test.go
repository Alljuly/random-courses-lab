package formas

import (
	"testing"
	"math"

)

func TestArea(t *testing.T){


	t.Run("Retangulo", func (t *testing.T){
		ret := Retangulo{10,20}
		areaEsperada := float64(200)

		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Error("Algo deu errado")
		}

	})


	t.Run("Circulo", func (t *testing.T){
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)

		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Fatal("Algo deu errado")
		}

	})
}