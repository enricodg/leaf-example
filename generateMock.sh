# =================================================
# Example
# =================================================
# mockgen
# --source=internal/usecases/apiKey/usecase.go
# -package=mock_usecase_apiKey
# --destination=mocks/usecases/apiKey/mock_usecase.go

# - UseCase
mkdir -p mocks/usecases
mockgen --source=internal/usecases/serviceParameter/usecase.go -package=mock_usecase_serviceParameter --destination=mocks/usecases/serviceParameter/mock_usecase.go

# - Outbound
# - Outbound.Repositories
mkdir -p mocks/outbound/repositories
mockgen --source=internal/outbound/repositories/serviceParameter/repository.go -package=mock_repository_serviceParameter --destination=mocks/outbound/repositories/serviceParameter/mock_repository.go

# ==========================================
# ================= RESOURCES ==============
# ==========================================
mkdir -p mocks/resource
# -> logger
mockgen -destination=mocks/resource/logger/mock_logger.go -package=mock_leafLogger github.com/paulusrobin/leaf-utilities/logger/logger Logger,StandardLogger
# -> sql
mockgen -destination=mocks/resource/mysql/mock_orm.go -package=mock_leafSql github.com/paulusrobin/leaf-utilities/database/sql/sql ORM
