module github.com/Hardw01f/goLeetWord

go 1.13

require (
	github.com/Hardw01f/goLeetWord/cmd v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/Hardw01f/goLeetWord/cmd => ./cmd
	github.com/Hardw01f/goLeetWord/pkg => ./pkg
)
