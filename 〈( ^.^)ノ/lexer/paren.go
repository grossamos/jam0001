package lexer

import "github.com/grossamos/jam0001/shared"

const cc_lparen = '('
const cc_rparen = ')'
const cc_openBock = '{'
const cc_closeBock = '}'

func (l *Lexer) make_paren_token() (shared.Token, error) {
	if l.current_char == cc_lparen {
		return shared.Token{Type: shared.TTlparen, Value: string(cc_lparen), Pos: l.pos}, nil
	} else if l.current_char == cc_rparen {
		return shared.Token{Type: shared.TTrparen, Value: string(cc_rparen), Pos: l.pos}, nil
	} else if l.current_char == cc_openBock {
		return shared.Token{Type: shared.TTopenBlock, Value: string(cc_openBock), Pos: l.pos}, nil
	} else if l.current_char == cc_closeBock {
		return shared.Token{Type: shared.TTcloseBlock, Value: string(cc_closeBock), Pos: l.pos}, nil
	}
	return shared.Token{}, &IllegalCharError{string(l.current_char), l.pos}
}
