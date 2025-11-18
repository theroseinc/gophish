# Browser-in-the-Browser (BitB) Attack Templates
## Educational Security Testing

⚠️ **FOR EDUCATIONAL LAB USE ONLY** - Isolated environments with explicit authorization

---

## What is Browser-in-the-Browser?

**Browser-in-the-Browser (BitB)** is an advanced phishing technique that creates fake OAuth/SSO popup windows using pure HTML, CSS, and JavaScript. Unlike traditional phishing pages, BitB:

- **Mimics browser chrome** - Fake address bar, SSL lock, window controls
- **Displays legitimate URLs** - Shows the real OAuth provider domain
- **Appears as a popup** - Looks like a separate browser window
- **Highly convincing** - Users trust OAuth popups more than embedded forms

### Why It's Effective

1. **Trust in OAuth** - Users expect OAuth popups when signing into third-party apps
2. **Visual authenticity** - Perfect replication of browser UI
3. **URL confidence** - Shows "https://accounts.google.com" in fake address bar
4. **Legitimate context** - Embedded within what appears to be a real application

---

## 📁 What's Included

### Core Library

**`lib/bitb.js`** - JavaScript library for creating fake browser windows
- Supports Chrome, Firefox, Safari, and Edge styles
- Draggable windows
- Customizable dimensions and content
- Event handling for close/submit

**`lib/bitb.css`** - Comprehensive styling
- Pixel-perfect browser chrome replication
- Multiple browser themes
- Responsive design
- Dark mode support

### Pre-built OAuth Templates

#### 1. **Google OAuth** (`google/google-oauth-bitb.html`)
- Mimics Google accounts.google.com OAuth flow
- Chrome browser style
- Full credential capture form
- "Sign in with Google" trigger button

#### 2. **Microsoft OAuth** (`microsoft/microsoft-oauth-bitb.html`)
- Mimics Microsoft login.microsoftonline.com flow
- Edge browser style
- Multi-step authentication flow
- "Sign in with Microsoft" trigger button

#### 3. **GitHub OAuth** (`github/github-oauth-bitb.html`)
- Mimics GitHub OAuth authorization flow
- Dark theme matching GitHub
- Permission scope display
- "Sign in with GitHub" trigger button

---

## 🚀 How to Use

### Method 1: Use Pre-built Templates

1. **Host the template** on your Gophish server or web server

2. **Send the link** in your phishing campaign:
   ```
   https://your-domain.com/templates/bitb/google/google-oauth-bitb.html
   ```

3. **User clicks "Sign in with Google"** - BitB window appears

4. **Credentials captured** when user submits the form

### Method 2: Embed in Existing Templates

Add BitB to any existing landing page:

```html
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="path/to/bitb.css">
</head>
<body>
    <h1>Welcome to Our Platform</h1>

    <!-- Trigger Button -->
    <button
        data-bitb-trigger
        data-bitb-title="Sign in - Google Accounts"
        data-bitb-url="https://accounts.google.com/signin/oauth"
        data-bitb-width="500"
        data-bitb-height="600"
        data-bitb-browser="chrome"
        data-bitb-content="oauth-content">
        Sign in with Google
    </button>

    <!-- Hidden Content -->
    <div id="oauth-content" style="display: none;">
        <!-- Your OAuth form HTML here -->
    </div>

    <script src="path/to/bitb.js"></script>
</body>
</html>
```

### Method 3: Programmatic Use

Create BitB windows dynamically:

```javascript
const bitb = new BrowserInBrowser({
    title: 'Sign in - Google Accounts',
    url: 'https://accounts.google.com/signin/oauth',
    width: 500,
    height: 600,
    browser: 'chrome', // chrome, firefox, safari, edge
    content: '<div>Your OAuth form HTML</div>',
    draggable: true,
    onClose: function() {
        console.log('Window closed');
    }
});

bitb.show();
```

---

## 🎨 Customization Options

### Data Attributes (HTML Method)

| Attribute | Description | Example |
|-----------|-------------|---------|
| `data-bitb-trigger` | Marks element as trigger | Required |
| `data-bitb-title` | Window title | "Sign in - Google Accounts" |
| `data-bitb-url` | Fake URL in address bar | "https://accounts.google.com" |
| `data-bitb-width` | Window width in pixels | 500 |
| `data-bitb-height` | Window height in pixels | 600 |
| `data-bitb-browser` | Browser style | "chrome", "firefox", "safari", "edge" |
| `data-bitb-content` | ID of content element | "oauth-content" |

### JavaScript Options

```javascript
{
    title: 'Window title',           // String
    url: 'https://example.com',      // String (shown in address bar)
    width: 500,                      // Number (pixels)
    height: 600,                     // Number (pixels)
    browser: 'chrome',               // 'chrome', 'firefox', 'safari', 'edge'
    content: '<div>HTML</div>',      // HTML string
    draggable: true,                 // Boolean
    theme: 'light',                  // 'light' or 'dark'
    onClose: function() {}           // Callback function
}
```

---

## 🎯 Integration with Gophish

### Step 1: Create Landing Page in Gophish

1. Go to **Landing Pages → New Page**
2. Name: "BitB - Google OAuth"
3. Import HTML: Copy content from `google/google-oauth-bitb.html`
4. Enable **Capture Credentials**
5. Enable **Capture Passwords**

### Step 2: Host Static Assets

The BitB library needs to be accessible:

**Option A: Host on Gophish server**
```bash
cp -r templates/bitb/lib /opt/gophish/static/bitb/
```

Update HTML paths:
```html
<link rel="stylesheet" href="/static/bitb/bitb.css">
<script src="/static/bitb/bitb.js"></script>
```

**Option B: Use CDN or external hosting**
```html
<link rel="stylesheet" href="https://your-cdn.com/bitb.css">
<script src="https://your-cdn.com/bitb.js"></script>
```

### Step 3: Create Campaign

1. **Email Template**: Use existing templates with OAuth prompt
2. **Landing Page**: Select your BitB landing page
3. **URL**: Your Gophish tracking URL
4. Launch campaign!

---

## 🔍 How to Detect BitB Attacks

As a defender, teach users to look for:

### Red Flags

1. **URL mismatch**: The page URL doesn't match the OAuth provider
   - Real: `https://accounts.google.com`
   - Fake: `https://evil-site.com` (but shows google.com in fake bar)

2. **Browser address bar**: Check the REAL browser address bar, not the popup
   - Press F11 to enter fullscreen - fake windows will overlay weirdly
   - Right-click and "Inspect Element" - real popups can't be inspected this way

3. **Window behavior**: Try to drag the window outside the browser
   - Real popups can move anywhere on screen
   - BitB windows are confined to the browser viewport

4. **Developer tools**: Press F12
   - Real OAuth popups are separate windows
   - BitB windows are just HTML divs - visible in DevTools

5. **URL copy test**: Try to select and copy the URL
   - Real address bars: URL is selectable
   - Fake address bars: Text might not be selectable or copies differently

### Technical Detection

For security teams:

```javascript
// Detect if page is trying to create fake browser windows
const suspiciousElements = document.querySelectorAll('.bitb-window, .fake-browser, [class*="popup-chrome"]');

if (suspiciousElements.length > 0) {
    console.warn('Possible BitB attack detected!');
}
```

Content Security Policy (CSP) headers:
```
Content-Security-Policy: frame-ancestors 'none';
```

---

## 🛡️ Defensive Measures

### For Users

1. **Check the real address bar** - Always verify the URL in your browser's address bar
2. **Avoid OAuth popups** - Use browser extensions or bookmark direct login pages
3. **Enable 2FA** - Even if credentials are stolen, 2FA provides protection
4. **Verify SSL certificate** - Click the padlock in the REAL address bar

### For Organizations

1. **Security Awareness Training** - Train users to recognize BitB attacks
2. **Email Filtering** - Block emails linking to suspicious domains
3. **Browser Extensions** - Deploy anti-phishing browser extensions
4. **FIDO2/WebAuthn** - Implement hardware-based authentication
5. **Monitor for abuse** - Track domains mimicking your OAuth flows

### For Developers

1. **Implement CSP** - Prevent framing and inline scripts
2. **CORS policies** - Restrict cross-origin requests
3. **Subresource Integrity** - Verify external resources haven't been tampered
4. **OAuth state parameter** - Prevent CSRF attacks
5. **Monitor OAuth callbacks** - Log and alert on unusual patterns

---

## 📖 Educational Use Cases

### Scenario 1: Security Awareness Training

**Objective**: Teach employees to recognize OAuth phishing

**Setup**:
1. Deploy BitB template targeting your organization
2. Send to test group with clear "training exercise" notice
3. Track who clicks and submits credentials
4. Provide immediate training to those who fail
5. Measure improvement over time

### Scenario 2: Penetration Testing

**Objective**: Test organizational defenses against advanced phishing

**Setup**:
1. Get written authorization from client
2. Deploy BitB attacking their SaaS applications
3. Document who falls for the attack
4. Test email filters, web proxies, user behavior
5. Provide detailed remediation report

### Scenario 3: Security Research

**Objective**: Study effectiveness of BitB vs traditional phishing

**Setup**:
1. Create isolated lab environment
2. Deploy both BitB and traditional phishing pages
3. Conduct A/B testing with volunteers
4. Measure success rates and detection rates
5. Publish findings to improve defenses

---

## 🎓 How BitB Works Technically

### Architecture

```
┌─────────────────────────────────────┐
│  Legitimate-looking Landing Page     │
│  (yoursite.com)                      │
│                                      │
│  ┌────────────────────────────────┐ │
│  │ "Sign in with Google" Button   │ │ ← User clicks
│  └────────────────────────────────┘ │
│                                      │
│  JavaScript triggers BitB window     │
│                                      │
│  ┌────────────────────────────────┐ │
│  │ Fake Browser Window (HTML/CSS)  │ │
│  │ ┌────────────────────────────┐ │ │
│  │ │ Fake Address Bar            │ │ │
│  │ │ https://accounts.google.com │ │ │
│  │ └────────────────────────────┘ │ │
│  │                                 │ │
│  │  [OAuth Login Form]             │ │ ← User enters creds
│  │  Email: _________________       │ │
│  │  Password: ______________       │ │
│  │  [Sign In]                      │ │
│  │                                 │ │
│  └────────────────────────────────┘ │
│                                      │
│  Credentials sent to attacker        │
└─────────────────────────────────────┘
```

### Key Components

1. **Backdrop** - Dark overlay blocking page interaction
2. **Fake Browser Chrome** - HTML/CSS replication of browser UI
3. **Fake Address Bar** - Shows legitimate OAuth provider URL
4. **Content Area** - OAuth login form
5. **Event Handlers** - Capture credentials on submit

### Why Users Fall for It

- **Visual Trust**: Looks exactly like real OAuth popup
- **URL Confidence**: Shows correct domain in (fake) address bar
- **Expected Behavior**: OAuth popups are normal/expected
- **SSL Icon**: Green padlock visible in fake chrome
- **Professional Design**: Pixel-perfect replication

---

## ⚠️ Legal & Ethical Considerations

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

### Best Practices

1. **Get written authorization** before testing
2. **Scope clearly defined** in testing agreements
3. **Notify users** after training exercises
4. **Secure data handling** - encrypt and delete captured credentials
5. **Follow laws** - GDPR, CCPA, CFAA compliance

---

## 🔬 Research & References

### Original Research

- **Browser-in-the-Browser (BITB) Attack** - [mr.d0x](https://mrd0x.com/browser-in-the-browser-phishing-attack/)
- **OAuth Phishing Techniques** - Various security researchers
- **UI Redressing Attacks** - Academic papers on clickjacking and similar

### Detection Research

- **FIDO Alliance** - Hardware-based authentication
- **W3C WebAuthn** - Phishing-resistant authentication
- **Google Advanced Protection** - Case studies

### Defense Frameworks

- **NIST Cybersecurity Framework**
- **OWASP Top 10**
- **MITRE ATT&CK** - T1566 Phishing

---

## 🛠️ Troubleshooting

### Issue: BitB window doesn't appear

**Solutions**:
- Check browser console for JavaScript errors
- Verify bitb.js and bitb.css are loading correctly
- Ensure data attributes are spelled correctly
- Check if content ID matches the hidden div

### Issue: Styling looks broken

**Solutions**:
- Verify bitb.css is loading before bitb.js
- Check for CSS conflicts with existing styles
- Clear browser cache
- Try different browser

### Issue: Window not draggable

**Solutions**:
- Check `draggable: true` option is set
- Verify titlebar element exists
- Check for JavaScript errors preventing event binding

### Issue: Forms not submitting

**Solutions**:
- Ensure form has proper action and method
- Check Gophish landing page has "Capture Credentials" enabled
- Verify form field names match expected values
- Test form submission in browser DevTools

---

## 📞 Support

### Questions?

1. Read this README thoroughly
2. Check TEMPLATE_SETUP_GUIDE.md in root directory
3. Review example templates in subdirectories
4. Test in isolated lab environment

### Reporting Issues

If you find bugs or have suggestions:
- Document the issue with screenshots
- Provide browser and OS information
- Include steps to reproduce
- Note any console errors

---

## 🎯 Quick Reference

### Supported Browsers Styles

- **Chrome** - Modern Chromium style
- **Firefox** - Mozilla Firefox style
- **Safari** - macOS Safari style
- **Edge** - Microsoft Edge style

### Recommended Dimensions

| Content Type | Width | Height |
|--------------|-------|--------|
| Simple Login | 450px | 500px |
| OAuth + Permissions | 500px | 650px |
| Multi-step Auth | 550px | 700px |

### Integration Checklist

- [ ] BitB library files (CSS + JS) uploaded to server
- [ ] HTML paths updated to reference library files
- [ ] Landing page created in Gophish
- [ ] Credential capture enabled
- [ ] Form fields use correct names (email, password)
- [ ] Test in multiple browsers
- [ ] Verify credentials are captured in Gophish

---

**Remember: Use these templates responsibly for education and defense, never for harm.**

*Last updated: 2024 - For educational security testing only*
