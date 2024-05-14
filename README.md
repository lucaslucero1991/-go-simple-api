# hp-dev-jobberwocky

You have been asked to implement **Jobberwocky** , a service that works as a store for job opportunities, where companies can share open positions.

## Initial setup

Install sql3
go get -u github.com/mattn/go-sqlite3@latest
configurar el gopath
ejecutar go mod init

configuracion inicial
installar homebrew, admin de paquetes default para mac
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

luego node
brew install node

verificamos la instalacion

node -v
npm -v

con esto ya podemos correr el inicailizador y el compilador

si nos sale el error que donde no tenemos configurado el gopath:
entonces ejecutamos
export GO111MODULE=on para permitir al proyecto crear un nombre explicito a la 
creación de inicializador.

go mod init v0
go mod tidy

con esto, ya podemos correr comandos como:
go get [libreria] para instalar dependecias.

Para correr los test instalamos gomock,
nos permite crear mocks de interfaces y setear
las expectativas en nuestros test.

Lo vamos a usar para validar nuestros servicios


## 1. Create a job posting service

Implement an application that exposes an API that lets users register new job opportunities.
- The app does not need to persist info on an external database service.
- Feel free to store jobs in memory or disk (CSV, SQLite,etc).
- Choose any API style: web‐based, REST, GraphQL, etc.

## 2. Create a job searching service

Extend your application to expose another endpoint that lets users find job opportunities from the service you have already created.

## 3. Create additional sources

In addition to our internal job service, we want our job-searching service to consume data from additional job opportunity sources using this service: [jobberwocky-extra-source](https://github.com/avatureta/jobberwocky-extra-source). Find the best way to return a response that combines the results from more than one source.

## 4. Create a Job Alerts service (optional)

Update your application so that it lets candidates subscribe to an email address and be notified whenever a new job is posted. An optional search pattern can be provided as a way to filter job posts.

## FAQ

### Do I need to create a UI?

We will only assess the backend, but you can buildone if you feel like it.

### Does the app require authentication?

No, it doesn't.

### What fields should I use for each entity?

As a developer, we expect that you design the proper structure for each entity, such as the job or the subscription entities.

### Can I use an external framework?

Yes, feel free to choose any framework that suitsyour needs.

### Which programming language should I use?

You may use: C++, C#, Python, Java/Kotlin, Javascript/Node/Typescript,PHP, Ruby. If you’d prefer to use a different language, please let us know before getting started.

### In which language should I program?

English please.

