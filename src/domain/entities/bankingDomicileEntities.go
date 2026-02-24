package domain

import "github.com/google/uuid"

type BankingDomicileEntities struct {
	Pv    uint64 `json."pv,omitempty"`
	token string `json."token,omitempty"`
	Guid  string `json."Guid,omitempty"`
	ReqId string `json."ReqId,omitempty"`
	email string `json."email,omitempty"`
}

func uuidRequest(user BankingDomicileEntities) BankingDomicileEntities {
	return BankingDomicileEntities{
		Pv:    user.Pv,
		token: user.token,
		Guid:  uuid.New().String(),
		ReqId: uuid.New().String(),
		email: user.email,
	}
}
