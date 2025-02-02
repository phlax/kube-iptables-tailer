package util

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
)

// required env vars
const (
	IptablesLogPrefix = "IPTABLES_LOG_PREFIX"
	IptablesLogPath   = "IPTABLES_LOG_PATH"
	JournalDirectory  = "JOURNAL_DIRECTORY"
)

// optional env vars
const (
	KubeApiServer = "KUBE_API_SERVER" // default value is empty string

	KubeEventDisplayReason        = "KUBE_EVENT_DISPLAY_REASON"
	DefaultKubeEventDisplayReason = "PacketDrop"

	KubeEventSourceComponentName        = "KUBE_EVENT_SOURCE_COMPONENT_NAME"
	DefaultKubeEventSourceComponentName = "kube-iptables-tailer"

	MetricsServerPort        = "METRICS_SERVER_PORT"
	DefaultMetricsServerPort = 9090

	LogLevel        = "LOG_LEVEL"
	DefaultLogLevel = "INFO"

	PacketDropChannelBufferSize         = "PACKET_DROP_CHANNEL_BUFFER_SIZE"
	DefaultPacketDropsChannelBufferSize = 100

	PacketDropLogTimeLayout        = "PACKET_DROP_LOG_TIME_LAYOUT"
	DefaultPacketDropLogTimeLayout = "Jan _2 15:04:05"
        // 2006-01-02T15:04:05.000000-07:00"

	PacketDropExpirationMinutes        = "PACKET_DROP_EXPIRATION_MINUTES"
	DefaultPacketDropExpirationMinutes = 10

	RepeatedEventIntervalMinutes        = "REPEATED_EVENTS_INTERVAL_MINUTES"
	DefaultRepeatedEventIntervalMinutes = 2

	WatchLogsIntervalSeconds       = "WATCH_LOGS_INTERVAL_SECONDS"
	DefaultWatchLogsIntervalSecond = 5

	PodIdentifier        = "POD_IDENTIFIER"
	DefaultPodIdentifier = "namespace"
	PodIdentifierLabel   = "POD_IDENTIFIER_LABEL"
)

func GetRequiredEnvString(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		zap.L().Fatal(fmt.Sprintf("Missing environment variable %v", key))
	}
	return val
}

func GetRequiredEnvInt(key string) int {
	stringVal := GetRequiredEnvString(key)
	intVal, err := strconv.Atoi(stringVal)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("Error converting environment variable %s to int: %v", stringVal, err))
	}
	return intVal
}

func GetEnvIntOrDefault(key string, def int) int {
	if env := os.Getenv(key); env != "" {
		val, err := strconv.Atoi(env)
		if err != nil {
			zap.L().Warn(fmt.Sprintf("Invalid value for %v: using default: %v", key, def))
			return def
		}
		return val
	}
	return def
}

func GetEnvStringOrDefault(key string, def string) string {
	if val := os.Getenv(key); len(val) > 0 {
		return val
	}
	return def
}
