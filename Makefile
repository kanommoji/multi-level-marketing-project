all :  unit-test build-artifact accept push-artifact

unit-test:
	mysql -umlm_dev -pmlm_dev < sql/prepared-data.sql
	TIME=20190701 go test ./... --cover

build-artifact:
	# build app
	# docker image build -t jinnerzaza/multilevel:1 .
	docker-compose build

accept:
	# run app
	# docker run
	docker-compose up -d
	# run ATDD
	mysql -umlm_dev -pmlm_dev < sql/prepared-data.sql
	newman run atdd/api/promote-member-level.json
	newman run atdd/api/demote-member-level.json
	# docker stop and rm
	docker-compose down

down:
	docker-compose down

