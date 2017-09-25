snapshot:
	cd devenv; goreleaser --snapshot --rm-dist
release:
	cd devenv; goreleaser --rm-dist
