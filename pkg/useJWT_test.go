package pkg

import (
	"testing"
)

func TestJWT(t *testing.T) {
	var tokenJWT, err = GenerateJWT()
	if err != nil {
		t.Fatal(err)
	}
	if tokenJWT == "" {
		t.Fatal("TokenString vazio")
	}

	var validate = ValidateJWT(tokenJWT)
	if validate == false {
		t.Fatal("não foi possível validar")
	}

	t.Log("Resultado esperado: \n", tokenJWT)
}
