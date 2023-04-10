# Table: workos_user

A Directory User represents an active Enterprise user. Developers can receive Webhooks as employees are added, updated or removed, allowing for provisioning and de-provisioning Users within an application.

## Examples

### Basic info

```sql
select
  id,
  user_name,
  state,
  directory_id,
  organization_id,
  created_at,
  first_name,
  last_name
from
  workos_user;
```

### List suspended users

```sql
select
  id,
  user_name,
  state,
  directory_id,
  organization_id,
  created_at,
  first_name,
  last_name
from
  workos_user
where
  state = 'suspended';
```

### List users of a particular group

```sql
select
  id,
  user_name,
  state,
  directory_id,
  organization_id,
  created_at,
  first_name,
  last_name
from
  workos_user,
  jsonb_array_elements(groups) as g
where
  g ->> 'Name' = 'test';
```

### List users of a particular organization

```sql
select
  u.id as user_id,
  u.user_name,
  u.state,
  u.directory_id,
  u.organization_id,
  u.created_at,
  u.first_name,
  u.last_name
from
  workos_user as u,
  workos_organization as o
where
  u.organization_id = o.id
  and o.name = 'test';
```

### List users of a particular directory

```sql
select
  u.id as user_id,
  u.user_name,
  u.state,
  u.directory_id,
  u.organization_id,
  u.created_at,
  u.first_name,
  u.last_name
from
  workos_user as u,
  workos_directory as d
where
  u.directory_id = d.id
  and d.name = 'test';
```
