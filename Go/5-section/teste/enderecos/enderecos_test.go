/*
// Teste unitário
package enderecos

import "testing"

func TestTipoEndereco(t *testing.T){
	enderecoTeste := "Avenida"

	tipoEsperado := "Avenida"

	tipoRecebido := TipoEndereco(enderecoTeste)

	if tipoRecebido != tipoEsperado {
		t.Error("O tipo recebido não é igual ao tipo esperado!")
	}


}


*/

// Teste unitário
package enderecos

import "testing"

type cenarioTeste struct {
	recebido string
	esperado string
}


func TestTipoEnderecoCenarios(t *testing.T){
	cenarios := []cenarioTeste{
		{"Rua da Liberdade", "Rua"},
		{"Avenida do Futuro", "Avenida"},
		{"Lagoa da canoa", "Tipo Invalido"},
		{"asdnajcna", "Tipo Invalido"},

	}

	for _, cenario := range cenarios{
		tipoRecebido := TipoEndereco(cenario.recebido)
		
		if tipoRecebido != cenario.esperado {
			t.Error("O tipo recebido não é igual ao tipo esperado!")
		}
	}


}

