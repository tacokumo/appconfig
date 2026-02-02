package apispec

import "github.com/go-playground/validator/v10"

type AppConfig struct {
	// AppName はアプリケーションの名前
	AppName string `json:"app_name" yaml:"app_name" validate:"required"`
	// Build はアプリケーションのビルド設定
	Build BuildConfig `json:"build" yaml:"build" validate:"required"`
	// Releases はアプリケーションのリリース設定
	Releases []ReleaseConfig `json:"releases" yaml:"releases" validate:"required"`
	// Service はアプリケーションのサービス設定
	Service ServiceConfig `json:"service" yaml:"service" validate:"required"`
	// Stages はアプリケーションのステージ設定
	// 何も定義されていない場合は、デフォルトで `production` ステージが作成される
	Stages []StageConfig `json:"stages,omitempty" yaml:"stages,omitempty"`
}

// Validate は AppConfig のバリデーションを行う
func (c *AppConfig) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

// BuildConfig はアプリケーションのビルド設定を表す
type BuildConfig struct {
	// Image はイメージビルドを行わず、既存のイメージを使用する場合に指定する
	Image string `json:"image" yaml:"image"`
	// Dockerfile はDockerイメージをビルドするためのDockerfileのパス
	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
	// DockerContextはDockerイメージをビルドするためのコンテキストのパス
	DockerContext string `json:"docker_context" yaml:"docker_context"`
	// BuildArgs はDockerビルド時に使用するビルド引数
	BuildArgs map[string]string `json:"build_args" yaml:"build_args"`
}

// ReleaseConfig はアプリケーションのリリース設定を表す
type ReleaseConfig struct {
	// Name はリリースの名前
	Name string `json:"name" yaml:"name" validate:"required"`
	// Resources はリリースのリソース設定
	Resources ResourceConfig `json:"resources" yaml:"resources" validate:"required"`
	// Action はリリースのアクション設定
	Action ReleaseActionConfig `json:"action" yaml:"action" validate:"required"`
}

// ResourceConfig はリソース設定を表す
type ResourceConfig struct {
	// CPU はCPUリソースの量
	CPU string `json:"cpu" yaml:"cpu" validate:"required"`
	// Memory はメモリリソースの量
	Memory string `json:"memory" yaml:"memory" validate:"required"`
}

// ReleaseActionConfig はリリースアクションの設定を表す
type ReleaseActionConfig struct {
	// Command はリリース時に実行されるコマンド
	Command []string `json:"command" yaml:"command" validate:"required,min=1,required"`
}

// ServiceConfig はアプリケーションのサービス設定を表す
type ServiceConfig struct {
	// Name はサービスの名前
	Name string `json:"name" yaml:"name" validate:"required"`
	// Command はサービスの起動コマンド
	Command []string `json:"command" yaml:"command" validate:"required,min=1,required"`
	// HTTP はサービスのHTTP設定
	HTTP []ServiceHTTPConfig `json:"http,omitempty" yaml:"http,omitempty"`
	// Healthcheck はサービスのヘルスチェック設定
	Healthcheck *HealthcheckConfig `json:"healthcheck,omitempty" yaml:"healthcheck,omitempty"`
	// Scale はサービスのスケーリング設定
	Scale *ServiceScaleConfig `json:"scale,omitempty" yaml:"scale,omitempty"`
	// MachineConfig はサービスのマシン設定
	MachineConfig *MachineConfig `json:"machine_config,omitempty" yaml:"machine_config,omitempty"`
}

// ServiceHTTPConfig はサービスのHTTP設定を表す
type ServiceHTTPConfig struct {
	// TargetPort はサービスがリッスンするポート
	TargetPort int `json:"target_port" yaml:"target_port" validate:"required,min=1"`
	// ForceHTTPS はHTTPリクエストをHTTPSにリダイレクトするかどうか
	ForceHTTPS bool `json:"force_https,omitempty" yaml:"force_https,omitempty"`
}

// HealthcheckConfig はサービスのヘルスチェック設定を表す
type HealthcheckConfig struct {
	// HTTP はHTTPヘルスチェックの設定
	HTTP *HealthcheckHTTPConfig `json:"http,omitempty" yaml:"http,omitempty" validate:"required_without=Process"`
	// Process はプロセスヘルスチェックの設定
	Process *HealthcheckProcessConfig `json:"process,omitempty" yaml:"process,omitempty" validate:"required_without=HTTP"`
}

// HealthcheckHTTPConfig はHTTPヘルスチェックの設定を表す
type HealthcheckHTTPConfig struct {
	// Path はヘルスチェックのエンドポイントパス
	Path string `json:"path" yaml:"path" validate:"required"`
}

// HealthcheckProcessConfig はプロセスヘルスチェックの設定を表す
type HealthcheckProcessConfig struct {
	// Command はヘルスチェックに使用するコマンド
	Command []string `json:"command" yaml:"command" validate:"required,min=1,required"`
}

// ServiceScaleConfig はサービスのスケーリング設定を表す
type ServiceScaleConfig struct {
	// Min はサービスの最小インスタンス数
	Min int `json:"min" yaml:"min" validate:"required,min=0"`
	// Max はサービスの最大インスタンス数
	Max int `json:"max" yaml:"max" validate:"required,min=1"`
	// Metric はスケーリングに使用するメトリクスの設定
	Metric ServiceMetricConfig `json:"metric" yaml:"metric" validate:"required"`
}

// ServiceMetricConfig はサービスのスケーリングメトリクスの設定を表す
type ServiceMetricConfig struct {
	// Type はメトリクスの種類
	Type string `json:"type" yaml:"type" validate:"required"`
	// Threshold はスケーリングの閾値
	Threshold int `json:"threshold" yaml:"threshold" validate:"required,min=1"`
}

// MachineConfig はサービスのマシン設定を表す
type MachineConfig struct {
	// CPU はマシンのCPUリソースの量
	CPU string `json:"cpu" yaml:"cpu" validate:"required"`
	// Memory はマシンのメモリリソースの量
	Memory string `json:"memory" yaml:"memory" validate:"required"`
	// Flavor はマシンのフレーバー
	Flavor string `json:"flavor,omitempty" yaml:"flavor,omitempty"`
}

// StageConfig はアプリケーションのステージ設定を表す
type StageConfig struct {
	// Name はステージの名前
	Name string `json:"name" yaml:"name" validate:"required"`
	// Policy はステージのポリシー設定
	Policy StagePolicyConfig `json:"policy" yaml:"policy" validate:"required"`
}

// StagePolicyConfig はステージのポリシー設定を表す
type StagePolicyConfig struct {
	// Type はポリシーの種類
	// 現在は `branch` のみサポート
	Type string `json:"type" yaml:"type" validate:"required"`
	// Branch はブランチポリシーの設定
	// Type が `branch` の場合に必須
	Branch *BranchConfig `json:"branch,omitempty" yaml:"branch,omitempty" validate:"required_if=Type branch"`
}

// BranchConfig はブランチポリシーの設定を表す
type BranchConfig struct {
	// Name は対象のブランチ名
	Name string `json:"name" yaml:"name" validate:"required"`
}
