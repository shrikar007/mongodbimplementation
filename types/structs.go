package types

type Request struct {
    Timestamp 			string			`json:"timestamp"`
	Name 				string			`json:"name"`
	Service 			string		    `json:"service"`
	Region 				string			`json:"region"`
	Details 			string			`json:"details"`
	Rperiod 			string			`json:"r_period"`
	ProjectName 		string			`json:"project_name"`
	Prioriy 			string			`json:"priority"`
	ApprovedBy 			string			`json:"approvedby"`
}
