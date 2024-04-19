# Onion-Architecture-Go

This example demonstrates the implementation of Onion Architecture in  Golang project.

Onion Architecture, introduced by [Jeffrey Palermo](https://jeffreypalermo.com/about/). It builds upon the core principles of Hexagonal Architecture, which takes a more granular approach by separating the core components into distinct layers: the business layer, domain logic layer, and models. This distinction provides a more nuanced structure for organizing application layers.

## Give a Star ⭐️

I'm glad that you're reading! If you found this project is useful or mindblow for you, do give it a star. Thanks!

## App Concept

### Architecture

![Onion Architecture Concept](https://raw.githubusercontent.com/yuntai229/Onion-Architecture-Go/master/Onion%20Architecture.drawio.png)

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
│   └── service      --> The domain logic
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

![Onion Architecture Relation](https://raw.githubusercontent.com/yuntai229/Onion-Architecture-Go/master/Onion%20Architecture-relation.drawio.png)

## How to Run

The build enrty point is **docker-compose.yaml**, and it based on **Dockerfile.local**. The online environment uat, product ...etc will be based on **Dockerfile.online**. In this project, you can just execute Make command in local, which defined in **Makefile**.

### Service

All there are executed in docker environment.

* Start

    Support hot reload by [Air](https://github.com/cosmtrek/air) package.

    ```bash
    make up
    ```

* Terminate the service

    ```bash
    make down
    ```

* Clean all the resource

    it will clean the images, volume and all the data stored in db

    ```bash
    make clean
    ```

### Code

* Unittest

    Support by [Convey](https://smartystreets.github.io/goconvey/) test framework

    ```bash
    make test
    ```

* Check potential error in your code

    ```bash
    make check
    ```

* Run this if there is any file changes under ports folder

    This feature supported by [mockgen](https://github.com/golang/mock) package, auto generated mock dependency for unittest

    **Requirement**: Need to install first by:

    ```bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

    Then you can execute command:

    ```bash
    make mock file=<filename>
    ```

     **Filename must to be exist under the ports folder**

### Db Migration

This feature supported by [atlas](https://atlasgo.io) package, auto generated mock dependency for unittest

**Requirement**: Need to install first, you can choose which install method you want in their website.

* Init Database

    ```bash
    make migrate-init
    ```

* check migrate status

    check the migrate file is staged or applied

    ```bash
    make migrate-status
    ```

* If you want make any schema changes in db, follow the step below

    1. Edit the file **schema.hcl**. You can refer the syntax from [HCL Schema](https://atlasgo.io/atlas-schema/hcl)
    2. Run this command to commit the changes to the **migrations** folder.

        ```bash
        make migrate-commit=<commit name>
        ```

    3. Run this command will apply the changes to your db

        ```bash
        make migrate-up
        ```

* If you want to revert the last changes in db, follow the step below

    1. Run this command will revert your db data to the previous version.

        ```bash
        make migrate-down
        ```

    2. Go to migrations folder to remove specify version file you roll back. Can be checked by **migrate status**.

* Clean all the db's data and schema

    It will also remove the **migrations** folder. If you want to recover your data, please run the **Init Database** command

    ```bash
    make migrate-clean
    ```

* Reverting all schema changes

    Will recover the status that defined in **migrations** folder

    ```bash
    make migrate-reload
    ```

## Notice

If you wnat to change the db config (database, port, password), you need to edit **/env/local-docker/rdb.yaml**, **docker-compose.yaml** file and **Makefile**. The other config changes, your check point will be in **cmd/config.go** file. About the online environment config including uat, product ...etc, we prefer you to store it in credentials services.
