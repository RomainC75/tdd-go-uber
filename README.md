# stryker 

## install
```npm i @stryker-mutator/jest-runner```
## use 
```npx stryker run ```

# UberLike SAAS!

L'application Uber-Like qui permet de réserver un Driver instantanément !

## **User Story - Réserver une course**

En tant que **Rider**,  
Je souhaite **réserver une course** pouvant m'amener à ma destination  
De sorte à assurer une alternative efficace aux transports en commun.

Un Rider peut réserver une course à tout moment pour n'importe quelle destination.  
La réservation est confirmée une fois qu'un Driver libre est assigné par le système.

---

## **Règles de Gestion de la Tarification**

### **1. Prix de Base par Direction**
- **Paris vers l'extérieur** : 20€.
- **Paris intra-muros (trajet à l'intérieur de Paris)** : 30€.
- **Extérieur vers Paris** : 10€.
- **Extérieur vers l'extérieur** : 50€.

### **2. Frais Kilométriques**
- Si forfait Basic, un supplément de **0,5€ par kilomètre** parcouru est ajouté au prix de base du trajet.
- Si forfait Premium, un supplément de **0.5€ par kilomètre** parcouru est ajouté au prix de base du trajet **à partir de 5km parcourus**.
- Les frais kilométriques sont calculés uniquement sur la distance parcourue,  
  sans tenir compte des bouchons ou autres ralentissements.

### **3. Changement de Trajet**
- Tout **changement de trajet** en cours de route entraîne un **surcoût de 10€**,  
  quelle que soit la formule initiale.

### **4. Option UberX**
- **Description** : UberX permet d'accéder à des véhicules de type berline haut de gamme.
- **Supplément** : Un supplément de **10€** est ajouté pour bénéficier du mode UberX.
- **Condition de Distance** : Le mode UberX est disponible uniquement pour les trajets de **3 kilomètres ou plus**.  
  Pour les trajets inférieurs à cette distance, l'option est refusée.
- **Offre Anniversaire** : Si le jour de la course correspond à l'anniversaire du Rider,  
  le supplément de 10€ pour l'option UberX est **offert** pour les trajets de 3 kilomètres ou plus.

### **5. Offre de Bienvenue**
- Une remise de **X €** est offerte aux Riders durant leur première année d'utilisation du service.

### **6. Restriction sur les Réservations**
- Un Rider **ne peut réserver qu'une seule course à la fois**.
- Il est impossible de réserver une nouvelle course tant que la course en cours n'est pas terminée ou annulée.

---

## **Règles d'assignation d'un Driver**
- Le système recherche un Driver disponible dans un rayon de **5 km** autour du Rider.
- Un Driver ne peut être assigné qu'à **une seule réservation** à la fois.

---

## **User Story - Lister toutes mes courses passées**
En tant que **Rider**,  
Je souhaite **lister tout l'historique de mes courses avec mention des Drivers respectifs**  
De sorte à pouvoir me figurer la fréquence de mon utilisation.

<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://nestjs.com/img/logo-small.svg" width="120" alt="Nest Logo" /></a>
</p>

[circleci-image]: https://img.shields.io/circleci/build/github/nestjs/nest/master?token=abc123def456
[circleci-url]: https://circleci.com/gh/nestjs/nest

  <p align="center">A progressive <a href="http://nodejs.org" target="_blank">Node.js</a> framework for building efficient and scalable server-side applications.</p>
    <p align="center">
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/v/@nestjs/core.svg" alt="NPM Version" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/l/@nestjs/core.svg" alt="Package License" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/dm/@nestjs/common.svg" alt="NPM Downloads" /></a>
<a href="https://circleci.com/gh/nestjs/nest" target="_blank"><img src="https://img.shields.io/circleci/build/github/nestjs/nest/master" alt="CircleCI" /></a>
<a href="https://discord.gg/G7Qnnhy" target="_blank"><img src="https://img.shields.io/badge/discord-online-brightgreen.svg" alt="Discord"/></a>
<a href="https://opencollective.com/nest#backer" target="_blank"><img src="https://opencollective.com/nest/backers/badge.svg" alt="Backers on Open Collective" /></a>
<a href="https://opencollective.com/nest#sponsor" target="_blank"><img src="https://opencollective.com/nest/sponsors/badge.svg" alt="Sponsors on Open Collective" /></a>
  <a href="https://paypal.me/kamilmysliwiec" target="_blank"><img src="https://img.shields.io/badge/Donate-PayPal-ff3f59.svg" alt="Donate us"/></a>
    <a href="https://opencollective.com/nest#sponsor"  target="_blank"><img src="https://img.shields.io/badge/Support%20us-Open%20Collective-41B883.svg" alt="Support us"></a>
  <a href="https://twitter.com/nestframework" target="_blank"><img src="https://img.shields.io/twitter/follow/nestframework.svg?style=social&label=Follow" alt="Follow us on Twitter"></a>
</p>
  <!--[![Backers on Open Collective](https://opencollective.com/nest/backers/badge.svg)](https://opencollective.com/nest#backer)
  [![Sponsors on Open Collective](https://opencollective.com/nest/sponsors/badge.svg)](https://opencollective.com/nest#sponsor)-->

## Description

[Nest](https://github.com/nestjs/nest) framework TypeScript starter repository.

## Project setup

```bash
$ npm install
```

## Compile and run the project

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## Run tests

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```

## Deployment

When you're ready to deploy your NestJS application to production, there are some key steps you can take to ensure it runs as efficiently as possible. Check out the [deployment documentation](https://docs.nestjs.com/deployment) for more information.

If you are looking for a cloud-based platform to deploy your NestJS application, check out [Mau](https://mau.nestjs.com), our official platform for deploying NestJS applications on AWS. Mau makes deployment straightforward and fast, requiring just a few simple steps:

```bash
$ npm install -g @nestjs/mau
$ mau deploy
```

With Mau, you can deploy your application in just a few clicks, allowing you to focus on building features rather than managing infrastructure.

## Resources

Check out a few resources that may come in handy when working with NestJS:

- Visit the [NestJS Documentation](https://docs.nestjs.com) to learn more about the framework.
- For questions and support, please visit our [Discord channel](https://discord.gg/G7Qnnhy).
- To dive deeper and get more hands-on experience, check out our official video [courses](https://courses.nestjs.com/).
- Deploy your application to AWS with the help of [NestJS Mau](https://mau.nestjs.com) in just a few clicks.
- Visualize your application graph and interact with the NestJS application in real-time using [NestJS Devtools](https://devtools.nestjs.com).
- Need help with your project (part-time to full-time)? Check out our official [enterprise support](https://enterprise.nestjs.com).
- To stay in the loop and get updates, follow us on [X](https://x.com/nestframework) and [LinkedIn](https://linkedin.com/company/nestjs).
- Looking for a job, or have a job to offer? Check out our official [Jobs board](https://jobs.nestjs.com).

## Support

Nest is an MIT-licensed open source project. It can grow thanks to the sponsors and support by the amazing backers. If you'd like to join them, please [read more here](https://docs.nestjs.com/support).

## Stay in touch

- Author - [Kamil Myśliwiec](https://twitter.com/kammysliwiec)
- Website - [https://nestjs.com](https://nestjs.com/)
- Twitter - [@nestframework](https://twitter.com/nestframework)

## License

Nest is [MIT licensed](https://github.com/nestjs/nest/blob/master/LICENSE).
