package security

import (
	"errors"
	"github.com/casbin/casbin/v2/effector"
)

const (
	AllowOverrideEffect   = "some(where (p_effect == allow))"
	DenyOverrideEffect    = "!some(where (p_effect == deny))"
	AllowAndDenyEffect    = "some(where (p_effect == allow)) && !some(where (p_effect == deny))"
	PriorityEffect        = "priority(p_effect) || deny"
	SubjectPriorityEffect = "subjectPriority(p_effect) || deny"
)

// Effector is default effector for Casbin.
type Effector struct {
}

// NewSecurityEffector is the constructor for SecurityEffector.
func NewSecurityEffector() *Effector {
	e := Effector{}
	return &e
}

// MergeEffects merges all matching results collected by the enforcer into a single decision.
func (e *Effector) MergeEffects(expr string, effects []effector.Effect, matches []float64, policyIndex int, policyLength int) (effector.Effect, int, error) {
	result := effector.Indeterminate
	explainIndex := -1

	switch expr {
	case AllowOverrideEffect:
		if matches[policyIndex] == 0 {
			break
		}
		// only check the current policyIndex
		if effects[policyIndex] == effector.Allow {
			result = effector.Allow
			explainIndex = policyIndex
			break
		}
	case DenyOverrideEffect:
		// only check the current policyIndex
		if matches[policyIndex] != 0 && effects[policyIndex] == effector.Deny {
			result = effector.Deny
			explainIndex = policyIndex
			break
		}
		// if no deny rules are matched  at last, then allow
		if policyIndex == policyLength-1 {
			result = effector.Allow
		}
	case AllowAndDenyEffect:
		// short-circuit if matched deny rule
		if matches[policyIndex] != 0 && effects[policyIndex] == effector.Deny {
			result = effector.Deny
			// set hit rule to the (first) matched deny rule
			explainIndex = policyIndex
			break
		}

		// short-circuit some effects in the middle
		if policyIndex < policyLength-1 {
			// choose not to short-circuit
			return result, explainIndex, nil
		}
		// merge all effects at last
		for i, effect := range effects {
			if matches[i] == 0 {
				continue
			}

			if effect == effector.Allow {
				result = effector.Allow
				// set hit rule to first matched allow rule
				explainIndex = i
				break
			}
		}
	case PriorityEffect, SubjectPriorityEffect:
		// reverse merge, short-circuit may be earlier
		for i := len(effects) - 1; i >= 0; i-- {
			if matches[i] == 0 {
				continue
			}

			if effects[i] != effector.Indeterminate {
				if effects[i] == effector.Allow {
					result = effector.Allow
				} else {
					result = effector.Deny
				}
				explainIndex = i
				break
			}
		}
	default:
		return effector.Deny, -1, errors.New("unsupported effect")
	}

	return result, explainIndex, nil
}
