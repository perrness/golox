package lox

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) scanTokens() []Token {
	var tokens []Token

	for {
		if s.isAtEnd() {
			break
		}

		s.start = s.current
		s.scanToken()
	}

	tokens = append(s.tokens, Token{EOF, "", nil, s.line})

	return tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
