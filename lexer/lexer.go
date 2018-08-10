// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/teslamotors/jsonql/token"
)

const (
	NoState    = -1
	NumStates  = 78
	NumSymbols = 66
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '_'
1: '.'
2: '.'
3: '"'
4: '\'
5: '"'
6: '"'
7: '''
8: '\'
9: '''
10: '''
11: 'n'
12: 'u'
13: 'l'
14: 'l'
15: 't'
16: 'r'
17: 'u'
18: 'e'
19: 'f'
20: 'a'
21: 'l'
22: 's'
23: 'e'
24: '.'
25: '['
26: ']'
27: '_'
28: 'e'
29: 'E'
30: '+'
31: '-'
32: '0'
33: '0'
34: 'x'
35: 'X'
36: '\'
37: 'x'
38: '\'
39: 'u'
40: 'b'
41: 'f'
42: 'n'
43: 'r'
44: 't'
45: 'v'
46: '\'
47: '\'
48: ' '
49: '\t'
50: '\f'
51: '\v'
52: \u00a0
53: \u202f
54: \u205f
55: \u3000
56: \ufeff
57: 'A'-'Z'
58: 'a'-'z'
59: '0'-'9'
60: '1'-'9'
61: '0'-'7'
62: 'a'-'f'
63: 'A'-'F'
64: \u2000-\u200a
65: .
*/
