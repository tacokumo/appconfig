package apispec

import "github.com/go-playground/validator/v10"

type AppConfig struct {
	AppName  string          `json:"app_name" yaml:"app_name" validate:"required"`
	Build    BuildConfig     `json:"build" yaml:"build" validate:"required"`
	Releases []ReleaseConfig `json:"releases" yaml:"releases" validate:"required"`
	Service ServiceConfig `json:"service" yaml:"service" validate:"required"`
}

func (c *AppConfig) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

type BuildConfig struct {
	Image string `json:"image" yaml:"image"`

	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
	// TODO: DockerContextが設定されているときは､Dockerfileが必要
	DockerContext string `json:"docker_context" yaml:"docker_context"`
	// TODO: BuildArgsが設定されているときは､Dockerfileが必要
	BuildArgs map[string]string `json:"build_args" yaml:"build_args"`
}

type ReleaseConfig struct {
	Name      string              `json:"name" yaml:"name" validate:"required"`
	Resources ResourceConfig      `json:"resources" yaml:"resources" validate:"required"`
	Action    ReleaseActionConfig `json:"action" yaml:"action" validate:"required"`
}

type ResourceConfig struct {
	CPU    string `json:"cpu" yaml:"cpu" validate:"required"`
	Memory string `json:"memory" yaml:"memory" validate:"required"`
}

type ReleaseActionConfig struct {
	Command []string `json:"command" yaml:"command" validate:"required,min=1,required"`
}

type ServiceConfig struct {
	Name          string              `json:"name" yaml:"name" validate:"required"`
	Command       []string            `json:"command" yaml:"command" validate:"required,min=1,required"`
	HTTP          *ServiceHTTPConfig  `json:"http,omitempty" yaml:"http,omitempty"`
	Healthcheck   *HealthcheckConfig  `json:"healthcheck,omitempty" yaml:"healthcheck,omitempty"`
	Scale         *ServiceScaleConfig `json:"scale,omitempty" yaml:"scale,omitempty"`
	MachineConfig *MachineConfig      `json:"machine_config,omitempty" yaml:"machine_config,omitempty"`
}

type ServiceHTTPConfig struct {
	TargetPort int  `json:"target_port" yaml:"target_port" validate:"required,min=1"`
	ForceHTTPS bool `json:"force_https,omitempty" yaml:"force_https,omitempty"`
}

type HealthcheckConfig struct {
	// httpかprocessのいずれかが必須
	HTTP    *HealthcheckHTTPConfig    `json:"http,omitempty" yaml:"http,omitempty" validate:"required_without=Process"`
	Process *HealthcheckProcessConfig `json:"process,omitempty" yaml:"process,omitempty" validate:"required_without=HTTP"`
}

type HealthcheckHTTPConfig struct {
	Path string `json:"path" yaml:"path" validate:"required"`
}

type HealthcheckProcessConfig struct {
	Command []string `json:"command" yaml:"command" validate:"required,min=1,required"`
}

type ServiceScaleConfig struct {
	Min    int                 `json:"min" yaml:"min" validate:"required,min=0"`
	Max    int                 `json:"max" yaml:"max" validate:"required,min=1"`
	Metric ServiceMetricConfig `json:"metric,omitempty" yaml:"metric,omitempty" validate:"required"`
}

type ServiceMetricConfig struct {
	Type      string `json:"type" yaml:"type" validate:"required"`
	Threshold int    `json:"threshold" yaml:"threshold" validate:"required,min=1"`
}

type MachineConfig struct {
	CPU    string `json:"cpu" yaml:"cpu" validate:"required"`
	Memory string `json:"memory" yaml:"memory" validate:"required"`
	Flavor string `json:"flavor,omitempty" yaml:"flavor,omitempty"`
}
