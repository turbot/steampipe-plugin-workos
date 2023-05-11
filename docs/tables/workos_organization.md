# Table: workos_organization

An Organization is a top-level resource in WorkOS. Each Connection, Directory, and Audit Trail Event belongs to an Organization. An Organization will usually represent one of your customers.

## Examples

### Basic info

```sql
select
  id,
  name,
  allow_profiles_outside_organization,
  created_at,
  updated_at
from
  workos_organization;
```

### List organizations which allowed outside profiles

```sql
select
  id,
  name,
  allow_profiles_outside_organization,
  created_at,
  updated_at
from
  workos_organization
where
  allow_profiles_outside_organization;
```

### List domains of a particular organization

```sql
select
  id as organization_id,
  name as organization_name,
  d ->> 'domain' as domain,
  d ->> 'id' as domain_id
from
  workos_organization,
  jsonb_array_elements(domains) as d
where
  name = 'test';
```

### List organizations created in the last 30 days

```sql
select
  id,
  name,
  allow_profiles_outside_organization,
  created_at,
  updated_at
from
  workos_organization
where
  created_at >= now() - interval '30' day;
```
