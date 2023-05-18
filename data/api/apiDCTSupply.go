package api

// DCTSupply represents the structure for dct supply that is returned by api routes
type DCTSupply struct {
	InitialMinted string `json:"initialMinted"`
	Supply        string `json:"supply"`
	Burned        string `json:"burned"`
	Minted        string `json:"minted"`
}
