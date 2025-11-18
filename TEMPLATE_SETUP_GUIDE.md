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
