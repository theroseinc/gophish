# Quick Start Guide - Gophish Setup

⚠️ **FOR EDUCATIONAL LAB USE ONLY**

This is a condensed setup guide. For detailed instructions, see [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md).

---

## 5-Minute Setup (Local Testing)

### Prerequisites
- Linux/macOS with Go 1.19+ installed
- Python 3.8+
- Git

### Step 1: Clone and Build (2 minutes)

```bash
# Clone repository
git clone https://github.com/YOUR_USERNAME/gophish.git
cd gophish

# Build Gophish
go build

# Generate admin SSL certificate
openssl req -newkey rsa:4096 -nodes -sha256 -keyout gophish_admin.key \
  -x509 -days 365 -out gophish_admin.crt \
  -subj "/C=US/ST=State/L=City/O=Org/CN=localhost"
```

### Step 2: Start Gophish (1 minute)

```bash
# Start Gophish
sudo ./gophish

# Copy the admin password shown in output
# Example: Please login with the username admin and the password abc123xyz
```

### Step 3: Access Admin Panel (1 minute)

1. Open browser: `https://127.0.0.1:3333`
2. Accept self-signed certificate warning
3. Login:
   - Username: `admin`
   - Password: [from step 2]
4. Change password in Settings

### Step 4: Install Templates (1 minute)

```bash
# Install Python dependencies
pip3 install requests colorama

# Edit setup script with your API key
cd scripts
nano setup_templates.py

# Get API key from Gophish: Settings → Account Settings → API Key
# Update this line:
GOPHISH_API_KEY = "YOUR_API_KEY_HERE"

# Run setup script
python3 setup_templates.py
```

**Output:**
```
✓ Installed: 9 email templates
✓ Installed: 9 landing page templates
```

### Step 5: Setup BitB Library

```bash
# Copy BitB to static directory
cd /home/user/gophish
cp -r templates/bitb static/

# Update paths in landing pages
cd scripts
./fix_bitb_paths.sh
```

---

## Test Campaign (5 minutes)

### Quick Test

1. **Create Group**:
   - Go to "Users & Groups" → "New Group"
   - Name: "Test"
   - Add your email

2. **Create SMTP Profile**:
   - Go to "Sending Profiles" → "New Profile"
   - Name: "Local"
   - Host: `localhost:25`
   - From: `test@test.com`

3. **Create Campaign**:
   - Go to "Campaigns" → "New Campaign"
   - Name: "Test Campaign"
   - Email Template: "Microsoft 365 - Security Alert"
   - Landing Page: "Microsoft 365 - Login Page"
   - URL: `http://localhost`
   - Groups: "Test"
   - SMTP: "Local"
   - Launch: Now

4. **Check Dashboard**:
   - View real-time stats
   - See captured data

---

## Production Setup (30 minutes)

For production deployment with SSL, domain, and email server:

### Quick Production Checklist

- [ ] VPS with Ubuntu 20.04+ (DigitalOcean, Linode, AWS)
- [ ] Domain name purchased and DNS configured
- [ ] SSL certificate (Let's Encrypt)
- [ ] Email service (SendGrid, Mailgun, or local Postfix)
- [ ] Firewall configured (UFW)
- [ ] Gophish running as systemd service

### Production Quick Commands

```bash
# 1. Install Let's Encrypt SSL
sudo apt install certbot
sudo certbot certonly --standalone -d your-domain.com

# 2. Update config.json for SSL
nano config.json
# Update phish_server section:
{
  "phish_server": {
    "listen_url": "0.0.0.0:443",
    "use_tls": true,
    "cert_path": "/etc/letsencrypt/live/your-domain.com/fullchain.pem",
    "key_path": "/etc/letsencrypt/live/your-domain.com/privkey.pem"
  }
}

# 3. Configure firewall
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow from YOUR_IP to any port 3333
sudo ufw enable

# 4. Create systemd service
sudo nano /etc/systemd/system/gophish.service
# Add service configuration (see COMPLETE_SETUP_GUIDE.md)

# 5. Start service
sudo systemctl enable gophish
sudo systemctl start gophish
```

---

## Template Features

### Email Templates (9 total)
- Microsoft 365 Security Alert
- Gmail Security Alert
- Stripe Payment Failed
- LinkedIn Connection Request
- Dropbox File Share
- DocuSign Signature Required
- PayPal Account Limitation
- Apple iCloud Storage Full
- Slack Workspace Invitation

### Landing Pages (9 total)
- Microsoft 365 Login (2-stage)
- Gmail Login
- Stripe Payment Update (full card capture)
- **LinkedIn Login** ✓ BitB Google OAuth
- **Dropbox Login** ✓ BitB Google + Apple OAuth
- **DocuSign Login** ✓ BitB Google + Apple OAuth
- PayPal Login
- Apple iCloud Login
- **Slack Login** ✓ BitB Google + Apple OAuth

### BitB (Browser-in-the-Browser)
- **What**: Fake OAuth popup windows showing legitimate URLs
- **Why**: Highly convincing, users trust OAuth popups
- **How**: Click SSO button → Fake browser window with real URL
- **Browsers**: Chrome, Firefox, Safari, Edge styles
- **Integrated in**: DocuSign, LinkedIn, Dropbox, Slack pages

---

## Common Issues & Quick Fixes

### Issue: Can't access admin panel
```bash
# Check if running
ps aux | grep gophish

# Check logs
tail -f gophish.log

# Restart
sudo pkill gophish
sudo ./gophish
```

### Issue: Landing pages show 404
```bash
# Check phish server config
grep phish_server config.json

# Should be 0.0.0.0:80 not 127.0.0.1:80
```

### Issue: BitB doesn't work
```bash
# Check BitB files exist
ls -lh static/bitb/lib/

# Test direct access
curl http://localhost/static/bitb/lib/bitb.js

# If 404, copy files:
cp -r templates/bitb static/
```

### Issue: Emails not sending
```bash
# Test SMTP
telnet smtp.gmail.com 587

# Check mail logs
sudo tail -f /var/log/mail.log

# Use external service (SendGrid, Mailgun)
```

---

## File Structure

```
gophish/
├── gophish                    # Binary (after build)
├── config.json                # Configuration
├── gophish.db                 # Database
├── templates/                 # All templates
│   ├── bitb/                 # BitB library & templates
│   ├── email/                # Email templates
│   ├── landing-pages/        # Landing pages
│   └── evilginx-configs/     # Evilginx configs
├── static/                    # Static files
│   └── bitb/                 # BitB library (copy here)
├── scripts/                   # Automation scripts
├── COMPLETE_SETUP_GUIDE.md   # Full detailed guide
├── QUICKSTART.md             # This file
└── README.md                  # Overview
```

---

## Important Commands

### Start/Stop Gophish
```bash
# Start
sudo ./gophish &

# Stop
sudo pkill gophish

# With systemd
sudo systemctl start gophish
sudo systemctl stop gophish
sudo systemctl restart gophish
```

### View Logs
```bash
# Real-time
tail -f gophish.log

# Last 50 lines
tail -50 gophish.log

# Search for errors
grep ERROR gophish.log
```

### Backup Database
```bash
# Simple backup
cp gophish.db gophish_backup_$(date +%Y%m%d).db

# Encrypted backup
gpg -c gophish.db
```

### Check Status
```bash
# Is Gophish running?
ps aux | grep gophish

# What ports are open?
sudo netstat -tlnp | grep gophish

# Check disk space
df -h
```

---

## Security Checklist

- [ ] Changed default admin password
- [ ] Admin panel only accessible from your IP
- [ ] SSL/TLS enabled for production
- [ ] Firewall configured (UFW)
- [ ] SSH key authentication enabled
- [ ] Root login disabled
- [ ] Database encrypted
- [ ] Backups automated
- [ ] Written authorization obtained
- [ ] Data retention policy in place

---

## Next Steps

1. **Test locally** - Run test campaign with your email
2. **Read full guide** - See [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md)
3. **Setup production** - Domain, SSL, email server
4. **Get authorization** - Written approval before production use
5. **Train users** - Educational follow-up after campaigns

---

## Resources

- **Full Setup Guide**: [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md)
- **Template Guide**: [EDUCATIONAL_TEMPLATES_README.md](EDUCATIONAL_TEMPLATES_README.md)
- **BitB Guide**: [templates/bitb/README.md](templates/bitb/README.md)
- **Gophish Docs**: https://docs.getgophish.com/

---

## Getting Help

1. Check [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) troubleshooting section
2. Check `gophish.log` for errors
3. Search GitHub issues
4. Ask in Gophish Slack community

---

**⚠️ Remember: For educational use only in authorized environments!**

*Quick Start Guide - Version 2.0.0*
