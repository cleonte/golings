test:
	@cd golings && go test -coverprofile=coverage.out -v $$(go list ./... | grep -v fixtures/error1)

# The packages of the tool: ./... would also load the exercises and fixtures/error1, which do not compile on purpose.
lint:
	@cd golings && golangci-lint run . ./cmd/... ./doctor/... ./exercises/... ./ui/... ./workspace/... ./tui/... ./fixtures/next/... ./fixtures/pending1/... ./fixtures/success1/...

watch:
	@go run golings/golings.go watch

run:
	@go run golings/golings.go run

# Installs golings into a temporary folder and runs the first commands a learner types.
smoke:
	@scripts/smoke.sh
