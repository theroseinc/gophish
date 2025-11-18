/**
 * Browser-in-the-Browser (BitB) Library
 * Creates fake OAuth/SSO popup windows for educational security testing
 *
 * Usage:
 *   const bitb = new BrowserInBrowser({
 *     title: 'Sign in - Google Accounts',
 *     url: 'https://accounts.google.com/signin/oauth',
 *     width: 500,
 *     height: 600,
 *     onClose: function() { console.log('Window closed'); }
 *   });
 *   bitb.show();
 */

class BrowserInBrowser {
    constructor(options) {
        this.options = {
            title: options.title || 'Sign in',
            url: options.url || 'https://example.com',
            width: options.width || 500,
            height: options.height || 600,
            content: options.content || '',
            onClose: options.onClose || function() {},
            draggable: options.draggable !== false,
            theme: options.theme || 'light', // light or dark
            browser: options.browser || 'chrome' // chrome, firefox, safari, edge
        };

        this.element = null;
        this.isDragging = false;
        this.dragOffset = { x: 0, y: 0 };

        this.init();
    }

    init() {
        // Create main container
        this.element = document.createElement('div');
        this.element.className = 'bitb-window';
        this.element.innerHTML = this.getTemplate();

        // Add to body
        document.body.appendChild(this.element);

        // Setup event listeners
        this.setupEventListeners();

        // Center window
        this.centerWindow();
    }

    getTemplate() {
        const browserChrome = this.getBrowserChrome();

        return `
            <div class="bitb-backdrop"></div>
            <div class="bitb-container" style="width: ${this.options.width}px;">
                ${browserChrome}
                <div class="bitb-content-wrapper" style="height: ${this.options.height}px;">
                    <div class="bitb-content">
                        ${this.options.content}
                    </div>
                </div>
            </div>
        `;
    }

    getBrowserChrome() {
        switch(this.options.browser) {
            case 'chrome':
                return this.getChromeChrome();
            case 'firefox':
                return this.getFirefoxChrome();
            case 'safari':
                return this.getSafariChrome();
            case 'edge':
                return this.getEdgeChrome();
            default:
                return this.getChromeChrome();
        }
    }

    getChromeChrome() {
        return `
            <div class="bitb-chrome bitb-chrome-light">
                <div class="bitb-titlebar">
                    <div class="bitb-tabs">
                        <div class="bitb-tab bitb-tab-active">
                            <div class="bitb-tab-favicon">
                                <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                                    <circle cx="8" cy="8" r="6" opacity="0.5"/>
                                </svg>
                            </div>
                            <div class="bitb-tab-title">${this.options.title}</div>
                            <div class="bitb-tab-close">×</div>
                        </div>
                        <div class="bitb-new-tab">+</div>
                    </div>
                    <div class="bitb-window-controls">
                        <div class="bitb-control bitb-minimize">−</div>
                        <div class="bitb-control bitb-maximize">□</div>
                        <div class="bitb-control bitb-close" data-bitb-close>×</div>
                    </div>
                </div>
                <div class="bitb-toolbar">
                    <div class="bitb-nav-buttons">
                        <button class="bitb-nav-btn" disabled>←</button>
                        <button class="bitb-nav-btn" disabled>→</button>
                        <button class="bitb-nav-btn">⟳</button>
                    </div>
                    <div class="bitb-addressbar">
                        <div class="bitb-security-icon">
                            <svg width="16" height="16" viewBox="0 0 16 16" fill="#0F9D58">
                                <path d="M8 2L3 4v3.5c0 3.1 2.1 6 5 6.5 2.9-.5 5-3.4 5-6.5V4l-5-2zm0 1.4l3.5 1.4v3.7c0 2.3-1.5 4.4-3.5 4.9-2-.5-3.5-2.6-3.5-4.9V4.8L8 3.4z"/>
                                <path d="M7 10.2L5.4 8.6l-.7.7L7 11.6l4.3-4.3-.7-.7z"/>
                            </svg>
                        </div>
                        <div class="bitb-url">${this.options.url}</div>
                        <div class="bitb-bookmark-icon">☆</div>
                    </div>
                    <div class="bitb-toolbar-icons">
                        <div class="bitb-icon">⋮</div>
                    </div>
                </div>
            </div>
        `;
    }

    getFirefoxChrome() {
        return `
            <div class="bitb-chrome bitb-chrome-firefox">
                <div class="bitb-titlebar">
                    <div class="bitb-tabs">
                        <div class="bitb-tab bitb-tab-active">
                            <div class="bitb-tab-favicon">🦊</div>
                            <div class="bitb-tab-title">${this.options.title}</div>
                            <div class="bitb-tab-close">×</div>
                        </div>
                    </div>
                    <div class="bitb-window-controls">
                        <div class="bitb-control bitb-minimize">−</div>
                        <div class="bitb-control bitb-maximize">□</div>
                        <div class="bitb-control bitb-close" data-bitb-close>×</div>
                    </div>
                </div>
                <div class="bitb-toolbar">
                    <div class="bitb-nav-buttons">
                        <button class="bitb-nav-btn">←</button>
                        <button class="bitb-nav-btn">→</button>
                        <button class="bitb-nav-btn">⟳</button>
                    </div>
                    <div class="bitb-addressbar">
                        <div class="bitb-security-icon">🔒</div>
                        <div class="bitb-url">${this.options.url}</div>
                    </div>
                    <div class="bitb-toolbar-icons">
                        <div class="bitb-icon">☰</div>
                    </div>
                </div>
            </div>
        `;
    }

    getSafariChrome() {
        return `
            <div class="bitb-chrome bitb-chrome-safari">
                <div class="bitb-titlebar">
                    <div class="bitb-window-controls bitb-controls-left">
                        <div class="bitb-control bitb-close bitb-macos" data-bitb-close></div>
                        <div class="bitb-control bitb-minimize bitb-macos"></div>
                        <div class="bitb-control bitb-maximize bitb-macos"></div>
                    </div>
                    <div class="bitb-title-center">${this.options.title}</div>
                </div>
                <div class="bitb-toolbar">
                    <div class="bitb-nav-buttons">
                        <button class="bitb-nav-btn">←</button>
                        <button class="bitb-nav-btn">→</button>
                    </div>
                    <div class="bitb-addressbar">
                        <div class="bitb-security-icon">🔒</div>
                        <div class="bitb-url">${this.options.url}</div>
                    </div>
                    <div class="bitb-toolbar-icons">
                        <div class="bitb-icon">⋯</div>
                    </div>
                </div>
            </div>
        `;
    }

    getEdgeChrome() {
        return `
            <div class="bitb-chrome bitb-chrome-edge">
                <div class="bitb-titlebar">
                    <div class="bitb-tabs">
                        <div class="bitb-tab bitb-tab-active">
                            <div class="bitb-tab-favicon">🌐</div>
                            <div class="bitb-tab-title">${this.options.title}</div>
                            <div class="bitb-tab-close">×</div>
                        </div>
                    </div>
                    <div class="bitb-window-controls">
                        <div class="bitb-control bitb-minimize">−</div>
                        <div class="bitb-control bitb-maximize">□</div>
                        <div class="bitb-control bitb-close" data-bitb-close>×</div>
                    </div>
                </div>
                <div class="bitb-toolbar">
                    <div class="bitb-nav-buttons">
                        <button class="bitb-nav-btn">←</button>
                        <button class="bitb-nav-btn">→</button>
                        <button class="bitb-nav-btn">⟳</button>
                    </div>
                    <div class="bitb-addressbar">
                        <div class="bitb-security-icon">🔒</div>
                        <div class="bitb-url">${this.options.url}</div>
                    </div>
                    <div class="bitb-toolbar-icons">
                        <div class="bitb-icon">⋯</div>
                    </div>
                </div>
            </div>
        `;
    }

    setupEventListeners() {
        const container = this.element.querySelector('.bitb-container');
        const titlebar = this.element.querySelector('.bitb-titlebar');
        const closeBtn = this.element.querySelector('[data-bitb-close]');
        const backdrop = this.element.querySelector('.bitb-backdrop');

        // Close button
        if (closeBtn) {
            closeBtn.addEventListener('click', () => this.close());
        }

        // Backdrop click
        if (backdrop) {
            backdrop.addEventListener('click', () => this.close());
        }

        // Draggable
        if (this.options.draggable && titlebar) {
            titlebar.style.cursor = 'move';

            titlebar.addEventListener('mousedown', (e) => {
                if (e.target.closest('[data-bitb-close]')) return;

                this.isDragging = true;
                const rect = container.getBoundingClientRect();
                this.dragOffset = {
                    x: e.clientX - rect.left,
                    y: e.clientY - rect.top
                };

                e.preventDefault();
            });

            document.addEventListener('mousemove', (e) => {
                if (!this.isDragging) return;

                const x = e.clientX - this.dragOffset.x;
                const y = e.clientY - this.dragOffset.y;

                container.style.left = x + 'px';
                container.style.top = y + 'px';
                container.style.transform = 'none';
            });

            document.addEventListener('mouseup', () => {
                this.isDragging = false;
            });
        }
    }

    centerWindow() {
        const container = this.element.querySelector('.bitb-container');
        const windowWidth = window.innerWidth;
        const windowHeight = window.innerHeight;

        const left = (windowWidth - this.options.width) / 2;
        const top = (windowHeight - this.options.height - 80) / 2; // 80 for chrome height

        container.style.left = left + 'px';
        container.style.top = top + 'px';
    }

    show() {
        if (this.element) {
            this.element.style.display = 'block';
            setTimeout(() => {
                this.element.classList.add('bitb-show');
            }, 10);
        }
    }

    hide() {
        if (this.element) {
            this.element.classList.remove('bitb-show');
            setTimeout(() => {
                this.element.style.display = 'none';
            }, 300);
        }
    }

    close() {
        this.hide();
        if (this.options.onClose) {
            this.options.onClose();
        }
        setTimeout(() => {
            if (this.element && this.element.parentNode) {
                this.element.parentNode.removeChild(this.element);
            }
        }, 300);
    }

    setContent(html) {
        const contentDiv = this.element.querySelector('.bitb-content');
        if (contentDiv) {
            contentDiv.innerHTML = html;
        }
    }
}

// Auto-init if data attributes are present
document.addEventListener('DOMContentLoaded', function() {
    const triggers = document.querySelectorAll('[data-bitb-trigger]');

    triggers.forEach(trigger => {
        trigger.addEventListener('click', function(e) {
            e.preventDefault();

            const contentId = this.getAttribute('data-bitb-content');
            const contentElement = document.getElementById(contentId);
            const content = contentElement ? contentElement.innerHTML : '';

            const bitb = new BrowserInBrowser({
                title: this.getAttribute('data-bitb-title') || 'Sign in',
                url: this.getAttribute('data-bitb-url') || 'https://accounts.google.com',
                width: parseInt(this.getAttribute('data-bitb-width')) || 500,
                height: parseInt(this.getAttribute('data-bitb-height')) || 600,
                browser: this.getAttribute('data-bitb-browser') || 'chrome',
                content: content
            });

            bitb.show();
        });
    });
});
