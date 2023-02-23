
[//]: # (Auto generated file from templates)

# Choose your own IT readventure! :books: :sunrise_over_mountains: :computer:

If you are a stubborn learner, eat books for breakfast, and want to discover new concepts or hunt new skills related to the IT world, I can assure you that you are in the right nook of the Internet. Prepare yourself to be amazed by a choose-your-own-adventure journey in the IT world. Hopefully, it will take you to interesting places thanks to the power and the fun of reading.

Finally, make sure you don’t forget about putting your newly acquired skills to the test, remember practice makes the master!

Starting your adventure is as easy as picking the set of skills that you want to learn or improve, check out the learning paths and choose your starting point.

:warning: *A minimum previous knowledge is required to start some of these adventures. Various of them have been assembled to enlarge or acquire expertise in topics considered, until certain point, advanced. If you are new to the IT world or if you want to follow a learning path of high difficulty my recommendation is that, before adventuring, you read first about the basics and foundational concepts of that particular topic.*


## :checkered_flag: Start your adventure, choose your learning path wisely

Stats:

{{- $badgesData := .BadgesData -}}
{{- $lpData := .LpData }}
- **{{ .Stats.TotalLearningPaths }}** learning paths :runner:
{{- range $ref, $num := .Stats.BooksInLearningPath -}}
{{- with $lp := get $lpData ($ref | toString) }}
{{- if ne $lp.Status "coming-soon" }}
{{- $icon := get $badgesData $lp.Status }}
  - **{{ $num }}** [*{{ $lp.Name }}*]({{ $.LearningPathsFolder | trimPrefix "." }}/{{ $ref }}.md) books :{{ $icon }}:
{{- end -}}
{{- end -}}
{{- end }}
- **{{ .Stats.TotalBooks }}** books :books: , discover them all in the [:scroll: book index]({{ .BookIndex }}).
- **{{ .Stats.TotalAuthors }}** authors :black_nib: , discover them all in the [:scroll: author index]({{ .AuthorIndex }}).

## :name_badge: Badges

All books have some badges associated to describe some aspect of them:

- You will read each book with the same amount of energy needed to fight these opponents, from lower to higher difficulty (these are totally subjective):

| Badge | Level |
| --- | --- |
| :ant: | Petty ant |
| :hatched_chick: | Naive chick |
| :dog2: | Brave dog |
| :tiger2: | Fierce tiger |
| :dragon: | Mighty dragon |

- Reading progress:

| Badge | Meaning |
| --- | --- |
| :green_book: | Read |
| :blue_book: | To be read soon |
| :orange_book: | Not read and I don't know when I will |

- My book rating and recommendation level from lower to higher:

| Badge  | Level     |
| ---    | ---       |
| :ok:   | Nice      |
| :up:   | Good      |
| :cool: | Very Good |
| :top:  | Excellent |

**Note**: there are just four levels because all books mentioned are good reads, bad books don't belong in this list :thumbsdown:.

- Other badges:

| Badge | Meaning |
| --- | --- |
| :arrows_counterclockwise: | The book it's somewhat 'old' and I think it would be great to have a new edition to refresh the content |
| :bookmark: | Recommended read no matter the learning path | 

## :alien: About me

I am a professional Cloud and Platform Engineer, DevOps practitioner, and Architect in construction. Also I develop things in my free time just for the fun of learning how not to do things. My main interests are the cloud, microservice architectures, EDA, serverless, system design, Kubernetes, and Golang.

This repo is conceived as a journal of my personal journey and a learning path to guide others who, like me, find in reading the best way to acquire deep knowledge in certain topics. I will put together a list that categorizes all the books I've read as well as those that I intend to read, or those that I consider of some importance in the IT field. The list is subjective, it will reflect my interests or the way I've decided to explore topics (most likely in a chaotic way).

If you think some books may fit better in other categories, or the category names are not accurate enough or have suggestions about possible improvements just let me know and I will gladly consider your comments. 

Of course, there will be books broadly considered a must-read that I have not listed, that's either because they don't fit in my career or simply because I don't know of their existence, in any case let me know!

:bowtie: Thanks for stopping by, enjoy!

## :pushpin: Worthy metions

- Icons:
  - [Skill icons](https://github.com/tandpfun/skill-icons)
  - [Tech icons](https://github.com/marwin1991/profile-technology-icons)
  - [Emoji icons](https://gist.github.com/kajal1106/b0bf3b9f93b4f484dc3703c8c64bbe1c)

[**⬆ top**](#choose-your-own-it-readventure-books-sunrise_over_mountains-computer)
