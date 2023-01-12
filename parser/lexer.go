package parser

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var digits = "0123456789"

const (
	itemNumber TokenType = iota
	itemOperator
	itemBracket
	itemIdentifier
	itemBool
	itemString
	itemEOF = -1
)

type TokenType int

type stateFn func(*Lexer) stateFn

type Token struct {
	tokenType TokenType
	val       string
	pos       int
}

type Lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens []Token
}

func (lexer *Lexer) word() string {
	return lexer.input[lexer.start:lexer.pos]
}

func (lexer *Lexer) emitValue(t TokenType, value string) {
	lexer.tokens = append(lexer.tokens, Token{
		tokenType: t,
		val:       value,
		pos:       lexer.start,
	})
	lexer.start = lexer.pos
}

func (lexer *Lexer) emit(tokenType TokenType) {
	lexer.emitValue(tokenType, lexer.word())
}

func (lexer *Lexer) backup() {
	lexer.pos -= lexer.width
}

func (lexer *Lexer) next() rune {
	if lexer.pos >= len(lexer.input) {
		lexer.width = 0
		return itemEOF
	}
	r, w := utf8.DecodeRuneInString(lexer.input[lexer.pos:])
	lexer.width = w
	lexer.pos += lexer.width
	return r
}

func (lexer *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, lexer.next()) {
		return true
	}
	lexer.backup()
	return false
}

func (lexer *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, lexer.next()) {
	}
	lexer.backup()
}

func (lexer *Lexer) peek() rune {
	r := lexer.next()
	lexer.backup()
	return r
}

func isAlphaNumeric(r rune) bool {
	return isAlphabetic(r) || isDigit(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isAlphabetic(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isNonToken(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func (lexer *Lexer) skip() {
	lexer.start = lexer.pos
}

func (lexer *Lexer) scanString(r rune) {
	ch := lexer.next()
	for ch != r {
		if ch == '\n' || ch == itemEOF {
			return
		}
		ch = lexer.next()
	}
	return
}

func (lexer *Lexer) scanNumber() bool {
	lexer.acceptRun(digits)
	if lexer.accept(".") {
		lexer.acceptRun(digits)
	}
	if lexer.accept("eE") {
		lexer.accept("+-")
		lexer.acceptRun(digits)
	}
	if isAlphaNumeric(lexer.peek()) {
		lexer.next()
		return false
	}
	return true
}

func lexNumber(lexer *Lexer) stateFn {
	if !lexer.scanNumber() {
		return nil
	}
	lexer.emit(itemNumber)
	return scan
}

func lexDot(lexer *Lexer) stateFn {
	lexer.next()
	if lexer.accept(digits) {
		lexer.backup()
		return lexNumber
	}
	return scan
}

func lexIdentifier(lexer *Lexer) stateFn {
loop:
	for {
		switch r := lexer.next(); {
		case isAlphaNumeric(r):
			// absorb
		default:
			lexer.backup()
			switch lexer.word() {
			case "or", "and":
				lexer.emit(itemOperator)
			case "true", "false":
				lexer.emit(itemBool)
			default:
				lexer.emit(itemIdentifier)
			}
			break loop
		}
	}
	return scan
}

func scan(lexer *Lexer) stateFn {
	switch r := lexer.next(); {
	case isNonToken(r):
		lexer.skip()
		return scan
	case r == itemEOF:
		lexer.emit(itemEOF)
		return nil
	case '0' <= r && r <= '9':
		lexer.backup()
		return lexNumber
	case r == '.':
		lexer.backup()
		return lexDot
	case strings.ContainsRune("([{", r):
		lexer.emit(itemBracket)
	case strings.ContainsRune(")]}", r):
		lexer.emit(itemBracket)
	case strings.ContainsRune("+-/%=><&|", r):
		lexer.emit(itemOperator)
	case isAlphaNumeric(r):
		lexer.backup()
		return lexIdentifier
	case r == '\'' || r == '"':
		lexer.scanString(r)
		str := lexer.word()
		lexer.emitValue(itemString, str[1:len(str)-1])
	default:
		return nil
	}
	return scan
}

func lex(input string) []Token {
	lexer := &Lexer{
		input:  input,
		tokens: make([]Token, 0),
	}

	for state := scan; state != nil; {
		state = state(lexer)
	}
	return lexer.tokens
}
