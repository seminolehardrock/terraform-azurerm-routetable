variable "create_resource_group" {
  type        = bool
  description = "Determine if a new resource group will be created or use an existing one."
  default     = false
}

variable "resource_group_name" {
  type        = string
  description = "Default resource group name that the network will be created in."
  default     = "myapp-rg"
}

variable "location" {
  type        = string
  description = "The location/region where the core network will be created. The full list of Azure regions can be found at https://azure.microsoft.com/regions"
  default     = "South Central US"
}

variable "route_table_name" {
  type        = string
  description = "The name of the RouteTable being created."
  default     = "routetable"
}

variable "disable_bgp_route_propagation" {
  type        = string
  description = "Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable."
  default     = "true"
}

variable "route_prefixes" {
  type        = list(string)
  description = "The list of address prefixes to use for each route."
  default     = ["10.0.1.0/24"]
}

variable "route_names" {
  type        = list(string)
  description = "A list of public subnets inside the vNet."
  default     = ["subnet1"]
}

variable "route_nexthop_types" {
  type        = list(string)
  description = "The type of Azure hop the packet should be sent to for each corresponding route.Valid values are 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'HyperNetGateway', 'VirtualAppliance', 'None'"
  default     = ["VnetLocal"]
}

variable "route_next_hop_in_ip_addresses" {
  type        = list(string)
  description = "The nexthop ip address for each corresponding route. Valid for VirtualAppliance nexthop type."
  default     = []
}

variable "tags" {
  type        = map(string)
  description = "The tags to associate with your network and subnets."

  default = {
    tag1 = ""
    tag2 = ""
  }
}
