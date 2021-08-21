package parser

import (
	"github.com/grossamos/jam0001/shared"
)

func GenerateAst(toks *[]shared.Token) []shared.Node {
	out := make([]shared.Node, 0, 256)

	for len(*toks) != 0 {
		tok := (*toks)[0]

		switch tok.Type {
		case shared.TTlparen, shared.TTopenBlock:
			*toks = (*toks)[1:]

			out[len(out)-1].Children = append(
				out[len(out)-1].Children,
				shared.Node{
					IsExpression: true,
					Children:     GenerateAst(toks)})

		case shared.TTrparen, shared.TTcloseBlock:
			return out

		case shared.TTinstruction: // (instruction (arguments))
			*toks = (*toks)[1:]

			out = append(out,
				shared.Node{
					IsExpression: true,
					Children: []shared.Node{{Val: tok}}})

			continue

		case shared.TTstring, shared.TTref:
			if len(*toks) > 1 &&
				(*toks)[1].Type == shared.TTlparen {

				*toks = (*toks)[1:]

				out = append(out,
					shared.Node{
						IsExpression: true,
						Children: []shared.Node{{Val: tok}}})

				continue
			} else {
				out = append(out, shared.Node{Val: tok})
			}

		case shared.TTwcomment:
			if len(out) == 0 {
				continue
			}

			index := len(out) - 1

			out[index] = shared.Node{
				IsExpression: true,
				Children: []shared.Node{
					{Val: tok},
					out[index]}}

		case shared.TTwhile:
			if len(*toks) <= 1 {
				break
			}

			out = append(out,
				shared.Node{
					IsExpression: true,
					Children: []shared.Node{{Val: tok}}})

		default:
			out = append(out, shared.Node{Val: tok})
		}

		if len(*toks) < 1 {
			break
		}

		*toks = (*toks)[1:]
	}

	return out
}
