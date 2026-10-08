package basaltic

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

// pagedFetcher serves a fixed set of pages, keyed by the marker that asks for
// them, and records what it was asked for.
type pagedFetcher struct {
	pages    []*Page[string]
	requests []string
	err      error
	errAt    int
}

func (f *pagedFetcher) fetch(ctx context.Context, marker string) (*Page[string], error) {
	f.requests = append(f.requests, marker)
	idx := len(f.requests) - 1
	if f.err != nil && idx == f.errAt {
		return nil, f.err
	}
	if idx >= len(f.pages) {
		return &Page[string]{}, nil
	}
	return f.pages[idx], nil
}

func TestPaginateWalksEveryPage(t *testing.T) {
	f := &pagedFetcher{pages: []*Page[string]{
		{Items: []string{"a", "b"}, Marker: "m1", HasMore: true},
		{Items: []string{"c", "d"}, Marker: "m2", HasMore: true},
		{Items: []string{"e"}, HasMore: false},
	}}

	var got []string
	for item, err := range Paginate(context.Background(), f.fetch) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got = append(got, item)
	}

	want := []string{"a", "b", "c", "d", "e"}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("items = %v, want %v", got, want)
	}
	wantMarkers := []string{"", "m1", "m2"}
	if fmt.Sprint(f.requests) != fmt.Sprint(wantMarkers) {
		t.Errorf("markers requested = %v, want %v", f.requests, wantMarkers)
	}
}

func TestPaginateStopsAtTheFirstError(t *testing.T) {
	sentinel := errors.New("boom")
	f := &pagedFetcher{
		pages: []*Page[string]{
			{Items: []string{"a"}, Marker: "m1", HasMore: true},
		},
		err:   sentinel,
		errAt: 1,
	}

	var got []string
	var gotErr error
	for item, err := range Paginate(context.Background(), f.fetch) {
		if err != nil {
			gotErr = err
			break
		}
		got = append(got, item)
	}
	if !errors.Is(gotErr, sentinel) {
		t.Errorf("error = %v, want the fetch error", gotErr)
	}
	if fmt.Sprint(got) != fmt.Sprint([]string{"a"}) {
		t.Errorf("items before the error = %v, want [a]", got)
	}
}

func TestPaginateStopsWhenTheCallerBreaks(t *testing.T) {
	f := &pagedFetcher{pages: []*Page[string]{
		{Items: []string{"a", "b"}, Marker: "m1", HasMore: true},
		{Items: []string{"c"}, Marker: "m2", HasMore: true},
	}}

	for item := range Paginate(context.Background(), f.fetch) {
		if item == "b" {
			break
		}
	}
	if len(f.requests) != 1 {
		t.Errorf("made %d requests after breaking on the first page, want 1", len(f.requests))
	}
}

func TestPaginateDoesNotSpinOnAMissingCursor(t *testing.T) {
	// A page that claims more but hands back no cursor would otherwise be
	// re-fetched forever.
	f := &pagedFetcher{pages: []*Page[string]{
		{Items: []string{"a"}, HasMore: true, Marker: ""},
	}}
	var count int
	for range Paginate(context.Background(), f.fetch) {
		count++
	}
	if count != 1 || len(f.requests) != 1 {
		t.Errorf("yielded %d items over %d requests, want 1 and 1", count, len(f.requests))
	}
}

func TestPaginateDoesNotSpinOnARepeatedCursor(t *testing.T) {
	// A service that keeps handing back the marker it was given would
	// otherwise be fetched in a loop that never ends.
	var requests int
	stuck := func(ctx context.Context, marker string) (*Page[string], error) {
		requests++
		if requests > 10 {
			t.Fatal("Paginate kept requesting a page whose cursor never advanced")
		}
		return &Page[string]{Items: []string{"a"}, Marker: "same", HasMore: true}, nil
	}

	var count int
	for range Paginate(context.Background(), stuck) {
		count++
	}
	// The first page advances the cursor from "" to "same"; the second
	// returns "same" again and the walk stops.
	if requests != 2 {
		t.Errorf("made %d requests, want 2 before noticing the cursor stopped moving", requests)
	}
	if count != 2 {
		t.Errorf("yielded %d items, want 2", count)
	}
}

func TestCollect(t *testing.T) {
	f := &pagedFetcher{pages: []*Page[string]{
		{Items: []string{"a", "b"}, Marker: "m1", HasMore: true},
		{Items: []string{"c"}},
	}}
	got, err := Collect(Paginate(context.Background(), f.fetch))
	if err != nil {
		t.Fatalf("Collect: %v", err)
	}
	if fmt.Sprint(got) != fmt.Sprint([]string{"a", "b", "c"}) {
		t.Errorf("Collect() = %v, want [a b c]", got)
	}
}

func TestCollectReturnsWhatItGotBeforeAnError(t *testing.T) {
	sentinel := errors.New("boom")
	f := &pagedFetcher{
		pages: []*Page[string]{{Items: []string{"a"}, Marker: "m1", HasMore: true}},
		err:   sentinel,
		errAt: 1,
	}
	got, err := Collect(Paginate(context.Background(), f.fetch))
	if !errors.Is(err, sentinel) {
		t.Errorf("err = %v, want the fetch error", err)
	}
	if fmt.Sprint(got) != fmt.Sprint([]string{"a"}) {
		t.Errorf("Collect() = %v, want the items gathered before the failure", got)
	}
}
