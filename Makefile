build-docker-image:
ifndef TAG
	ERR = $(error TAG is undefined)
	$(ERR)
endif
	docker build --no-cache -t healthz:${TAG} .

push-docker-image:
ifndef REGISTRY
	ERR = $(error REGISTRY is undefined)
	$(ERR)
endif
	docker tag healthz:${TAG} ${REGISTRY}/healthz:${TAG}
	docker push ${REGISTRY}/healthz:${TAG}

release: build-docker-image push-docker-image
	echo "healthz:${TAG} is published"
