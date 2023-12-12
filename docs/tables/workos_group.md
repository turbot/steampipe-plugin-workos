---
title: "Steampipe Table: workos_group - Query WorkOS Groups using SQL"
description: "Allows users to query WorkOS Groups, providing details such as group ID, name, description, and associated users."
---

# Table: workos_group - Query WorkOS Groups using SQL

WorkOS is a service that allows developers to implement enterprise-level features into their applications. It provides functionalities such as Single Sign-On (SSO), Directory Sync, and Audit Trail, making it easier to manage and secure applications. A WorkOS Group is a collection of users from an organization's identity provider that can be managed as a single entity.

## Table Usage Guide

The `workos_group` table provides insights into WorkOS Groups within WorkOS. As a developer or IT administrator, explore details about each group through this table, including group ID, name, description, and associated users. Utilize it to manage and organize users effectively, ensuring appropriate access control and security within your application.

## Examples

### Basic info
Explore which groups have been created within your organization, including their unique identifiers and the time of creation. This can help you keep track of group formation and understand the structure of your organization over time.

```sql+postgres
select
  id,
  name,
  directory_id,
  organization_id,
  created_at
from
  workos_group;
```

```sql+sqlite
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
Explore which groups belong to a specific organization. This is useful to understand the structure and distribution of groups within that organization, providing insights for management and organizational planning.

```sql+postgres
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

```sql+sqlite
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
Explore which groups belong to a specific directory to better manage and organize your resources. This can be particularly useful in identifying and sorting groups for administrative or security purposes.

```sql+postgres
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

```sql+sqlite
select
  g.id as group_id,
  g.name as group_name,
  g.directory_id,
  g.organization_id,
  g.created_at
from
  workos_group as g
join
  workos_directory as d
on
  g.directory_id = d.id
where
  d.name = 'test';
```

### List groups created in the last 30 days
Explore groups that have been established within the past month. This can be beneficial for understanding recent organizational changes or additions.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  directory_id,
  organization_id,
  created_at
from
  workos_group
where
  created_at >= datetime('now', '-30 day');
```