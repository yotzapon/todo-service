migrate-create:
	# $ make migrate-create name=create_payment_code_owner_mapping_table
	go get github.com/golang-migrate/migrate/v4/cmd/migrate
	go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir migrations $(name)

_bindata:
	go get -d github.com/go-bindata/go-bindata/go-bindata
	go run github.com/go-bindata/go-bindata/go-bindata -nocompress -prefix "./migrations/" -pkg "migrations" -o "internal/bindata/migrations/migrations.go" "migrations"
	go mod tidy

generate_spec:
	swag init

generate_mock:
	go run github.com/golang/mock/mockgen -destination=./mocks/mock_db.go -package=mocks github.com/yotzapon/todo-service/internal/database DB
	go run github.com/golang/mock/mockgen -destination=./mocks/mock_db_todo.go -package=mocks github.com/yotzapon/todo-service/internal/database TodoRepositoryInterface
	go run github.com/golang/mock/mockgen -destination=./mocks/mock_db_user.go -package=mocks github.com/yotzapon/todo-service/internal/database UserRepositoryInterface
	go run github.com/golang/mock/mockgen -destination=./mocks/mock_service_auth.go -package=mocks github.com/yotzapon/todo-service/internal/services AuthServiceInterface
	go run github.com/golang/mock/mockgen -destination=./mocks/mock_service_todo.go -package=mocks github.com/yotzapon/todo-service/internal/services TodoServiceInterface

e2e:
	sh ./tests/automated/run.sh

commit_check:
	pre-commit run
