package runtime

import "testing"

func TestParseRuntimeMapEntryPreservesCommaInValue(t *testing.T) {
	response := `type=str, case=sensitive, found=yes, idx=tree, key="example.test", value="[{"h":"websrv2","cn":"EU"},{"h":"websrv3","cn":"NA"}]", type="str"`

	entry := parseRuntimeMapEntry(response)
	if entry == nil {
		t.Fatal("parseRuntimeMapEntry() = nil, want entry")
	}
	if entry.Key != "example.test" {
		t.Errorf("Key = %q, want %q", entry.Key, "example.test")
	}
	wantValue := `[{"h":"websrv2","cn":"EU"},{"h":"websrv3","cn":"NA"}]`
	if entry.Value != wantValue {
		t.Errorf("Value = %q, want %q", entry.Value, wantValue)
	}
}

func TestParseRuntimeMapEntryReturnsNilForMissingEntry(t *testing.T) {
	response := `type=str, case=sensitive, found=no, idx=tree`
	if entry := parseRuntimeMapEntry(response); entry != nil {
		t.Fatalf("parseRuntimeMapEntry() = %#v, want nil", entry)
	}
}
