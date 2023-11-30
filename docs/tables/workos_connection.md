---
title: "Steampipe Table: workos_connection - Query WorkOS Connections using SQL"
description: "Allows users to query WorkOS Connections, which are used to configure and manage Single Sign-On (SSO) connections to identity providers."
---

# Table: workos_connection - Query WorkOS Connections using SQL

WorkOS Connections are a service within WorkOS that allows administrators to configure and manage Single Sign-On (SSO) connections to various identity providers. It provides a centralized way to set up and manage SSO connections, streamlining the process of integrating SSO into your application. WorkOS Connections help you leverage existing enterprise identity platforms and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `workos_connection` table provides insights into SSO connections within WorkOS. As a system administrator or developer, explore connection-specific details through this table, including the connection type, domain, status, and associated metadata. Utilize it to uncover information about connections, such as those with specific identity providers, the status of each connection, and the verification of connection details.

## Examples

### Basic info
Explore which connections are active within your organization by identifying the specific instances and their creation dates. This can help assess the overall usage and manage resources effectively.

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
Determine the areas in which your WorkOS connections are inactive. This can help you manage resources more efficiently by identifying unused connections and potentially reassigning them.

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
Explore which connections are associated with a specific organization to better manage and monitor your organizational resources. This can be particularly useful in identifying potential bottlenecks or inefficiencies within your organization's network.

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
Explore which connections are based on Azure within your organization. This can be useful for assessing the prevalence and usage of Azure services in your network.

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
Discover the connections that were established in the recent month. This is useful for monitoring recent activity and understanding your organization's interaction patterns.

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