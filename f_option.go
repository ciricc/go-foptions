package foptions

type Opt[S any] func(settings *S) error

// Use invokes each option and if the error occured on some of them, returns that error.
// Any change of the settings after use of any option will not be canceled if error occured.
func Use[S any, O Opt[S]](settings *S, opts ...O) (*S, error) {
	for _, opt := range opts {
		err := opt(settings)
		if err != nil {
			return settings, err
		}
	}

	return settings, nil
}
