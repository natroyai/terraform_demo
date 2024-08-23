// PLUGIN para conectarse por SSH y mandar comandos

package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/scrapli/scrapligo/driver/generic"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/util"
)
type Config struct {
    Host     string
    Username string
    Password string
    Port     int
}


func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "host": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The Cisco IOS device host to connect to",
            },
            "username": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The username for authentication",
            },
            "password": {
                Type:        schema.TypeString,
                Required:    true,
                Sensitive:   true,
                Description: "The password for authentication",
            },
            "port": {
                Type:        schema.TypeInt,
                Optional:    true,
                Default:     22,
                Description: "The SSH port to connect to",
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "ciscoios_ssh_command": resourceCiscoSSH(),
        },
        ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
            return &Config{
                Host:     d.Get("host").(string),
                Username: d.Get("username").(string),
                Password: d.Get("password").(string),
                Port:     d.Get("port").(int),
            }, nil
        },
    }
}

func resourceCiscoSSH() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCiscoSSHCreate,
		ReadContext:   schema.NoopContext,
		UpdateContext: resourceCiscoSSHUpdate,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"commands": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"result": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The result of the command execution",
			},
		},
	}
}

func resourceCiscoSSHCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCiscoSSHUpdate(ctx, d, m)
}

func resourceCiscoSSHUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	
    config := m.(*Config)

    host := config.Host
    username := config.Username
    password := config.Password
    port := config.Port
    commands := d.Get("commands").([]interface{})

	extraKexs := []string{
		"diffie-hellman-group14-sha1",
		"diffie-hellman-group-exchange-sha256",
		"diffie-hellman-group14-sha",
	}
	promptPattern := regexp.MustCompile(`(?im)^[a-z\d.\-@_()/:]{1,48}[#>$]\s*$`)
	extraCiphers := []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-cbc", "aes192-cbc", "aes256-cbc"}

	opts := []util.Option{
		options.WithAuthNoStrictKey(),
		options.WithTransportType("standard"),
		options.WithStandardTransportExtraKexs(extraKexs),
		options.WithStandardTransportExtraCiphers(extraCiphers),
		options.WithTimeoutSocket(15 * time.Second),
		options.WithTimeoutOps(300 * time.Second),
		options.WithPromptPattern(promptPattern),
		options.WithPort(port),
		options.WithAuthUsername(username),
		options.WithAuthPassword(password),
	}

	driver, err := generic.NewDriver(host, opts...)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create driver: %w", err))
	}

	if err := driver.Open(); err != nil {
		return diag.FromErr(fmt.Errorf("failed to open connection: %w", err))
	}
	defer driver.Close()

	var resultBuilder strings.Builder
	for _, cmd := range commands {
		command := cmd.(string)
		resp, err := driver.SendCommand(command)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error sending command: %w", err))
		}

		resultBuilder.WriteString(fmt.Sprintf("Command: %s\n", command))
		resultBuilder.WriteString(fmt.Sprintf("Response:\n%s\n", resp.Result))
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error sending command: %w", err))
	}
	d.SetId(fmt.Sprintf("%s:%s", host, username))
	d.Set("result", resultBuilder.String())

	return diags
}


