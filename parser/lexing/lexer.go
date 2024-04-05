package lexing

import (
	"experiment/parser/schema"
	"fmt"
	"regexp"
)

func Lex(source string, nodeType string, lexer schema.Lexer) schema.AstNode {

	fmt.Println(lexer[nodeType])

	result := new(schema.AstNode)

	// this lexer

	//fmt.Println(abstractSyntaxTree)

	/*
		run the node through its tests
		if there is a match perform it's split
			then loop through the splits
			and lex those nodes
		else if there is no match eject or record error
	*/

	/*
		astNode = &schema.AstNode{
			NodeType:   "",               // type of node
			Text:       "",               // the raw text of the node
			LineNumber: 1,                // the source line of this node
			CharStart:  1,                // start character of this node
			CharEnd:    1,                // end character of this node
			Child:      [], // zero or more child nodes
		}
	*/

	return *result
}

func splitter(text string, regex string) []string {
	re := regexp.MustCompile(regex)
	result := re.Split(text, -1)
	return result
}
