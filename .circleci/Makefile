ci:
	docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd):/root/project -w /root/project circleci/circleci-cli:alpine@arm64 circleci config validate
