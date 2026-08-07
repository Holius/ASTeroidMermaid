awk:
	awk -f go_metadata.awk $$(find zarf -name '*.go' -type f | xargs) > filedata/expected_go_filedata.go

mini_awk:
	awk -f go_metadata.awk zarf/src/internal/template/template.go > filedata/expected_go_filedata.go

setup:
	git submodule update --init --recursive -v

test:
	go test -v $$(find . -path "./.*" -type d -prune -o -path "./zarf*" -type d -prune -o -type d -print) 

mini_test:
	go test -v $$(find . -path "./.*" -type d -prune -o -path "./zarf*" -type d -prune -o -type d -print) -run ^TestGoPackageDataFromGoMetadataWithoutTests$