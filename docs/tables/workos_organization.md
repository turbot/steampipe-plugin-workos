---
title: "Steampipe Table: workos_organization - Query WorkOS Organizations using SQL"
description: "Allows users to query WorkOS Organizations, providing crucial data regarding the organizations managed by WorkOS."
---

# Table: workos_organization - Query WorkOS Organizations using SQL

WorkOS is a platform that helps developers to quickly implement enterprise-ready features into their applications. It provides features such as Single Sign-On, Directory Sync, and more, enabling seamless integration with various enterprise environments. An organization in WorkOS represents a group of users, typically corresponding to a company, that uses the same WorkOS settings and features.

## Table Usage Guide

The `workos_organization` table provides insights into organizations managed by WorkOS. As a developer or system administrator, you can explore organization-specific details through this table, including domain, name, and associated metadata. Utilize it to manage and monitor the organizations, their settings, and to understand the usage of WorkOS features across different organizations.

## Examples

### Basic info
Explore the settings of your organization to understand whether profiles from outside the organization are permitted, and assess when these settings were last updated. This is useful for maintaining security and ensuring only authorized profiles have access.

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

### List organizations that allow outside profiles
Explore which organizations permit profiles from outside their own, providing insights into their openness to external collaboration. This can assist in identifying potential partners or assessing the openness of your competitive landscape.

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
Explore which domains are associated with a specific organization to better manage or monitor the organization's online presence and activities. This is especially useful for IT administrators or cybersecurity professionals who need to keep track of all the domains under an organization's purview.

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
Discover the organizations that have been established in the past month. This is useful to keep track of new additions and ensure all recent organizations are accounted for.

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