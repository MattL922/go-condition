package condition

// Condition is a function that evaluates to true or false.
type Condition func() bool

// All returns true only if all of the provided conditions are true, false
// otherwise.
func All(conditions ...Condition) bool {
    for _, condition := range conditions {
        if !condition() { return false }
    }
    return true
}

// Any returns true if at least one of the provided conditions is true, false
// otherwise.
func Any(conditions ...Condition) bool {
    for _, condition := range conditions {
        if condition() { return true }
    }
    return false
}

// None returns true if none of the provided conditions are true, false
// otherwise.
func None(conditions ...Condition) bool {
    for _, condition := range conditions {
        if condition() { return false }
    }
    return true
}

