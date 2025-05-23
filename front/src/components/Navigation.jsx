import React from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';

export default function Navigation({ isAuthenticated, setIsAuthenticated }) {
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
    navigate('/login');
  };

  return (
    <nav>
      {isAuthenticated ? (
        <>
          <Link to="/main" className={location.pathname === '/main' ? 'active' : ''}>
            Main
          </Link>

          <Link to="/profile" className={location.pathname === '/profile' ? 'active' : ''}>
            Profile
          </Link>

          <Link to="/categories" className={location.pathname === '/categories' ? 'active' : ''}>
            Categories
          </Link>

          <button onClick={handleLogout} className="logout-button">
            Logout
          </button>
        </>
      ) : (
        <>
          <Link to="/login" className={location.pathname === '/login' ? 'active' : ''}>
            Login
          </Link>

          <Link to="/register" className={location.pathname === '/register' ? 'active' : ''}>
            Register
          </Link>
        </>
      )}
    </nav>
  );
}
