# Onion-Architecture-Go

This example demonstrates the implementation of Onion Architecture in  Golang project.

Onion Architecture, introduced by [Jeffrey Palermo](https://jeffreypalermo.com/about/). It builds upon the core principles of Hexagonal Architecture, which takes a more granular approach by separating the core components into distinct layers: the business layer, domain logic layer, and models. This distinction provides a more nuanced structure for organizing application layers.

## Give a Star ⭐️

I'm glad that you're reading! If you found this project is useful or mindblow for you, do give it a star. Thanks!

## App Concept

### Architecture

![Onion Architecture Concept](https://raw.githubusercontent.com/yuntai229/Onion-Architecture-Go/develop/Onion%20Architecture.drawio.png)

* The left side is input side, where the data resource from. The **Presentation** layer is acted as input.
* The right side is output side, where the side app drive. The **Infrastructure** layer is acted as output.
* **Ports** is the place that interface written. The **Infrastructure** must to go througth it, and the **Presentation** is optional, but I make all of them interact with **Ports** to increase App Core isolation.
* App Core is a scope contain all your app logic, rule, model code, not a layer.
* It's OK to has some rule defined in the **Model** layer.

### Folder Structure

```text
├── app              --> The business logic
├── cmd              --> The app init step
├── domain
│   ├── model
│   └── service      --> Domain Service
├── dto
├── extend           --> App Utils/Tools
├── infrastructure
│   └── rdb
├── mocks            --> Test obj, auto generated
├── ports
└── presentation
    └── api
        ├── handler
        └── middleware
```

All of these are folders. The app entry point is **main.go**

### Relation

![Onion Architecture Relation](https://raw.githubusercontent.com/yuntai229/Onion-Architecture-Go/develop/Onion%20Architecture-relation.drawio.png)

## How to Run

Working on it...
