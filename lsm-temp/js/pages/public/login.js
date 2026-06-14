import { loadHTMLPage } from "../../helpers/helpers.js";
import Auth from "../../api/auth.js";
import router from "../../router/router.js";

console.log("login.js loaded");

export default function LoginPage() {
  return {
    async render() {
      console.log("Rendering login page");
      return await loadHTMLPage("public", "login");
    },

    mount() {
      console.log("Login page mounted");
      const form = document.getElementById("login-form");
      const errorDiv = document.getElementById("error-message");

      // Setup password visibility toggle
      this.setupPasswordToggle();

      if (form) {
        form.addEventListener("submit", async (e) => {
          e.preventDefault();

          const nationalId = document.getElementById("national_id")?.value;
          const password = document.getElementById("password")?.value;

          if (!nationalId || !password) {
            this.showError("لطفاً تمام فیلدها را پر کنید", errorDiv);
            return;
          }

          // Disable button and show loading
          const submitBtn = form.querySelector('button[type="submit"]');
          const originalText = submitBtn.textContent;
          submitBtn.disabled = true;
          submitBtn.textContent = "در حال ورود...";

          try {
            const result = await Auth.login(nationalId, password);

            if (result.success) {
              this.showError(
                "ورود موفقیت‌آمیز! در حال انتقال...",
                errorDiv,
                true,
              );

              // Store user info
              sessionStorage.setItem(
                "current_user",
                JSON.stringify(result.user),
              );

              // Redirect to first of the user
              // Check for superuser latter
              const roles = result.user.roles;
              console.log("roles list", roles);
              if (roles.length > 0) {
                const firstRole = roles[0];
                console.log("First role", firstRole);
                router.navigate(`/${firstRole}`);
              }
            } else {
              // TODO: no rule, maybe showing error
              router.navigate("/");
            }
          } catch (error) {
            this.showError(
              "خطایی رخ داده است. لطفاً مجدداً تلاش کنید.",
              errorDiv,
            );
            console.log("Error on login.js", error);
          } finally {
            submitBtn.disabled = false;
            submitBtn.textContent = originalText;
          }
        });
      }
    },

    setupPasswordToggle() {
      const toggleBtn = document.getElementById("toggle-password");
      const passwordInput = document.getElementById("password");

      if (toggleBtn && passwordInput) {
        toggleBtn.addEventListener("click", () => {
          const type =
            passwordInput.getAttribute("type") === "password"
              ? "text"
              : "password";
          passwordInput.setAttribute("type", type);

          // Toggle eye icon
          toggleBtn.classList.toggle("active");

          // Change emoji based on visibility
          toggleBtn.textContent = type === "password" ? "👁️" : "👁️‍🗨️";
        });
      }
    },

    showError(message, errorDiv, isSuccess = false) {
      if (errorDiv) {
        errorDiv.textContent = message;
        errorDiv.classList.remove("hidden");
        errorDiv.classList.remove("success");

        if (isSuccess) {
          errorDiv.classList.add("success");
        }

        // Auto-hide after 5 seconds
        setTimeout(() => {
          errorDiv.classList.add("hidden");
        }, 5000);
      }
    },
  };
}
