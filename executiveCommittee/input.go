package executiveCommittee

type CompanyIDInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateExecutiveCommitteeInput struct {
	Name string `json:"name" binding:"required"`
}
