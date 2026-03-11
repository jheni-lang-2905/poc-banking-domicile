package domain

import "github.com/google/uuid"

type BankingDomicileEntities struct {
	Pv    uint64 `json."pv,omitempty"`
	Token string `json."token,omitempty"`
	Guid  string `json."Guid,omitempty"`
	ReqId string `json."ReqId,omitempty"`
	Email string `json."email,omitempty"`
}

func (bank *BankingDomicileEntities) BankUuid(user BankingDomicileEntities) *BankingDomicileEntities {
	return &BankingDomicileEntities{
		Pv:    user.Pv,
		Token: user.Token,
		Guid:  uuid.New().String(),
		ReqId: uuid.New().String(),
		Email: user.Email,
	}
}
