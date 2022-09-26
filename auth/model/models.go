package model

type User struct {
	ID        string `bson:"_id,omitempty"`
	Firstname string `bson:"firstname"`
	Lastname  string `bson:"lastname"`
	Username  string `bson:"userName"`
	Password  string `bson:"password"`
	Activated bool   `bson:"activated"`
}

type RabbitMessageFormat struct {
	Type          string
	DateTimeStamp string
	Body          any
}

type GrpcServer struct {
}
