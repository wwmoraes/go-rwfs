.PHONY: hooks/%
hooks/pre-bump:: check

hooks/post-bump::
	git push
	$(if ${VERSION},git push ${VERSION})
