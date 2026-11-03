test:
	@cd golings && go test -coverprofile=coverage.out -v $$(go list ./... | grep -v fixtures/error1)

watch:
	@go run golings/golings.go watch

run:
	@go run golings/golings.go run

# Installs golings into a temporary folder and runs the first commands a learner types.
smoke:
	@scripts/smoke.sh
