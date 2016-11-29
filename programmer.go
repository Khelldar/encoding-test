package people

import "time"

type Programmer struct {
	Name             string
	Birthday         time.Time
	FavoriteLanguage Language
	LastScore        int
}

type Language int

const (
	Go Language = iota
	JavaScript
	CSharp
	CPlusPlus
)
