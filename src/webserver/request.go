package webserver

// Request Type for a structure of an alertmanager POST request
type Request struct {
	Alerts []Alert `json:"alerts"`
}

// Alert Type for a structure of an alerts array
type Alert struct {
	Labels struct {
		ClusterDiscriminator string `json:"cluster_discriminator"`
	} `json:"labels"`
}
