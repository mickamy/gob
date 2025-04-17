APP_NAME = gob
VERSION ?= dev
BUILD_DIR = bin
GORELEASER ?= go tool goreleaser
VERSION_VARIABLE = github.com/mickamy/gob/internal/cli/version/version.version

.PHONY: all up down build install uninstall clean version test fmt

all: build

up: down
	docker compose up -d

down:
	docker compose down

build:
	@echo "🔨 Building $(APP_NAME)..."
	go build -ldflags "-X $(VERSION_VARIABLE)=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/gob

install:
	@echo "📦 Installing $(APP_NAME)..."
	go install -ldflags "-X $(VERSION_VARIABLE)=$(VERSION)" ./cmd/gob

uninstall:
	@echo "🗑️  Uninstalling $(APP_NAME)..."
	@bin_dir=$$(go env GOBIN); \
	if [ -z "$$bin_dir" ]; then \
		bin_dir=$$(go env GOPATH)/bin; \
	fi; \
	echo "Removing $$bin_dir/$(APP_NAME)"; \
	rm -f $$bin_dir/$(APP_NAME)

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BUILD_DIR)

version:
	@echo "🔖 Version: $(VERSION)"

test:
	@echo "🧪 Running tests..."
	go test ./...

test-panic:
	@echo "🧪 Testing: expected panic (gob run)..."
	@$(APP_NAME) run ./testdata/panic.go --open || echo "💥 Panic detected and reported"

test-ok:
	@echo "🧪 Testing: no panic expected (gob run)..."
	@$(APP_NAME) run ./testdata/no_panic.go --open

test-lib:
	@echo "🧪 Testing: panic using gob.Handle() (embedded)..."
	@go run ./testdata/handle.go || echo "💥 Panic detected and reported"

fmt:
	@echo "📝 Formatting code..."
	gofmt -w -l .

release:
	@echo "🚀 Running release..."
	$(GORELEASER) release --clean

snapshot:
	@echo "🔍 Running snapshot release (dry run)..."
	$(GORELEASER) release --snapshot --clean
