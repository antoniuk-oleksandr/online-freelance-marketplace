#  Online Freelance Marketplace

**Online Freelance Marketplace** is a full-stack web application that connects freelancers with clients. It enables users to register, browse services, place orders, chat in real-time, and manage profiles through a modern and user-friendly interface.

---

##  Features

- **User Authentication**: Secure login with Google OAuth.
- **Freelancer Profiles**: Customizable user pages with avatars, bios, skills, reviews, and more.
- **Service Listings**: Freelancers can offer services in multiple tiers (Basic, Standard, Premium).
- **Order Management**: Clients can place, track, and manage service orders.
- **Real-Time Chat**: Communication between users via an integrated chat system.
- **Media Support**: Upload images and videos for posts, services, and messages.
- **Responsive Design**: Fully responsive UI built with modern web technologies.
- **Dockerized Architecture**: Easy to deploy with Docker and Docker Compose.

---

##  Tech Stack

**Frontend**  
- Svelte  
- TailwindCSS  
- TypeScript  

**Backend**  
- Go (Fiber)  
- PostgreSQL  
- OAuth 2.0 (Google)

**Others**  
- Docker & Docker Compose  
- Static File Server  
- REST API  

---

## Project Structure

```
online-freelance-marketplace/
├── backend/             # Go server for API logic and database interaction
├── file-server/         # Static file server for media storage
├── frontend/            # Svelte frontend for client interaction
└── docker-compose.yml   # Docker-based orchestration
```

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Node.js](https://nodejs.org/)
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

### Run Locally (Docker)

Clone the repository:

```bash
git clone https://github.com/antoniuk-oleksandr/online-freelance-marketplace.git
cd online-freelance-marketplace
```

Build and run the app:

```bash
docker-compose up --build
```

This command launches:
- the backend server
- the frontend UI
- the file server
- the PostgreSQL database

---

## Environment Variables

You need to create `.env` files for the backend and frontend.

### Backend `.env`

```
# JWT configuration
JWT_SECRET=jwt_secret_here
ENCRYPTION_KEY=encryption_key_here

# Database configuration
DB_NAME=database_name
DB_USER=database_user
DB_PASSWORD=database_password
DB_PORT=database_port
DB_HOST=database_host
DB_SSL_MODE=db_ssl_mode

# Redis configuration
REDIS_HOST=redis_host
REDIS_PORT=redis_port
REDIS_PASSWORD=redis_password

# Email configuration
EMAIL_PASSWORD=email_password
SENDER_EMAIL=sender_email

# Frontend configuration
FRONTEND_HOST=frontend_host

# Google OAuth configuration
GOOGLE_CLIENT_ID=google_client_id
GOOGLE_CLIENT_SECRET=google_client_secret
GOOGLE_REDIRECT_URI=google_redirect_uri

# Application settings
MAX_SEARCH_RESULTS=maximum_search_results_value
MAX_CONNECTIONS=maximum_connections_value
MAX_FREELANCE_BY_ID_REVIEWS=maximum_freelance_by_id_reviews_value
MAX_USER_BY_ID_REVIEWS=maximum_user_by_id_reviews_value
MAX_USER_BY_ID_SERVICES=maximum_user_by_id_services_value
SERVICE_FEE=service_fee_value

# Payment configuration
RSA_PUBLIC_KEY_PATH=rsa_public_key_path
RSA_PRIVATE_KEY_PATH=rsa_private_key_path

# File server configuration
FILE_SERVER_HOST=file_server_host

# My profile settings
MAX_MY_PROFILE_ORDER_RESULTS=maximum_my_profile_order_results_value
MAX_MY_PROFILE_SERVICE_RESULTS=maximum_my_profile_service_results_value
MAX_MY_PROFILE_REQUEST_RESULTS=maximum_my_profile_request_results_value
```

### Frontend `.env`

```
VITE_GOOGLE_CLIENT_ID: '__VITE_GOOGLE_CLIENT_ID__',
VITE_MAX_FREELANCE_BY_ID_REVIEWS: '__VITE_MAX_FREELANCE_BY_ID_REVIEWS__',
VITE_SERVICE_FEES: '__VITE_SERVICE_FEES__',
VITE_FILE_SERVER_HOST: '__VITE_FILE_SERVER_HOST__',
VITE_WEBSOCKET_HOST: '__VITE_WEBSOCKET_HOST__',
VITE_BACKEND_HOST: '__VITE_BACKEND_HOST__'
```

---

## Roadmap

- [x] Google OAuth authentication
- [x] Profile pages and editing
- [x] Orders and payment flow
- [x] Chat system
- [x] Reviews and ratings
- [*] Service creation and editing
- [ ] Admin panel for moderation
- [ ] Notifications and alerts

---

## Contributing

Contributions are welcome! Please fork the repo and submit a pull request with your changes. For major updates, open an issue first to discuss what you’d like to change.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

## Author

Created by [Oleksandr Antoniuk](https://github.com/antoniuk-oleksandr)

Feel free to reach out for questions, ideas, or collaborations.
