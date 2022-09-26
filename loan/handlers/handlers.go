package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"loan/clients"
	"loan/db"
	"loan/models"
	"loan/repositories"
)

type Handler struct {
	loanRepo *repositories.LoanRepo
}

func NewHandler(loanRepo *repositories.LoanRepo) *Handler {
	return &Handler{
		loanRepo: loanRepo,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("health handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("loan service OK"))
}

func (h *Handler) Apply(w http.ResponseWriter, r *http.Request) {

	loanRequest := &models.LoanApplicationReqeust{}

	log.Println("loanRequest", loanRequest)

	log.Println("r.Body", r.Body)

	err := json.NewDecoder(r.Body).Decode(&loanRequest)
	if err != nil {
		log.Println("err decoding loanRequest", err)
	}
	log.Println("recvd loan request", loanRequest)

	loanStatus := models.LoanStatus{
		StatusName: "applied",
		Reason:     "none",
	}
	log.Println("loanStatus", loanStatus)

	loanRequest.Status = loanStatus
	loanRequest.LoanAppliedDate = fmt.Sprint(time.Now().Unix())

	loanCreateErr := h.loanRepo.CreateLoan(db.Ctx, loanRequest)

	resp := make(map[string]string)

	if loanCreateErr != nil {
		resp["message"] = "Loan request not accepted"
	} else {
		resp["message"] = "Loan request accepted"
	}

	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {

	var req map[string]string

	log.Println(r.Body)

	errs := json.NewDecoder(r.Body).Decode(&req)
	if errs != nil {
		log.Println("err decoding username")
	}
	log.Println("checking loan status for ", req)
	status, err := h.loanRepo.GetLoanStatus(db.Ctx, req["username"])
	if err != nil {
		log.Println("err", err)
	}
	log.Println("hstatus #v", status.StatusName)

	resp := status.StatusName

	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (h *Handler) Loans(w http.ResponseWriter, r *http.Request) {
	loans := h.loanRepo.GetAllLoans(context.TODO())
	log.Println("Loans", loans)

	jsonResp, err := json.Marshal(loans)
	if err != nil {
		log.Println("err marshalling loans list")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (h *Handler) ApproveLoan(w http.ResponseWriter, r *http.Request) {
	log.Println("approve loan")
	var request map[string]string
	var userName string

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("err while decoding request", err)
	}
	userName = request["userName"]
	log.Println("userName", userName)
	approved := h.loanRepo.ApproveLoan(context.TODO(), userName)

	if approved {
		messageSent := clients.PublishUserCommunicationMessage(userName, "Your loan is approved")
		log.Println("messageSent", messageSent)
	}
}

func (h *Handler) RejectLoan(w http.ResponseWriter, r *http.Request) {
	log.Println("reject loan")
	var request map[string]string

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("reject loan : error decoding r.Body", err)
	}

	userName := request["userName"]

	log.Println("rejecting laon for ", userName)
	h.loanRepo.RejectLoan(context.TODO(), userName)
}
