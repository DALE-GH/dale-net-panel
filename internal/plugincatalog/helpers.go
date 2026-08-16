package plugincatalog

func compactVersion(v string) []string {
	if v == "" {
		return nil
	}
	return []string{v}
}

func tail[T any](in []T, n int) []T {
	if len(in) <= n {
		return in
	}
	return in[len(in)-n:]
}
