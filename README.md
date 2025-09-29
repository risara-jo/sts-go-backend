# Smart Transit System - Backend API

The official backend API service for the Smart Transit System - an advanced bus booking and fleet management platform built on WSO2 Choreo Platform.

## System Overview

The Smart Transit System is a comprehensive platform that includes:

- **Passenger Mobile App** - Advanced booking with lounge integration and real-time tracking
- **Bus Owner/Fleet Management App** - Comprehensive fleet and route management with analytics  
- **Driver/Conductor Mobile App** - Trip operations, QR scanning, and emergency reporting
- **Lounge Management App** - Premium lounge services with route-based discovery
- **Admin Web Dashboard** - Advanced system monitoring and comprehensive analytics

## Backend Features

- ✅ Health check and monitoring endpoints (`/health`, `/healthz`, `/ready`, `/health/db`)
- ✅ PostgreSQL database connection with connection pooling optimization
- ✅ Asgardeo (WSO2 Identity Server) authentication integration
- ✅ WSO2 Choreo Platform deployment ready
- ✅ Environment-based configuration for multiple deployment environments
- ✅ Production-ready logging and error handling

## Architecture

- **Platform**: WSO2 Choreo (Cloud-native platform)
- **Database**: PostgreSQL with Supabase (connection pooling optimized)
- **Authentication**: Asgardeo OAuth2/OpenID Connect
- **Architecture**: Monolith with modular design for simplified deployment and maintenance

## Quick Start

### Local Development

```bash
# Navigate to the project
cd minimal-backend

# Install dependencies
go mod tidy

# Configure environment variables in .env file
DATABASE_URL="postgresql://username:password@host:port/database"
PORT="8080"

# Run the application
go run main.go
```

### Docker Deployment

```bash
# Build the Docker image
docker build -t smart-transit-backend .

# Run the container
docker run -p 8080:8080 -e DATABASE_URL="your_supabase_url" smart-transit-backend
```

## API Endpoints

| Endpoint | Description | Use Case |
|----------|-------------|----------|
| `GET /health` | Basic service health check | Load balancer health check |
| `GET /healthz` | Container health check | Kubernetes liveness probe |
| `GET /ready` | Readiness probe with database connectivity | Kubernetes readiness probe |
| `GET /health/db` | Database connectivity and pool status | Database monitoring |

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | `""` |

## Choreo Deployment

This project is configured for WSO2 Choreo Platform:

1. **Dockerfile**: Choreo-compliant with user ID 10001
2. **Health Checks**: Multiple endpoints for different monitoring needs
3. **Configuration**: Environment-based setup
4. **Security**: Non-root user execution

### Deploy to Choreo

1. Push code to GitHub repository
2. Create new component in Choreo
3. Configure environment variables
4. Deploy and test endpoints

## Project Structure

```
minimal-backend/
├── .choreo/
│   └── component.yaml      # Choreo configuration
├── config/
│   └── config.go          # Configuration management
├── models/
│   └── database.go        # Database connection
├── main.go                # Application entry point
├── Dockerfile             # Container configuration
├── go.mod                 # Go modules
└── README.md             # This file
```

## License

Private project for internal use.
