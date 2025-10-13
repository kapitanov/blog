build:
	hugo

preview:
	hugo server -buildDrafts --buildExpired --buildFuture --disableFastRender --baseURL http://localhost:1313/personal-blog/
