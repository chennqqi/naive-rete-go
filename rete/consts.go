package rete

const (
	CLASS_NAME = iota
	IDENTIFIER
	ATTRIBUTE
	VALUE
	NO_TEST
)

const (
	BETA_MEMORY_NODE = "beta_memory"
	JOIN_NODE        = "join_node"
	NEGATIVE_NODE    = "negative_node"
	NCC_NODE         = "ncc_node"
	NCC_PARTNER_NODE = "ncc_parter_node"
	FILTER_NODE      = "filter_node"
)

var FIELDS = []int{CLASS_NAME, IDENTIFIER, ATTRIBUTE, VALUE}
