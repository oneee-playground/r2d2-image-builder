module github.com/oneee-playground/r2d2-image-builder

go 1.22

require (
	github.com/GoogleContainerTools/kaniko v1.22.0
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go v1.53.12
	github.com/aws/aws-sdk-go-v2 v1.27.0
	github.com/aws/aws-sdk-go-v2/config v1.27.11
	github.com/aws/aws-sdk-go-v2/service/sqs v1.32.3
	github.com/go-git/go-billy/v5 v5.5.0
	github.com/go-git/go-git/v5 v5.12.0
	github.com/google/go-containerregistry v0.19.1
	github.com/google/uuid v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.9.0
)

replace github.com/GoogleContainerTools/kaniko => github.com/oneee-playground/kaniko v0.0.0-20240505131118-0b8b2ed26525

require (
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20230306123547-8075edf89bb0 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.12 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Microsoft/hcsshim v0.12.3 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.11 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.27.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.23.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.6 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/awslabs/amazon-ecr-credential-helper/ecr-login v0.0.0-20240419161514-af205d85bb44 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chrismellard/docker-credential-acr-env v0.0.0-20230304212654-82a0ddb27589 // indirect
	github.com/cilium/ebpf v0.11.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/containerd/cgroups/v3 v3.0.3 // indirect
	github.com/containerd/containerd v1.7.16 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/containerd/errdefs v0.1.0 // indirect
	github.com/containerd/fifo v1.1.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/containerd/ttrpc v1.2.3 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v26.1.0+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v26.1.0+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.1 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/ePirat/docker-credential-gitlabci v1.0.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/go-uuid v1.0.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/moby/buildkit v0.13.1 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/swarmkit/v2 v2.0.0-20240415162501-c1c857e2dca1 // indirect
	github.com/moby/sys/mount v0.3.3 // indirect
	github.com/moby/sys/mountinfo v0.7.1 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/sys/symlink v0.2.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/otiai10/copy v1.14.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/skeema/knownhosts v1.2.2 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	go.etcd.io/etcd/raft/v3 v3.5.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/oauth2 v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
)
