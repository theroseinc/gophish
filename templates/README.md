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
