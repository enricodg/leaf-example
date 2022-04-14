# =================================================
# Example
# =================================================
# mockgen
# --source=internal/usecases/apiKey/usecase.go
# -package=mock_usecase_apiKey
# --destination=mocks/usecases/apiKey/mock_usecase.go

# - UseCase
mkdir -p mocks/usecases

# - Outbound
# - Outbound.Repositories
mkdir -p mocks/outbound/repositories

# - Outbound.Cache
mkdir -p mocks/outbound/cache

# - Outbound.Messaging
mkdir -p mocks/outbound/messaging

# - Outbound.Webservices
mkdir -p mocks/outbound/webservices

# ==========================================
# ================= RESOURCES ==============
# ==========================================
mkdir -p mocks/resource
