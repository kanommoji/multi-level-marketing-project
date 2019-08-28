all :  test build-artifact accept push-artifact

build-artifact:
	# build app
	# docker image build -t jinnerzaza/multilevel:1 .
	docker-compose build


test:
	docker-compose up -d
	go test ./... --cover
	docker-compose down

accept:
	# run app
	# docker run
	docker-compose up -d
	# run ATDD
	newman run atdd/api/promote-member-level.json
	newman run atdd/api/demote-member-level.json
	# docker stop and rm
	docker-compose down

down:
	docker-compose down


