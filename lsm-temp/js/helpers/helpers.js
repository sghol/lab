const cache = {};

export async function loadHTML(path) {
    
    if (cache[path]) {
        console.log('Loading from cache:', path);
        return cache[path];
    }
    
    try {
        console.log('Fetching html:', path);
        const response = await fetch(`/html/${path}`);
        
        if (!response.ok) {
            throw new Error(`HTML not found: ${path}`);
        }
        
        const html = await response.text();
        cache[path] = html;
        console.log('html loaded successfully:', path);
        return html;
        
    } catch (error) {
        console.error('html loading error:', error);
        return '<div class="error">Failed to load page</div>';
    }
}

export async function loadHTMLPage(role, page) {
    return loadHTML(`pages/${role}/${page}.html`);
}

export async function loadHTMLComponent(component) {
    return loadHTML(`components/${component}.html`);
}

export function learHTMLCache() {
    Object.keys(cache).forEach(key => delete cache[key]);
}
