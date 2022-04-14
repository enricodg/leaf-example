package pkgConfig

import leafConfig "github.com/paulusrobin/leaf-utilities/config"

type (
	ProfilerConfig struct {
		// - Toggle
		GoogleProfilerEnable    bool   `envconfig:"GOOGLE_PROFILER_ENABLE" default:"false"`
		GoogleProfilerProjectID string `envconfig:"GOOGLE_PROFILER_PROJECT_ID"`
	}
)

func NewProfilerConfig() (ProfilerConfig, error) {
	configuration := ProfilerConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return ProfilerConfig{}, err
	}
	return configuration, nil
}
