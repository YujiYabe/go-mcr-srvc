package model

//Todo is TodoModel
type Todo struct {
	ID        int    `json:"id"`        //TaskのID
	Task      string `json:"task"`      //Task自体
	LimitDate string `json:"limitDate"` //Taskの完了期限
	Status    bool   `json:"status"`    //Taskの状態(0=未済,1=済)
}
