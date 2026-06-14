class Router {
  constructor() {
    this.routes = {};
    this.currentPage = null;

    // Listen for back/forward button
    window.addEventListener("popstate", () => this.handleRoute());

    // Prevent default behaviour of every link with data-link attribute
    document.addEventListener("click", (e) => {
      if (e.target.matches("[data-link]")) {
        e.preventDefault();
        this.navigate(e.target.getAttribute("href"));
      }
    });
  }

  // Register a route
  register(path, handler) {
    this.routes[path] = handler;
  }

  navigate(path) {
    history.pushState(null, null, path);
    this.handleRoute();
  }

  async handleRoute() {
    const path = window.location.pathname;
    const handler = this.routes[path] || this.routes["/404"];
    
    console.log('Current path:', path);
    console.log('Handler found:', !!handler);

    if (handler) {
      const app = document.getElementById("app");
      
      // Clear previous content
      if (this.currentPage && this.currentPage.cleanup) {
        this.currentPage.cleanup();
      }

      // Get the page object
      this.currentPage = await handler();
      
      // Render the page (AWAIT the async render method)
      const html = await this.currentPage.render();
      app.innerHTML = html;

      // Run page initialization
      if (this.currentPage.mount) {
        this.currentPage.mount();
      }
    }
  }

  // Start the router
  start() {
    this.handleRoute();
  }
}

const router = new Router();
export default router;
