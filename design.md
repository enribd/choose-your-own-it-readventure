# Design

Book information:
  - cover
  - title
  - author
  - release date
  - pages
  - url to official info
  - my own description (where does it fit in the learning path and why)
  - learning paths associated (icons?)
  - badges:
    - level: beginner, intermediate, advanced (icons?)
    - tags icons

## TODO

- At the end of each learning path give some ideas about how to practice the skills acquired.

## App

Features:
- auto generate toc in README
- auto generate book index page
- book index page: learning paths are links to each path
- auto generate learning paths page


## Data file

books:
  - cover: building-microservices.jpeg
    title: Building Microservices
    subtitle: Designing Fine-Grained Systems
    url: https://learning.oreilly.com/library/view/-/9781492034018/
    authors: Sam Newman
    release: 2022
    pages: 616
    description: |
      One of the most important books in the field. Far from advocating for the monolithic architectures exile, the book offers useful insights to help you identify use cases for monoliths, or when to turn to microservices. It will teach you what microservices really are, their evolutionary origin, principles, characteristics and all the new challenges they bring to the table. Finally, the author explains how organizations should evolve to adapt their internal structure and vision in order to efficiently deliver value using microservices architectures.
    learning_paths:
      - system_design
      - microservices
    badges:
      difficulty: low
      progress: read
      rating: top
      relevant: true
paths:
  - system_design:
    - apis
    - microservices
    - event_driven_architecture
    - serverless
  - golang:
    - cloud_native_applications
  - kubernetes
  - software_architecture:
    - domain_driven_design
  - management:
    - devops
    - team_management
badges:
  difficulty:
    - name: beginner
      icon: ant
    - name: easy
      icon: hatched_chick
    - name: intermediate
      icon: dog2
    - name: hard
      icon: tiger2
    - name: expert
      icon: dragon
  progress:
    - name: read
      icon: green_book
    - name: scheduled
      icon: blue_book
    - name: not_scheduled
      icon: orange_book
  rating:
    - name: nice
      icon: ok
    - name: good
      icon: up
    - name: very_good
      icon: cool
    - name: excellent
      icon: top
  other:
    - name: must_read
      icon: bookmark
    - name: old
      icon: arrows_counterclockwise
