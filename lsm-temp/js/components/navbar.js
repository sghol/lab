import { loadHTMLComponent } from '../helpers/helpers.js';

export async function renderNavbar() {
    return await loadHTMLComponent('navbar');
}

export function initNavbar() {
    console.log('Navbar initialized');
    
    // Example: highlight active link
    const currentPath = window.location.pathname;
    const links = document.querySelectorAll('.navbar a[data-link]');
    
    links.forEach(link => {
        if (link.getAttribute('href') === currentPath) {
            link.classList.add('active');
        }
    });
}
