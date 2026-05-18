module github.com/grokify/prism

go 1.26.2

replace (
	github.com/grokify/prism-capability => ../prism-capability
	github.com/grokify/prism-core => ../prism-core
	github.com/grokify/prism-execution => ../prism-execution
	github.com/grokify/prism-intelligence => ../prism-intelligence
)

require (
	github.com/grokify/prism-capability v0.0.0
	github.com/grokify/prism-execution v0.11.0
	github.com/grokify/prism-intelligence v0.0.0
)

require (
	github.com/grokify/prism-core v0.0.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)
