package utils

import (
	"log"

	"google.golang.org/grpc"
	"v2ray.com/core/app/proxyman/command"
	stats_command "v2ray.com/core/app/stats/command"
)

var handler_service_map map[string]*command.HandlerServiceClient = make(map[string]*command.HandlerServiceClient)
var stats_service_map map[string]*stats_command.StatsServiceClient = make(map[string]*stats_command.StatsServiceClient)

func initClient(target string) error {
	if handler_service_map[target] != nil && stats_service_map[target] != nil {
		return nil
	}
	cmdCon, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Println("Init client failure.")
		return err
	}
	if handler_service_map[target] == nil {
		handler_service := command.NewHandlerServiceClient(cmdCon)
		handler_service_map[target] = &handler_service
	}
	if stats_service_map[target] == nil {
		stats_service := stats_command.NewStatsServiceClient(cmdCon)
		stats_service_map[target] = &stats_service
	}
	return nil
}

func GetProxymanClient(target string) *command.HandlerServiceClient {
	err := initClient(target)
	if err != nil {
		return nil
	}
	return handler_service_map[target]
}

func GetStatsClient(target string) *stats_command.StatsServiceClient {
	err := initClient(target)
	if err != nil {
		return nil
	}
	return stats_service_map[target]
}
