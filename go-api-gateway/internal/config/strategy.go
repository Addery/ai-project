package config

const (
	StrategyRoundRobin         = "round_robin"          // 轮询
	StrategyWeightedRoundRobin = "weighted_round_robin" // 平滑加权轮询
	/*
			StrategyLeastConnections 最少连接（Least Connections）是一种动态负载均衡策略，其核心思想是：将新请求分配给当前活跃连接数
		最少的后端实例。这种策略特别适用于请求处理时间差异较大的场景（例如有的请求只需 10ms，有的需要 2s），能有效
		避免慢请求“拖垮”某个节点。
	*/
	StrategyLeastConnections  = "least_connections"  // 最少连接数
	StrategyRandom            = "random"             // 随机
	StrategyIPHash            = "ip_hash"            // 基于客户端 IP 的哈希
	StrategyConsistentHashing = "consistent_hashing" // 一致性哈希
)

var SupportedStrategies = map[string]bool{
	StrategyRoundRobin:         true,
	StrategyWeightedRoundRobin: true,
	StrategyLeastConnections:   true,
}
