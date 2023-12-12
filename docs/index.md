---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/workos.svg"
brand_color: "#6363F1"
display_name: "WorkOS"
short_name: "workos"
description: "Steampipe plugin to query directories, groups and more from WorkOS."
og_description: "Query WorkOS with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/workos-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# WorkOS + Steampipe

[WorkOS](https://workos.com/) is a modern API platform that empowers any developer to quickly build and ship enterprise features.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List your WorkOS organizations:

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

## Quick start

### Install

Download and install the latest WorkOS plugin:

```bash
steampipe plugin install workos
```

### Credentials

| Item        | Description                                                                                                                                                                 |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | WorkOS requires an [API Key](https://workos.com/docs/reference/api-keys) for all requests.                                                                                  |
| Permissions | API keys have the same permission as the user who creates them, and if the user permissions change, the API key permissions also change.                                   |
| Radius      | Each connection represents a single WorkOS Installation.                                                                                                                    |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/workos.spc`)<br />2. Credentials specified in environment variables, e.g., `WORKOS_API_KEY`. |

### Configuration

Installing the latest workos plugin will create a config file (`~/.steampipe/config/workos.spc`) with a single connection named `workos`:

```hcl
connection "workos" {
  plugin = "workos"

  # `api_key` - API key for your WorkOS account. (Required)
  # For more information on the API Key, please see https://workos.com/docs/reference/api-keys.
  # Can also be set with the WORKOS_API_KEY environment variable.
  # api_key = "sk_test_a2V5XzAxR1g1QjNDRTFCU1NYSEhZMktINjVWTUFSLHBUSm1pWGpVMnV6dDNK"
}
```

Alternatively, you can also use the standard WorkOS environment variables to obtain credentials **only if other argument (`api_key`) is not specified** in the connection:

```sh
export WORKOS_API_KEY=sk_test_a2V5XzAxR1g1QjNDRTFCU1NYSEhZMktINjVWTUFSLHBUSm1pWGpVMnV6dDNK
```


