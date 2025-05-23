import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const inputStyle = { width: '100%', padding: '8px', marginBottom: '10px' };
const buttonStyle = { width: '100%', padding: '10px' };

export default function Login({ setIsAuthenticated }) {
  const [form, setForm] = useState({ username: '', password: '' });
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage('');
    setLoading(true);

    try {
      const response = await axios.post('http://localhost:8081/login', form);
      localStorage.setItem('token', response.data.token);
      setMessage('Login successful!');
      setIsAuthenticated(true); 
      setTimeout(() => {
        navigate('/main');
      }, 1000);
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Login failed. Try again.';
      setMessage(errorMsg);
      setIsAuthenticated(false);
    }
    setLoading(false);
  };

  return (
    <div style={{ maxWidth: 400, margin: 'auto' }}>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <input
          name="username"
          placeholder="Username"
          value={form.username}
          onChange={handleChange}
          required
          style={inputStyle}
        />
        <input
          name="password"
          type="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          required
          style={inputStyle}
        />
        <button type="submit" disabled={loading} style={buttonStyle}>
          {loading ? 'Logging in...' : 'Login'}
        </button>
      </form>
      {message && (
        <p style={{ marginTop: 10, color: message.includes('successful') ? 'green' : 'red' }}>
          {message}
        </p>
      )}
    </div>
  );
}
