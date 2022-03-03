resource "azurerm_virtual_network" "example_vn_azure" {
  resource_group_name = azurerm_resource_group.vn-backend-rg.name
  name                = "example_vn"
  location            = "ukwest"
  address_space       = ["10.0.0.0/16"]
}
resource "azurerm_route_table" "example_vn_azure" {
  resource_group_name = azurerm_resource_group.vn-backend-rg.name
  name                = "example_vn"
  location            = "ukwest"

  route {
    name           = "local"
    address_prefix = "0.0.0.0/0"
    next_hop_type  = "VnetLocal"
  }
}
resource "azurerm_resource_group" "vn-backend-rg" {
  name     = "vn-backend-rg"
  location = "ukwest"
}
provider "azurerm" {
  features {}
}
