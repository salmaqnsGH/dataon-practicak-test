package company

import (
	// executivecommittee "dataon/executiveCommittee"
	subdivision "dataon/subdivision"
)

type CompanyOutput struct {
	ID                 int
	Name               string
	Executivecommittee []ECOutput
}

type ECOutput struct {
	ID       int
	Name     string
	Division []Division
}

type Division struct {
	ID          int
	Name        string
	SubDivision []subdivision.SubDivision
}
