![image](https://github.com/AlphaDecodeX/nomise-ragnarok/assets/91979252/3cbd7409-f917-40dc-af1d-9e65e72cae0c)

# NoMise Ragnarok:- Order Checkout and Payment Service

## Tentative Tech Stack

</br>

![Golang](https://img.shields.io/badge/-Golang-00ADD8?logo=go&logoColor=white&style=flat-square) ![Gin](https://img.shields.io/badge/-Gin-00ADD8?logo=go&logoColor=white&style=flat-square) ![SQL](https://img.shields.io/badge/-SQL-FFD700?logo=mysql&logoColor=white&style=flat-square) ![Razorpay](https://img.shields.io/badge/-Razorpay-FF5722?logo=razorpay&logoColor=white&style=flat-square) ![Apache Kafka](https://img.shields.io/badge/-Apache_Kafka-231F20?logo=apache-kafka&logoColor=white&style=flat-square)



## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This is our Backend service where We'll get items from the cart that user has added and then we'll process the information and after getting the payment we'll send a notification to the user using notification service.

## Features

- Order to be processed and stored in the DB as order history
- Payment to receive in NoMise cash pool first
- Able to do the Payment from Razorpay API
- Able to publish an event to Apache Kafka to send the notification to the user (async)

## Usage

This service will be used maily for checkout, payment collection, order processing and storing relative order information

## Contributing

This project has been started while keeping beginners in focus. You will receive a detailed blog each time on the progress of the project and how it is going. Blogs will be much detailed on how and what are the changes made. It would be beginner-friendly, keeping in mind learning open-source and enterprise-level software development.

- Contribute to Other Services:-
  - [NoMise Frontent](https://github.com/AlphaDecodeX/NoMise_Store)
  - [Inventory Service](https://github.com/AlphaDecodeX/nomise_inventory)

## License

This project is licensed under the Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0) license.

This license allows others to share and adapt the work for non-commercial purposes, as long as they give appropriate credit and license their new creations under the same terms. For more information, please see the [License.md](LICENSE.md).
