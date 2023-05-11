# Table: workos_connection

A Connection represents the relationship between WorkOS and any collection of application users. This collection of application users may include personal or enterprise Identity Providers, or passwordless authentication methods like Magic Link. As a layer of abstraction, a WorkOS Connection rests between an application and its users, separating an application from the implementation details required by specific standards like OAuth 2.0 and SAML.

## Examples

### Basic info

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  connection_type
from
  workos_connection;
```

### List inactive connections

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  connection_type
from
  workos_connection
where
  state = 'inactive';
```

### List connections of a particular organization

```sql
select
  c.id as connection_id,
  c.name as connection_name,
  c.state,
  c.organization_id,
  c.created_at,
  c.connection_type
from
  workos_connection as c,
  workos_organization as o
where
  c.organization_id = o.id
  and o.name = 'test';
```

### List azure based connections

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  connection_type
from
  workos_connection
where
  connection_type like 'Azure%';
```

### List connections created in the last 30 days

```sql
select
  id,
  name,
  state,
  organization_id,
  created_at,
  connection_type
from
  workos_connection
where
  created_at >= now() - interval '30' day;
```
