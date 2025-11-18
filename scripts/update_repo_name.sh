#!/bin/bash

# Script to update repository name in all documentation
# Usage: ./update_repo_name.sh [GITHUB_USERNAME]

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Repository Name Update Script${NC}"
echo -e "${BLUE}========================================${NC}"
echo

# Get GitHub username
if [ -z "$1" ]; then
    echo -e "${YELLOW}Enter your GitHub username:${NC}"
    read GITHUB_USER
else
    GITHUB_USER="$1"
fi

if [ -z "$GITHUB_USER" ]; then
    echo "Error: GitHub username cannot be empty"
    exit 1
fi

echo -e "${GREEN}✓${NC} Using GitHub username: ${BLUE}$GITHUB_USER${NC}"
echo -e "${GREEN}✓${NC} Repository name: ${BLUE}rollgophish${NC}"
echo

# Confirm
echo -e "${YELLOW}This will update all documentation files. Continue? (y/n)${NC}"
read -r response
if [[ ! "$response" =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 0
fi

echo

# Navigate to repository root
cd "$(dirname "$0")/.."

# Files to update
FILES=(
    "README.md"
    "COMPLETE_SETUP_GUIDE.md"
    "QUICKSTART.md"
    "REPOSITORY_SETUP.md"
)

# Update each file
for file in "${FILES[@]}"; do
    if [ -f "$file" ]; then
        echo -e "${BLUE}Updating${NC} $file..."

        # Create backup
        cp "$file" "$file.bak"

        # Replace YOUR_USERNAME with actual username
        sed -i "s|YOUR_USERNAME|$GITHUB_USER|g" "$file"

        # Replace /gophish.git with /rollgophish.git
        sed -i "s|/gophish\\.git|/rollgophish.git|g" "$file"

        # Replace /gophish with /rollgophish in URLs (but not in titles)
        sed -i "s|github\\.com/[^/]*/gophish\"|github.com/$GITHUB_USER/rollgophish\"|g" "$file"

        # Replace generic gophish repo references
        sed -i "s|git clone https://github\\.com/[^/]*/gophish|git clone https://github.com/$GITHUB_USER/rollgophish|g" "$file"

        echo -e "${GREEN}  ✓ Updated${NC}"
    else
        echo -e "${YELLOW}  ⚠ File not found: $file${NC}"
    fi
done

echo
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}Update Complete!${NC}"
echo -e "${GREEN}========================================${NC}"
echo
echo "Updated files:"
for file in "${FILES[@]}"; do
    if [ -f "$file" ]; then
        echo "  • $file"
    fi
done

echo
echo "Backup files created with .bak extension"
echo
echo "Next steps:"
echo "  1. Review changes: git diff"
echo "  2. If satisfied, commit: git add . && git commit -m 'Update repository name to rollgophish'"
echo "  3. Remove backups: rm *.bak"
echo "  4. Follow REPOSITORY_SETUP.md to create GitHub repo"
echo
echo "Repository URL will be:"
echo -e "  ${BLUE}https://github.com/$GITHUB_USER/rollgophish${NC}"
echo
