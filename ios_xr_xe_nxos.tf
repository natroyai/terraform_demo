/*
terraform {
   required_version = ">= 1.9"
   // VsCode tira errores si no usas el Terraform v.1.9.5
  required_providers {
    iosxr = {
      source  = "CiscoDevNet/iosxr"
      version = "0.5.0"
    }
    iosxe = {
      source  = "CiscoDevNet/iosxe"
      version = "0.5.6"
    }
    nxos = {
      source  = "CiscoDevNet/nxos"
      version = "0.5.4"
    }
  }
}

/*
provider "iosxr" {
  alias    = "ROUTER-1-XR"
  username = "admin"
  password = "admin"
  host     = "100.100.1.141"
}

// It communicates with IOS-XE devices via the RESTCONF API,
// which requires the following device configuration.
// > ip http secure-server
// > restconf

provider "iosxe" {
  alias    = "ROUTER-1-XE"
  username = "admin"
  password = "password"
  url      = "https://10.1.1.1"
}

resource "iosxe_cli" "ROUTER-1-XE" {
  provider = iosxe.ROUTER-1-XE
  cli = <<-EOT
  interface Loopback123
  description configured-via-restconf-cli
  EOT
}

resource "iosxr_gnmi" "ROUTER-1" {
  provider = iosxr.ROUTER-1-XR
  path     = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}
*/

