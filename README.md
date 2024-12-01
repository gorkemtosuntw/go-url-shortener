# URL Shortener Service

A modern URL shortening service built with Go, featuring PostgreSQL storage, click tracking, and Docker support.

## Features
- URL shortening with unique codes
- URL validation
- Click tracking
- Access timestamp logging
- PostgreSQL database storage
- Docker and Docker Compose support
- Production-ready configuration
- Environment-based configuration
- Clean architecture implementation

## Requirements
- Go 1.21+
- PostgreSQL
- Docker and Docker Compose (optional)

## Installation

### Local Development
1. Clone the repository
2. Install dependencies:
```bash
go mod download
```
3. Configure your PostgreSQL database
4. Update environment variables or config.go with your settings
5. Run the application:
```bash
go run .
```

### Docker Development
1. Clone the repository
2. Run with Docker Compose:
```bash
docker-compose up --build
```

### Production Deployment
1. Configure your domain and SSL certificates
2. Set up your production environment variables
3. Run with production configuration:
```bash
# Create secrets directory
mkdir -p secrets
# Set database password
echo "your-secure-password" > secrets/db_password.txt
# Start services
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

## API Usage

### Shorten URL
```bash
POST /shorten
Content-Type: application/json

{
    "original_url": "https://example.com"
}
```

Response:
```json
{
    "short_url": "http://yourdomain.com/abc123"
}
```

### Access Shortened URL
```bash
GET /{short_code}
```
Automatically redirects to the original URL and records the click.

## Project Structure
```
.
├── internal/
│   ├── config/      # Configuration management
│   ├── handler/     # HTTP handlers
│   ├── model/       # Data models
│   ├── repository/  # Database operations
│   └── service/     # Business logic
├── deployment/      # Deployment configurations
├── docker-compose.yml
├── docker-compose.prod.yml
├── Dockerfile
└── main.go
```

## Architecture
The project follows clean architecture principles:
- **Models**: Core business objects
- **Repository**: Data access layer
- **Service**: Business logic layer
- **Handler**: HTTP interface layer
- **Config**: Configuration management

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License
MIT License
