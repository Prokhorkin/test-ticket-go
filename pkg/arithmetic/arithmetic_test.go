package arithmetic

import "testing"

func TestAdd(t *testing.T) {

res:=Add(4,4)

if res!=8 {
	t.Errorf("Результат операции %d, должно быть %d",res,8)
}
}

func TestSub(t *testing.T) {
res:=Sub(8,4)

	if res!=4 {
		t.Errorf("Результат операции %d, должно быть %d",res,4)
	}
}
func TestMul(t *testing.T) {
	res:=Mul(3,3)

	if res!=9 {
		t.Errorf("Результат операции %d, должно быть %d",res,9)
	}
}
func TestDiv(t *testing.T) {
	res:=Div(8,4)

	if res!=2 {
		t.Errorf("Результат операции %d, должно быть %d",res,2)
	}
}