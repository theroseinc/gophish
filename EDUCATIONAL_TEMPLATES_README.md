# Gophish Advanced Templates & Evilginx Integration
## Educational Cybersecurity Lab Setup

[![Educational Use](https://img.shields.io/badge/Use-Educational%20Only-red.svg)](https://github.com)
[![Lab Environment](https://img.shields.io/badge/Environment-Isolated%20Lab-orange.svg)](https://github.com)
[![Security Research](https://img.shields.io/badge/Purpose-Defensive%20Security-green.svg)](https://github.com)

---

## ⚠️ CRITICAL NOTICE - EDUCATIONAL USE ONLY

**This repository contains powerful security testing tools designed EXCLUSIVELY for:**

✅ Personal isolated lab environments
✅ Authorized security awareness training
✅ Defensive security research
✅ Understanding attack techniques to build better defenses
✅ CTF competitions with explicit permission

❌ **PROHIBITED USES:**
- Unauthorized access to systems or accounts
- Phishing real users without explicit written consent
- Credential theft or unauthorized data collection
- Any illegal or malicious activity

**You are solely responsible for ensuring your use complies with all applicable laws and regulations.**

---

## 🎯 What's Included

This educational package provides a complete phishing simulation environment for learning defensive security:

### 📧 Professional Email Templates
- **Microsoft 365 Security Alert** - Realistic M365 security notification
- **Gmail Security Alert** - Google account security warning
- **Stripe Payment Failed** - Payment processing notification with full banking simulation

### 🌐 Realistic Landing Pages
- **Microsoft 365 Login** - Two-stage authentication flow
- **Gmail Login** - Material Design login interface
- **Stripe Payment Update** - Complete payment form with billing details

### 🔧 Evilginx Integration
- **Session Manager** - Capture and manage authentication sessions
- **Phishlet Configurations** - Pre-built configs for Microsoft, Google, Stripe
- **Proxy Manager** - Reverse proxy for session hijacking simulations

### 🤖 Automation Scripts
- **Template Installer** - One-command template deployment
- **Phishlet Generator** - Automated Evilginx configuration
- **Quick Start Script** - Complete environment setup

---

## 🚀 Quick Start (5 Minutes)

### Step 1: Run Setup Script

```bash
cd gophish
./scripts/quick_start.sh
```

This generates:
- ✅ 3 email templates
- ✅ 3 landing pages
- ✅ 3 Evilginx phishlets
- ✅ Complete documentation

### Step 2: Install Templates to Gophish

```bash
cd scripts
python3 setup_templates.py
```

Enter your Gophish API key when prompted (found in Settings > Account Settings).

### Step 3: Create Your First Campaign

1. Log into Gophish at `https://localhost:3333`
2. Create a **Sending Profile** (SMTP settings)
3. Create a **User Group** (test email addresses)
4. Create a **Campaign**:
   - Email Template: "Microsoft 365 - Security Alert"
   - Landing Page: "Microsoft 365 - Login Page"
   - URL: Your landing page URL
5. Click **Launch Campaign**

### Step 4: View Results

Monitor your dashboard to see:
- Email opened rates
- Link clicked rates
- Credentials submitted
- Timeline of events

---

## 📁 Project Structure

```
gophish/
├── templates/
│   ├── email/                          # Email templates
│   │   ├── microsoft365-security-alert.html
│   │   ├── gmail-security-alert.html
│   │   └── stripe-payment-failed.html
│   ├── landing-pages/                  # Landing pages
│   │   ├── microsoft365-login.html
│   │   ├── gmail-login.html
│   │   └── stripe-payment-update.html
│   ├── evilginx-configs/               # Evilginx phishlets
│   │   ├── microsoft365.yaml
│   │   ├── gmail.yaml
│   │   └── stripe.yaml
│   └── README.md                       # Template documentation
├── evilginx/                           # Evilginx Go modules
│   ├── helpers.go                      # URL encryption utilities
│   ├── session_manager.go              # Session capture & management
│   ├── phishlet_config.go              # Phishlet configuration
│   └── proxy_manager.go                # Reverse proxy management
├── scripts/                            # Automation scripts
│   ├── setup_templates.py              # Template installer
│   ├── generate_phishlets.py           # Phishlet generator
│   └── quick_start.sh                  # Complete setup automation
├── TEMPLATE_SETUP_GUIDE.md             # Detailed setup guide
└── EDUCATIONAL_TEMPLATES_README.md     # This file
```

---

## 🎓 What You'll Learn

### Understanding Attack Techniques

1. **Email Spoofing**
   - How attackers forge sender addresses
   - SPF, DKIM, DMARC bypass techniques
   - Visual deception in email design

2. **Credential Harvesting**
   - Form-based credential capture
   - Multi-stage authentication flows
   - Data exfiltration methods

3. **Session Hijacking**
   - Cookie theft and replay attacks
   - Token-based authentication bypass
   - Man-in-the-middle proxy techniques

4. **Social Engineering**
   - Urgency and fear tactics
   - Authority impersonation
   - Pretext development

### Building Better Defenses

Use this knowledge to:
- ✅ Train users to recognize phishing attempts
- ✅ Implement email filtering rules
- ✅ Develop phishing detection systems
- ✅ Configure security headers (CSP, HSTS, etc.)
- ✅ Deploy multi-factor authentication
- ✅ Create incident response procedures

---

## 🔬 Advanced Features

### Email Template Capabilities

**Dynamic Variables:**
```html
{{.FirstName}}      <!-- Recipient first name -->
{{.LastName}}       <!-- Recipient last name -->
{{.Email}}          <!-- Recipient email -->
{{.Position}}       <!-- Recipient position -->
{{.URL}}            <!-- Phishing link with tracking -->
{{.Tracker}}        <!-- Tracking pixel -->
{{.CurrentDate}}    <!-- Current date/time -->
```

**Professional Design:**
- Responsive HTML (mobile + desktop)
- Authentic branding and styling
- Realistic content and messaging
- Proper email formatting

### Landing Page Features

**Multi-Stage Capture:**
- JavaScript-driven multi-step forms
- Client-side validation
- Realistic error handling
- Auto-submission to Gophish

**Credential Types:**
- Username/password combinations
- Email addresses and phone numbers
- Full payment card details (Stripe template)
- Billing addresses and personal info

### Evilginx Integration

**Session Management:**
```go
// Create session manager
sm := evilginx.NewSessionManager()

// Capture session
session := &evilginx.Session{
    Username:  "user@example.com",
    Password:  "password123",
    Cookies:   map[string]string{...},
}
sm.AddSession(session)

// Export session
jsonData, _ := sm.ExportSession(session.ID)
```

**Phishlet Configuration:**
```go
// Load phishlet
phishlet := evilginx.Microsoft365Phishlet()

// Generate YAML config
yaml := phishlet.GeneratePhishletYAML()

// Register with proxy manager
pm := evilginx.NewProxyManager()
pm.RegisterPhishlet(phishlet)
```

**Reverse Proxy:**
```go
// Get proxy for phishlet
proxy, _ := pm.GetProxy("microsoft365", "login")

// Handle HTTP requests
http.Handle("/", proxy)
```

---

## 📚 Detailed Documentation

### For Template Details
👉 **[templates/README.md](templates/README.md)**
- Template specifications
- Customization guide
- Usage examples
- Defensive indicators

### For Complete Setup
👉 **[TEMPLATE_SETUP_GUIDE.md](TEMPLATE_SETUP_GUIDE.md)**
- Prerequisites
- Installation steps
- Campaign creation
- Troubleshooting
- Best practices

### For Evilginx Setup
1. Install Evilginx 2.3+
2. Copy phishlets: `cp templates/evilginx-configs/*.yaml /path/to/evilginx/phishlets/`
3. Configure domain and IP
4. Enable phishlets
5. Create lures
6. Use lure URLs in Gophish campaigns

---

## 🛠️ Technical Architecture

### How It Works

```
┌─────────────────┐
│  Target User    │
└────────┬────────┘
         │ 1. Receives email
         ▼
┌─────────────────┐
│   Gophish       │──── Email Template ────┐
│   Mail Server   │                         │
└────────┬────────┘                         │
         │ 2. Clicks link                   │
         ▼                                   ▼
┌─────────────────┐              ┌──────────────────┐
│   Evilginx      │◄─────────────│  Landing Page    │
│   Proxy         │  3. Proxies  │  (Credential     │
└────────┬────────┘     traffic  │   Capture)       │
         │                        └──────────────────┘
         │ 4. Captures session
         ▼
┌─────────────────┐
│  Real Service   │ (Microsoft, Google, etc.)
│  (Office 365,   │
│   Gmail, etc.)  │
└─────────────────┘
         │
         │ 5. Returns auth tokens
         ▼
┌─────────────────┐
│ Session Manager │
│ (Stores cookies,│
│  tokens, creds) │
└─────────────────┘
```

### Data Flow

1. **Email Sent**: Gophish sends templated email to target
2. **Link Clicked**: User clicks tracking link
3. **Proxy Intercept**: Evilginx proxies request to landing page
4. **Credential Entry**: User enters credentials on landing page
5. **Form Submit**: Credentials sent to Gophish & proxied to real service
6. **Session Capture**: Evilginx captures session cookies/tokens
7. **Data Storage**: All data stored in Gophish database + session manager

---

## 🔐 Security & Privacy

### Lab Environment Best Practices

#### Network Isolation
```bash
# Use VM with host-only networking
# No internet access for test VMs
# Separate network segment
```

#### Data Protection
- ❌ Never use real credentials
- ❌ Never target real users without authorization
- ✅ Use test accounts only
- ✅ Encrypt all stored data
- ✅ Delete data after testing

#### Access Control
- 🔒 Change default Gophish password
- 🔒 Use strong API keys
- 🔒 Enable TLS with valid certificates
- 🔒 Restrict admin access
- 🔒 Monitor access logs

### Legal Compliance

**Before using these tools, ensure:**
1. ✅ You have explicit written authorization
2. ✅ You're in a controlled environment
3. ✅ You comply with local laws (CFAA, GDPR, etc.)
4. ✅ You have incident response procedures
5. ✅ You maintain audit logs

---

## 🎯 Use Cases

### 1. Security Awareness Training

**Objective:** Train employees to recognize phishing

**Setup:**
1. Deploy templates in Gophish
2. Target test users with simulated campaigns
3. Track who clicks/submits credentials
4. Provide immediate training to those who fail
5. Measure improvement over time

### 2. Red Team Exercises

**Objective:** Test organizational defenses

**Setup:**
1. Get authorization from management
2. Configure Evilginx for session capture
3. Launch realistic campaign
4. Document defensive gaps
5. Provide remediation recommendations

### 3. Security Research

**Objective:** Understand attack techniques

**Setup:**
1. Create isolated lab environment
2. Deploy templates and Evilginx
3. Analyze traffic and behavior
4. Document findings
5. Develop detection signatures

### 4. Product Security Testing

**Objective:** Test email filters and web protections

**Setup:**
1. Deploy templates against test environment
2. Verify email filters catch phishing
3. Test web application firewalls
4. Validate user training effectiveness
5. Improve security controls

---

## 🐛 Troubleshooting

### Common Issues

#### Templates Not Installing
```bash
# Verify Gophish is running
curl -k https://localhost:3333

# Check API key
# Settings > Account Settings in Gophish UI

# Install manually via UI if script fails
```

#### Landing Pages Not Capturing
```bash
# Enable credential capture in page settings
# Verify "Capture Credentials" is checked
# Verify "Capture Passwords" is checked
# Check browser console for errors
```

#### Evilginx Not Working
```bash
# Verify DNS points to Evilginx server
nslookup login.yourdomain.com

# Check phishlet is enabled
phishlets

# Review Evilginx logs
# Look for proxy errors
```

#### Email Not Sending
```bash
# Test SMTP connection
telnet localhost 25

# Check Gophish logs
tail -f gophish.log

# Verify sending profile configuration
```

---

## 📖 Resources

### Official Documentation
- [Gophish Docs](https://docs.getgophish.com/)
- [Evilginx Docs](https://help.evilginx.com/)
- [OWASP Phishing Guide](https://owasp.org/www-community/attacks/Phishing)

### Learning Resources
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
- [MITRE ATT&CK - Phishing](https://attack.mitre.org/techniques/T1566/)
- [Anti-Phishing Working Group](https://apwg.org/)

### Security Research
- [PhishLabs Blog](https://www.phishlabs.com/blog/)
- [KnowBe4 Research](https://www.knowbe4.com/phishing)

---

## 🤝 Contributing

This is an educational project. If you'd like to contribute:

### Adding Templates
1. Create HTML template following existing format
2. Test thoroughly in isolated environment
3. Document template features
4. Include defensive indicators

### Improving Code
1. Follow Go best practices
2. Add comprehensive comments
3. Include unit tests
4. Update documentation

### Reporting Issues
- Use GitHub issues
- Provide detailed reproduction steps
- Include environment details
- Suggest potential fixes

---

## 📝 License

This educational project is provided for learning purposes. Use responsibly and ethically.

**Disclaimer:** The authors are not responsible for misuse of these tools. Users are solely responsible for ensuring their use complies with all applicable laws and regulations.

---

## 🙏 Acknowledgments

Built on top of:
- **Gophish** - Jordan Wright & contributors
- **Evilginx** - Kuba Gretzky
- **Go Community** - Amazing tools and libraries

---

## 📞 Support

### Questions?
1. Read [TEMPLATE_SETUP_GUIDE.md](TEMPLATE_SETUP_GUIDE.md)
2. Check [templates/README.md](templates/README.md)
3. Review troubleshooting section above

### Found a Bug?
- Create detailed issue report
- Include steps to reproduce
- Provide environment details

### Want to Learn More?
- Study the template code
- Analyze network traffic
- Read security research papers
- Practice in safe lab environment

---

## 🎓 Educational Objectives

By using these templates, you should learn:

✅ How phishing emails deceive users
✅ How credential harvesting works
✅ How session hijacking bypasses 2FA
✅ How to recognize phishing attempts
✅ How to implement email security
✅ How to detect and prevent phishing

**Use this knowledge to build better defenses and protect organizations from real threats.**

---

**Remember: With great power comes great responsibility. Use these tools ethically and legally.**

---

## 🚦 Quick Reference

### Start Gophish
```bash
./gophish
# Access: https://localhost:3333
```

### Install Templates
```bash
python3 scripts/setup_templates.py
```

### Generate Phishlets
```bash
python3 scripts/generate_phishlets.py
```

### Full Setup
```bash
./scripts/quick_start.sh
```

### View Logs
```bash
tail -f gophish.log
```

---

**Happy Learning! 🎓 Stay Ethical! 🛡️**
