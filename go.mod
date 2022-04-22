module github.com/enricodg/leaf-example

go 1.18

require (
	github.com/getsentry/sentry-go v0.13.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/golang/mock v1.6.0
	github.com/labstack/echo/v4 v4.7.2
	github.com/labstack/gommon v0.3.1
	github.com/paulusrobin/leaf-utilities/appRunner v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/cache/cache v0.0.0-20220420070753-f0289b1e394b
	github.com/paulusrobin/leaf-utilities/cache/integrations/memcache v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/cache/integrations/redis v0.0.0-20220413144204-972ab9e3b19d
	github.com/paulusrobin/leaf-utilities/common v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/config v0.0.0-20220413144204-972ab9e3b19d
	github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm/mysql v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm/plugin/softDelete v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/database/sql/sql v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/encoding/json v0.0.0-20220420071653-80e4cae9f4d2
	github.com/paulusrobin/leaf-utilities/logger/integrations/logrus v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/logger/logger v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/mandatory v0.0.0-20220420081700-6d04b27d3e05
	github.com/paulusrobin/leaf-utilities/messageQueue/integrations/kafka v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/time v0.0.0-20220420071653-80e4cae9f4d2
	github.com/paulusrobin/leaf-utilities/token v0.0.0-20220413144204-972ab9e3b19d
	github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic v0.0.0-20220420071653-80e4cae9f4d2
	github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/tracer/tracer v0.0.0-20220420071653-80e4cae9f4d2
	github.com/paulusrobin/leaf-utilities/translator v0.0.0-20220413144204-972ab9e3b19d
	github.com/paulusrobin/leaf-utilities/validator/integrations/v10 v10.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/validator/validator v0.0.0-20220420070019-7843b97718b6
	github.com/paulusrobin/leaf-utilities/webClient/integrations/heimdall v0.0.0-20220422095901-cbdbd676731a
	github.com/paulusrobin/leaf-utilities/webClient/webClient v0.0.0-20220420071653-80e4cae9f4d2
	github.com/stretchr/testify v1.7.0
	go.uber.org/dig v1.14.1
	gorm.io/gorm v1.23.4
	gorm.io/plugin/optimisticlock v1.0.7
)

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/compute v0.1.0 // indirect
	cloud.google.com/go/profiler v0.2.0 // indirect
	github.com/DataDog/datadog-go v3.7.1+incompatible // indirect
	github.com/Shopify/sarama v1.32.0 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/cactus/go-statsd-client/statsd v0.0.0-20200423205355-cb0885a1018c // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gojek/heimdall/v7 v7.0.2 // indirect
	github.com/gojek/valkyrie v0.0.0-20180215180059-6aee720afcdf // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/pprof v0.0.0-20220113144219-d25a53d42d00 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/newrelic/go-agent/v3 v3.15.2 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrlogrus v1.0.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm v0.0.0-20220420081227-c2758aaa663c // indirect
	github.com/paulusrobin/leaf-utilities/slack v0.0.0-20220420071653-80e4cae9f4d2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/thoas/go-funk v0.9.2 // indirect
	github.com/ua-parser/uap-go v0.0.0-20211112212520-00c877edfe0f // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/api v0.65.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220111164026-67b88f271998 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
)
