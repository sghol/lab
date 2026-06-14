// Main application entry point
import router from "./router/router.js";

// pages 
import HomePage from './pages/public/home.js';
import LoginPage from './pages/public/login.js';
import Error404Page from './pages/public/404.js';

// dashboards
import AdminDashboard from './pages/admin/dashboard.js';
import TeacherDashboard from './pages/teacher/dashboard.js';
import StudentDashboard from './pages/student/dashboard.js';
import SuperuserDashboard from './pages/superuser/dashboard.js';

// register
router.register("/", HomePage);
router.register("/login", LoginPage);
router.register("/admin", AdminDashboard);
router.register("/teacher", TeacherDashboard);
router.register("/student", StudentDashboard);
router.register("/superuser", SuperuserDashboard);
router.register('/404', Error404Page);

console.log("Starting router...");
router.start();
