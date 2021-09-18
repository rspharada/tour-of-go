module go-sample

go 1.16

replace go-sample/a => ./a

replace go-sample/b => ./b

require (
	go-sample/a v0.0.0-00010101000000-000000000000
	go-sample/b v0.0.0-00010101000000-000000000000 // indirect
)
