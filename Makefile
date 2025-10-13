build:
	[ -d public ] && rm -rf public
	hugo --minify --buildFuture

preview:
	hugo server -buildDrafts --buildExpired --buildFuture --disableFastRender --baseURL http://localhost:1313/personal-blog/
