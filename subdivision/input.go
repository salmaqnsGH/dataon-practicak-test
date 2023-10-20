package suubdivision

type DivisionIDInput struct {
	ID int `uri:"id" binding:"required"`
}

type SubDivisionIDInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateSubDivisionInput struct {
	Name string `json:"name" binding:"required"`
}
