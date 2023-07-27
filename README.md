![image](https://hub.steampipe.io/images/plugins/turbot/workos-social-graphic.png)

# WorkOS Plugin for Steampipe

Use SQL to query directories, groups and more from WorkOS.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/workos)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/workos/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-workos/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install workos
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/workos#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/workos#configuration).

Configure the API key in `~/.steampipe/config/workos.spc`:

```hcl
connection "workos" {
  # Authentication information
  api_key   = "sk_test_a2V5XzAxR1g1QjNDRTFCU1NYSEhZMktINjVWTUFSLHBUSm1pWGpVMnV6dDNK"
}
```

or through environment variables

```sh
export WORKOS_API_KEY="sk_test_a2V5XzAxR1g1QjNDRTFCU1NYSEhZMktINjVWTUFSLHBUSm1pWGpVMnV6dDNK"
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-workos.git
cd steampipe-plugin-workos
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/workos.spc
```

Try it!

```
steampipe query
> .inspect workos
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-workos/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [WorkOS Plugin](https://github.com/turbot/steampipe-plugin-workos/labels/help%20wanted)
