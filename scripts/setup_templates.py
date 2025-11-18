#!/usr/bin/env python3
"""
Gophish Template Installer
Automatically imports email and landing page templates into Gophish
"""

import os
import json
import requests
import sys
from pathlib import Path

# Configuration
GOPHISH_API_URL = "https://localhost:3333/api"
API_KEY = ""  # Set this or pass as environment variable

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

def print_error(text):
    print(f"{Colors.FAIL}✗ {text}{Colors.END}")

def print_info(text):
    print(f"{Colors.CYAN}ℹ {text}{Colors.END}")

def get_api_key():
    """Get API key from environment or user input"""
    global API_KEY

    if os.getenv('GOPHISH_API_KEY'):
        API_KEY = os.getenv('GOPHISH_API_KEY')
        return

    if not API_KEY:
        print_info("Please enter your Gophish API key")
        print_info("(Found in Settings > Account Settings in Gophish UI)")
        API_KEY = input(f"{Colors.CYAN}API Key: {Colors.END}").strip()

def create_email_template(name, subject, html_path):
    """Create an email template in Gophish"""

    try:
        with open(html_path, 'r') as f:
            html_content = f.read()

        template_data = {
            "name": name,
            "subject": subject,
            "text": "Please view this email in HTML mode",
            "html": html_content,
            "attachments": []
        }

        headers = {
            "Authorization": f"Bearer {API_KEY}",
            "Content-Type": "application/json"
        }

        response = requests.post(
            f"{GOPHISH_API_URL}/templates/",
            json=template_data,
            headers=headers,
            verify=False
        )

        if response.status_code == 201:
            print_success(f"Created email template: {name}")
            return True
        else:
            print_error(f"Failed to create {name}: {response.text}")
            return False

    except Exception as e:
        print_error(f"Error creating {name}: {str(e)}")
        return False

def create_landing_page(name, html_path, capture_credentials=True, capture_passwords=True):
    """Create a landing page in Gophish"""

    try:
        with open(html_path, 'r') as f:
            html_content = f.read()

        page_data = {
            "name": name,
            "html": html_content,
            "capture_credentials": capture_credentials,
            "capture_passwords": capture_passwords,
            "redirect_url": ""
        }

        headers = {
            "Authorization": f"Bearer {API_KEY}",
            "Content-Type": "application/json"
        }

        response = requests.post(
            f"{GOPHISH_API_URL}/pages/",
            json=page_data,
            headers=headers,
            verify=False
        )

        if response.status_code == 201:
            print_success(f"Created landing page: {name}")
            return True
        else:
            print_error(f"Failed to create {name}: {response.text}")
            return False

    except Exception as e:
        print_error(f"Error creating {name}: {str(e)}")
        return False

def main():
    """Main installation function"""

    print_header("Gophish Template Installer")

    # Disable SSL warnings
    import urllib3
    urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

    # Get API key
    get_api_key()

    # Get script directory
    script_dir = Path(__file__).parent.parent
    template_dir = script_dir / "templates"

    # Email Templates
    print_header("Installing Email Templates")

    email_templates = [
        {
            "name": "Microsoft 365 - Security Alert",
            "subject": "Unusual sign-in activity detected",
            "path": template_dir / "email" / "microsoft365-security-alert.html"
        },
        {
            "name": "Gmail - Security Alert",
            "subject": "Security alert: New sign-in",
            "path": template_dir / "email" / "gmail-security-alert.html"
        },
        {
            "name": "Stripe - Payment Failed",
            "subject": "Payment Failed - Action Required",
            "path": template_dir / "email" / "stripe-payment-failed.html"
        }
    ]

    for template in email_templates:
        if template["path"].exists():
            create_email_template(
                template["name"],
                template["subject"],
                template["path"]
            )
        else:
            print_error(f"Template file not found: {template['path']}")

    # Landing Pages
    print_header("Installing Landing Pages")

    landing_pages = [
        {
            "name": "Microsoft 365 - Login Page",
            "path": template_dir / "landing-pages" / "microsoft365-login.html"
        },
        {
            "name": "Gmail - Login Page",
            "path": template_dir / "landing-pages" / "gmail-login.html"
        },
        {
            "name": "Stripe - Payment Update",
            "path": template_dir / "landing-pages" / "stripe-payment-update.html"
        }
    ]

    for page in landing_pages:
        if page["path"].exists():
            create_landing_page(
                page["name"],
                page["path"],
                capture_credentials=True,
                capture_passwords=True
            )
        else:
            print_error(f"Landing page file not found: {page['path']}")

    print_header("Installation Complete!")
    print_info("Templates have been installed in your Gophish instance")
    print_info("You can now create campaigns using these templates")
    print_info("\nNext steps:")
    print_info("  1. Create a sending profile in Gophish")
    print_info("  2. Create a user group with target emails")
    print_info("  3. Create a campaign using the installed templates")

if __name__ == "__main__":
    main()
