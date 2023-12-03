# Full-Stack Project: Go Backend and JavaScript Frontend

Welcome to my full-stack project! The backend is crafted in Go and communicates with a MySQL database, while the frontend is developed using JavaScript, HTML, and CSS. The project includes a user-friendly login page that securely stores user data in the database. User credentials are encrypted using Go's bcrypt library for robust security.

## Getting Started

To deploy the application effortlessly, I've containerized it using Docker. Utilize the provided `docker-compose.yml` and `Dockerfile` to kickstart the project:

```bash
docker-compose up
```
To stop the running container use

```bash
docker-compose down
```
this will also remove image and running container.

Once the app is up and runnig go to

```php
http://localhost:8080
```
to access th UI.
