# TODO

- [ ] fix image sizes
- [ ] update data yamls
  - [ ] lps
  - [ ] books
  - [ ] tabs (improve descriptions)
  - [ ] tags urls
- [ ] MVP
  - [ ] lps: eda, golang, k8s, microservices, software arch, sre, system design, team management 
  - [ ] tag links
  - [ ] tabs descriptions
- [x] fix books loaded counter in main.go logs
- [x] fix if the lp has no tabs declared but it has books the book list appears empty -> made tabs mandatory
- [x] Sort books when loaded in aux data structures
- [x] Fix case when there are lp tabs that are not declared in lps (the tabs are ignored)
- [x] Get better order icons (function to parse from 12 to ":material-numeric-1-box::material-numeric-2-box:)
- [x] Fix github learning paths
- [x] Use generics to replace loadEntity() funcs
- [x] Load tags
- [x] Include tabs in lp templates
- [x] Load badges with the new structure
- [x] Load tabs (they are integrated in the data directory of the lps, must differentiate them)
- [x] Define and create auxiliar data structures
- [x] Sort lp tabs when loading learning-paths.yaml
- [x] Override mechanisms for lp tabs (merge with tabs.yaml)
- [x] Check that lps don't have duplicated tabs
- [x] Check that books and lps don't have duplicated tags
- [x] Check that a book does not appears twice in the same lp tab
- [x] tabs
    - [x] tabs: show icon in tab name
    - [x] tabs: show in tab name the number of books between parenthesis len(LearningPathTabBooks[b.lpRef][b.tabRef])
- [x] tags index

## Roadmap

- [ ] Buy me a coffee button
- [ ] Use mark with an icon 'new' for lps, tabs or books added (in lps, book index)
- [?] Show book TOC when hovering the  cover o add a new element (instant preview)
