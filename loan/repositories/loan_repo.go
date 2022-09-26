package repositories

import (
	"context"
	"loan/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanRepo struct {
	db *mongo.Database
}

func NewLoanRepo(db *mongo.Database) *LoanRepo {
	return &LoanRepo{
		db: db,
	}
}

func (loanRepo *LoanRepo) CreateLoan(ctx context.Context, loan *models.LoanApplicationReqeust) error {
	log.Println("creating loan ", loan)

	res, err := loanRepo.db.Collection("loans").InsertOne(ctx, loan)
	if err != nil {
		log.Println("err creating loan", err)
		return err
	}
	log.Println("created loan", res)
	return nil
}

func (loanRepo *LoanRepo) GetLoanStatus(ctx context.Context, username string) (models.LoanStatus, error) {

	var doc models.LoanApplicationReqeust
	// var loanStatus models.LoanStatus
	filter := bson.D{{Key: "userName", Value: username}}

	err := loanRepo.db.Collection("loans").FindOne(ctx, filter).Decode(&doc)

	// log.Println("loanStatus result", doc)

	loanStatus := doc.Status

	if err != nil {
		log.Println("err finding doc", err)
	}

	return loanStatus, nil
}

func (loanRepo *LoanRepo) GetAllLoans(ctx context.Context) []models.LoanApplicationReqeust {

	loans := []models.LoanApplicationReqeust{}
	cur, err := loanRepo.db.Collection("loans").Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println("err finding all loans")
	}
	loan := models.LoanApplicationReqeust{}
	for cur.Next(context.TODO()) {
		err := cur.Decode(&loan)
		if err != nil {
			log.Println("err decoding loan")
		}
		loans = append(loans, loan)
	}
	return loans
}

func (LoanRepo *LoanRepo) ApproveLoan(ctx context.Context, userName string) bool {

	log.Println("loanrepo", userName)
	filter := bson.D{{Key: "userName", Value: userName}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status.status", Value: "approved"},
		{Key: "status.reason", Value: "none_approved"}}}}

	result, err := LoanRepo.db.Collection("loans").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println("err updating", err)
	}

	log.Println("update res", result.ModifiedCount)

	return true
}

func (LoanRepo *LoanRepo) RejectLoan(ctx context.Context, userName string) bool {
	filter := bson.D{
		{
			Key:   "userName",
			Value: userName,
		},
	}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{{
				Key:   "status.status",
				Value: "rejected",
			},
				{
					Key:   "status.reason",
					Value: "ineligible",
				}},
		},
	}

	result, err := LoanRepo.db.Collection("loans").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("err saving reject loan")
	}
	log.Println(result)

	return true
}
