import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';

import Login from './components/Login';
import Register from './components/Register';
import Navigation from './components/Navigation';
import Main from './components/Main';
import Profile from './components/Profile';
import CategoryPage from './components/CategoryPage';
import './App.css'; 
export default function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) setIsAuthenticated(true);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
  };

  return (
    <Router>
      <div style={{ padding: 20 }}>
        <Navigation isAuthenticated={isAuthenticated} setIsAuthenticated={setIsAuthenticated} />
        <hr />
        <Routes>
          <Route
            path="/login"
            element={
              isAuthenticated ? <Navigate to="/main" replace /> : <Login setIsAuthenticated={setIsAuthenticated} />
            }
          />
          <Route
            path="/register"
            element={
              isAuthenticated ? <Navigate to="/main" replace /> : <Register />
            }
          />
          <Route
            path="/main"
            element={
              isAuthenticated ? <Main /> : <Navigate to="/login" replace />
            }
          />
          <Route
            path="/profile"
            element={
              isAuthenticated ? <Profile /> : <Navigate to="/login" replace />
            }
          />
        
          <Route
            path="/categories"
            element={
              isAuthenticated ? <CategoryPage /> : <Navigate to="/login" replace />
            }
          />
          <Route
            path="/"
            element={
              isAuthenticated
                ? <Navigate to="/main" replace />
                : <h2>Welcome! Please choose Login or Register.</h2>
            }
          />
          <Route path="*" element={<Navigate to={isAuthenticated ? "/main" : "/login"} replace />} />
        </Routes>
      </div>
    </Router>
  );
}