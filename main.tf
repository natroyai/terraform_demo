terraform {
  required_providers {
    ciscoios = {
      source  = "terraform.local/local/ciscoios"
      version = "1.0.0"
    }
  }
}

// VS tira error unexpected attribute pero esta funciona bien
provider "ciscoios" {
  host     = "100.100.1.141"
  username = "natroy2"
  password = "natroy2"
  port     = 22
}

resource "ciscoios_ssh_command" "example" {
  commands = [
    "enable",
    "conf t",
    "hostname TERRAFORM",
    "exit",
    "wr"
  ]
}

output "command_result" {
  value = ciscoios_ssh_command.example.result
}
