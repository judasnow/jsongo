package main

import (
	"fmt"
)
type state int

type FARule struct {
	state state
	nextState state
	char string
}
func (r FARule) appliesTo(state state, char string) bool {
	return r.state == state && r.char == char
}
func (r FARule) follow() state {
	return r.nextState
}
func (r FARule) inspect() {
	fmt.Printf("%s->%s->%s", r.state, r.char, r.nextState)
}

type FARulebook struct {
	rules []FARule
}
// GO 的默认初始化为 0 值，感觉就像是

//func NewFARulebook(rules []FARule) FARulebook {
//	return FARulebook{
//		status: rules,
//	}
//}
func (book FARulebook) nextState(state state, char string) state {
	return book.ruleFor(state, char).follow()
}
func (book FARulebook) ruleFor(state state, char string) FARule {
	for _, r := range book.rules {
		if r.appliesTo(state, char) {
			return r
		}
	}
	// 原代码这里返回 nil
	panic("no rule")
}

type DFA struct {
	crtState state
	acceptStats []state
	ruleBook FARulebook
}
// 当前状态是不是刻意接收的状态
func (dfa DFA) accepting () bool {
	for _, s := range dfa.acceptStats {
		if s == dfa.crtState {
			return true
		}
	}
	return false
}
func NewDFA(crtState state, acceptStats []state, ruleBook FARulebook) DFA {
	return DFA{
		crtState: crtState,
		acceptStats: acceptStats,
		ruleBook: ruleBook,
	}
}

func main() {
	b := FARulebook{
		rules: []FARule{
			FARule{state: 1, char: "a", nextState: 2},
			FARule{state: 1, char: "b", nextState: 1},
			FARule{state: 2, char: "a", nextState: 2},
			FARule{state: 2, char: "b", nextState: 3},
			FARule{state: 3, char: "a", nextState: 3},
			FARule{state: 3, char: "b", nextState: 3},
		},
	}

	dfa := NewDFA(1, []state{1,3}, b)
	fmt.Println(dfa.accepting())
}
