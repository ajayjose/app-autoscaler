package config_test

import (
	. "autoscaler/eventgenerator/config"

	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	var (
		conf        *Config
		err         error
		configBytes []byte
	)

	Describe("LoadConfig", func() {

		JustBeforeEach(func() {
			conf, err = LoadConfig(configBytes)
		})

		Context("with valid yaml", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"

`)
			})

			It("returns the config", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(conf).To(Equal(&Config{Server: ServerConfig{Port: 8081},
					Logging: LoggingConfig{Level: "info"},
					DB:      DBConfig{PolicyDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable", AppMetricDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"},
					Aggregator: AggregatorConfig{
						AggregatorExecuteInterval: 30 * time.Second,
						PolicyPollerInterval:      30 * time.Second,
						MetricPollerCount:         10,
						AppMonitorChannelSize:     100},
					Evaluator: EvaluatorConfig{
						EvaluationManagerInterval: 30 * time.Second,
						EvaluatorCount:            10,
						TriggerArrayChannelSize:   100},
					ScalingEngine: ScalingEngineConfig{
						ScalingEngineUrl: "http://localhost:8082"},
					MetricCollector: MetricCollectorConfig{
						MetricCollectorUrl: "http://localhost:8083"},
				}))
			})
		})
		Context("with invalid yaml", func() {
			BeforeEach(func() {
				configBytes = []byte(`
  server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
		`)
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(MatchRegexp(".*did not find expected <document start>.*")))
			})
		})

		Context("when it gives a non integer server port", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: "NOT-INTEGER-VALUE"
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal.*into int")))
			})
		})

		Context("when it gives a non integer aggregator_execute_interval", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: "NOT-INTEGER-VALUE"
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal .* into time.Duration")))
			})
		})

		Context("when it gives a non integer policy_poller_interval", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: "NOT-INTEGER-VALUE"
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal .* into time.Duration")))
			})
		})

		Context("when it gives a non integer metric_poller_count", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: "NOT-INTEGER-VALUE"
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal.*into int")))
			})
		})

		Context("when it gives a non integer app_monitor_channel_size", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: "NOT-INTEGER-VALUE"
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal.*into int")))
			})
		})

		Context("when it gives a non integer evaluation_manager_execute_interval", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: "NOT-INTEGER-VALUE"
  evaluator_count: 10
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal .* into time.Duration")))
			})
		})

		Context("when it gives a non integer evaluator_count", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: "NOT-INTEGER-VALUE"
  trigger_array_channel_size: 100
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal.*into int")))
			})
		})
		Context("when it gives a non integer trigger_array_channel_size", func() {
			BeforeEach(func() {
				configBytes = []byte(`
server:
  port: 8081
logging:
  level: "info"
db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
aggregator: 
  aggregator_execute_interval: 30s
  policy_poller_interval: 30s
  metric_poller_count: 10
  app_monitor_channel_size: 100
evaluator:
  evaluation_manager_execute_interval: 30s
  evaluator_count: 10
  trigger_array_channel_size: "NOT-INTEGER-VALUE"
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("cannot unmarshal.*into int")))
			})
		})

		Context("with partial config", func() {
			BeforeEach(func() {
				configBytes = []byte(`

db:
  policy_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
  app_metrics_db_url: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"
scalingEngine:
  scaling_engine_url: "http://localhost:8082"
metricCollector:
  metric_collector_url: "http://localhost:8083"
`)
			})

			It("returns default values", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(conf.Aggregator.PolicyPollerInterval).To(Equal(time.Duration(DefaultPolicyPollerInterval)))

				Expect(err).NotTo(HaveOccurred())
				Expect(conf).To(Equal(&Config{Server: ServerConfig{Port: 8080},
					Logging: LoggingConfig{Level: "info"},
					DB:      DBConfig{PolicyDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable", AppMetricDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"},
					Aggregator: AggregatorConfig{AggregatorExecuteInterval: DefaultAggregatorExecuteInterval,
						PolicyPollerInterval:  DefaultPolicyPollerInterval,
						MetricPollerCount:     DefaultMetricPollerCount,
						AppMonitorChannelSize: DefaultAppMonitorChannelSize},
					Evaluator: EvaluatorConfig{EvaluationManagerInterval: DefaultEvaluationExecuteInterval,
						EvaluatorCount:          DefaultEvaluatorCount,
						TriggerArrayChannelSize: DefaultTriggerArrayChannelSize},
					ScalingEngine: ScalingEngineConfig{
						ScalingEngineUrl: "http://localhost:8082"},
					MetricCollector: MetricCollectorConfig{
						MetricCollectorUrl: "http://localhost:8083"}}))
			})
		})

	})

	Describe("Validate", func() {
		BeforeEach(func() {
			conf = &Config{Server: ServerConfig{Port: 8081},
				Logging: LoggingConfig{Level: "info"},
				DB:      DBConfig{PolicyDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable", AppMetricDBUrl: "postgres://postgres:password@localhost/autoscaler?sslmode=disable"},
				Aggregator: AggregatorConfig{
					AggregatorExecuteInterval: 30 * time.Second,
					PolicyPollerInterval:      30 * time.Second,
					MetricPollerCount:         10,
					AppMonitorChannelSize:     100},
				Evaluator: EvaluatorConfig{
					EvaluationManagerInterval: 30 * time.Second,
					EvaluatorCount:            10,
					TriggerArrayChannelSize:   100},
				ScalingEngine: ScalingEngineConfig{
					ScalingEngineUrl: "http://localhost:8082"},
				MetricCollector: MetricCollectorConfig{
					MetricCollectorUrl: "http://localhost:8083"}}
		})

		JustBeforeEach(func() {
			err = conf.Validate()
		})

		Context("when server port <= 0", func() {
			BeforeEach(func() {
				conf.Server.Port = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: server port is less-equal than 0 or more than 65535")))
			})
		})

		Context("when server port > 65535", func() {
			BeforeEach(func() {
				conf.Server.Port = 65536
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: server port is less-equal than 0 or more than 65535")))
			})
		})

		Context("when policy db url is not set", func() {

			BeforeEach(func() {
				conf.DB.PolicyDBUrl = ""
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: Policy DB url is empty")))
			})
		})

		Context("when appmetric db url is not set", func() {

			BeforeEach(func() {
				conf.DB.AppMetricDBUrl = ""
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: AppMetric DB url is empty")))
			})
		})
		Context("when scaling engine url is not set", func() {

			BeforeEach(func() {
				conf.ScalingEngine.ScalingEngineUrl = ""
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: Scaling engine url is empty")))
			})
		})
		Context("when metric collector url is not set", func() {

			BeforeEach(func() {
				conf.MetricCollector.MetricCollectorUrl = ""
			})

			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: Metric collector url is empty")))
			})
		})

		Context("when AggregatorExecuateInterval <= 0", func() {
			BeforeEach(func() {
				conf.Aggregator.AggregatorExecuteInterval = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: aggregator execute interval is less-equal than 0")))
			})
		})

		Context("when PolicyPollerInterval:  <= 0", func() {
			BeforeEach(func() {
				conf.Aggregator.PolicyPollerInterval = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: policy poller interval is less-equal than 0")))
			})
		})

		Context("when MetricPollerCount <= 0", func() {
			BeforeEach(func() {
				conf.Aggregator.MetricPollerCount = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: metric poller count is less-equal than 0")))
			})
		})

		Context("when AppMonitorChannelSize <= 0", func() {
			BeforeEach(func() {
				conf.Aggregator.AppMonitorChannelSize = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: appMonitor channel size is less-equal than 0")))
			})
		})

		Context("when EvaluationManagerInterval <= 0", func() {
			BeforeEach(func() {
				conf.Evaluator.EvaluationManagerInterval = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: evalution manager execeute interval is less-equal than 0")))
			})
		})

		Context("when EvaluatorCount <= 0", func() {
			BeforeEach(func() {
				conf.Evaluator.EvaluatorCount = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: evaluator count is less-equal than 0")))
			})
		})

		Context("when TriggerArrayChannelSize <= 0", func() {
			BeforeEach(func() {
				conf.Evaluator.TriggerArrayChannelSize = 0
			})
			It("should error", func() {
				Expect(err).To(MatchError(MatchRegexp("Configuration error: trigger-array channel size is less-equal than 0")))
			})
		})

	})
})
