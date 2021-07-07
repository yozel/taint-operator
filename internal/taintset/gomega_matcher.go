package taintset

import (
	"github.com/onsi/gomega/types"

	"fmt"
)

func Contains(expected interface{}) types.GomegaMatcher {
	return &contains{
		expected: expected,
	}
}

type contains struct {
	expected interface{}
}

func (matcher *contains) Match(actual interface{}) (success bool, err error) {
	actualo, ok := actual.(*TaintSet)
	if !ok {
		return false, fmt.Errorf("Contains matcher expects an TaintSet")
	}

	expectedo, ok := matcher.expected.(*TaintSet)
	if !ok {
		return false, fmt.Errorf("Contains matcher expects an TaintSet")
	}

	return actualo.IsSupersetOf(expectedo), nil
}

func (matcher *contains) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto contain \n\t%#v", actual, matcher.expected)
}

func (matcher *contains) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to contain \n\t%#v", actual, matcher.expected)
}
