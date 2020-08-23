package money_overseer

type MoneyOverseer struct {
}

func NewMoneyOverseer() *MoneyOverseer {
	return &MoneyOverseer{}
}

func (mo *MoneyOverseer) SetToken(token string) {
	mo.token = token
}

func (mo *MoneyOverseer) GetToken() string {
	return mo.token
}
