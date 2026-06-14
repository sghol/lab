// Home page controller
import { loadHTMLPage } from "../../helpers/helpers.js";
import auth from "../../api/auth.js";

console.log("home.js loaded");

export default function HomePage() {
  return {
    // render
    async render() {
      const homeHTML = await loadHTMLPage("public", "home");
      return `${homeHTML}`;
    },

    // interactivity
    mount() {
      console.log("Home page mounted");
      // Check if user is logged in
      const isLoggedIn = auth.isLoggedIn();
      console.log(isLoggedIn);
      const userData = sessionStorage.getItem('current_user');
      let user = null;

      if (userData) {
        user = JSON.parse(userData);
      }

      // Update navigation if logged in
      if (isLoggedIn && user) {
        this.updateNavigation(user);
      }

      // Setup logout button if exists
      this.setupLogout();
    },

    cleanup() { },

    // Update navigation for logged in users
    updateNavigation(user) {
      const nav = document.querySelector('.hero-nav ul');
      if (!nav) return;

      nav.innerHTML = `
        <li><a href="/dashboard" data-link>داشبورد</a></li>
        <li><a href="/profile" data-link>پروفایل</a></li>
        <li><a href="#" id="logout-btn">خروج</a></li>
      `;

      // Add welcome message
      const heroContent = document.querySelector('.hero-content');
      if (heroContent) {
        const existingWelcome = heroContent.querySelector('.welcome-msg');
        if (existingWelcome) existingWelcome.remove();

        const welcomeMsg = document.createElement('div');
        welcomeMsg.className = 'welcome-msg';
        welcomeMsg.innerHTML = `
          <div class="user-welcome">
            <p>👋 خوش آمدید <strong>${user.first_name} ${user.last_name}</strong></p>
          </div>
        `;
        heroContent.appendChild(welcomeMsg);
      }
    },

    // Setup logout functionality
    setupLogout() {
      const logoutBtn = document.getElementById('logout-btn');
      if (!logoutBtn) return;
      
      logoutBtn.addEventListener('click', (e) => {
        e.preventDefault();
        console.log("Logging out...");
        
        // Clear authentication
        auth.clearToken();
        sessionStorage.removeItem('current_user');
        
        // Refresh page to show logged out state
        window.location.reload();
      });
    }

  };
}
