package models

type AuthRequest struct {
	JWTToken string
}

type AuthResponse struct {
	Valid bool
}

type LoanStatus struct {
	StatusName string `bson:"status"`
	Reason     string `bson:"reason"`
}

type LoanApplicationReqeust struct {
	ID              string     `bson:"_id,omitempty"`
	Username        string     `bson:"userName,omitempty"`
	FirstName       string     `bson:"firstName"`
	Lastname        string     `bson:"lastName"`
	MonthlySalary   int        `bson:"monthlySalary"`
	LoanAmount      int        `bson:"loanAmount"`
	Status          LoanStatus `bson:"status"`
	LoanAppliedDate string     `bson:"loanAppliedDateTime"`
}

type LoansList struct {
	loans []LoanApplicationReqeust
}

type UserCommunication struct {
	UserEmail string
	Message   string
	DateTime  string
}
