# Functional options

Usage:

```go
import (
	"github.com/ciricc/go-foptions"
)

type Settings struct {
	Limit int
}

func WithLimit(l int) foptions.Opt[Settings] {
	return func(s *Settings) error {
		if l <= 0 {
			return errors.New("invalid limit")
		}

		s.Limit = l
		
		return nil
	}
}

type Service struct{
	s *Settings
}

func NewService(opts ...foptions.Opt[Settings]) (*Service, error) {
	settings, err := foptions.Use(&Settings{
		Limit: 100
	}, opts...) 
	if err != nil {
		return nil, err
	}
	
	return &Service{s: settings}
}
```

Maybe later i will find out how to do this more simply.