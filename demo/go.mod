module demo

go 1.16

// 本地包
replace github.com/fjs-icu/gd => ../

replace github.com/fjs-icu/win => ../../win

require (
	github.com/fjs-icu/gd v0.0.0-00010101000000-000000000000
	github.com/fjs-icu/win v0.0.0-00010101000000-000000000000 // indirect
)
