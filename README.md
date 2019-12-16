[![CircleCI](https://circleci.com/gh/giantswarm/sudo-azure-operator.svg?&style=shield)](https://circleci.com/gh/giantswarm/sudo-azure-operator) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/sudo-azure-operator/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/sudo-azure-operator)

# sudo-azure-operator

Idea: `sudo-azure-operator` manages Azure resources that require higher-level roles, so that permissions
required by the `azure-operator` can be reduced to resource group level.

Goals:
- Reduce level of permissions required by `azure-operator`
- Enable alternative workflow of creating tenant clusters where `azure-operator` gets to use existing
  Azure resources that require higher-level roles (e.g. resource groups that require subscription-level
  permissions to be created)

Non-goals:
- Change default user experience for creating tenant clusters on Microsoft Azure