

[//]: # (Auto generated file from templates)

# ![img](https://user-images.githubusercontent.com/25181517/182534006-037f08b5-8e7b-4e5f-96b6-5d2a5558fa85.png){: style="height:35px"} Kubernetes Learning Path (14 :books:)

After mastering building and running containers at small scale, orchestration is the next step in the road. Kubernetes is the most popular orchestrator, backed by the [CNCF](https://www.cncf.io/), a variety os small and big companies, and open source projects. Thanks to its capabilities and the ecosystem of open source projects built around it, Kubernetes, is a the industry facto standard for running microservice platforms at scale.

=== ":material-dots-vertical: Foundational &nbsp; 3:material-bookshelf:"
    Basic concepts required to continue the journey of this learning path :material-ray-start-arrow:.

    | Order | Cover | Info | Description |
    | :---: | :---: | :--- | :--- |
    | **:material-numeric-1:{.order-icon}** |![img](/assets/books/covers/kubernetes-up-and-running.jpeg)| [**Kubernetes: Up & Running**](https://learning.oreilly.com/library/view/-/9781098110192/) <br> *Brendan Burns, Joe Beda, Kelsey Hightower* <br> *Published in 2022* <br> *326 pages* <br> :hatched_chick:{ title="Easy" } :green_book:{ title="Read" } :cool:{ title="Very Good" } | Start your Kubernetes journey from the very basics. Learn the building blocks to get a solid knowledge base that will allow you advance in the fascinating world of container orchestration and microservice platforms. Recommended to every developer, software architect, infrastructure engineer or platform engineer due to the Kubernetes relevance nowadays.<br><br><br>[**#pod**]()&nbsp;&nbsp;[**#deployment**]()&nbsp;&nbsp;[**#statefulset**]()&nbsp;&nbsp;[**#confimap**]()&nbsp;&nbsp;[**#secret**]()&nbsp;&nbsp;|
    | **:material-numeric-2:{.order-icon}** |![img](/assets/books/covers/kubernetes-best-practices.jpeg)| [**Kubernetes Best Practices: Kubernetes Best Practices**](https://learning.oreilly.com/library/view/-/9781492056461/) <br> *Brendan Burns, Eddie Villalba, Dave Strebel, Lachlan Evenson* <br> *Published in 2019* <br> *268 pages* <br> :hatched_chick:{ title="Easy" } :orange_book:{ title="Not Scheduled" } | Usually, after being familiar with the basics of any technology, learning the best practices is a nice step, it will help you to avoid common mistakes and open new ways of thinking and doing things. This book will do exaclty that, from basic to more advanced topics.<br><br><br>[**#probes**]()&nbsp;&nbsp;[**#init-container**]()&nbsp;&nbsp;[**#service-account**]()&nbsp;&nbsp;[**#rbac**]()&nbsp;&nbsp;|
    | **:material-numeric-3:{.order-icon}** |![img](/assets/books/covers/kubernetes-patterns.jpeg)| [**Kubernetes Patterns: Reusable Elements for Designing Cloud Native Applications**](https://learning.oreilly.com/library/view/-/9781098131678/) <br> *Bilgin Ibryam, Roland Huss* <br> *Published in 2023 (2nd edition)* <br> *300 pages* <br> :hatched_chick:{ title="Easy" } :green_book:{ title="Read" } :cool:{ title="Very Good" } :soon:{ title="Coming-Soon" } | After a few years of using Kubernetes teams started to detect recognizable patterns that are commonly replicated. From the famous sidecar pattern and init containers to the necessary configuration patterns and beyond. This book will teach you how to solve common challenges in cloud native environments. <br><br><br>[**#adapter-pattern**]()&nbsp;&nbsp;[**#ambassador-pattern**]()&nbsp;&nbsp;[**#init-container**]()&nbsp;&nbsp;[**#sidecar-container**]()&nbsp;&nbsp;|



=== ":material-dots-grid: Advanced &nbsp; 2:material-bookshelf:"
    Avanced concepts

    | Order | Cover | Info | Description |
    | :---: | :---: | :--- | :--- |
    | **:material-numeric-1:{.order-icon}** |![img](/assets/books/covers/production-kubernetes.jpeg)| [**Production Kubernetes: Production Kubernetes**](https://learning.oreilly.com/library/view/-/9781492092292/) <br> *Josh Rosso, Rich Lander, Alex Brand, John Harris* <br> *Published in 2021* <br> *506 pages* <br> :tiger2:{ title="Hard" } :green_book:{ title="Read" } :star:{ title="Excellent" } | Production Kubernetes will guide you from the view of an amateur to the vision of all you need to know and implement in order to build a professional and productive microservice platform. It also demonstrates the Kubernetes API extensibility through drivers for networking, storage and more. Other interesting topis are admission controllers, operators, service meshes and security. Definitely this is a book I really enjoyed, I recommend to read it at the same time with [Design Patterns for Cloud Native Applications](https://www.oreilly.com/library/view/design-patterns-for/9781492090700/), it will broaden your vision as a developer as well as your perspective as platform engineer.<br><br><br>[**#platform-engineering**]()&nbsp;&nbsp;[**#admission-controller**]()&nbsp;&nbsp;[**#spiffe**]()&nbsp;&nbsp;[**#csi-driver**]()&nbsp;&nbsp;[**#cni-driver**]()&nbsp;&nbsp;|
    | **:material-numeric-2:{.order-icon}** |![img](/assets/books/covers/kubernetes-operators.jpeg)| [**Kubernetes Operators: Automating the Container Orchestration Platform**](https://learning.oreilly.com/library/view/-/9781492048039/) <br> *Jason Dobies, Joshua Wood* <br> *Published in 2020* <br> *154 pages* <br> :tiger2:{ title="Hard" } :orange_book:{ title="Not Scheduled" } :arrows_counterclockwise:{ title="Old" } | Operators are another way of leveraging Kubernetes itself, they offer you the possibility to customize and automate your workloads management. Its a very popular pattern, in this book you will find what you need to get started to develop your own operator and all the tooling around the process. This book is a little bit 'old', many things have changed in this topic since 2020, it needs a second edition to adapt its contents to more updated practices.<br><br><br>[**#controller**]()&nbsp;&nbsp;[**#feedback-loop**]()&nbsp;&nbsp;|



=== ":material-function-variant: FaaS &nbsp; 1:material-bookshelf:"
    Function as a Service or serverless

    | Order | Cover | Info | Description |
    | :---: | :---: | :--- | :--- |
    | **:material-numeric-1:{.order-icon}** |![img](/assets/books/covers/knative-in-action.jpeg)| [**Knative in Action**](https://learning.oreilly.com/library/view/-/9781617296642/) <br> *Jacques Chester* <br> *Published in 2021* <br> *272 pages* <br> :tiger2:{ title="Hard" } :green_book:{ title="Read" } :ok:{ title="Nice" } | Knative is, along [OpenFaaS](https://www.openfaas.com), one of the most popular ways to build a serverless platform over Kubernetes. The author starts from zero explaining the motivation to adopt a serverless paradigm, then presents Knative in a technical demonstration of all its basic and some of the more advanced features like traffic management and application delivery. There are not many books on the topic yet as Knative is still a relatively young project.<br><br><br>[**#cold-start**]()&nbsp;&nbsp;[**#scale-to-zero**]()&nbsp;&nbsp;[**#faas**]()&nbsp;&nbsp;|


The following paths are opened to you now, choose wisely:

- [Microservices :construction:](/learning-paths/microservices): Study the pinnacle of distributed systems architectures, learn its tenets, and foremost, when and how to implement it.


Want to change the subject? Here are some suggestions about other paths you can explore:

- [System Design :ballot_box_with_check:](/learning-paths/system-design): Acquire the skill needed to design and build systems, no matter if simple or complex. Learn how to identify the elements needed to create systems, to resolve scalability problems, detect possible points of failure, when to use an API, where to place a cache, when to use a NoSql database, and more.


??? tip "Learn more related concepts! :round_pushpin: :beginner: :gem:"

    <sub>[**#container-runtime**]()&nbsp;&nbsp;[**#ingress**]()&nbsp;&nbsp;[**#controller**]()&nbsp;&nbsp;[**#crd**]()&nbsp;&nbsp;[**#operator**]()&nbsp;&nbsp;[**#csi-driver**]()&nbsp;&nbsp;[**#admission-controller**]()&nbsp;&nbsp;[**#service-mesh**]()&nbsp;&nbsp;[**#platform-engineering**]()&nbsp;&nbsp;</sub>

[**⬆ back to top**](#kubernetes-learning-path-14)
