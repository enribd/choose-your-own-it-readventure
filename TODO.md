# TODO

- [-] Load tags
- [ ] Finish load tags tests
- [-] Load badges with the new structure
- [-] Load tabs (they are integrated in the data directory of the lps, must differentiate them)
- [?] Create a filters package to filter and sort models
- [ ] Define and create auxiliar data structures
- [-] Sort books when loaded in aux data structures
- [ ] Check that books or lps don't have duplicated tags
- [ ] Sort lp tabs when loading learning-paths.yaml
- [ ] Override mechanisms for lp tabs (merge with tabs.yaml)
- [ ] Check that lps don't have duplicated tabs
- [ ] make templates
    - [ ] tabs: show icon in tab name
    - [ ] tabs: show in tab name the number of books between parenthesis len(LearningPathTabBooks[b.lpRef][b.tabRef])
- [ ] tags index

## Book declaration example

```yaml
- title: Building Microservices
  subtitle: Designing Fine-Grained Systems
  cover: building-microservices.jpeg
  draft: false
  url: https://learning.oreilly.com/library/view/-/9781492034018/
  authors:
    - Sam Newman
  release: 2022
  pages: 616
  desc: |-
    One of the most important books in the field. Far from advocating for the monolithic architectures exile, the book offers useful insights to help you identify use cases for monoliths, or when to turn to microservices. It will teach you what microservices really are, their evolutionary origin, principles, characteristics and all the new challenges they bring to the table. Finally, the author explains how organizations should evolve to adapt their internal structure and vision in order to efficiently deliver value using microservices architectures.
  learning_paths:
    - lp_ref: microservices
      tab_ref: foundational
      order: 1
      weight: 1
    - lp_ref: microservices
      tab_ref: certification
      order: 3
      weight: 2
  badges:
    - intermediate
    - read
    - excellent
    - must-read
```

## LP declaration example

```yaml
- name: Microservices
  ref: microservices
  status: in-progress
  desc: |
    Distributed systems are not new but the way they are built nowadays is. Monolithic architectures need to evolve to leverage the cloud and the many advantages that microservices offer (scalability, fast releases, high-availability, resilience, and more). As usually happen in life, nothing is just benefits, and microservices architectures are not different, they bring many challenges with them like a more complicated management or debugging, economic costs and the necessary knowledge to build and run them. However, if this kind of architecture fit your needs or if you are interested in finding out what all the fuss about microservices is about don't hesitate and dive in!.
  summary: |
    Study the pinnacle of distributed systems architectures, learn its tenets, and foremost, when and how to implement it.
  tabs:
    # TODO merge with tabs.yaml data (this takes priority)
    # TODO sort by order when loaded
    # TODO check no orders are equal
    - ref: foundational
      data:
        # override order
        order: 1
    - ref: intermediate
      order: 2
    - ref: advanced
      order: 3
  related:
    - system-design
    - kubernetes
    - apis
    - event-driven-architecture
  suggested:
    - serverless
    - golang
    - docker
  tags:
  logo:
    source: /assets/learning-paths/icons/microservices.png
```
