package basaltic

import (
	"context"
	"iter"
)

// Page is one page of a list operation.
//
// Pages are cursor-based: carry [Page.Marker] into the next request to get
// the page after this one. Do not construct or parse a marker — its internal
// form varies by endpoint and is not stable across releases.
//
// Page until HasMore is false, not until a page looks short. The platform
// clamps an over-large limit rather than refusing it, so a page shorter than
// the one you asked for is normal and says nothing about whether more exist.
type Page[T any] struct {
	// Items are this page's results.
	Items []T
	// Total is how many items exist in all, where the operation reports it.
	Total int
	// Limit is the page size the platform actually applied, which may be
	// smaller than the one requested.
	Limit int
	// Marker is the cursor for the next page. Empty when there is none.
	Marker string
	// HasMore reports whether another page exists. Operations that do not
	// paginate always report false.
	HasMore bool
}

// Len returns the number of items on this page.
func (p *Page[T]) Len() int {
	if p == nil {
		return 0
	}
	return len(p.Items)
}

// pageFetcher retrieves the page following marker. An empty marker asks for
// the first page.
type pageFetcher[T any] func(ctx context.Context, marker string) (*Page[T], error)

// paginate walks every page, yielding items one at a time.
//
// It stops at the first error, yielding the zero item alongside it, so a
// caller that checks err on each step cannot mistake a truncated walk for a
// complete one:
//
//	for item, err := range seq {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
func paginate[T any](ctx context.Context, fetch pageFetcher[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		marker := ""
		for {
			page, err := fetch(ctx, marker)
			if err != nil {
				yield(zero, err)
				return
			}
			for _, item := range page.Items {
				if !yield(item, nil) {
					return
				}
			}
			// A page that claims more but hands back no cursor would spin
			// forever re-fetching the same page. Stop instead.
			if !page.HasMore || page.Marker == "" || page.Marker == marker {
				return
			}
			marker = page.Marker
		}
	}
}

// Paginate walks a cursor-paginated operation, yielding every item across
// every page.
//
// Generated "All" methods are built on this. Use it directly only for a
// list endpoint the SDK does not yet generate.
func Paginate[T any](ctx context.Context, fetch func(ctx context.Context, marker string) (*Page[T], error)) iter.Seq2[T, error] {
	return paginate(ctx, pageFetcher[T](fetch))
}

// Collect drains an iterator into a slice, stopping at the first error.
//
// Convenient for small result sets. It holds every item in memory, so prefer
// ranging over the iterator where the set could be large.
func Collect[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var out []T
	for item, err := range seq {
		if err != nil {
			return out, err
		}
		out = append(out, item)
	}
	return out, nil
}
