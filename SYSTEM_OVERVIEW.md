# Smart Transit System - Backend API Service

## Overview

This is the official backend API service for the **Smart Transit System** - an advanced bus booking and fleet management platform built on WSO2 Choreo Platform.

## System Architecture

- **Platform**: WSO2 Choreo (Cloud-native serverless platform)  
- **Database**: PostgreSQL with Supabase (optimized connection pooling)
- **Authentication**: Asgardeo OAuth2/OpenID Connect integration
- **Architecture**: Monolith design for simplified deployment and maintenance

## Smart Transit System Components

### 1. Passenger Mobile App

- Advanced booking system with lounge integration
- Real-time bus tracking and notifications  
- Digital wallet with escrow payment system
- Route discovery and schedule management

### 2. Bus Owner/Fleet Management App  

- Comprehensive fleet management and analytics
- Route planning and optimization
- Staff management and scheduling
- Revenue and performance dashboards

### 3. Driver/Conductor Mobile App

- Trip operations and passenger management
- QR code scanning for ticket validation
- Emergency reporting and communication
- Performance tracking and trip logs

### 4. Lounge Management App

- Premium lounge services management
- Route-based discovery and booking integration
- Amenity and service management
- Customer experience optimization

### 5. Admin Web Dashboard

- Advanced system monitoring and analytics
- Government pricing compliance management
- Comprehensive reporting and insights
- User and system administration

## Backend Features

- ✅ Health monitoring and readiness probes
- ✅ PostgreSQL database with intelligent connection pooling
- ✅ WSO2 Choreo Platform deployment ready
- ✅ Production-grade logging and error handling
- ✅ Container orchestration support (Kubernetes/Docker)

## API Endpoints

| Endpoint | Purpose | Description |
|----------|---------|-------------|
| `/health` | System Health | Basic service health check with version info |
| `/healthz` | Container Health | Kubernetes/Choreo liveness probe |
| `/ready` | Readiness Check | Database connectivity and service readiness |
| `/health/db` | Database Health | PostgreSQL connection and pool status |

## Deployment

The service is designed for deployment on WSO2 Choreo Platform with:

- Automatic scaling and load balancing
- Integrated monitoring and logging
- Security compliance and access control
- CI/CD pipeline integration

## Technology Stack

- **Language**: Go 1.24
- **Framework**: Gin (HTTP web framework)
- **ORM**: GORM (Go Object-Relational Mapping)
- **Database**: PostgreSQL via Supabase
- **Platform**: WSO2 Choreo
- **Authentication**: Asgardeo
- **Containerization**: Docker with multi-stage builds

---

*Part of the Smart Transit System ecosystem by BusLounge*
