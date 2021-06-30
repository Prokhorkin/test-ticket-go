package management

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"test-project/pkg/arithmetic"
)
type parametrsJson struct {
	Success bool
	ErrCode string
	Value int
}

var Port string

func init() {
	Port= "5505"
}

func Handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"метод": r.Method,
			"путь":  r.URL.Path,
		}).Info("Получен управляющий запрос")

		p1:=r.URL.Query().Get("a")
		p1Int,_:=strconv.Atoi(p1)

		p2:=r.URL.Query().Get("b")
		p2Int,_:=strconv.Atoi(p2)

	path := r.URL.Path
	switch path {
	case "/ping":
		log.Info("OK")
	case "/add":
		z:= arithmetic.Add(p1Int,p2Int)
		fmt.Println(SetFormatResponse(z))
		fmt.Fprintf(w,SetFormatResponse(z))
	case "/sub":
		z:= arithmetic.Sub(p1Int,p2Int)
		fmt.Println(SetFormatResponse(z))
		fmt.Fprintf(w,SetFormatResponse(z))
	case "/mul":
		z:= arithmetic.Mul(p1Int,p2Int)
		fmt.Println(SetFormatResponse(z))
		fmt.Fprintf(w,SetFormatResponse(z))
	case "/div":
		z:= arithmetic.Div(p1Int,p2Int)
		fmt.Println(SetFormatResponse(z))
		fmt.Fprintf(w,SetFormatResponse(z))
	default:
		log.Info("метод не найден")
	}
	}
}

func SetFormatResponse(z int) string {
	//сформировать ответ по структуре
	param:= parametrsJson{
		Success: true,
		ErrCode: "",
		Value: z,
	}

	//сформировать JSON формат
	b,err:=json.Marshal(param)
	if err!=nil {
		log.Error(err)
		return ""
	}
	return string(b)
}