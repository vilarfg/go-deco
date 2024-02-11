package deco_test

import (
	"strings"
	"testing"

	"github.com/vilarfg/go-deco"
)

func Test(t *testing.T) {
	var (
		stop      = func(s string) string { return s + ". Stop it." }
		name      = func(s string) string { return "Ford, " + s }
		turning   = func(s string) string { return "you're turning into " + s }
		uppercase = func(s string) string { return strings.ToUpper(s) }
		identity  = func(s string) string { return s }

		chain0 = deco.Chain[string]()
		chain1 = deco.Chain(stop, name)
		chain2 = chain1.Extend()
		chain3 = chain1.Extend(turning)
		chain4 = chain3.Extend(uppercase)
		chain5 = chain4.Extend(nil, identity)
		chain6 = deco.Chain(uppercase)
		chain7 = chain6.Extend(turning, stop)
	)
	for _, tc := range []struct {
		chain    deco.Decorator[string]
		input    string
		expected string
	}{
		{chain0, "whale", "whale"},
		{chain1, "bowl of petunias", "Ford, bowl of petunias. Stop it."},
		{chain2, "bowl of petunias", "Ford, bowl of petunias. Stop it."},
		{chain3, "a sofa", "Ford, you're turning into a sofa. Stop it."},
		{chain4, "yarn", "Ford, you're turning into YARN. Stop it."},
		{chain5, "a penguin", "Ford, you're turning into A PENGUIN. Stop it."},
		{chain6, "whale", "WHALE"},
		{chain7, "a sofa", "YOU'RE TURNING INTO A SOFA. STOP IT."},
	} {
		var got = tc.chain(tc.input)
		if got != tc.expected {
			t.Errorf(`expected: "%s", got: "%s"`, tc.expected, got)
		}
		if got = tc.chain.Apply(tc.input); got != tc.expected {
			t.Errorf(`expected: "%s", got: "%s"`, tc.expected, got)
		}
	}
}
