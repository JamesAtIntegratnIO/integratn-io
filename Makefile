buildprod:
	hugo --gc --minify
buildpreview:
	hugo --gc --minify --buildDrafts --buildFuture -b $DEPLOY_PRIME_URL
buildgo:
	cd netlify/functions/gozette
	pwd
	ls
	go get .
	go build -o  ../gozette-function .