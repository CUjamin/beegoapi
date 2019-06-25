package dao
import(
	"github.com/astaxie/beego/logs"
)

type AnythingDo struct {
	Id       string
	Username string
	Password string
}

func insertAnythingDo(anythingDo AnythingDo){
	logs.Info("")
}

func deleteAnythingDo(anythingDo AnythingDo){
	logs.Info("")
}


func selectAnythingDo(AnythingId string) (anything AnythingDo){
	anythingDo := AnythingDo{"chebenid", "chebenname", "chebenpw"}
	logs.Info("")
	return anythingDo
}