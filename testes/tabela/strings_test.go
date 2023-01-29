package tabela

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> escontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Go é legal", "Go", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Carlos", "r", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
