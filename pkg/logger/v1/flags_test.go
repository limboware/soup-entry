package loggerv1

import "testing"

func TestThatSerializationIsOk_001(t *testing.T) {
	expected := "logs,warns,errs"
	val := PrintLogs.With(PrintWarns).With(PrintErrs)

	if result := val.String(); result != expected {
		t.Errorf("expected: %s, but got %s", expected, result)
	}
}

func TestThatSerializationIsOk_002(t *testing.T) {
	expected := ""
	val := None

	if result := val.String(); result != expected {
		t.Errorf("expected: %s, but got %s", expected, result)
	}
}

func TestThatSerializationIsOk_003(t *testing.T) {
	expected := "debugs,sys,warns,errs"
	val := PrintDebug.With(PrintSystem).With(PrintWarns.With(PrintErrs))

	if result := val.String(); result != expected {
		t.Errorf("expected: %s, but got %s", expected, result)
	}
}