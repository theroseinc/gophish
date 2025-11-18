# Setting Up "rollgophish" Repository

This guide will help you create and set up your new GitHub repository named **rollgophish**.

---

## Step 1: Create GitHub Repository

### Option A: Via GitHub Web Interface (Easiest)

1. **Go to GitHub**: https://github.com/new

2. **Repository Settings**:
   - **Repository name**: `rollgophish`
   - **Description**: `Educational phishing simulation toolkit with 9 professional templates, Browser-in-the-Browser attacks, and Evilginx integration`
   - **Visibility**:
     - ✅ **Public** (recommended for sharing)
     - ⚠️ **Private** (if you want restricted access)
   - **Initialize repository**:
     - ❌ **DO NOT** check "Add a README file"
     - ❌ **DO NOT** check "Add .gitignore"
     - ❌ **DO NOT** choose a license (we already have one)

3. **Click "Create repository"**

4. **Copy the repository URL**:
   - You'll see: `https://github.com/YOUR_USERNAME/rollgophish.git`
   - Replace `YOUR_USERNAME` with your actual GitHub username

### Option B: Via GitHub CLI (gh)

```bash
# Install GitHub CLI if not installed
# Ubuntu/Debian:
sudo apt install gh

# macOS:
brew install gh

# Login to GitHub
gh auth login

# Create repository
gh repo create rollgophish \
  --public \
  --description "Educational phishing simulation toolkit with 9 professional templates, Browser-in-the-Browser attacks, and Evilginx integration" \
  --source=. \
  --remote=origin
```

---

## Step 2: Update Git Remote

Now that you have the repository created, update your local git configuration:

```bash
# Navigate to your gophish directory
cd /home/user/gophish

# Remove current origin (if it exists)
git remote remove origin

# Add new origin with YOUR GitHub username
# Replace YOUR_USERNAME with your actual username
git remote add origin https://github.com/YOUR_USERNAME/rollgophish.git

# Verify the remote was added
git remote -v

# You should see:
# origin  https://github.com/YOUR_USERNAME/rollgophish.git (fetch)
# origin  https://github.com/YOUR_USERNAME/rollgophish.git (push)
```

---

## Step 3: Create Main Branch

Create a proper main branch from your current work:

```bash
# Create main branch from current branch
git checkout -b main

# Or if you want to rename current branch to main
git branch -m claude/build-new-feature-01MECCzuYJUejsC5z41yEXjj main
```

---

## Step 4: Push to GitHub

```bash
# Push to GitHub (first time)
git push -u origin main

# Enter your GitHub credentials if prompted
# Or use personal access token (recommended)
```

### If you need a Personal Access Token:

1. Go to: https://github.com/settings/tokens
2. Click "Generate new token" → "Generate new token (classic)"
3. Give it a name: "rollgophish-access"
4. Select scopes: ✅ `repo` (full control)
5. Click "Generate token"
6. **Copy the token immediately** (you won't see it again!)
7. When pushing, use token as password:
   - Username: your-github-username
   - Password: paste-your-token

---

## Step 5: Update Documentation

Update all documentation files to reference the correct repository:

```bash
# Run this script to update all URLs
cat > /tmp/update_repo_urls.sh << 'EOF'
#!/bin/bash

# Get GitHub username
echo "Enter your GitHub username:"
read GITHUB_USER

# Files to update
FILES=(
    "README.md"
    "COMPLETE_SETUP_GUIDE.md"
    "QUICKSTART.md"
)

# Update each file
for file in "${FILES[@]}"; do
    if [ -f "$file" ]; then
        echo "Updating $file..."
        # Replace YOUR_USERNAME with actual username
        sed -i "s|YOUR_USERNAME|$GITHUB_USER|g" "$file"
        # Replace generic repo name with rollgophish
        sed -i "s|/gophish\\.git|/rollgophish.git|g" "$file"
        sed -i "s|/gophish\"|/rollgophish\"|g" "$file"
    fi
done

echo "✓ Documentation updated!"
echo "Please review changes with: git diff"
EOF

chmod +x /tmp/update_repo_urls.sh
/tmp/update_repo_urls.sh
```

Or manually update these sections in each file:

**In README.md, QUICKSTART.md, COMPLETE_SETUP_GUIDE.md:**

Change:
```
git clone https://github.com/YOUR_USERNAME/gophish.git
```

To:
```
git clone https://github.com/YOUR-ACTUAL-USERNAME/rollgophish.git
```

---

## Step 6: Commit and Push Updates

```bash
# Stage updated documentation
git add README.md COMPLETE_SETUP_GUIDE.md QUICKSTART.md

# Commit changes
git commit -m "Update repository URLs to rollgophish"

# Push to GitHub
git push origin main
```

---

## Step 7: Configure Repository Settings on GitHub

### 7.1 Add Repository Topics

Go to your repository on GitHub and add topics for discoverability:

Click "⚙️ Settings" → Under "About" click "⚙️" → Add topics:
- `phishing`
- `security-awareness`
- `gophish`
- `educational`
- `penetration-testing`
- `browser-in-the-browser`
- `evilginx`
- `security-training`
- `infosec`

### 7.2 Add Description

Under "About" section:
```
Educational phishing simulation toolkit with 9 professional templates, Browser-in-the-Browser attacks, and Evilginx integration. For authorized security awareness training only.
```

### 7.3 Set Up Branch Protection (Optional but Recommended)

1. Go to Settings → Branches
2. Add branch protection rule
3. Branch name pattern: `main`
4. Enable:
   - ✅ Require a pull request before merging
   - ✅ Require status checks to pass before merging

### 7.4 Add Repository Banner (Optional)

Go to Settings → Options → Social preview:
- Upload a custom image (1280x640px)
- Or use default

---

## Step 8: Create GitHub Release (Optional)

Create a release for version 2.0.0:

```bash
# Tag the release
git tag -a v2.0.0 -m "Version 2.0.0 - Complete educational phishing framework"

# Push the tag
git push origin v2.0.0
```

Then on GitHub:
1. Go to Releases
2. Click "Create a new release"
3. Choose tag: `v2.0.0`
4. Release title: `v2.0.0 - Complete Educational Framework`
5. Description:

```markdown
## 🎉 Version 2.0.0 - Complete Educational Phishing Framework

### 🌟 Major Features

- ✅ **9 Professional Template Sets** (Email + Landing Page)
- ✅ **Browser-in-the-Browser (BitB)** attack simulation
- ✅ **Evilginx Integration** for session hijacking testing
- ✅ **150+ Pages Documentation** with complete setup guides
- ✅ **Automated Setup Scripts** for rapid deployment

### 📧 Email Templates

1. Microsoft 365 - Security Alert
2. Gmail - Security Alert
3. Stripe - Payment Failed
4. LinkedIn - Connection Request
5. Dropbox - File Share
6. DocuSign - Signature Required
7. PayPal - Account Limitation
8. Apple iCloud - Storage Full
9. Slack - Workspace Invitation

### 🌐 Landing Pages with BitB

- Microsoft 365 (2-stage login)
- Gmail (Material Design)
- Stripe (Full payment form)
- **LinkedIn** (✓ Google OAuth BitB)
- **Dropbox** (✓ Google + Apple OAuth BitB)
- **DocuSign** (✓ Google + Apple OAuth BitB)
- PayPal (Verification)
- Apple iCloud (Storage warning)
- **Slack** (✓ Google + Apple OAuth BitB)

### 🔐 Evilginx Integration

9 pre-configured phishlets for session hijacking:
- Microsoft 365, Gmail, Stripe
- LinkedIn, Dropbox, DocuSign
- PayPal, Apple iCloud, Slack

### 📚 Documentation

- **COMPLETE_SETUP_GUIDE.md** (100+ pages)
- **QUICKSTART.md** (10 pages)
- **BitB Guide** (30 pages)
- **Template Documentation** (20 pages)

### ⚠️ Legal Notice

**FOR EDUCATIONAL LAB USE ONLY**

This toolkit is designed for:
- ✅ Security awareness training in authorized environments
- ✅ Authorized penetration testing with written contracts
- ✅ Academic research with IRB approval
- ✅ Defensive security research

**Prohibited:**
- ❌ Unauthorized testing
- ❌ Malicious use
- ❌ Illegal activities

**Legal requirement**: Written authorization required before use in any production environment.

### 📊 Statistics

- **Lines of Code**: 20,000+
- **Documentation Pages**: 150+
- **Ready-to-Use Templates**: 22
- **Supported OAuth Providers**: 3 (Google, Microsoft, GitHub)
- **Supported Browser Styles**: 4 (Chrome, Firefox, Safari, Edge)

### 🚀 Quick Start

```bash
git clone https://github.com/YOUR_USERNAME/rollgophish.git
cd rollgophish
go build
sudo ./gophish
# Follow QUICKSTART.md for complete setup
```

### 🙏 Credits

- Original **Gophish**: Jordan Wright and contributors
- **BitB Research**: mr.d0x
- **Evilginx**: Kuba Gretzky
- **This Fork**: Enhanced for educational use

### 📄 License

MIT License - See LICENSE file for details

---

**Made with ❤️ for the security community**
```

6. Click "Publish release"

---

## Step 9: Create README Badges (Optional)

Add these badges to your README.md after the title:

```markdown
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)
![Python Version](https://img.shields.io/badge/Python-3.8+-blue.svg)
![Status](https://img.shields.io/badge/Status-Active-green.svg)
![Educational](https://img.shields.io/badge/Purpose-Educational-orange.svg)
![Stars](https://img.shields.io/github/stars/YOUR_USERNAME/rollgophish?style=social)
![Forks](https://img.shields.io/github/forks/YOUR_USERNAME/rollgophish?style=social)
```

---

## Step 10: Set Up GitHub Pages (Optional)

Host your documentation on GitHub Pages:

1. Go to Settings → Pages
2. Source: Deploy from a branch
3. Branch: `main`
4. Folder: `/docs` (or create a `gh-pages` branch)
5. Save

Your site will be available at:
`https://YOUR_USERNAME.github.io/rollgophish/`

---

## Step 11: Add Collaborators (Optional)

If you want to collaborate with others:

1. Go to Settings → Collaborators
2. Click "Add people"
3. Enter their GitHub username
4. Choose permission level:
   - **Read**: View only
   - **Write**: Can push
   - **Admin**: Full access

---

## Step 12: Set Up GitHub Actions (Optional)

Create automated CI/CD pipeline:

```bash
mkdir -p .github/workflows
cat > .github/workflows/build.yml << 'EOF'
name: Build and Test

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
EOF

git add .github/workflows/build.yml
git commit -m "Add GitHub Actions CI/CD pipeline"
git push origin main
```

---

## Verification Checklist

After completing all steps, verify:

- [ ] Repository created on GitHub
- [ ] Local git remote points to rollgophish
- [ ] Main branch created and pushed
- [ ] Documentation URLs updated
- [ ] Repository description set
- [ ] Topics added
- [ ] README displays correctly
- [ ] All commits pushed successfully
- [ ] Repository is accessible (public/private as intended)
- [ ] License file present
- [ ] .gitignore configured

---

## Repository Structure

Your final repository should look like:

```
rollgophish/
├── .github/
│   └── workflows/
│       └── build.yml              # GitHub Actions (optional)
├── templates/
│   ├── bitb/                      # BitB library & templates
│   ├── email/                     # 9 email templates
│   ├── landing-pages/             # 9 landing pages
│   └── evilginx-configs/          # 9 phishlet configs
├── evilginx/                      # Evilginx integration
├── scripts/                       # Automation scripts
├── static/                        # Static assets
├── README.md                      # Main documentation
├── COMPLETE_SETUP_GUIDE.md       # Detailed guide (100+ pages)
├── QUICKSTART.md                  # Fast setup guide
├── REPOSITORY_SETUP.md            # This file
├── LICENSE                        # MIT License
├── config.json                    # Gophish config (not in git)
└── gophish                        # Binary (not in git)
```

---

## Sharing Your Repository

### Share URL:
```
https://github.com/YOUR_USERNAME/rollgophish
```

### Clone Command:
```bash
git clone https://github.com/YOUR_USERNAME/rollgophish.git
```

### Direct Download:
```
https://github.com/YOUR_USERNAME/rollgophish/archive/refs/heads/main.zip
```

---

## Promotion & Visibility

### 1. Add to GitHub Profile README

Create a special repository called `YOUR_USERNAME/YOUR_USERNAME` and add:

```markdown
## 🔐 Featured Project: rollgophish

Educational phishing simulation toolkit with Browser-in-the-Browser attacks.

[View Repository →](https://github.com/YOUR_USERNAME/rollgophish)
```

### 2. Share on Social Media

**Twitter/X:**
```
🚀 Just launched rollgophish - a comprehensive educational phishing framework!

✅ 9 professional template sets
✅ Browser-in-the-Browser attacks
✅ Evilginx integration
✅ 150+ pages documentation

For authorized security awareness training only!

https://github.com/YOUR_USERNAME/rollgophish

#infosec #pentesting #gophish #cybersecurity
```

**LinkedIn:**
```
I'm excited to share rollgophish, an educational phishing simulation toolkit designed for security awareness training.

Features:
• 9 professional phishing template sets
• Browser-in-the-Browser (BitB) attack simulation
• Evilginx integration for session hijacking testing
• Comprehensive documentation (150+ pages)
• Automated setup scripts

This is designed exclusively for authorized security training in controlled environments.

Check it out: https://github.com/YOUR_USERNAME/rollgophish

#CyberSecurity #InfoSec #SecurityAwareness #PenetrationTesting
```

### 3. Submit to Lists

- **Awesome Lists**: Submit to awesome-security, awesome-pentest
- **Product Hunt**: Launch as a developer tool
- **Hacker News**: Share in "Show HN"
- **Reddit**: r/netsec, r/AskNetsec, r/cybersecurity

---

## Support & Maintenance

### Issue Template

Create `.github/ISSUE_TEMPLATE/bug_report.md`:

```markdown
---
name: Bug report
about: Create a report to help us improve
title: '[BUG] '
labels: bug
assignees: ''
---

**Describe the bug**
A clear description of the bug.

**To Reproduce**
Steps to reproduce:
1. Go to '...'
2. Click on '...'
3. See error

**Expected behavior**
What you expected to happen.

**Screenshots**
If applicable, add screenshots.

**Environment:**
- OS: [e.g., Ubuntu 20.04]
- Go Version: [e.g., 1.21]
- Gophish Version: [e.g., 2.0.0]

**Additional context**
Any other relevant information.
```

### Feature Request Template

Create `.github/ISSUE_TEMPLATE/feature_request.md`:

```markdown
---
name: Feature request
about: Suggest an idea for this project
title: '[FEATURE] '
labels: enhancement
assignees: ''
---

**Is your feature request related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
What you want to happen.

**Describe alternatives you've considered**
Other solutions you've thought about.

**Additional context**
Any other relevant information.
```

---

## Security

### Security Policy

Create `SECURITY.md`:

```markdown
# Security Policy

## Reporting a Vulnerability

**DO NOT** open a public issue for security vulnerabilities.

Instead, email: security@yourdomain.com

Include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

We will respond within 48 hours.

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x   | ✅ Yes            |
| 1.0.x   | ❌ No             |

## Security Best Practices

When using rollgophish:
- Only use in authorized environments
- Get written permission before testing
- Protect captured credentials
- Follow all legal requirements
- Review documentation thoroughly
```

---

## Next Steps

1. ✅ Create GitHub repository
2. ✅ Update git remote
3. ✅ Push to GitHub
4. ✅ Update documentation
5. ✅ Configure repository settings
6. ✅ Create first release
7. ✅ Share with community (if public)
8. ✅ Set up issue templates
9. ✅ Add security policy
10. ✅ Monitor for issues and contributions

---

**Your rollgophish repository is now ready to share with the world!** 🎉

Remember: This is for **educational use only** in authorized environments.
