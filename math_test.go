package main

func TestSoma(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Error("Resultado da soma é invalido: Resultado %d. Esperado %d", total, 30)
	}
}