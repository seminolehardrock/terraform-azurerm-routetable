resource "random_id" "rg_name" {
  byte_length = 8
}

module "routetable" {
  source                         = "../../"
  create_resource_group          = true
  resource_group_name            = random_id.rg_name.hex
  location                       = var.location
  route_prefixes                 = ["10.0.1.0/24"]
  route_nexthop_types            = ["VirtualAppliance"]
  route_names                    = ["route1"]
  route_next_hop_in_ip_addresses = ["192.168.1.0/26"]

  tags = {
    environment = "dev"
    costcenter  = "it"
  }
}
