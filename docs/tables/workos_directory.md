---
title: "Steampipe Table: workos_directory - Query WorkOS Directories using SQL"
description: "Allows users to query WorkOS Directories, specifically the directory ID, name, type, and state, providing insights into the organization's directory structure and status."
---

# Table: workos_directory - Query WorkOS Directories using SQL

WorkOS Directories is a feature within WorkOS that enables you to manage and sync your organization's directory data. It provides a centralized way to handle directory services, including users, groups, and more. WorkOS Directories assists you in maintaining an updated and organized view of your directory structure across various platforms.

## Table Usage Guide

The `workos_directory` table provides insights into directories within WorkOS. As an IT administrator, explore directory-specific details through this table, including directory ID, name, type, and state. Utilize it to uncover information about directories, such as those that are active or inactive, the type of directory (e.g., Google Workspace, Azure AD), and the verification of directory names.

## Examples

### Basic info
Explore which directories are active within your organization, when they were created, and their associated domain details. This can be beneficial for understanding the structure and timeline of your organizational directories.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which specific organizational directories are located and when they were created. This is particularly useful for tracking and managing directories within a certain organization.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that contain unlinked directories to maintain data integrity and ensure all directories are properly connected. This is beneficial in identifying potential issues that could affect data accessibility and organization.

```sql+postgres
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

```sql+sqlite
The PostgreSQL query provided does not use any PostgreSQL-specific functions or data types, so it can be used as it is in SQLite.

Here is the SQLite query:

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
```

### List gsuite directories
Explore the Gsuite directories in your organization to understand their current state and creation date. This can be useful for auditing purposes or to manage the directories more effectively.

```sql+postgres
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

```sql+sqlite
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

### List directories created in the last 30 days
Explore recent organizational changes by identifying directories that were established within the last month. This allows for a timely review and management of new additions to the system.

```sql+postgres
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
  created_at >= now() - interval '30' day;
```

```sql+sqlite
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
  created_at >= datetime('now', '-30 day');
```