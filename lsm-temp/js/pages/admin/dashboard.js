// /js/pages/admin/dashboard.js
export default function AdminDashboard() {
  return {
    async render() {
      return `
        <div class="dashboard">
          <h1>📊 Admin Dashboard</h1>
          <p>Welcome to the admin panel</p>
        </div>
      `;
    },
  };
}
