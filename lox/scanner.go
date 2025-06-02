package lox

import (
	"fmt"
)

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

func (s *Scanner) scanToken() {
	s.advance()
	b := s.source[s.current]

	switch b {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for {
				if s.peek() == '\n' && !s.isAtEnd() {
					break
				}
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(b) {
			s.number()
		} else {
			error(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) advance() {
	s.current++
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal fmt.Stringer) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType, text, literal, s.line})
}

func (s *Scanner) addTokenWithString(tokenType TokenType, value string) {
	s.tokens = append(s.tokens, Token{tokenType, value, nil, s.line})
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) string() {
	for {
		if s.peek() == '"' && !s.isAtEnd() {
			break
		} else if s.peek() == '\n' {
			s.line++
			s.advance()
		}

		if s.isAtEnd() {
			error(s.line, "Unterminated string.")
		}

		s.advance()
		value := s.source[s.start+1 : s.current-1]
		s.addTokenWithString(STRING, value)
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()

		if s.peek() == '.' && isDigit(s.peekNext()) {
			s.advance()

			for isDigit(s.peek()) {
				s.advance()
			}
		}
	}

	s.addTokenWithString(NUMBER, s.source[s.start:s.current])
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}
