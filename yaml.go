package yamal

import (
	"io"
	"strings"
)

type TokenType int

const (
	bBreak TokenType = iota
)

type Token struct{
	Type TokenType
	Value string
}

func (tok Token) String() string {
	return string(tok.Value)
}

func LexicalScanner(rrd io.RuneReader, tokens *[]Token) error {
	const (
		inital = iota
		bBreakState
		cEscape
	)

	var (
		builder strings.Builder

		state = inital
	)

	readNextCharacter:
	for {
		char, _, err := rrd.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		nextState:

		switch state {
		case inital:
			switch {
			case char == '-': // c-sequence-entry
			case char == '?': // c-mapping-key
			case char == ':': // c-mapping-value
			case char == ',': // c-collect-entry
			case char == '[': // c-sequence-start
			case char == ']': // c-sequence-end
			case char == '{': // c-mapping-start
			case char == '}': // c-mapping-end
			case char == '#': // c-comment
			case char == '&': // c-anchor
			case char == '*': // c-alias
			case char == '!': // c-tag
			case char == '|': // c-literal
			case char == '>': // c-folded
			case char == '\'': // c-single-quote'
				state = cEscape
				builder.WriteRune(char)

			case char == '"': // c-double-quote
			case char == '%': // c-directive
			case char == '@'|| char == '`': // c-reserved

			case char == '\n':
				builder.WriteRune(char)

				*tokens  = append(*tokens, Token{
					Type: bBreak,
					Value: builder.String(),
				})
				builder.Reset()

			case char == '\r' :
				state = bBreakState
				builder.WriteRune(char)
				continue readNextCharacter
			case char == ' ':
			case char == '\t':

			case char >= '0' && char <= '9': // ns-dec-digit

			}

		case bBreakState:
			state = inital

			if char != '\n' {
				*tokens  = append(*tokens, Token{
					Type: bBreak,
					Value: builder.String(),
				})
				builder.Reset()

				continue nextState
			}

			builder.WriteRune(char)

			*tokens  = append(*tokens, Token{
				Type: bBreak,
				Value: builder.String(),
			})
			builder.Reset()

			continue readNextCharacter

		case cEscape:

			switch char {
			case '0', 'a', 'b', 't', '\x09' ,'n',
				'v', 'r', 'e', '\x20', '"', '/', '\\',
				'N', '_', 'L', 'P':

			}

		default:
		}


	}

	return nil
}