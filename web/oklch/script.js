const switchModeCheck = document.getElementById("switch-mode");
switchModeCheck.addEventListener("change", () => {
  const isDark = switchModeCheck.checked;
  document.documentElement.setAttribute("data-theme", isDark ? "dark": "light");
});
