import React, { useState } from 'react';
import axios from 'axios';

const inputStyle = { width: '100%', padding: '8px', marginBottom: '10px' };
const buttonStyle = { width: '100%', padding: '10px' };

export default function Register() {
  const [form, setForm] = useState({ username: '', password: '' });
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage('');
    setLoading(true);

    try {
      const response = await axios.post('http://localhost:8081/register', form);
      setMessage(response.data.message || 'Registration successful!');
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Registration failed. Try again.';
      setMessage(errorMsg);
    }
    setLoading(false);
  };

  return (
    <div style={{ maxWidth: 400, margin: 'auto' }}>
      <h2>Register</h2>
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
          {loading ? 'Registering...' : 'Register'}
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
