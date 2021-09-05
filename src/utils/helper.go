package utils

import (
	"context"
	"fmt"

	proxyman "v2ray.com/core/app/proxyman/command"
	stats "v2ray.com/core/app/stats/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
)

type VmessUser struct {
	Level        uint32                `json:"level"`
	Id           string                `json:"id"`
	Email        string                `json:"email"`
	AlterId      uint32                `json:"alterId"`
	SecurityType protocol.SecurityType `json:"securityType"`
}

type rpc_helper struct {
	handler_service proxyman.HandlerServiceClient
	stats_service   stats.StatsServiceClient
}

func NewRPCHelper(target string) *rpc_helper {
	helper := new(rpc_helper)
	handler_service := GetProxymanClient(target)
	stats_service := GetStatsClient(target)
	helper.handler_service = *handler_service
	helper.stats_service = *stats_service
	return helper
}

func (helper rpc_helper) QueryStats(pattern string, reset bool) (stats.QueryStatsResponse, error) {
	resp, err := helper.stats_service.QueryStats(context.Background(), &stats.QueryStatsRequest{
		Pattern: pattern,
		Reset_:  reset,
	})
	return *resp, err
}

func (helper rpc_helper) GetUserTraffic(email string, traffic_type string, reset bool) (*stats.GetStatsResponse, error) {
	name := fmt.Sprintf("user>>>%s>>>traffic>>>%s", email, traffic_type)
	return helper.getStat(name, reset)
}

func (helper rpc_helper) getStat(name string, reset bool) (*stats.GetStatsResponse, error) {
	resp, err := helper.stats_service.GetStats(context.Background(), &stats.GetStatsRequest{
		Name:   name,
		Reset_: reset,
	})
	return resp, err
}

func (helper rpc_helper) GetInboundTraffic(tag string, traffic_type string, reset bool) (*stats.GetStatsResponse, error) {
	name := fmt.Sprintf("inbound>>>%s>>>traffic>>>%s", tag, traffic_type)
	return helper.getStat(name, reset)
}

func (helper rpc_helper) GetInboundTrafficAndReset(tag string, traffic_type string) (*stats.GetStatsResponse, error) {
	name := fmt.Sprintf("inbound>>>%s>>>traffic>>>%s", tag, traffic_type)
	return helper.getStat(name, true)
}

func (helper rpc_helper) AddVmessUser(inboundTag string, user VmessUser) error {
	_, err := helper.handler_service.AlterInbound(context.Background(), &proxyman.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&proxyman.AddUserOperation{
			User: &protocol.User{
				Level: user.Level,
				Email: user.Email,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:      user.Id,
					AlterId: user.AlterId,
					SecuritySettings: &protocol.SecurityConfig{
						Type: user.SecurityType,
					},
				}),
			},
		}),
	})
	return err
}

func (helper rpc_helper) RemoveUser(inboundTag string, email string) string {
	_, err := helper.handler_service.AlterInbound(context.Background(), &proxyman.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&proxyman.RemoveUserOperation{
			Email: email,
		}),
	})
	success := err == nil
	if success {
		return "ok"
	} else {
		return err.Error()
	}
}
