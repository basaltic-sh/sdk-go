package basaltic

import "time"

// Optional fields in request bodies are pointers, so that sending a zero
// value and leaving a field alone stay distinguishable — an update that set
// a field to false or 0 would otherwise be indistinguishable from one that
// did not mention it, and would silently leave the old value in place.
//
// These helpers exist because Go has no literal for the address of a
// constant. Response and model types use plain values; only request bodies
// need this.

// Ptr returns a pointer to v. Use it to set an optional request field:
//
//	req := &compute.InstanceUpdateRequest{
//		Name: basaltic.Ptr("web-02"),
//	}
func Ptr[T any](v T) *T { return &v }

// Deref returns the value p points at, or the zero value when p is nil.
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// DerefOr returns the value p points at, or fallback when p is nil.
func DerefOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}

// String returns a pointer to s. Shorthand for [Ptr].
func String(s string) *string { return &s }

// Bool returns a pointer to b. Shorthand for [Ptr].
func Bool(b bool) *bool { return &b }

// Int returns a pointer to i. Shorthand for [Ptr].
func Int(i int) *int { return &i }

// Int64 returns a pointer to i. Shorthand for [Ptr].
func Int64(i int64) *int64 { return &i }

// Float64 returns a pointer to f. Shorthand for [Ptr].
func Float64(f float64) *float64 { return &f }

// Time returns a pointer to t. Shorthand for [Ptr].
func Time(t time.Time) *time.Time { return &t }
