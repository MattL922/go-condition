package condition

import (
    "testing"
)

func alwaysTrue() bool {
    return true
}

func alwaysFalse() bool {
    return false
}

func TestAll(t *testing.T) {
    if !All(alwaysTrue, alwaysTrue) {
        t.Fatal("All(true, true) failed")
    }
    if All(alwaysTrue, alwaysFalse) {
        t.Fatal("All(true, false) failed")
    }
    if All(alwaysFalse, alwaysFalse) {
        t.Fatal("All(false, false) failed")
    }
}

func TestAny(t *testing.T) {
    if !Any(alwaysTrue, alwaysTrue) {
        t.Fatal("Any(true, true) failed")
    }
    if !Any(alwaysTrue, alwaysFalse) {
        t.Fatal("Any(true, false) failed")
    }
    if Any(alwaysFalse, alwaysFalse) {
        t.Fatal("Any(false, false) failed")
    }
}

func TestNone(t *testing.T) {
    if None(alwaysTrue, alwaysTrue) {
        t.Fatal("None(true, true) failed")
    }
    if None(alwaysTrue, alwaysFalse) {
        t.Fatal("None(true, false) failed")
    }
    if !None(alwaysFalse, alwaysFalse) {
        t.Fatal("None(false, false) failed")
    }
}
