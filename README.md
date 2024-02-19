# Contributors

- Magali DESLOUS-PAOLI (Magali-D) <magali.paoli@tutanota.com>
- Kylian BARTHELEMY (KyDev25) <kylian.barthelemy2003@gmail.com>
- Cédric GAUTIER (cedricgautier) <gautiercedric07@gmail.com>
- Barnabé PILLIAUDIN (BarnabePILLIAUDIN) <barnabe.pilliaudin@gmail.com>
- Alix BERTHELOT (hallycsse) <alixberthelot03@gmail.com>

# Jira link

[Tracker - Kanban board](https://kylianbarthelemy.atlassian.net/jira/software/projects/AIRNEIS/boards/2)


## Introduction

_ÀIRNEIS is a Scottish company and they are resellers of furniture designed by Scottish designers._

We want to create our own e-commerce solution for an international clientele with products they already have locally.
To do this, we created an e-commerce site as well an application, and a Backoffice to manage the content.
All with a secure and maintainable payment system over time.

## Getting started

This project is configured to run with Docker through `make` command.<br>
To deploy the project in dev mode:
```
make dev
```
To deploy the project in prod mode:
```
make up
``` 
To shut down the application:
```
make down
``` 
## What does our app contain ?

### The Home Page

The home page provides an overview of the categories as well as the "Highlanders of the moment" which will contain products selected in the Backoffice.
So, this page contains all categories information.

### The Category Page

The category page is the main access point to the product catalogue of a category.
So, this page contains information about products in a single category.

### The Product Page

The product page is the most important page on which the customer can purchase a product.
So, this page contains all information about a single product.

### The Search Page

The search page is accessible everywhere from the site header.
This page allows them to research a specific product or type of product. The user also be able to sort the result by price or newness.

### The Shopping Cart Page

The shopping cart page is accessible to all users, logged or not, they can manage their shopping cart.
This page contains the list of added products and presents the total to be paid.

### The Checkout Page

The checkout page comes after the shopping cart page.
This page contains all the information about the delivery or order.
He contains several steps like:

- Delivery.
- Order.
- Payment accepted.

### The Registration Page

The registration page allows all users to register on the site or application.
This page contains mandatory information for a user before registering on the site or application.

### The Login Page

The login page allows all users who already have an account to login to the site or application.
This page therefore contains mandatory information for full access to the site or application.

### The Account Settings Page

The account settings allow a user to modify their account information and profile.
This page contains all the information of the logged in user.

### My Orders Page

My orders page allows a user to see their past orders, grouped by year.
This page contains a list of all orders with the status like:

- In progress
- Delivered
- Cancelled

### The Contact Page

The contact page is a simple page where all users can write to administrators with a simple contact form.
This page contains minimal information if a user has a problem with the site or a recommendation.

## The Backoffice

The back office allows the administrator to see all the information on a product, a category or an user.
An administrator can see, create, update, or delete an item.
It must be accessible to administrators only.
The back office is as simple and functional as possible.

It also contains a dashboard with histograms and graphs.

## Miscellaneous

### Menu

Users have different access to the burger menu if they are logged in or not.

### Design System

We use a design system with a library of  graphics component.
Because if the site evolves, it is possible that the design evolves too.

### i18n

[!image](./docs/functional/i18n_Schema.png)

We use the i18n.
i18n makes it possible to adapt computer software to the different languages, regional peculiarities, and technical requirements of a target local.
So, our site and application is delivered in multilingual.

### a11y

We use also a11y.
a11y represents web accessibility.
Is the inclusive practice of ensuring there are no barriers that prevent interaction with, or access to, websites by people with physical disabilities, situational disabilities, and socio-economic restrictions on bandwidth and speed.
So, we insist on compliance with modern accessibility standards for all applications.

## Main Roles

- Mme. Magali DESLOUS-PAOLI [ SRE ]
- Mr. Kylian BARTHELEMY [ Product Owner ]
- Mr Cédric GAUTIER [ Security Engineer ]
- Mr. Barnabé PILLIAUDIN [ Developer Full-Stack ]
- Mr. Alix BERTHELOT [ Developer Full-Stack ]
