#!/bin/bash

# Check if running in GitHub Actions
if [ -n "$GITHUB_ACTIONS" ]; then
    # GitHub Actions environment
    echo "DATABASE_URL=${{ secrets.DATABASE_URL }}" > app.env
    echo "API_KEY=${{ secrets.API_KEY }}" >> app.env
    # Add more secrets as needed
else
    # Local environment
    # You can have a default app.env file for local development
    # or implement logic to populate it based on your local settings
    cp default_app.env app.env
fi
