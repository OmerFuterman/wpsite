<!--
*** Thanks for checking out my project. If you have a question or suggestion,
*** please email me at omerfu@gmail.com, or fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again!
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** github_username, repo_name, twitter_handle, email, project_title, project_description
-->




<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/OmerFuterman/wpsite">
    <img src="/resources/images/go-logo.png" alt="Logo" width="auto" height="80">
  </a>

  <h3 align="center">Cool People API</h3>

  <p align="center">
    An API to interact with a database of cool people. Worried you aren't on the database? Add yourself now!
    <br />
    <a href="https://github.com/OmerFuterman/wpsite"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/OmerFuterman/wpsite/issues">Report Bug</a>
    ·
    <a href="https://github.com/OmerFuterman/wpsite/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

As a web developer with a strong passion for the back-end, Golang is exciting to me, this project was meant for me to learn, develop, and prove my skills with Golang. Have a feature you'd like to see me add here? Email me at: omerfu@gmail.com, I love finding new areas to expand my strengths!


### Built With

* Go



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Golang
  ```sh
  guide to install Golang on your personal device: https://golang.org/doc/install
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/OmerFuterman/wpsite.git
   ```
2. Install gorilla mux and necessary packages
   ```sh
   go get -u github.com/gorilla/mux
   go get github.com/subosito/env
   go get github.com/gorilla/handlers
   go get github.com/lib/pq
   ```
3. If you recieved an error on step #2 do the following
   ```sh
   go mod init your_directory_name
   go mod tidy
   now finish step number 2 again
   ```



<!-- USAGE EXAMPLES -->
## Usage

Please keep in mind that if you are running a local copy you will have to start up your server to get a response. Do this with the following command:
  go run main.go (make sure you are in the project directory)
  
If you are spinning up this project on a local device your url for the API will be: http://localhost:8080/<br>
If you are sending requests to my public API the url will be: https://wpsite-e2xqf.ondigitalocean.app/


### List of routes

#### GET /people?limit={limit}&offset={offset}

This route will return a list of all records in the database

{limit} will limit the amount of records to the amount given

{offset} will skip over the amount of records given before giving back responses


#### GET /search?q={name}

This route will query the database by name of record

{name} will be the name to search


#### POST /add

This route will add a new record to the database

This route expects the following information in the body:
<ul>
  <li><strong>description</strong> string (a description of the person)</li>
  <li><strong>gender</strong> string (the gender of the person) Accepts: male, female, other, or prefer not to say)</li>
  <li><strong>coollevel</strong> boolean (how cool the person is) Accepts: true or false)</li>
  <li><strong>name</strong> string (the name of the person)</li>
</ul>


#### PUT /update

This route will update an exitsting record on the database

This route expects the following information in the body:
<ul>
  <li><strong>id</strong> int (the id of the record being altered) Accepts: any positive int</li>
  <li><strong>descrption</strong> string (a description of the person)</li>
  <li><strong>gender</strong> string (the gender of the person being altered) Accepts: male, female, other, or prefer not to say)</li>
  <li><strong>coollevel</strong> boolean (how cool the person is) Accepts: true or false)</li>
  <li><strong>name</strong> string (the name of the person)</li>
</ul>


#### DELETE /remove?id={id}

This route will delete one record from the database, yes, that means that you are currently able to delete your friend if you so choose, I am not liable to arguments caused by this

{id} will be the id of the record to be deleted in the database



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- CONTACT -->
## Contact

Omer Futerman - omerfu@gmail.com - 949-400-9703 - Website: https://rimomir.com

Project Link: https://github.com/OmerFuterman/wpsite/

