const API_BASE = 'http://localhost:8080';


// Simple auth helper to manage tokens
const Auth = {
    // Store token in localStorage
    setToken(token) {
        localStorage.setItem('lms_token', token);
    },

    // Get token from localStorage
    getToken() {
        return localStorage.getItem('lms_token');
    },

    // Remove token (logout)
    clearToken() {
        localStorage.removeItem('lms_token');
    },

    // Check if user is logged in
    isLoggedIn() {
        return !!this.getToken();
    },

    // Get token for API headers
    getAuthHeader() {
        const token = this.getToken();
        return token ? { 'Authorization': `Bearer ${token}` } : {};
    },

    // Login function
    async login(nationalId, password) {
        try {
            const response = await fetch(`${API_BASE}/auth/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    national_id: nationalId,
                    password: password
                })
            });

            const data = await response.json();

            if (response.ok) {
                // Store the token
                this.setToken(data.token);
                return { success: true, user: data.user };
            } else {
                return {
                    success: false,
                    error: data.message || 'Login failed'
                };
            }
        } catch (error) {
            console.error('Login error:', error);
            return {
                success: false,
                error: 'Network error. Please try again.'
            };
        }
    },

    // Get current user info (protected endpoint)
    async getCurrentUser() {
        try {
            const response = await fetch(`${API_BASE}/auth/me`, {
                headers: {
                    'Content-Type': 'application/json',
                    ...this.getAuthHeader()
                }
            });

            if (response.ok) {
                return await response.json();
            }
            return null;
        } catch (error) {
            console.error('Get user error:', error);
            return null;
        }
    }
};

export default Auth;