patch: build-patch git-push
minor: build-minor git-push
major: build-major git-push

build-patch:
	bumpversion patch

build-minor:
	bumpversion minor

build-major:
	bumpversion major

git-push:
	git push --all origin
	git push --tags origin