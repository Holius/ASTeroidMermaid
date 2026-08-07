awk:
	awk -f go_metadata.awk $$(find zarf -name '*.go' -type f | xargs) > expected_go_metadata.go

mini_awk:
	awk -f go_metadata.awk zarf/src/internal/template/template.go > expected_go_metadata.go

setup:
	git submodule update --init --recursive -v

test:
	go test . -v