module github.com/nanofuxion/kik_go_api

go 1.15

require (
	github.com/google/uuid v1.1.2
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
)

replace github.com/nanofuxion/kik_go_api/send => ./send

replace github.com/nanofuxion/kik_go_api/utils => ./utils
