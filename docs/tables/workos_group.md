# Table: workos_group

A Directory Group represents an Enterprise organizational unit of users. Developers can receive Webhooks as groups are added, updated, or removed, allowing for group-based resource access.

## Examples

### Basic info

```sql
select
  id,
  name,
  directory_id,
  organization_id,
  created_at
from
  workos_group;
```

### List groups of a particular organization

```sql
select
  g.id as group_id,
  g.name as group_name,
  g.directory_id,
  g.organization_id,
  g.created_at
from
  workos_group as g,
  workos_organization as o
where
  g.organization_id = o.id
  and o.name = 'test';
```

### List groups of a particular directory

```sql
select
  g.id as group_id,
  g.name as group_name,
  g.directory_id,
  g.organization_id,
  g.created_at
from
  workos_group as g,
  workos_directory as d
where
  g.directory_id = d.id
  and d.name = 'test';
```

### List groups created in the last 30 days

```sql
select
  id,
  name,
  directory_id,
  organization_id,
  created_at
from
  workos_group
where
  created_at >= now() - interval '30' day;
```
