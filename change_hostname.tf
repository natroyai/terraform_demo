terraform {
  required_providers {
    iosxr = {
      source  = "CiscoDevNet/iosxr"
      version = "0.5.0"
    }
  }
}

provider "iosxr" {
  alias    = "ROUTER-1"
  username = "natroy2"
  password = "natroy2"
  host     = "100.100.1.141"
}

resource "iosxr_gnmi" "ROUTER-1" {
  provider = iosxr.ROUTER-1
  path     = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}
