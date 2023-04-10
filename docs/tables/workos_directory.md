# Table: workos_directory

A Directory stores information about an Enterprise Client’s employee management system. Synchronizing with a Directory enables Developers to receive changes to an Enterprise Client’s User and Group structure.

Directory Providers vary in implementation details and may require different sets of fields for integration, such as API keys, subdomains, endpoints, usernames, etc. Where available, the WorkOS API will provide these fields when fetching Directory records.

## Examples

### Basic info

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  domain
from
  workos_directory;
```

### List directories of a particular organization

```sql
select
  d.id as directory_id,
  d.name as directory_name,
  d.state,
  d.organization_id,
  d.created_at,
  d.domain
from
  workos_directory as d,
  workos_organization as o
where
  d.organization_id = o.id
  and o.name = 'test';
```

### List unlinked directories

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  domain
from
  workos_directory
where
  state = 'unlinked';
```

### List gsuite directories

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  domain
from
  workos_directory
where
  type = 'gsuite directory';
```
