---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/workos.svg"
brand_color: "#6495ED"
display_name: "WorkOS"
short_name: "workos"
description: "Steampipe plugin to query directories, groups and more from WorkOS."
og_description: "Query WorkOS with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/workos-social-graphic.png"
---

# WorkOS + Steampipe

[WorkOS](https://workos.com/) is a modern API platform that empowers any developer to quickly build and ship enterprise features.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Get WorkOS organization details:

```sql
select
  id,
  name,
  allow_profiles_outside_organization,
  created_at
from
  workos_organization;
```

```
+--------------------------------+-------------+-------------------------------------+---------------------------+
| id                             | name        | allow_profiles_outside_organization | created_at                |
+--------------------------------+-------------+-------------------------------------+---------------------------+
| org_01GX5EDRXZV7GTF3SQHXKHBGFE | turbot      | true                                | 2023-04-04T11:44:18+05:30 |
| org_01GX5SQ6CBFKP7X1A804PH4FT8 | turbot-dev  | true                                | 2023-04-04T15:01:41+05:30 |
+--------------------------------+-------------+-------------------------------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/workos/tables)**

## Get started

### Install

Download and install the latest WorkOS plugin:

```bash
steampipe plugin install workos
```

### Configuration

Installing the latest workos plugin will create a config file (`~/.steampipe/config/workos.spc`) with a single connection named `workos`:

```hcl
connection "workos" {
  plugin = "workos"

  # API key for your WorkOS account.
  # For more information on the APIKey, please see https://workos.com/docs/reference/api-keys.
  # Can also be set with the WORKOS_API_KEY environment variable.
  # api_key = "sk_test_a2V5XzAxR1g1QjNDRTFCU1NYSEhZMktINjVWTUFSLHBUSm1pWGpVMnV6dDNK"
}
```

- `api_key` - WorkOS API key. Can also be set with the `WORKOS_API_KEY` environment variable.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-workos
- Community: [Slack Channel](https://steampipe.io/community/join)
