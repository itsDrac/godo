package tokens

type Tokenizer interface {
	GenerateToken(data any) (string, error)
	ValidateToken(token string) (any, error)
}
