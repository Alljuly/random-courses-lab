package enderecos

import "strings"

func TipoEndereco(endereco string) string{
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
		return strings.Title(logadouro) 
	}

	return "Tipo Invalido"

}