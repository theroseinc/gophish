#!/bin/bash

# Gophish Template Quick Start Script
# Sets up templates and configurations for educational lab use

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m' # No Color

print_header() {
    echo -e "\n${BOLD}${BLUE}========================================${NC}"
    echo -e "${BOLD}${BLUE}$1${NC}"
    echo -e "${BOLD}${BLUE}========================================${NC}\n"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
}

print_info() {
    echo -e "${CYAN}ℹ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

# Get script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"

print_header "Gophish Template Quick Start"

print_warning "Educational Lab Use Only!"
print_info "This setup is for learning defensive security in isolated environments"
echo ""

# Check if Python 3 is installed
if command -v python3 &> /dev/null; then
    print_success "Python 3 found"
else
    print_error "Python 3 is required but not installed"
    exit 1
fi

# Check if required Python packages are installed
print_info "Checking Python dependencies..."
python3 -c "import requests" 2>/dev/null || {
    print_warning "Installing requests library..."
    pip3 install requests --quiet
}

# Generate phishlets
print_header "Step 1: Generating Evilginx Phishlets"
python3 "$SCRIPT_DIR/generate_phishlets.py"

# Ask if user wants to install templates to Gophish
print_header "Step 2: Install Templates to Gophish (Optional)"
print_info "Do you want to install templates to a running Gophish instance?"
read -p "$(echo -e ${CYAN}Enter y/n: ${NC})" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    print_info "Starting template installation..."
    python3 "$SCRIPT_DIR/setup_templates.py"
else
    print_info "Skipping template installation"
    print_info "You can run 'python3 scripts/setup_templates.py' later"
fi

# Create README for templates
print_header "Step 3: Creating Documentation"

cat > "$PROJECT_DIR/templates/README.md" << 'EOF'
# Gophish Templates for Educational Lab Use

This directory contains professional phishing templates for use in isolated educational environments.

## ⚠️ EDUCATIONAL USE ONLY

These templates are designed for:
- **Personal home labs** for learning defensive security
- **Isolated environments** with no external connectivity
- **Security research** and understanding attack techniques
- **Defensive training** to recognize phishing attempts

**NEVER use against real-world targets without explicit written authorization!**

## Directory Structure

```
templates/
├── email/                          # Email templates
│   ├── microsoft365-security-alert.html
│   ├── gmail-security-alert.html
│   └── stripe-payment-failed.html
├── landing-pages/                  # Landing page templates
│   ├── microsoft365-login.html
│   ├── gmail-login.html
│   └── stripe-payment-update.html
└── evilginx-configs/               # Evilginx phishlet configs
    ├── microsoft365.yaml
    ├── gmail.yaml
    └── stripe.yaml
```

## Quick Start

### 1. Install Templates to Gophish

```bash
cd scripts
python3 setup_templates.py
```

You'll need your Gophish API key (found in Settings > Account Settings).

### 2. Use with Evilginx (Optional)

Copy phishlet configurations to Evilginx:

```bash
cp templates/evilginx-configs/*.yaml /path/to/evilginx/phishlets/
```

In Evilginx:
```
phishlets hostname microsoft365 your-domain.com
phishlets enable microsoft365
lures create microsoft365
lures get-url 0
```

### 3. Create a Campaign in Gophish

1. **Create Sending Profile**: Configure SMTP settings
2. **Create User Group**: Add test email addresses
3. **Select Template**: Choose from installed templates
4. **Select Landing Page**: Choose matching landing page
5. **Launch Campaign**: Start your educational test

## Template Details

### Microsoft 365 Security Alert
- **Email**: Security alert about unusual sign-in
- **Landing**: Two-stage login (email then password)
- **Captures**: Email, password
- **Evilginx**: Session token capture enabled

### Gmail Security Alert
- **Email**: Google security notification
- **Landing**: Material Design login flow
- **Captures**: Email, password
- **Evilginx**: Multi-cookie session capture

### Stripe Payment Failed
- **Email**: Payment failure notification
- **Landing**: Full payment form with billing
- **Captures**: Email, phone, full card details, billing address
- **Evilginx**: Session and payment data capture

## Advanced Features

### Email Templates
- Responsive HTML design
- Gophish template variables: `{{.FirstName}}`, `{{.Email}}`, `{{.URL}}`
- Professional styling matching real services
- Current date injection: `{{.CurrentDate}}`

### Landing Pages
- Multi-stage credential capture
- Client-side form validation
- Realistic error messages
- Auto-submit to Gophish

### Evilginx Integration
- Session token capture
- Cookie hijacking simulation
- Reverse proxy configuration
- Custom JavaScript injection

## Security Research Use

These templates demonstrate:
- **Visual deception techniques**
- **Email spoofing indicators**
- **URL obfuscation methods**
- **Credential harvesting workflows**
- **Session hijacking mechanics**

Use this knowledge to:
- Train users to recognize phishing
- Implement better email filters
- Develop detection mechanisms
- Understand attacker TTPs

## Defensive Indicators

When analyzing these templates, look for:
1. **URL mismatches** in email links vs. real domains
2. **Sender address** spoofing in email headers
3. **SSL certificate** warnings on landing pages
4. **Form action** targets in HTML source
5. **JavaScript** credential stealing code

## Lab Setup Recommendations

1. **Isolated Network**: Use VM network or air-gapped environment
2. **Test Accounts**: Create dedicated test email accounts
3. **Local SMTP**: Set up local mail server (no external sending)
4. **Documentation**: Keep logs of all activities
5. **Cleanup**: Remove templates after training complete

## Troubleshooting

### Templates not importing
- Check Gophish API key is correct
- Verify Gophish is running on localhost:3333
- Check SSL certificate settings

### Landing pages not capturing
- Ensure "Capture Credentials" is enabled in page settings
- Verify form fields match Gophish expectations
- Check browser console for JavaScript errors

### Evilginx not proxying
- Verify phishlet hostname configuration
- Check DNS settings point to Evilginx server
- Review Evilginx logs for errors

## Legal and Ethical Notice

**These templates are powerful tools that must be used responsibly:**

✅ **Acceptable Use:**
- Personal isolated lab environments
- Authorized security awareness training
- Defensive security research
- CTF competitions with permission

❌ **Prohibited Use:**
- Unauthorized access to systems
- Phishing real users without consent
- Credential theft
- Any illegal activity

**You are responsible for ensuring your use complies with all applicable laws and regulations.**

## Resources

- [Gophish Documentation](https://docs.getgophish.com/)
- [Evilginx Documentation](https://help.evilginx.com/)
- [OWASP Phishing Resources](https://owasp.org/www-community/attacks/Phishing)

---

**Remember: Use for education and defense, never for harm.**
EOF

print_success "Created templates/README.md"

# Create main documentation
cat > "$PROJECT_DIR/TEMPLATE_SETUP_GUIDE.md" << 'EOF'
# Template Setup Guide - Educational Lab

Complete guide for setting up Gophish with advanced templates and Evilginx integration.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Template Installation](#template-installation)
4. [Evilginx Integration](#evilginx-integration)
5. [Creating Your First Campaign](#creating-your-first-campaign)
6. [Advanced Features](#advanced-features)
7. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Software
- **Gophish** (latest version)
- **Python 3.6+**
- **Evilginx 2.3+** (optional, for session hijacking)
- **Go 1.16+** (if building from source)

### Python Dependencies
```bash
pip3 install requests
```

## Installation

### 1. Clone/Download Templates

Templates are located in:
- `templates/email/` - Email templates
- `templates/landing-pages/` - Landing pages
- `templates/evilginx-configs/` - Evilginx phishlets

### 2. Start Gophish

```bash
# Build and run Gophish
go build
./gophish
```

Access at: `https://localhost:3333`

Default credentials:
- Username: `admin`
- Password: (shown in terminal on first run)

### 3. Get API Key

1. Log in to Gophish UI
2. Navigate to: Settings → Account Settings
3. Copy your API key
4. Save for template installation

## Template Installation

### Automated Installation

```bash
cd scripts
python3 setup_templates.py
```

When prompted, enter your Gophish API key.

### Manual Installation

#### Email Templates

1. In Gophish UI: Email Templates → New Template
2. Name: `Microsoft 365 - Security Alert`
3. Subject: `Unusual sign-in activity detected`
4. HTML: Copy from `templates/email/microsoft365-security-alert.html`
5. Click "Save Template"

Repeat for Gmail and Stripe templates.

#### Landing Pages

1. In Gophish UI: Landing Pages → New Page
2. Name: `Microsoft 365 - Login Page`
3. HTML: Copy from `templates/landing-pages/microsoft365-login.html`
4. Enable: ✓ Capture Credentials
5. Enable: ✓ Capture Passwords
6. Click "Save Page"

Repeat for Gmail and Stripe pages.

## Evilginx Integration

### Setup Evilginx

```bash
# Install Evilginx
git clone https://github.com/kgretzky/evilginx2.git
cd evilginx2
make
sudo ./bin/evilginx -p ./phishlets
```

### Install Phishlets

```bash
# Generate phishlets
cd /path/to/gophish/scripts
python3 generate_phishlets.py

# Copy to Evilginx
cp ../templates/evilginx-configs/*.yaml /path/to/evilginx2/phishlets/
```

### Configure Phishlet

```
# In Evilginx console
config domain your-domain.com
config ip YOUR_SERVER_IP

phishlets hostname microsoft365 login.your-domain.com
phishlets enable microsoft365

lures create microsoft365
lures edit 0 redirect_url https://office.com
lures get-url 0
```

### Integrate with Gophish

Use the Evilginx lure URL as the `{{.URL}}` in your Gophish campaign.

## Creating Your First Campaign

### 1. Create Sending Profile

Settings → Sending Profiles → New Profile

```
Name: Test SMTP
From: Security Team <security@company.com>
Host: localhost:25
Username: (if required)
Password: (if required)
```

### 2. Create User Group

Users & Groups → New Group

```
Name: Test Users
Add user:
  First Name: Test
  Last Name: User
  Email: test@yourdomain.com
  Position: Employee
```

### 3. Create Campaign

Campaigns → New Campaign

```
Name: Microsoft 365 Security Test
Email Template: Microsoft 365 - Security Alert
Landing Page: Microsoft 365 - Login Page
URL: http://localhost/ (or your Evilginx URL)
Sending Profile: Test SMTP
Groups: Test Users
```

### 4. Launch Campaign

Click "Launch Campaign" and monitor results in the dashboard.

## Advanced Features

### Custom Email Variables

Available template variables:
- `{{.FirstName}}` - Recipient first name
- `{{.LastName}}` - Recipient last name
- `{{.Email}}` - Recipient email
- `{{.Position}}` - Recipient position
- `{{.URL}}` - Phishing link with tracking
- `{{.Tracker}}` - Tracking pixel
- `{{.From}}` - Sender address
- `{{.RId}}` - Result ID

### Session Capture with Evilginx

The Evilginx integration captures:
- **Session cookies** - Authentication tokens
- **Local storage** - OAuth tokens
- **API calls** - Real-time requests
- **2FA bypass** - Session replay

View captured sessions:
```
sessions
```

Export session:
```
sessions <id>
```

### Webhook Integration

Configure webhooks to receive real-time notifications:

```go
// In Gophish settings
Webhook URL: https://your-server.com/webhook
Secret: your-secret-key
```

### API Usage

```python
import requests

API_KEY = "your-api-key"
BASE_URL = "https://localhost:3333/api"

headers = {
    "Authorization": f"Bearer {API_KEY}"
}

# Get campaign results
response = requests.get(
    f"{BASE_URL}/campaigns/1/results",
    headers=headers,
    verify=False
)

results = response.json()
for result in results:
    if result['status'] == 'Submitted Data':
        print(f"Captured: {result}")
```

## Troubleshooting

### Templates Not Importing

**Problem**: Script fails with API error

**Solution**:
- Verify Gophish is running
- Check API key is correct
- Ensure using HTTPS (not HTTP)
- Disable SSL verification for self-signed certs

### Landing Page Not Capturing Credentials

**Problem**: No data captured when submitting forms

**Solution**:
- Ensure "Capture Credentials" is enabled
- Check "Capture Passwords" is enabled
- Verify form fields have `name` attributes
- Check browser console for errors

### Evilginx Not Proxying

**Problem**: Phishlet not loading correctly

**Solution**:
- Verify DNS points to Evilginx server
- Check phishlet is enabled: `phishlets`
- Review hostname configuration
- Check firewall allows ports 80/443

### Email Not Sending

**Problem**: Campaign stuck on "Sending"

**Solution**:
- Verify SMTP settings in sending profile
- Check SMTP server is accessible
- Review Gophish logs: `tail -f gophish.log`
- Test SMTP with: `telnet smtp-server 25`

### Sessions Not Capturing

**Problem**: Evilginx not capturing sessions

**Solution**:
- Verify phishlet auth_tokens configuration
- Check cookie domain matches
- Review Evilginx debug logs
- Ensure HTTPS is being used

## Best Practices

### Lab Environment Setup

1. **Use isolated network**: VMs with host-only networking
2. **No external email**: Local SMTP server only
3. **Test accounts only**: Never use real credentials
4. **Document everything**: Keep detailed logs
5. **Clean up**: Remove templates after training

### Realistic Simulations

1. **Timing**: Send during business hours
2. **Pretext**: Use relevant scenarios
3. **Follow-up**: Provide training after tests
4. **Metrics**: Track and analyze results
5. **Improvement**: Iterate based on feedback

### Security Hardening

1. **Change default passwords**: Gophish admin account
2. **Use strong API keys**: Rotate regularly
3. **Enable TLS**: Use valid certificates
4. **Restrict access**: Firewall rules
5. **Monitor logs**: Watch for anomalies

## Additional Resources

- **Gophish Docs**: https://docs.getgophish.com/
- **Evilginx Docs**: https://help.evilginx.com/
- **Phishing Framework**: NIST Cybersecurity Framework
- **Legal Guidelines**: Check local laws before testing

---

**Educational Use Only - Use Responsibly**
EOF

print_success "Created TEMPLATE_SETUP_GUIDE.md"

print_header "Setup Complete!"

print_success "Templates are ready to use!"
echo ""
print_info "What's been created:"
print_info "  • 3 professional email templates"
print_info "  • 3 realistic landing pages"
print_info "  • 3 Evilginx phishlet configurations"
print_info "  • Complete documentation and guides"
echo ""
print_info "Next steps:"
print_info "  1. Read TEMPLATE_SETUP_GUIDE.md for detailed setup"
print_info "  2. Read templates/README.md for template details"
print_info "  3. Run 'python3 scripts/setup_templates.py' to install"
print_info "  4. Configure Evilginx (optional) for session capture"
echo ""
print_warning "Remember: Educational lab use only!"
print_info "Use these tools to learn defensive security and protect against attacks"
echo ""
