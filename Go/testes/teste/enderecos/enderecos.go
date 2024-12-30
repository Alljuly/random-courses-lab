package enderecos

import (
	"golang.org/x/text/cases"
    "golang.org/x/text/language"
	"strings"
)

func TipoEndereco(endereco string)  string{
	tipos := []string{"rua", "avenida", "estrada", "rodovia"}

	endereco = strings.ToLower(endereco)

	logadouro := strings.Split(endereco, " ")[0]


	logadouroDeEnderecoValido := false

	for _, tipo := range tipos{
		if logadouro == tipo{
			logadouroDeEnderecoValido = true
		}
	}
	if logadouroDeEnderecoValido {
		return cases.Title(language.Und).String(logadouro) 
	}

	return "Tipo Invalido"
}