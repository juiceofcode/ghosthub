package profile

type Profile struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	SSHKeyPath string `json:"sshKeyPath"`
}

type Profiles map[string]Profile 