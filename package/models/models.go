package models

type Analysis struct {
	Score     Score       `json:"score"`
	Evidence  []Evidence  `json:"evidence"`
	Kpi       []KPI       `json:"kpi"`
	Gap       []Gap       `json:"gap"`
	Questions []Questions `json:"questions"`
}

type Score struct {
	Score         int    `json:"value"`
	Label         string `json:"label"`
	Band          string `json:"band"`
	Justification string `json:"justification"`
	Confidence    string `json:"confidence"`
}

type KPI struct {
	Kpi_map        string `json:"kpimap"`
	Supervisor_map string `json:"supervisormap"`
	Layer          string `json:"layer"`
}

type Evidence struct {
	Quote          string `json:"quote"`
	Signal         string `json:"signal"`
	Dimension      string `json:"dimension"`
	Interpretation string `json:"interpretation"`
}

type Gap struct {
	Dimension string `json:"gap"`
	Absent    string `json:"absent"`
}

type Questions struct {
	Questions string `json:"questions"`
	Gap       string `json:"gap"`
}
