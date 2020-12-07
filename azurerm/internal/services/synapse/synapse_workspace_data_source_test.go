package synapse_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
)

type SynapseWorkspaceDataSource struct{}

func TestAccDataSourceSynapseWorkspace_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_synapse_workspace", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: SynapseWorkspaceDataSource{}.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("resource_group_name").Exists(),
				check.That(data.ResourceName).Key("connectivity_endpoints.%").Exists(),
			),
		},
	})
}

func (d SynapseWorkspaceDataSource) basic(data acceptance.TestData) string {
	config := SynapseWorkspaceResource{}.basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_synapse_workspace" "test" {
  name                = azurerm_synapse_workspace.test.name
  resource_group_name = azurerm_synapse_workspace.test.resource_group_name
}
`, config)
}