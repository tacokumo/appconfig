package apispec

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAppConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  AppConfig
		wantErr bool
	}{
		{
			name: "必須フィールドがすべて設定されている場合、エラーにならない",
			config: AppConfig{
				AppName: "myapp",
				Build: BuildConfig{
					Image: "myapp:latest",
				},
				Releases: []ReleaseConfig{
					{
						Name: "release-v1",
						Resources: ResourceConfig{
							CPU:    "500m",
							Memory: "256Mi",
						},
						Action: ReleaseActionConfig{
							Command: []string{"echo", "deploy"},
						},
					},
				},
				Service: ServiceConfig{
					Name:    "web",
					Command: []string{"npm", "start"},
				},
			},
			wantErr: false,
		},
		{
			name: "Serviceが設定されていない場合、エラーになる",
			config: AppConfig{
				AppName: "myapp",
				Build: BuildConfig{
					Image: "myapp:latest",
				},
				Releases: []ReleaseConfig{
					{
						Name: "release-v1",
						Resources: ResourceConfig{
							CPU:    "500m",
							Memory: "256Mi",
						},
						Action: ReleaseActionConfig{
							Command: []string{"echo", "deploy"},
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("AppConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  BuildConfig
		wantErr bool
	}{
		{
			name: "Imageが設定されている場合、エラーにならない",
			config: BuildConfig{
				Image: "myapp:latest",
			},
			wantErr: false,
		},
		{
			name: "BuildArgsとDockerfileがともに設定されている場合､エラーにならない",
			config: BuildConfig{
				Dockerfile: "Dockerfile",
				BuildArgs: map[string]string{
					"ARG1": "value1",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReleaseConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ReleaseConfig
		wantErr bool
	}{
		{
			name: "NameとAction.Commandが設定されている場合、エラーにならない",
			config: ReleaseConfig{
				Name: "release-v1",
				Action: ReleaseActionConfig{
					Command: []string{"echo", "deploy"},
				},
				Resources: ResourceConfig{
					CPU:    "500m",
					Memory: "256Mi",
				},
			},
			wantErr: false,
		},
		{
			name: "Nameが設定されていない場合、エラーになる",
			config: ReleaseConfig{
				Action: ReleaseActionConfig{
					Command: []string{"echo", "deploy"},
				},
				Resources: ResourceConfig{
					CPU:    "500m",
					Memory: "256Mi",
				},
			},
			wantErr: true,
		},
		{
			name: "Action.Commandが設定されていない場合、エラーになる",
			config: ReleaseConfig{
				Name: "release-v1",
				Action: ReleaseActionConfig{
					Command: []string{},
				},
				Resources: ResourceConfig{
					CPU:    "500m",
					Memory: "256Mi",
				},
			},
			wantErr: true,
		},
		{
			name: "Resourcesが設定されていない場合、エラーになる",
			config: ReleaseConfig{
				Name: "release-v1",
				Action: ReleaseActionConfig{
					Command: []string{"echo", "deploy"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReleaseConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestResourceConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ResourceConfig
		wantErr bool
	}{
		{
			name: "CPUとMemoryが設定されている場合、エラーにならない",
			config: ResourceConfig{
				CPU:    "500m",
				Memory: "256Mi",
			},
			wantErr: false,
		},
		{
			name: "CPUが設定されていない場合、エラーになる",
			config: ResourceConfig{
				Memory: "256Mi",
			},
			wantErr: true,
		},
		{
			name: "Memoryが設定されていない場合、エラーになる",
			config: ResourceConfig{
				CPU: "500m",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReleaseActionConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ReleaseActionConfig
		wantErr bool
	}{
		{
			name: "Commandが設定されている場合、エラーにならない",
			config: ReleaseActionConfig{
				Command: []string{"echo", "deploy"},
			},
			wantErr: false,
		},
		{
			name:    "Commandが設定されていない場合、エラーになる",
			config:  ReleaseActionConfig{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReleaseActionConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ServiceConfig
		wantErr bool
	}{
		{
			name: "NameとCommandが設定されている場合、エラーにならない",
			config: ServiceConfig{
				Name:    "web",
				Command: []string{"npm", "start"},
			},
			wantErr: false,
		},
		{
			name: "Nameが設定されていない場合、エラーになる",
			config: ServiceConfig{
				Command: []string{"npm", "start"},
			},
			wantErr: true,
		},
		{
			name: "Commandが設定されていない場合、エラーになる",
			config: ServiceConfig{
				Name: "web",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceHTTPConfig_Validate(t *testing.T) {
	t.Parallel()
	{
		tests := []struct {
			name    string
			config  ServiceHTTPConfig
			wantErr bool
		}{
			{
				name: "TargetPortが設定されている場合、エラーにならない",
				config: ServiceHTTPConfig{
					TargetPort: 80,
				},
				wantErr: false,
			},
			{
				name:    "TargetPortが設定されていない場合、エラーになる",
				config:  ServiceHTTPConfig{},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				err := validator.New().Struct(tt.config)
				if (err != nil) != tt.wantErr {
					t.Errorf("ServiceHTTPConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestHealthcheckConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  HealthcheckConfig
		wantErr bool
	}{
		{
			name: "HTTPが設定されている場合、エラーにならない",
			config: HealthcheckConfig{
				HTTP: &HealthcheckHTTPConfig{
					Path: "/",
				},
			},
			wantErr: false,
		},
		{
			name: "Processが設定されている場合、エラーにならない",
			config: HealthcheckConfig{
				Process: &HealthcheckProcessConfig{
					Command: []string{"echo", "healthcheck"},
				},
			},
			wantErr: false,
		},
		{
			name:    "HTTPもProcessも設定されていない場合、エラーになる",
			config:  HealthcheckConfig{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthcheckConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHealthcheckHTTPConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  HealthcheckHTTPConfig
		wantErr bool
	}{
		{
			name: "Pathが設定されている場合、エラーにならない",
			config: HealthcheckHTTPConfig{
				Path: "/",
			},
			wantErr: false,
		},
		{
			name:    "Pathが設定されていない場合、エラーになる",
			config:  HealthcheckHTTPConfig{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthcheckHTTPConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHealthcheckProcessConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  HealthcheckProcessConfig
		wantErr bool
	}{
		{
			name: "Commandが設定されている場合、エラーにならない",
			config: HealthcheckProcessConfig{
				Command: []string{"echo", "hello"},
			},
			wantErr: false,
		},
		{
			name:    "Commandが設定されていない場合、エラーになる",
			config:  HealthcheckProcessConfig{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthcheckProcessConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceScaleConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ServiceScaleConfig
		wantErr bool
	}{
		{
			name: "MinとMaxが設定されている場合、エラーにならない",
			config: ServiceScaleConfig{
				Min: 1,
				Max: 5,
				Metric: ServiceMetricConfig{
					Type:      "cpu",
					Threshold: 80,
				},
			},
			wantErr: false,
		},
		{
			name: "Minが設定されていない場合、エラーになる",
			config: ServiceScaleConfig{
				Max: 5,
				Metric: ServiceMetricConfig{
					Type:      "cpu",
					Threshold: 80,
				},
			},
			wantErr: true,
		},
		{
			name: "Maxが設定されていない場合、エラーになる",
			config: ServiceScaleConfig{
				Min: 1,
				Metric: ServiceMetricConfig{
					Type:      "cpu",
					Threshold: 80,
				},
			},
			wantErr: true,
		},
		{
			name: "Metricが設定されていない場合、エラーになる",
			config: ServiceScaleConfig{
				Min: 1,
				Max: 5,
			},
			wantErr: true,
		},
		{
			name: "Minが0未満の場合、エラーになる",
			config: ServiceScaleConfig{
				Min: -1,
				Max: 5,
				Metric: ServiceMetricConfig{
					Type:      "cpu",
					Threshold: 80,
				},
			},
			wantErr: true,
		},
		{
			name: "Maxが1未満の場合、エラーになる",
			config: ServiceScaleConfig{
				Min: 1,
				Max: 0,
				Metric: ServiceMetricConfig{
					Type:      "cpu",
					Threshold: 80,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceScaleConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceMetricConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  ServiceMetricConfig
		wantErr bool
	}{
		{
			name: "TypeとThresholdが設定されている場合、エラーにならない",
			config: ServiceMetricConfig{
				Type:      "cpu",
				Threshold: 80,
			},
			wantErr: false,
		},
		{
			name: "Typeが設定されていない場合、エラーになる",
			config: ServiceMetricConfig{
				Threshold: 80,
			},
			wantErr: true,
		},
		{
			name: "Thresholdが設定されていない場合、エラーになる",
			config: ServiceMetricConfig{
				Type: "cpu",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceMetricConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMachineConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		config  MachineConfig
		wantErr bool
	}{
		{
			name: "CPUとMemoryが設定されている場合、エラーにならない",
			config: MachineConfig{
				CPU:    "500m",
				Memory: "256Mi",
			},
			wantErr: false,
		},
		{
			name: "CPUが設定されていない場合、エラーになる",
			config: MachineConfig{
				Memory: "256Mi",
			},
			wantErr: true,
		},
		{
			name: "Memoryが設定されていない場合、エラーになる",
			config: MachineConfig{
				CPU: "500m",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.New().Struct(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("MachineConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
