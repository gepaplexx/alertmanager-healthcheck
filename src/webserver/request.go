package webserver

// Type for a structure of an alertmanager POST request
type Request struct {
	Alerts   []Alert `json:"alerts"`
}

// Type for a structure of an alerts array
type Alert struct {
	Labels struct {
        	GepardecCluster        string `json:"gepardec_cluster"`
	} `json:"labels"`
}