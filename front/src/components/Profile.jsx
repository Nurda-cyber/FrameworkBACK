import React, { useEffect, useState } from 'react';
import axios from 'axios';

export default function Profile() {
  const [user, setUser] = useState(null);
  const [message, setMessage] = useState('');
  const [isUpdating, setIsUpdating] = useState(false);
  const [usernameInput, setUsernameInput] = useState('');

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = () => {
    const token = localStorage.getItem('token');
    if (!token) {
      setMessage('You are not authorized. Please log in.');
      return;
    }

    axios
      .get('http://localhost:8081/user/me', {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        setUser(res.data);
        setUsernameInput(res.data.username);
        setMessage('');
      })
      .catch((err) => setMessage(err.response?.data?.error || 'Failed to fetch profile'));
  };

  const handleDelete = () => {
    const token = localStorage.getItem('token');
    if (!window.confirm('Are you sure you want to delete your profile?')) return;

    axios
      .delete('http://localhost:8081/user/me', {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then(() => {
        alert('Profile deleted successfully.');
        setUser(null);
        setMessage('Profile deleted. Please log in again.');
        localStorage.removeItem('token');
      })
      .catch((err) => setMessage(err.response?.data?.error || 'Failed to delete profile'));
  };

  const handleUpdate = () => {
    const token = localStorage.getItem('token');
    if (!usernameInput.trim()) {
      alert('Username cannot be empty.');
      return;
    }

    axios
      .put(
        'http://localhost:8081/user/me',
        { username: usernameInput },
        { headers: { Authorization: `Bearer ${token}` } }
      )
      .then((res) => {
        alert('Profile updated successfully!');
        setUser(res.data);
        setIsUpdating(false);
        setMessage('');
      })
      .catch((err) => setMessage(err.response?.data?.error || 'Failed to update profile'));
  };

  if (message) return <p className="message">{message}</p>;
  if (!user) return <p className="message">Loading profile...</p>;

  return (
    <div className="container">
      <h2>👤 Profile</h2>
      <div className="profile-info">
        <p><strong>ID:</strong> {user.id}</p>

        {!isUpdating ? (
          <p><strong>Name:</strong> {user.username}</p>
        ) : (
          <input
            type="text"
            value={usernameInput}
            onChange={(e) => setUsernameInput(e.target.value)}
          />
        )}

        <div style={{ marginTop: 20 }}>
          {!isUpdating ? (
            <>
              <button onClick={() => setIsUpdating(true)} style={{ marginRight: 10 }}>
                Жаңарту
              </button>
              <button onClick={handleDelete} style={{ backgroundColor: '#dc3545', color: 'white' }}>
                Өшіру
              </button>
            </>
          ) : (
            <>
              <button onClick={handleUpdate} style={{ marginRight: 10 }}>
                Сақтау
              </button>
              <button onClick={() => setIsUpdating(false)}>
                Болдырмау
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
