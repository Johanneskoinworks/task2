build :
	docker build -t asgard-service-example . --build-arg GITHUB_USERNAME=${GITHUB_USERNAME} --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}