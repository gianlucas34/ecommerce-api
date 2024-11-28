package ports

type PasswordHashing interface {
	Generate(plaintext string) (string, error)
	Compare(plaintext string, digest string) (bool, error)
}
