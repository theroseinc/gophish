# Gophish Educational Phishing Framework

<p align="center">
  <img src="https://raw.githubusercontent.com/gophish/gophish/master/static/images/gophish_purple.png" alt="Gophish Logo" width="200"/>
</p>

<p align="center">
  <strong>Comprehensive phishing simulation toolkit for security awareness training</strong>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#documentation">Documentation</a> •
  <a href="#templates">Templates</a> •
  <a href="#legal">Legal</a>
</p>

---

## ⚠️ CRITICAL NOTICE

**FOR EDUCATIONAL LAB USE ONLY**

This repository contains advanced phishing simulation tools designed for:
- ✅ Security awareness training in authorized environments
- ✅ Authorized penetration testing with written contracts
- ✅ Academic research with IRB approval
- ✅ Defensive security research in isolated labs

**PROHIBITED USE:**
- ❌ Unauthorized testing of third-party systems
- ❌ Credential theft for malicious purposes
- ❌ Any illegal activity or harm

**Legal Requirement:** Written authorization required before use in any production environment.

---

## Overview

This is an enhanced fork of [Gophish](https://github.com/gophish/gophish) - the open-source phishing toolkit - extended with:

- **9 Professional Phishing Template Sets** (Email + Landing Page)
- **Browser-in-the-Browser (BitB) Attack Simulation**
- **Evilginx Integration** for session hijacking testing
- **Automated Setup Scripts** for rapid deployment
- **Comprehensive Documentation** with step-by-step guides

### What Makes This Different?

| Feature | Standard Gophish | This Repository |
|---------|-----------------|-----------------|
| Email Templates | Basic | 9 Professional Sets |
| Landing Pages | Basic | 9 Realistic Pages |
| BitB Support | ❌ | ✅ 4 Pages Integrated |
| Evilginx Configs | ❌ | ✅ 9 Phishlets |
| Automation Scripts | ❌ | ✅ Python + Bash |
| Setup Documentation | Basic | 100+ Page Guide |
| OAuth Simulation | ❌ | ✅ BitB Popups |

---

## Features

### 🎨 Professional Templates

**9 Complete Template Sets:**

1. **Microsoft 365** - Security alert with 2-stage login
2. **Gmail** - Google security notification
3. **Stripe** - Payment failure with full card capture
4. **LinkedIn** - Connection request with Google OAuth (BitB)
5. **Dropbox** - File sharing with Google/Apple OAuth (BitB)
6. **DocuSign** - Signature required with Google/Apple OAuth (BitB)
7. **PayPal** - Account limitation notice
8. **Apple iCloud** - Storage full notification
9. **Slack** - Workspace invitation with Google/Apple OAuth (BitB)

Each set includes:
- Professionally designed email template
- Matching landing page with credential capture
- Responsive mobile-friendly design
- Realistic branding and styling

### 🌐 Browser-in-the-Browser (BitB)

**What is BitB?**
- Creates fake OAuth popup windows using HTML/CSS/JavaScript
- Displays legitimate URLs (e.g., `accounts.google.com`) in fake address bar
- Shows fake SSL padlock and browser chrome
- Highly convincing - users trust OAuth popups

**BitB Features:**
- ✅ 4 Browser styles: Chrome, Firefox, Safari, Edge
- ✅ Draggable fake windows
- ✅ Pre-built OAuth templates: Google, Microsoft, GitHub
- ✅ Integrated into 4 landing pages
- ✅ Data-attribute driven (easy to use)
- ✅ Comprehensive documentation

**Integrated Landing Pages:**
- DocuSign (Google + Apple OAuth)
- LinkedIn (Google OAuth)
- Dropbox (Google + Apple OAuth)
- Slack (Google + Apple OAuth)

### 🔐 Evilginx Integration (Advanced)

**Session Hijacking Capabilities:**
- Man-in-the-middle reverse proxy
- Captures authentication cookies and tokens
- Bypasses 2FA/MFA
- 9 pre-configured phishlets included

**Phishlet Configurations:**
- Microsoft 365, Gmail, Stripe
- LinkedIn, Dropbox, DocuSign
- PayPal, Apple iCloud, Slack

**Note:** Requires separate Evilginx installation. See [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) for details.

---

## Quick Start

### Prerequisites

- **OS**: Linux (Ubuntu 20.04+) or macOS
- **Go**: Version 1.19+
- **Python**: Version 3.8+
- **Git**: Latest version

### 5-Minute Local Setup

```bash
# 1. Clone repository
git clone https://github.com/YOUR_USERNAME/gophish.git
cd gophish

# 2. Build Gophish
go build

# 3. Generate admin certificate
openssl req -newkey rsa:4096 -nodes -sha256 -keyout gophish_admin.key \
  -x509 -days 365 -out gophish_admin.crt \
  -subj "/C=US/ST=State/L=City/O=Org/CN=localhost"

# 4. Start Gophish
sudo ./gophish
# Copy the admin password from output

# 5. Access admin panel
# Open: https://127.0.0.1:3333
# Login: admin / [password from step 4]

# 6. Install templates
pip3 install requests colorama
cd scripts
nano setup_templates.py  # Add your API key from Gophish Settings
python3 setup_templates.py

# 7. Setup BitB
cd ..
cp -r templates/bitb static/
```

**That's it!** You now have a fully functional phishing simulation environment.

### Quick Test Campaign

1. Create a test group with your email
2. Select any email + landing page template
3. Launch campaign
4. Check your email and click the link
5. View results in dashboard

---

## Documentation

### 📚 Available Guides

| Document | Purpose | Pages | Audience |
|----------|---------|-------|----------|
| [QUICKSTART.md](QUICKSTART.md) | Fast setup | 10 | Everyone |
| [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) | Detailed walkthrough | 100+ | Detailed setup |
| [EDUCATIONAL_TEMPLATES_README.md](EDUCATIONAL_TEMPLATES_README.md) | Template overview | 20 | Template users |
| [templates/bitb/README.md](templates/bitb/README.md) | BitB guide | 30 | Security researchers |

### 🎯 Choose Your Path

**New to Gophish?**
→ Start with [QUICKSTART.md](QUICKSTART.md)

**Need detailed instructions?**
→ Read [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md)

**Want to understand BitB?**
→ See [templates/bitb/README.md](templates/bitb/README.md)

**Setting up production?**
→ Follow [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) production section

---

## Templates

### Template Sets (9 Total)

Each set includes both email and landing page:

| # | Service | Email Template | Landing Page | BitB |
|---|---------|----------------|--------------|------|
| 1 | Microsoft 365 | Security alert | 2-stage login | ❌ |
| 2 | Gmail | Security alert | Material design | ❌ |
| 3 | Stripe | Payment failed | Full card form | ❌ |
| 4 | LinkedIn | Connection request | Professional | ✅ Google |
| 5 | Dropbox | File share | Cloud storage | ✅ Google + Apple |
| 6 | DocuSign | Signature needed | Document platform | ✅ Google + Apple |
| 7 | PayPal | Account limitation | Payment verification | ❌ |
| 8 | Apple iCloud | Storage full | Apple ID login | ❌ |
| 9 | Slack | Workspace invite | Team collaboration | ✅ Google + Apple |

**Total: 9 email templates + 9 landing pages + 4 BitB integrations = 22 components**

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Gophish Framework                        │
├─────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ Admin Panel  │  │ Phish Server │  │ SMTP Server  │      │
│  │ Port 3333    │  │ Port 80/443  │  │ Port 25/587  │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
├─────────────────────────────────────────────────────────────┤
│                     Template System                          │
│  • 9 Email Templates                                         │
│  • 9 Landing Pages (4 with BitB)                            │
│  • BitB Library (JavaScript + CSS)                          │
│  • 9 Evilginx Phishlet Configs                             │
├─────────────────────────────────────────────────────────────┤
│                  Optional: Evilginx                          │
│  • Reverse Proxy + Session Hijacking                        │
│  • Cookie/Token Capture                                      │
│  • 2FA Bypass                                               │
└─────────────────────────────────────────────────────────────┘
```

---

## BitB (Browser-in-the-Browser) Explained

### What is BitB?

Browser-in-the-Browser is an advanced phishing technique that creates fake OAuth/SSO popup windows using HTML, CSS, and JavaScript.

**How It Works:**
1. User clicks "Sign in with Google" button
2. Fake browser window appears (pure HTML/CSS)
3. Fake address bar shows `https://accounts.google.com`
4. Fake SSL padlock and browser chrome displayed
5. User enters credentials thinking it's real OAuth
6. Credentials captured by your phishing page

**Why It's Effective:**
- ✅ Users trust OAuth popups more than embedded forms
- ✅ Shows legitimate URLs in fake address bar
- ✅ Pixel-perfect browser UI replication
- ✅ Appears as separate window
- ✅ Highly convincing visual authenticity

**Integrated Pages:**
- DocuSign, LinkedIn, Dropbox, Slack

**Documentation:** [templates/bitb/README.md](templates/bitb/README.md)

---

## Use Cases

### 1. Security Awareness Training

Train employees to recognize phishing attacks.

**Process:**
1. Deploy in isolated environment
2. Select relevant templates
3. Send to test group
4. Track clicks and submissions
5. Provide immediate training
6. Measure improvement

### 2. Penetration Testing

Test organizational defenses with written authorization.

**Process:**
1. Get signed authorization
2. Deploy with BitB for advanced simulation
3. Test email filters and user awareness
4. Document findings
5. Provide remediation report

### 3. Research & Development

Study phishing effectiveness in controlled environment.

**Process:**
1. Create isolated lab
2. Test different techniques
3. A/B test traditional vs BitB
4. Measure effectiveness
5. Publish findings

---

## Security Considerations

### Server Security

- ✅ Firewall configured (UFW)
- ✅ SSL/TLS encryption
- ✅ Admin panel access restricted
- ✅ SSH hardened
- ✅ Database encrypted

### Campaign Security

- ✅ Written authorization required
- ✅ Immediate debrief after campaign
- ✅ No disciplinary action
- ✅ Data encryption
- ✅ Legal compliance (GDPR, CCPA, CFAA)

### Data Protection

- ✅ Secure file permissions
- ✅ Encrypted backups
- ✅ Proper data retention
- ✅ Delete credentials after use

---

## Legal & Ethical Use

### ✅ Acceptable Use

- Personal isolated lab environments
- Authorized penetration testing with signed contracts
- Security awareness training for your own organization
- Academic research with IRB approval
- Defensive security research

### ❌ Prohibited Use

- Phishing real users without authorization
- Credential theft for malicious purposes
- Testing third-party systems without permission
- Violating computer fraud and abuse laws (CFAA, etc.)
- Any illegal activity

### Legal Requirements

**Before ANY phishing simulation:**
1. Get written approval from organization leadership
2. Define scope (who, what, when, how)
3. Define data handling procedures
4. Ensure legal compliance
5. Have incident response plan

**⚠️ Consult a lawyer before conducting phishing simulations in production.**

---

## Installation Methods

### Method 1: Quick Local Setup (5 minutes)

```bash
git clone https://github.com/YOUR_USERNAME/gophish.git
cd gophish
go build
sudo ./gophish
```

See [QUICKSTART.md](QUICKSTART.md) for details.

### Method 2: Production Setup (30 minutes)

Full setup with SSL, domain, email server.

See [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) production section.

### Method 3: Docker Deployment

```bash
docker build -t gophish .
docker run -d -p 3333:3333 -p 80:80 gophish
```

See [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) Docker section.

---

## Troubleshooting

### Common Issues

**Can't access admin panel:**
```bash
ps aux | grep gophish  # Check if running
tail -f gophish.log    # Check logs
```

**Landing pages show 404:**
- Ensure `phish_server.listen_url` is `0.0.0.0:80` not `127.0.0.1:80`

**BitB doesn't work:**
```bash
ls -lh static/bitb/lib/  # Check files exist
cp -r templates/bitb static/  # Copy if missing
```

**Emails not sending:**
- Test SMTP profile in Gophish UI
- Check DNS records (SPF, DKIM, DMARC)
- Verify email service configuration

**Full troubleshooting:** See [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) troubleshooting section.

---

## Contributing

We welcome contributions!

### Ways to Contribute

- 📧 **New Templates** - Add more phishing templates
- 🐛 **Bug Fixes** - Report and fix bugs
- 📖 **Documentation** - Improve guides
- ✨ **Features** - Propose new capabilities
- 🧪 **Testing** - Test and provide feedback

### Process

1. Fork the repository
2. Create feature branch: `git checkout -b feature/new-template`
3. Make changes and test
4. Commit: `git commit -m "Add template for X"`
5. Push: `git push origin feature/new-template`
6. Create Pull Request

---

## Roadmap

### Version 2.1 (Planned)

- [ ] Additional BitB browser styles
- [ ] More OAuth provider simulations (Facebook, Twitter)
- [ ] Web-based BitB configurator
- [ ] Mobile app phishing templates

### Version 3.0 (Future)

- [ ] AI-powered template generation
- [ ] Multi-language support
- [ ] Advanced analytics dashboard
- [ ] Automated campaign optimization

---

## Credits & Acknowledgments

### Original Project
- **Gophish**: [Jordan Wright](https://github.com/jordan-wright) and contributors
- **License**: MIT

### Enhancements
- **BitB Research**: [mr.d0x](https://mrd0x.com/)
- **Evilginx**: [Kuba Gretzky](https://github.com/kgretzky)
- **This Fork**: Enhanced by security community

### Resources
- OWASP Testing Guide
- Security research papers
- Community contributions

---

## Support

### Documentation
- [QUICKSTART.md](QUICKSTART.md) - Fast setup
- [COMPLETE_SETUP_GUIDE.md](COMPLETE_SETUP_GUIDE.md) - Detailed guide
- [templates/bitb/README.md](templates/bitb/README.md) - BitB guide

### Community
- **GitHub Issues**: Bug reports and feature requests
- **Gophish Slack**: General questions
- **Official Docs**: https://docs.getgophish.com/

### Professional Support
- Custom template development
- Training and consulting
- Paid support available

---

## License

```
Gophish - Open-Source Phishing Framework

The MIT License (MIT)

Copyright (c) 2013 - 2020 Jordan Wright

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

**Additional Content (Templates, BitB, Documentation):**
- Created for educational purposes
- Free to use in authorized environments
- Attribution appreciated

**⚠️ DISCLAIMER**: The authors are not responsible for misuse of this tool. Only use in authorized environments with explicit written permission.

---

## Badges

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)
![Python Version](https://img.shields.io/badge/Python-3.8+-blue.svg)
![Status](https://img.shields.io/badge/Status-Active-green.svg)
![Educational](https://img.shields.io/badge/Purpose-Educational-orange.svg)

---

## What's Included?

```
✓ 9 email templates
✓ 9 landing pages
✓ 4 BitB-integrated pages
✓ 9 Evilginx phishlet configs
✓ BitB library (JS + CSS)
✓ 3 pre-built BitB OAuth templates
✓ Python automation scripts
✓ 100+ page setup guide
✓ Quick start guide
✓ Comprehensive BitB documentation
```

**Total Lines of Code**: 20,000+
**Documentation Pages**: 150+
**Ready-to-Use Templates**: 22

---

<p align="center">
  <strong>⚠️ FOR EDUCATIONAL LAB USE ONLY ⚠️</strong><br>
  Use responsibly. Get authorization. Train ethically.
</p>

<p align="center">
  Made with ❤️ for the security community
</p>

---

**Version 2.0.0** | **Last Updated: 2024** | **Educational Use Only**
