# Complete Setup Guide - Gophish Educational Phishing Framework

⚠️ **CRITICAL: FOR EDUCATIONAL LAB USE ONLY**
This repository contains advanced phishing simulation tools for security awareness training in isolated environments with explicit authorization.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [System Requirements](#system-requirements)
3. [Repository Setup](#repository-setup)
4. [Gophish Installation](#gophish-installation)
5. [Directory Structure](#directory-structure)
6. [Template Installation](#template-installation)
7. [Browser-in-the-Browser (BitB) Setup](#browser-in-the-browser-bitb-setup)
8. [Evilginx Integration (Optional)](#evilginx-integration-optional)
9. [SSL/TLS Certificate Setup](#ssltls-certificate-setup)
10. [Domain Configuration](#domain-configuration)
11. [Email Server Setup](#email-server-setup)
12. [Testing Your Setup](#testing-your-setup)
13. [Running Your First Campaign](#running-your-first-campaign)
14. [Troubleshooting](#troubleshooting)
15. [Security Considerations](#security-considerations)
16. [Legal & Ethical Guidelines](#legal--ethical-guidelines)

---

## Prerequisites

### Required Knowledge
- Basic Linux command line usage
- Understanding of DNS and web servers
- Basic networking concepts
- HTML/CSS fundamentals (optional but helpful)

### Required Software
- **Operating System**: Linux (Ubuntu 20.04+ recommended) or macOS
- **Go**: Version 1.19 or higher
- **Git**: For version control
- **Python**: Version 3.8 or higher (for automation scripts)
- **Text Editor**: VS Code, nano, vim, or your preference

### Optional Software (for advanced features)
- **Docker**: For containerized deployment
- **Nginx**: For reverse proxy
- **Certbot**: For Let's Encrypt SSL certificates
- **Postfix** or **SendGrid**: For email delivery

---

## System Requirements

### Minimum Requirements
- **CPU**: 2 cores
- **RAM**: 2GB
- **Storage**: 10GB free space
- **Network**: Static IP or dynamic DNS

### Recommended Requirements
- **CPU**: 4 cores
- **RAM**: 4GB
- **Storage**: 20GB free space
- **Network**: Static IP with proper DNS configuration

### Firewall Ports
You'll need to open these ports:
- **80/TCP**: HTTP (will redirect to HTTPS)
- **443/TCP**: HTTPS (for phishing pages)
- **3333/TCP**: Gophish admin panel (restrict to your IP only)
- **25/TCP, 587/TCP**: SMTP (if hosting your own mail server)

---

## Repository Setup

### Step 1: Fork or Clone This Repository

**Option A: Fork to Your Own GitHub Account**

1. Go to the repository on GitHub
2. Click the "Fork" button in the top right
3. Clone your forked repository:

```bash
# Replace YOUR_USERNAME with your GitHub username
git clone https://github.com/YOUR_USERNAME/gophish.git
cd gophish
```

**Option B: Clone Directly**

```bash
git clone https://github.com/theroseinc/gophish.git
cd gophish
```

### Step 2: Create Your Own Branch

```bash
# Create and switch to your own branch
git checkout -b main

# If you want to keep the existing work
git checkout -b production
```

### Step 3: Set Up Git Configuration

```bash
# Configure your git user
git config user.name "Your Name"
git config user.email "your.email@example.com"

# Set up remote (if you forked)
git remote add upstream https://github.com/theroseinc/gophish.git

# Verify remotes
git remote -v
```

---

## Gophish Installation

### Step 1: Install Go (if not already installed)

**On Ubuntu/Debian:**

```bash
# Download Go (check for latest version at https://go.dev/dl/)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

# Remove any previous Go installation
sudo rm -rf /usr/local/go

# Extract the archive
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Add Go to your PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

**On macOS:**

```bash
# Using Homebrew
brew install go

# Verify installation
go version
```

### Step 2: Build Gophish

```bash
# Navigate to the gophish directory
cd /home/user/gophish

# Install dependencies
go get ./...

# Build Gophish
go build

# Verify the binary was created
ls -lh gophish
```

You should see a `gophish` binary file created.

### Step 3: Configure Gophish

```bash
# Create a config file if it doesn't exist
cp config.json config.json.backup

# Edit the config file
nano config.json
```

**Minimal config.json:**

```json
{
  "admin_server": {
    "listen_url": "127.0.0.1:3333",
    "use_tls": true,
    "cert_path": "gophish_admin.crt",
    "key_path": "gophish_admin.key"
  },
  "phish_server": {
    "listen_url": "0.0.0.0:80",
    "use_tls": false
  },
  "db_name": "sqlite3",
  "db_path": "gophish.db",
  "migrations_prefix": "db/db_",
  "contact_address": "your-email@example.com",
  "logging": {
    "filename": "gophish.log",
    "level": "info"
  }
}
```

**Important Configuration Notes:**

- `admin_server.listen_url`: Admin panel address (keep on 127.0.0.1 for security)
- `phish_server.listen_url`: Phishing page server (0.0.0.0 to accept from any IP)
- `use_tls`: Set to `true` in production, `false` for testing
- `db_path`: SQLite database location

### Step 4: Generate Admin Panel SSL Certificates

```bash
# Generate self-signed certificate for admin panel
openssl req -newkey rsa:4096 -nodes -sha256 -keyout gophish_admin.key \
  -x509 -days 365 -out gophish_admin.crt \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost"
```

### Step 5: First Run

```bash
# Make the binary executable
chmod +x gophish

# Run Gophish
sudo ./gophish
```

**First time output will show:**
```
Please login with the username admin and the password [RANDOM_PASSWORD]
```

**IMPORTANT**: Copy this password immediately! You'll need it to log in.

### Step 6: Access Admin Panel

1. Open your browser
2. Navigate to: `https://127.0.0.1:3333`
3. Accept the self-signed certificate warning
4. Login with:
   - Username: `admin`
   - Password: [The password from the output]

### Step 7: Change Default Password

1. Click on "Settings" in the top right
2. Click "Account Settings"
3. Change your password to something secure
4. Save changes

---

## Directory Structure

Your repository should have this structure:

```
gophish/
├── templates/
│   ├── bitb/                          # Browser-in-the-Browser templates
│   │   ├── lib/
│   │   │   ├── bitb.js               # BitB JavaScript library
│   │   │   └── bitb.css              # BitB styling
│   │   ├── google/
│   │   │   └── google-oauth-bitb.html
│   │   ├── microsoft/
│   │   │   └── microsoft-oauth-bitb.html
│   │   ├── github/
│   │   │   └── github-oauth-bitb.html
│   │   └── README.md                  # BitB documentation
│   ├── email/                         # Email templates
│   │   ├── microsoft365-security-alert.html
│   │   ├── gmail-security-alert.html
│   │   ├── stripe-payment-failed.html
│   │   ├── linkedin-connection-request.html
│   │   ├── dropbox-file-share.html
│   │   ├── docusign-signature-required.html
│   │   ├── paypal-account-limitation.html
│   │   ├── apple-icloud-storage-full.html
│   │   └── slack-workspace-invitation.html
│   ├── landing-pages/                 # Landing page templates
│   │   ├── microsoft365-login.html
│   │   ├── gmail-login.html
│   │   ├── stripe-payment-update.html
│   │   ├── linkedin-login.html        # ✓ BitB integrated
│   │   ├── dropbox-login.html         # ✓ BitB integrated
│   │   ├── docusign-login.html        # ✓ BitB integrated
│   │   ├── paypal-login.html
│   │   ├── apple-icloud-login.html
│   │   └── slack-login.html           # ✓ BitB integrated
│   ├── evilginx-configs/              # Evilginx YAML configs
│   │   ├── microsoft365.yaml
│   │   ├── gmail.yaml
│   │   └── stripe.yaml
│   └── README.md
├── evilginx/                          # Evilginx integration code
│   ├── session_manager.go
│   ├── phishlet_config.go
│   └── proxy_manager.go
├── scripts/                           # Automation scripts
│   ├── setup_templates.py            # Auto-install templates
│   ├── generate_phishlets.py         # Generate Evilginx configs
│   └── quick_start.sh                # Complete setup automation
├── static/                            # Static files served by Gophish
│   ├── endpoint/                     # Gophish UI files
│   └── bitb/                         # BitB library (symlink or copy)
├── db/                                # Database migrations
├── config.json                        # Gophish configuration
├── gophish                           # Gophish binary (after build)
├── gophish.db                        # SQLite database (created on first run)
├── COMPLETE_SETUP_GUIDE.md           # This file
├── EDUCATIONAL_TEMPLATES_README.md   # Template documentation
├── TEMPLATE_SETUP_GUIDE.md           # Template setup guide
└── README.md                          # Main README
```

### Step 1: Create Missing Directories

```bash
# Navigate to your gophish directory
cd /home/user/gophish

# Create necessary directories if they don't exist
mkdir -p static/bitb
mkdir -p templates/bitb/lib
mkdir -p templates/email
mkdir -p templates/landing-pages
mkdir -p templates/evilginx-configs
mkdir -p scripts
mkdir -p logs
```

### Step 2: Verify Directory Structure

```bash
# Check that all directories exist
tree -L 2 -d
```

---

## Template Installation

### Method 1: Automatic Installation Using Python Script

This is the **recommended method** for most users.

**Step 1: Install Python Dependencies**

```bash
# Install pip if not already installed
sudo apt install python3-pip  # Ubuntu/Debian
brew install python3          # macOS

# Install required Python packages
pip3 install requests colorama
```

**Step 2: Configure the Setup Script**

```bash
# Navigate to scripts directory
cd /home/user/gophish/scripts

# Edit the setup script
nano setup_templates.py
```

**Update these variables at the top of the file:**

```python
# Gophish API Configuration
GOPHISH_API_URL = "https://127.0.0.1:3333"  # Your Gophish admin URL
GOPHISH_API_KEY = "YOUR_API_KEY_HERE"       # Get this from Gophish settings

# Template directories (usually no need to change)
EMAIL_TEMPLATES_DIR = "../templates/email"
LANDING_TEMPLATES_DIR = "../templates/landing-pages"
```

**Step 3: Get Your Gophish API Key**

1. Log into Gophish admin panel
2. Click "Settings" → "Account Settings"
3. Find the "API Key" section
4. Copy your API key
5. Paste it into the `setup_templates.py` file

**Step 4: Run the Setup Script**

```bash
# Make the script executable
chmod +x setup_templates.py

# Run the script
python3 setup_templates.py
```

**Expected Output:**

```
🚀 Gophish Template Setup Script
================================

✓ Connected to Gophish API successfully
✓ Found 9 email templates
✓ Found 9 landing page templates

📧 Installing Email Templates...
✓ Installed: Microsoft 365 Security Alert
✓ Installed: Gmail Security Alert
✓ Installed: Stripe Payment Failed
... [continues for all templates]

🌐 Installing Landing Page Templates...
✓ Installed: Microsoft 365 Login
✓ Installed: Gmail Login
✓ Installed: Stripe Payment Update
... [continues for all templates]

✅ Installation Complete!
   - 9 email templates installed
   - 9 landing page templates installed
```

### Method 2: Manual Installation via Gophish UI

If you prefer manual installation or the script doesn't work:

**For Each Email Template:**

1. Log into Gophish admin panel
2. Click "Email Templates" in the sidebar
3. Click "New Template"
4. Fill in the form:
   - **Name**: `Microsoft 365 - Security Alert` (descriptive name)
   - **Import Email**: Click "Import Email"
   - **Envelope Sender**: `security@microsoft.com` (spoofed from address)
   - **Subject**: Copy from template file
   - **HTML**: Paste HTML from template file
5. Enable "Add Tracking Image"
6. Click "Save Template"
7. Repeat for all 9 email templates

**For Each Landing Page:**

1. Click "Landing Pages" in the sidebar
2. Click "New Page"
3. Fill in the form:
   - **Name**: `Microsoft 365 - Login Page`
   - **HTML**: Paste HTML from landing page file
4. Enable "Capture Submitted Data"
5. Enable "Capture Passwords"
6. Click "Save Page"
7. Repeat for all 9 landing pages

**Template Files Location:**
- Email templates: `/home/user/gophish/templates/email/*.html`
- Landing pages: `/home/user/gophish/templates/landing-pages/*.html`

---

## Browser-in-the-Browser (BitB) Setup

BitB creates fake OAuth popup windows that display legitimate URLs in a fake browser address bar - highly effective for training users.

### Step 1: Verify BitB Files Exist

```bash
# Check that BitB library files exist
ls -lh /home/user/gophish/templates/bitb/lib/

# You should see:
# bitb.js  - JavaScript library
# bitb.css - Styling
```

### Step 2: Copy BitB Library to Gophish Static Directory

Gophish serves static files from the `static` directory, so we need to make BitB accessible:

**Option A: Create Symlink (Recommended)**

```bash
# Navigate to static directory
cd /home/user/gophish/static

# Create symlink to BitB library
ln -s ../templates/bitb bitb

# Verify symlink
ls -lh bitb
```

**Option B: Copy Files**

```bash
# Copy entire BitB directory
cp -r /home/user/gophish/templates/bitb /home/user/gophish/static/

# Verify files were copied
ls -lh /home/user/gophish/static/bitb/lib/
```

### Step 3: Update Landing Pages to Use Correct BitB Path

The landing pages reference BitB as `../bitb/lib/`, but when served by Gophish, they need to reference `/static/bitb/lib/`.

**Option A: Automated Update (Recommended)**

```bash
# Create a script to update paths
cat > /home/user/gophish/scripts/fix_bitb_paths.sh << 'EOF'
#!/bin/bash

# Update BitB CSS paths in landing pages
find ../templates/landing-pages -name "*.html" -type f -exec \
    sed -i 's|href="../bitb/lib/bitb.css"|href="/static/bitb/lib/bitb.css"|g' {} +

# Update BitB JS paths in landing pages
find ../templates/landing-pages -name "*.html" -type f -exec \
    sed -i 's|src="../bitb/lib/bitb.js"|src="/static/bitb/lib/bitb.js"|g' {} +

echo "✓ BitB paths updated successfully"
EOF

# Make executable
chmod +x /home/user/gophish/scripts/fix_bitb_paths.sh

# Run the script
cd /home/user/gophish/scripts
./fix_bitb_paths.sh
```

**Option B: Manual Update**

Edit each landing page file that uses BitB:
- `docusign-login.html`
- `linkedin-login.html`
- `dropbox-login.html`
- `slack-login.html`

Change:
```html
<link rel="stylesheet" href="../bitb/lib/bitb.css">
```
To:
```html
<link rel="stylesheet" href="/static/bitb/lib/bitb.css">
```

And:
```html
<script src="../bitb/lib/bitb.js"></script>
```
To:
```html
<script src="/static/bitb/lib/bitb.js"></script>
```

### Step 4: Test BitB Standalone

Before using in a campaign, test the BitB templates:

```bash
# Option 1: Use Python's built-in HTTP server
cd /home/user/gophish/templates/bitb
python3 -m http.server 8000

# Open browser to:
# http://localhost:8000/google/google-oauth-bitb.html
# http://localhost:8000/microsoft/microsoft-oauth-bitb.html
# http://localhost:8000/github/github-oauth-bitb.html
```

**What to Test:**
1. Click the "Sign in with Google/Microsoft/GitHub" button
2. Verify a fake browser window appears
3. Check that it shows the correct URL in the fake address bar
4. Try dragging the window (should be draggable)
5. Try submitting the form (should capture input)
6. Click the X button (should close the window)

### Step 5: Configure BitB in Landing Pages

The following landing pages have BitB pre-integrated:

1. **DocuSign Login** - Google & Apple OAuth popups
2. **LinkedIn Login** - Google OAuth popup
3. **Dropbox Login** - Google & Apple OAuth popups
4. **Slack Login** - Google & Apple OAuth popups

**No additional configuration needed!** Just use these landing pages in your campaigns.

---

## Evilginx Integration (Optional)

⚠️ **ADVANCED FEATURE**: Evilginx enables man-in-the-middle session hijacking to bypass 2FA. Only use in authorized environments.

### Prerequisites

- **Evilginx 2**: Must be installed separately
- **VPS**: With public IP and proper DNS
- **Domain**: Properly configured domain name
- **SSL Certificate**: Valid SSL/TLS certificate

### Step 1: Install Evilginx

```bash
# Install Go (if not already)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Clone Evilginx
git clone https://github.com/kgretzky/evilginx2.git
cd evilginx2

# Build Evilginx
make

# Install
sudo make install

# Verify installation
evilginx version
```

### Step 2: Generate Phishlet Configurations

```bash
# Navigate to scripts directory
cd /home/user/gophish/scripts

# Generate YAML phishlet configs
python3 generate_phishlets.py

# Check generated files
ls -lh ../templates/evilginx-configs/
```

**You should see:**
- `microsoft365.yaml`
- `gmail.yaml`
- `stripe.yaml`
- `linkedin.yaml`
- `dropbox.yaml`
- `docusign.yaml`
- `paypal.yaml`
- `apple.yaml`
- `slack.yaml`

### Step 3: Install Phishlets to Evilginx

```bash
# Copy phishlets to Evilginx directory
sudo cp /home/user/gophish/templates/evilginx-configs/*.yaml \
    /usr/share/evilginx/phishlets/

# Or if using custom installation directory
cp /home/user/gophish/templates/evilginx-configs/*.yaml \
    ~/evilginx2/phishlets/
```

### Step 4: Configure Evilginx

```bash
# Start Evilginx
sudo evilginx

# In Evilginx console:
config domain your-domain.com
config ip YOUR_PUBLIC_IP

# Enable a phishlet
phishlets hostname microsoft365 login.your-domain.com
phishlets enable microsoft365

# Create a lure
lures create microsoft365
lures get-url 0
```

### Step 5: Integration with Gophish

This repository includes Go code for integration, but it requires custom compilation:

**Files involved:**
- `evilginx/session_manager.go` - Manages captured sessions
- `evilginx/phishlet_config.go` - Phishlet configurations
- `evilginx/proxy_manager.go` - Proxy management

**To compile with Evilginx integration:**

```bash
# This is advanced - requires modifying Gophish core
# Contact for advanced integration support
```

---

## SSL/TLS Certificate Setup

For production use, you **must** use proper SSL/TLS certificates.

### Option 1: Let's Encrypt (Free, Recommended)

**Prerequisites:**
- Domain name pointing to your server
- Ports 80 and 443 open

**Step 1: Install Certbot**

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install certbot

# macOS
brew install certbot
```

**Step 2: Generate Certificate**

```bash
# Stop Gophish temporarily
sudo pkill gophish

# Generate certificate
sudo certbot certonly --standalone -d your-domain.com

# Certificates will be saved to:
# /etc/letsencrypt/live/your-domain.com/fullchain.pem
# /etc/letsencrypt/live/your-domain.com/privkey.pem
```

**Step 3: Update Gophish Configuration**

```bash
nano /home/user/gophish/config.json
```

Update the `phish_server` section:

```json
{
  "phish_server": {
    "listen_url": "0.0.0.0:443",
    "use_tls": true,
    "cert_path": "/etc/letsencrypt/live/your-domain.com/fullchain.pem",
    "key_path": "/etc/letsencrypt/live/your-domain.com/privkey.pem"
  }
}
```

**Step 4: Set Up Auto-Renewal**

```bash
# Test renewal
sudo certbot renew --dry-run

# Add cron job for auto-renewal
sudo crontab -e

# Add this line:
0 0 * * * certbot renew --quiet --post-hook "systemctl restart gophish"
```

### Option 2: Self-Signed Certificate (Testing Only)

```bash
# Generate self-signed certificate
openssl req -newkey rsa:4096 -nodes -sha256 \
  -keyout gophish_phish.key \
  -x509 -days 365 -out gophish_phish.crt \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=your-domain.com"

# Update config.json to use these certificates
```

**⚠️ WARNING**: Self-signed certificates will show browser warnings and reduce campaign effectiveness!

---

## Domain Configuration

Proper domain configuration is **critical** for successful campaigns.

### Step 1: Purchase a Domain

**Recommended Registrars:**
- Namecheap
- Google Domains
- Cloudflare Registrar

**Domain Selection Tips:**
- Use a domain similar to but not identical to the target brand
- Examples:
  - `microsoft-auth.com` (NOT `microsoft.com`)
  - `google-secure.net` (NOT `google.com`)
  - `paypal-verify.com` (NOT `paypal.com`)

### Step 2: Configure DNS Records

**A Record (Required):**

```
Type: A
Name: @
Value: YOUR_SERVER_IP
TTL: 3600
```

**A Record for Subdomain (if using www):**

```
Type: A
Name: www
Value: YOUR_SERVER_IP
TTL: 3600
```

**MX Record (for email):**

```
Type: MX
Name: @
Value: mail.your-domain.com
Priority: 10
TTL: 3600
```

**SPF Record (prevents email spoofing detection):**

```
Type: TXT
Name: @
Value: v=spf1 ip4:YOUR_SERVER_IP ~all
TTL: 3600
```

**DKIM Record (email authentication):**

```bash
# Generate DKIM key
opendkim-genkey -t -s mail -d your-domain.com

# Add the generated TXT record to DNS
Type: TXT
Name: mail._domainkey
Value: [Generated DKIM public key]
TTL: 3600
```

**DMARC Record (email policy):**

```
Type: TXT
Name: _dmarc
Value: v=DMARC1; p=none; rua=mailto:dmarc@your-domain.com
TTL: 3600
```

### Step 3: Verify DNS Propagation

```bash
# Check A record
dig your-domain.com +short

# Check MX record
dig MX your-domain.com +short

# Check TXT records
dig TXT your-domain.com +short

# Full DNS check
nslookup your-domain.com
```

**Wait 24-48 hours for full DNS propagation.**

---

## Email Server Setup

You have several options for sending phishing emails.

### Option 1: Third-Party Email Service (Easiest)

**Recommended Services:**
- **SendGrid** (requires verification)
- **Amazon SES** (requires verification)
- **Mailgun** (requires verification)

**SendGrid Setup Example:**

1. Sign up at https://sendgrid.com
2. Verify your domain
3. Create an API key
4. Configure Gophish:

```bash
nano /home/user/gophish/config.json
```

Add SMTP configuration:

```json
{
  "smtp_config": {
    "host": "smtp.sendgrid.net:587",
    "username": "apikey",
    "password": "YOUR_SENDGRID_API_KEY",
    "from": "security@your-domain.com"
  }
}
```

### Option 2: Local Postfix Server (Advanced)

**Step 1: Install Postfix**

```bash
sudo apt update
sudo apt install postfix mailutils
```

During installation:
- Choose "Internet Site"
- Set system mail name to your domain

**Step 2: Configure Postfix**

```bash
sudo nano /etc/postfix/main.cf
```

Update these settings:

```
myhostname = mail.your-domain.com
mydomain = your-domain.com
myorigin = $mydomain
inet_interfaces = all
mydestination = $myhostname, localhost.$mydomain, localhost, $mydomain
```

**Step 3: Restart Postfix**

```bash
sudo systemctl restart postfix
sudo systemctl enable postfix
```

**Step 4: Test Email**

```bash
echo "Test email body" | mail -s "Test Subject" your-email@gmail.com
```

**Step 5: Configure in Gophish**

Create SMTP profile in Gophish UI:
- **Name**: Local Postfix
- **Host**: localhost:25
- **From**: security@your-domain.com
- **Ignore Certificate Errors**: Checked (for testing)

---

## Testing Your Setup

### Pre-Flight Checklist

Before running campaigns, verify everything works:

- [ ] Gophish binary runs without errors
- [ ] Admin panel accessible at https://127.0.0.1:3333
- [ ] Can log in with admin credentials
- [ ] All 9 email templates imported
- [ ] All 9 landing page templates imported
- [ ] SSL certificate configured (for production)
- [ ] Domain resolves to server IP
- [ ] Phishing server accessible on port 80/443
- [ ] Email server can send mail
- [ ] BitB library accessible at /static/bitb/lib/

### Test 1: Access Admin Panel

```bash
# Start Gophish
cd /home/user/gophish
sudo ./gophish &

# Check logs
tail -f gophish.log
```

Open browser: `https://127.0.0.1:3333`

**Expected**: Login page loads successfully

### Test 2: Test Landing Page

1. In Gophish admin, go to "Landing Pages"
2. Click on any landing page
3. Copy the preview URL
4. Open in new browser tab

**Expected**: Landing page displays correctly

### Test 3: Test BitB Functionality

1. Choose a landing page with BitB (DocuSign, LinkedIn, Dropbox, or Slack)
2. Preview the page
3. Click the SSO button (e.g., "Sign in with Google")

**Expected**: Fake OAuth popup window appears with realistic browser chrome

### Test 4: Test Credential Capture

1. Preview a landing page
2. Enter test credentials: `test@example.com` / `password123`
3. Submit the form
4. Go to "Dashboard" in Gophish

**Expected**: Dashboard shows "1 submitted data"

### Test 5: Test Email Sending

1. Create a test group:
   - Name: "Test Group"
   - Email: your-personal-email@gmail.com
   - First Name: "Test"
   - Last Name: "User"

2. Create a test SMTP profile:
   - Name: "Test SMTP"
   - Host: `localhost:25` (or your SMTP server)
   - From: `test@your-domain.com`

3. Create a test campaign:
   - Name: "Test Campaign"
   - Email Template: Any template
   - Landing Page: Any landing page
   - URL: `http://your-domain.com`
   - Launch Date: Now
   - Send By: 1 hour from now
   - Groups: Test Group
   - SMTP Profile: Test SMTP

4. Launch the campaign

**Expected**:
- Campaign shows in dashboard
- Email arrives in your inbox
- Clicking link opens landing page
- Dashboard updates with "Email Opened" and "Clicked Link"

---

## Running Your First Campaign

Now that everything is tested, let's run a proper campaign.

### Step 1: Create User Groups

1. Click "Users & Groups" → "New Group"
2. Name: "Security Awareness Training - Q1"
3. Add users:

**Option A: Manual Entry**
- Click "Add" to add users one by one

**Option B: CSV Import**
- Create CSV file:

```csv
First Name,Last Name,Email,Position
John,Doe,john.doe@company.com,Manager
Jane,Smith,jane.smith@company.com,Developer
```

- Click "Bulk Import Users"
- Upload CSV

### Step 2: Create SMTP Profile

1. Click "Sending Profiles" → "New Profile"
2. Fill in details:
   - **Name**: Production SMTP
   - **From**: `security@your-domain.com`
   - **Host**: Your SMTP server
   - **Username**: (if required)
   - **Password**: (if required)
3. Click "Send Test Email" to verify
4. Save Profile

### Step 3: Create Campaign

1. Click "Campaigns" → "New Campaign"
2. Fill in details:

**Basic Information:**
- **Name**: Q1 2024 - Microsoft Phishing Simulation
- **Email Template**: Microsoft 365 - Security Alert
- **Landing Page**: Microsoft 365 - Login Page
- **URL**: `https://login-microsoft.your-domain.com`
- **Launch Date**: [Select date/time]
- **Send By**: [Select end date/time]

**Advanced Options:**
- **Sending Profile**: Production SMTP
- **Groups**: Select your target group
- **Send Emails By**: Spread over hours (reduces suspicion)

3. Click "Launch Campaign"

### Step 4: Monitor Campaign

Dashboard shows real-time stats:
- **Total Emails Sent**: Number of targets
- **Emails Opened**: Users who opened email (tracking pixel)
- **Clicked Link**: Users who clicked the phishing link
- **Submitted Data**: Users who entered credentials
- **Email Reported**: Users who reported as phishing (via plugin)

### Step 5: Export Results

1. Go to campaign details
2. Click "Export CSV"
3. Analyze results in spreadsheet

**Metrics to track:**
- Open rate: `(Emails Opened / Total Sent) * 100`
- Click rate: `(Clicked Link / Total Sent) * 100`
- Submission rate: `(Submitted Data / Total Sent) * 100`

### Step 6: Follow-Up Training

**For users who clicked/submitted:**
1. Send educational email explaining what happened
2. Provide security awareness training resources
3. Schedule 1-on-1 training if needed

---

## Troubleshooting

### Issue: Can't Access Admin Panel

**Symptoms**: Browser can't connect to https://127.0.0.1:3333

**Solutions**:

1. Check Gophish is running:
```bash
ps aux | grep gophish
```

2. Check port is listening:
```bash
sudo netstat -tlnp | grep 3333
```

3. Check firewall:
```bash
sudo ufw status
sudo ufw allow 3333/tcp
```

4. Check logs:
```bash
tail -f /home/user/gophish/gophish.log
```

### Issue: Landing Pages Show 404 Error

**Symptoms**: Clicking phishing link shows "404 Not Found"

**Solutions**:

1. Verify phish server is running:
```bash
sudo netstat -tlnp | grep :80
```

2. Check config.json:
```bash
grep phish_server /home/user/gophish/config.json
```

3. Ensure `listen_url` is `0.0.0.0:80` not `127.0.0.1:80`

4. Restart Gophish:
```bash
sudo pkill gophish
sudo ./gophish &
```

### Issue: BitB Windows Don't Appear

**Symptoms**: Clicking SSO button does nothing

**Solutions**:

1. Check browser console (F12) for JavaScript errors

2. Verify BitB files are accessible:
```bash
curl http://your-domain.com/static/bitb/lib/bitb.js
curl http://your-domain.com/static/bitb/lib/bitb.css
```

3. Check file paths in landing page HTML:
```bash
grep bitb /home/user/gophish/templates/landing-pages/docusign-login.html
```

Should show:
```html
<link rel="stylesheet" href="/static/bitb/lib/bitb.css">
<script src="/static/bitb/lib/bitb.js"></script>
```

4. Verify symlink or files exist:
```bash
ls -lh /home/user/gophish/static/bitb/lib/
```

### Issue: Emails Not Sending

**Symptoms**: Campaign shows "Sending" but emails never arrive

**Solutions**:

1. Test SMTP profile:
   - Go to "Sending Profiles"
   - Click "Send Test Email"
   - Check if test email arrives

2. Check SMTP logs:
```bash
sudo tail -f /var/log/mail.log
```

3. Verify DNS records:
```bash
dig MX your-domain.com +short
dig TXT your-domain.com +short
```

4. Check if port 25/587 is open:
```bash
telnet smtp.gmail.com 587
```

5. Verify SPF/DKIM/DMARC:
   - Use https://mxtoolbox.com/SuperTool.aspx
   - Enter your domain
   - Check all email authentication records

### Issue: Emails Go to Spam

**Symptoms**: Emails arrive but in spam folder

**Solutions**:

1. **Warm up your domain**: Send legitimate emails first
2. **Configure proper DNS records**: SPF, DKIM, DMARC
3. **Use proper from address**: Match your domain
4. **Avoid spam trigger words**: "Free", "Urgent", "Act Now"
5. **Use good email content**: Proper formatting, no excessive links
6. **Start with low volume**: Don't send 1000 emails at once

### Issue: SSL Certificate Errors

**Symptoms**: Browser shows "Your connection is not private"

**Solutions**:

1. **For self-signed certificates**: This is expected, users must accept the warning

2. **For Let's Encrypt certificates**:
```bash
# Check certificate validity
sudo openssl x509 -in /etc/letsencrypt/live/your-domain.com/fullchain.pem -noout -dates

# Renew certificate
sudo certbot renew

# Restart Gophish
sudo pkill gophish
sudo ./gophish &
```

3. **Verify certificate matches domain**:
```bash
sudo openssl x509 -in /etc/letsencrypt/live/your-domain.com/fullchain.pem -noout -subject
```

### Issue: High CPU/Memory Usage

**Symptoms**: Server slow, Gophish using lots of resources

**Solutions**:

1. Check database size:
```bash
ls -lh /home/user/gophish/gophish.db
```

2. Clean up old campaigns:
   - Delete completed campaigns from UI
   - Archive old data

3. Optimize database:
```bash
sqlite3 /home/user/gophish/gophish.db "VACUUM;"
```

4. Increase server resources:
   - Upgrade VPS plan
   - Add more RAM

### Issue: Port 80/443 Already in Use

**Symptoms**: Error: "bind: address already in use"

**Solutions**:

1. Check what's using the port:
```bash
sudo lsof -i :80
sudo lsof -i :443
```

2. Common culprits:
   - **Apache**: `sudo systemctl stop apache2`
   - **Nginx**: `sudo systemctl stop nginx`
   - **Another Gophish instance**: `sudo pkill gophish`

3. Change Gophish port (temporary solution):
```json
{
  "phish_server": {
    "listen_url": "0.0.0.0:8080"
  }
}
```

---

## Security Considerations

### Server Security

**1. Firewall Configuration**

```bash
# Install UFW
sudo apt install ufw

# Default deny incoming
sudo ufw default deny incoming
sudo ufw default allow outgoing

# Allow SSH (change 22 to your SSH port)
sudo ufw allow 22/tcp

# Allow Gophish admin (restrict to your IP)
sudo ufw allow from YOUR_IP to any port 3333

# Allow phishing server
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# Enable firewall
sudo ufw enable
```

**2. Change Default SSH Port**

```bash
sudo nano /etc/ssh/sshd_config

# Change line:
Port 2222

# Restart SSH
sudo systemctl restart sshd
```

**3. Disable Root Login**

```bash
sudo nano /etc/ssh/sshd_config

# Change line:
PermitRootLogin no

# Restart SSH
sudo systemctl restart sshd
```

**4. Install Fail2Ban**

```bash
sudo apt install fail2ban

# Create jail configuration
sudo nano /etc/fail2ban/jail.local
```

Add:
```
[sshd]
enabled = true
port = 2222
maxretry = 3
```

```bash
# Restart Fail2Ban
sudo systemctl restart fail2ban
```

**5. Keep System Updated**

```bash
# Update packages
sudo apt update && sudo apt upgrade -y

# Enable automatic security updates
sudo apt install unattended-upgrades
sudo dpkg-reconfigure -plow unattended-upgrades
```

**6. Restrict Admin Panel Access**

Only allow access from your IP:

```bash
sudo ufw allow from YOUR_IP to any port 3333
```

Or use SSH tunnel:
```bash
# On your local machine
ssh -L 3333:localhost:3333 user@your-server

# Then access admin panel at:
# https://localhost:3333
```

### Campaign Security

**1. Use Proper Disclosure**

Add this to every email footer:
```
This is a simulated phishing exercise conducted by [Your Organization] Security Team
for educational purposes. If you have questions, contact security@yourcompany.com
```

**2. Immediate Debrief**

After campaign, send debrief email explaining:
- This was a test
- What red flags to look for
- How to report real phishing
- No disciplinary action

**3. Data Protection**

**Encrypt database:**
```bash
# Backup database
cp gophish.db gophish.db.backup

# Encrypt with GPG
gpg -c gophish.db.backup
```

**Secure file permissions:**
```bash
chmod 600 gophish.db
chmod 600 config.json
chmod 600 *.key
```

**4. Audit Logs**

Enable detailed logging:
```json
{
  "logging": {
    "filename": "gophish.log",
    "level": "debug"
  }
}
```

Rotate logs:
```bash
# Install logrotate config
sudo nano /etc/logrotate.d/gophish
```

Add:
```
/home/user/gophish/gophish.log {
    daily
    rotate 7
    compress
    missingok
    notifempty
}
```

---

## Legal & Ethical Guidelines

### ⚠️ CRITICAL: Legal Requirements

**1. Written Authorization Required**

Before ANY phishing simulation:
- Get written approval from organization leadership
- Define scope (who, what, when)
- Define acceptable use
- Define data handling procedures

**2. Prohibited Actions**

**NEVER:**
- Test third-party systems without written permission
- Collect sensitive data (SSN, credit cards, health info)
- Use against real users without authorization
- Sell or distribute captured data
- Cause harm or disruption
- Violate computer fraud laws (CFAA in USA)

**3. Data Handling**

- Delete captured credentials immediately after campaign
- Don't store passwords in plain text
- Use encryption for sensitive data
- Have data retention policy
- Follow GDPR/CCPA if applicable

### Ethical Guidelines

**1. Educational Purpose Only**

This tool is for:
- Security awareness training
- Authorized penetration testing
- Academic research with IRB approval
- Defensive security research

**2. Respect Privacy**

- Minimize data collection
- Don't read personal emails
- Don't access personal accounts
- Respect boundaries

**3. No Harm Principle**

- No disciplinary action for failing
- Provide immediate education
- Support users who are concerned
- Foster learning environment

**4. Transparency**

- Disclose testing after campaign
- Explain red flags users missed
- Provide resources for improvement
- Answer questions honestly

### Legal Resources

**United States:**
- Computer Fraud and Abuse Act (CFAA)
- CAN-SPAM Act
- State privacy laws

**Europe:**
- GDPR (General Data Protection Regulation)
- ePrivacy Directive

**International:**
- Budapest Convention on Cybercrime
- Local computer crime laws

**Consult a lawyer before conducting phishing simulations in production environments.**

---

## Advanced Topics

### Running Gophish as a Service

Create systemd service:

```bash
sudo nano /etc/systemd/system/gophish.service
```

Add:
```ini
[Unit]
Description=Gophish Phishing Framework
After=network.target

[Service]
Type=simple
User=gophish
Group=gophish
WorkingDirectory=/home/user/gophish
ExecStart=/home/user/gophish/gophish
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl daemon-reload
sudo systemctl enable gophish
sudo systemctl start gophish
sudo systemctl status gophish
```

### Using Nginx Reverse Proxy

Install Nginx:
```bash
sudo apt install nginx
```

Configure:
```bash
sudo nano /etc/nginx/sites-available/gophish
```

Add:
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

Enable:
```bash
sudo ln -s /etc/nginx/sites-available/gophish /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

### Docker Deployment

Create Dockerfile:
```dockerfile
FROM golang:1.21-alpine
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build
EXPOSE 3333 80
CMD ["./gophish"]
```

Build and run:
```bash
docker build -t gophish .
docker run -d -p 3333:3333 -p 80:80 gophish
```

---

## Maintenance & Backups

### Daily Maintenance

```bash
# Check Gophish status
sudo systemctl status gophish

# Check disk space
df -h

# Check logs for errors
tail -f /home/user/gophish/gophish.log | grep ERROR
```

### Weekly Maintenance

```bash
# Backup database
cp /home/user/gophish/gophish.db \
   /home/user/backups/gophish_$(date +%Y%m%d).db

# Clean up old campaigns (via UI)

# Update system packages
sudo apt update && sudo apt upgrade -y
```

### Monthly Maintenance

```bash
# Rebuild Gophish with latest updates
cd /home/user/gophish
git pull
go build

# Vacuum database
sqlite3 gophish.db "VACUUM;"

# Review and rotate logs
logrotate -f /etc/logrotate.d/gophish

# Review SSL certificate expiry
sudo certbot certificates
```

### Backup Strategy

**Automated Backup Script:**

```bash
cat > /home/user/gophish/backup.sh << 'EOF'
#!/bin/bash
BACKUP_DIR="/home/user/backups"
DATE=$(date +%Y%m%d_%H%M%S)
mkdir -p $BACKUP_DIR

# Backup database
cp /home/user/gophish/gophish.db $BACKUP_DIR/gophish_$DATE.db

# Backup config
cp /home/user/gophish/config.json $BACKUP_DIR/config_$DATE.json

# Backup templates
tar -czf $BACKUP_DIR/templates_$DATE.tar.gz /home/user/gophish/templates/

# Delete backups older than 30 days
find $BACKUP_DIR -name "gophish_*.db" -mtime +30 -delete
find $BACKUP_DIR -name "config_*.json" -mtime +30 -delete
find $BACKUP_DIR -name "templates_*.tar.gz" -mtime +30 -delete

echo "Backup completed: $DATE"
EOF

chmod +x /home/user/gophish/backup.sh

# Add to crontab
crontab -e

# Add line:
0 2 * * * /home/user/gophish/backup.sh >> /home/user/backups/backup.log 2>&1
```

---

## Additional Resources

### Documentation
- **Gophish Official Docs**: https://docs.getgophish.com/
- **Evilginx2 Wiki**: https://github.com/kgretzky/evilginx2/wiki
- **BitB Original Research**: https://mrd0x.com/browser-in-the-browser-phishing-attack/

### Security Awareness Training Resources
- **NIST Cybersecurity Framework**: https://www.nist.gov/cyberframework
- **SANS Security Awareness**: https://www.sans.org/security-awareness-training/
- **PhishMe**: https://cofense.com/

### Community
- **Gophish GitHub**: https://github.com/gophish/gophish
- **Gophish Slack**: https://gophish.org/slack

### Reporting Issues
- **Gophish Issues**: https://github.com/gophish/gophish/issues
- **This Repository**: [Your GitHub Issues Link]

---

## Changelog

### Version 2.0.0 (Current)
- ✓ Added Browser-in-the-Browser (BitB) attack templates
- ✓ Enhanced 4 landing pages with BitB OAuth popups
- ✓ Created comprehensive BitB documentation
- ✓ Added 6 additional phishing template sets (LinkedIn, Dropbox, DocuSign, PayPal, Apple, Slack)
- ✓ Added 6 new Evilginx phishlet configurations
- ✓ Complete setup guide created

### Version 1.0.0
- Initial release with 3 template sets
- Basic Evilginx integration
- Core automation scripts

---

## Support

### Getting Help

1. **Check this guide first** - Most questions are answered here
2. **Check logs** - `tail -f gophish.log`
3. **Check GitHub Issues** - Someone may have had the same problem
4. **Ask in community** - Gophish Slack channel
5. **Professional support** - Contact for paid support

### Contributing

If you want to contribute to this repository:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-template`
3. Make your changes
4. Test thoroughly
5. Commit: `git commit -m "Add new template for X"`
6. Push: `git push origin feature/new-template`
7. Create Pull Request

---

## Acknowledgments

- **Gophish Team** - For the excellent phishing framework
- **Kuba Gretzky** - For Evilginx2
- **mr.d0x** - For Browser-in-the-Browser research
- **Security Community** - For ongoing research and improvements

---

## License

This project is for **educational purposes only**. Use responsibly and ethically.

**⚠️ DISCLAIMER**: The authors are not responsible for misuse of this tool. Only use in authorized environments with explicit written permission.

---

**End of Complete Setup Guide**

*Last Updated: 2024*
*Version: 2.0.0*
