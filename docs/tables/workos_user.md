---
title: "Steampipe Table: workos_user - Query WorkOS Users using SQL"
description: "Allows users to query WorkOS Users, specifically the user's details, providing insights into user data and potential anomalies."
---

# Table: workos_user - Query WorkOS Users using SQL

WorkOS is a platform that provides a set of APIs to help developers quickly add enterprise-ready features to their applications. It includes capabilities for Single Sign-On, Directory Sync, and Audit Trail, among others. A WorkOS User is an individual with access to WorkOS, with details including user ID, email, first and last names, and related organization data.

## Table Usage Guide

The `workos_user` table provides insights into user details within WorkOS. As a developer or security analyst, explore user-specific details through this table, including user ID, email, first and last names, and related organization data. Utilize it to uncover information about users, such as their associated organizations and roles, aiding in user management and security analysis.

## Examples

### Basic info
Explore the user profiles within your organization to gain insights into their status and creation date. This can help in assessing the user activity and managing the user database effectively.

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
Identify users whose accounts are currently suspended. This is useful for account management and to ensure that any unexpected or unauthorized suspensions are immediately addressed.

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
Explore which users belong to a specific group, allowing for efficient management and organization of user access and permissions. This is particularly beneficial in large organizations where grouping users can simplify administrative tasks.

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
Explore which users belong to a specific organization, gaining insights into their user ID, username, and other relevant details. This can be useful for managing user access and understanding user distribution across different organizations.

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
Explore which users are associated with a specific directory to manage access and permissions efficiently. This is particularly useful for administrators seeking to maintain security and organization within their system.

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