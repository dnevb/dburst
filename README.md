# DBURST

Database tool for developers, SQL programmers, database administrators and analysts.

## Installation

The fastest way to get started with dburst is with the docker image

### Docker

pull the image

```sh
docker pull dnevb/dburst:latest
```

and run the image

```sh
docker run -p 8080:8080 dnevb/dburst

# or attaching a volume
docker run -v dburst_data:/data -p 8080:8080 dnevb/dburst
```

## Documentation

TODO

## Roadmap

- [ ] schema inspector
  - columns
  - indexes
  - triggers
  - table preview
  - object definition
  - schema editor
  - metadata editor
- [ ] Sql editor
  - connector selection
  - autocomplete
  - quick database schema view
  - object D&D
  - export to worksheet
- [ ] ERD
  - generation
  - viewer
- [ ] Worksheets
  - sql sheets
  - python sheets
  - prql sheets
  - variable definitions
  - version history
  - realtime editor
- [ ] Data export/import
  - allow sql/csv/parquet formats
- [ ] Metadata search engine
  - full text search
  - metadata indexing
- [ ] Migration manager
  - schema migration
  - up/down migrations
  - stages manager
- [ ] Task scheduling
  - schedule worksheets
  - workflows
- [ ] Git integrations
  - export/import worksheets to github
- [ ] User management
  - permissions
  - user teams/groups
- [ ] Execution plan viewer
- [ ] DBT core
  - itegration with worksheets and connectors
  - git management
  - scheduling
