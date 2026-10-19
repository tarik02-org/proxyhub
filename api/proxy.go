package api

type Proxy struct {
	ID              string   `json:"id"`
	Version         string   `json:"version"`
	Port            int32    `json:"port"`
	Started         int64    `json:"started"`
	EgressWhitelist []string `json:"egress_whitelist"`
	// ClientIP is the address the node joined from, as the hub saw it through its ingress.
	ClientIP string `json:"client_ip,omitempty"`
}
