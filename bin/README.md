Este es el binario de un plugin para hacer una conexion mediante SSH y mandar un comando.

Para usar un plugin local (en Linux por lo menos):
+ Pegar el binario en /home/USER/.terraform.d/plugins/terraform.local/local/ciscoios/1.0.0/linux_amd64. (Puede ser cualquier lado pero tiene que tener esa arquitectura las carpetas)
+ Crear un archivo en /root que se llama `.terraformrc` y pegarle el path a la carpeta plugins:

```
provider_installation {
  filesystem_mirror {
    path = "/home/USER/.terraform.d/plugins"
  }
}
```
+ Correr en terraform_demo/ : `rm -rf .terraform .terraform.lock.hcl`
+ Despues: `terraform init -reconfigure`
+ Listo. `terraform apply` y deberia andar.