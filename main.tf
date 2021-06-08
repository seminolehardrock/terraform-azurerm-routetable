resource "azurerm_resource_group" "routetablerg" {
  count = var.create_resource_group ? 1 : 0

  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_route_table" "rtable" {
  name                          = var.route_table_name
  location                      = var.location
  resource_group_name           = var.resource_group_name
  disable_bgp_route_propagation = var.disable_bgp_route_propagation

  depends_on = [
    azurerm_resource_group.routetablerg
  ]
}

resource "azurerm_route" "route" {
  name                = var.route_names[count.index]
  resource_group_name = var.resource_group_name
  route_table_name    = azurerm_route_table.rtable.name
  address_prefix      = var.route_prefixes[count.index]
  next_hop_type       = var.route_nexthop_types[count.index]
  count               = length(var.route_names)

  depends_on = [
    azurerm_route_table.rtable
  ]
}
