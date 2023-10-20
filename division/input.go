package division

type ExecutiveCommitteeIDInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateDivisionInput struct {
	Name string `json:"name" binding:"required"`
}
