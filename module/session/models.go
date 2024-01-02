package sessionsvc

type CredentialModel struct {
	Scheme   string
	Host     string
	Username string
	Password string
	Database string
	Options  string
}
