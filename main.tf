terraform {
  required_providers {
    ciscoios = {
      source  = "terraform.local/local/ciscoios"
      version = "1.0.0"
    }
  }
}

// VS tira error unexpected attribute pero esta bien esto
provider "ciscoios" {
  host     = "100.100.1.141"
  username = "natroy2"
  password = "natroy2"
  port     = 22
}

resource "ciscoios_ssh_command" "example" {
  command = "show conf"
}

output "command_result" {
  value = ciscoios_ssh_command.example.result
}