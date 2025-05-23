import React, { useEffect, useState } from 'react';
import axios from 'axios';

export default function Main() {
  const [toys, setToys] = useState([]);
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState(''); 

  const fetchToys = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      setMessage('You are not authorized. Please log in.');
      setLoading(false);
      return;
    }

    try {
      const response = await axios.get('http://localhost:8081/user/toys', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setToys(response.data);
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Failed to fetch toys';
      setMessage(errorMsg);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchToys();
  }, []);

  const handleDelete = async (id) => {
    const token = localStorage.getItem('token');
    try {
      await axios.delete(`http://localhost:8081/user/toys/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setToys(toys.filter((toy) => toy.id !== id));
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Failed to delete toy';
      setMessage(errorMsg);
    }
  };

  const handleEdit = async (id) => {
    const toy = toys.find((t) => t.id === id);
    if (!toy) return;

    const newName = prompt('Enter new name:', toy.name);
    if (newName === null || newName.trim() === '') return;

    const newDescription = prompt('Enter new description:', toy.description);
    if (newDescription === null || newDescription.trim() === '') return;

    const token = localStorage.getItem('token');

    try {
      const updatedToy = {
        ...toy,
        name: newName,
        description: newDescription,
      };

      await axios.put(`http://localhost:8081/user/toys/${id}`, updatedToy, {
        headers: { Authorization: `Bearer ${token}` },
      });

      setToys(
        toys.map((t) => (t.id === id ? { ...t, name: newName, description: newDescription } : t))
      );
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Failed to update toy';
      setMessage(errorMsg);
    }
  };

  const handleAddToy = async () => {
    const name = prompt('Enter toy name:');
    if (name === null || name.trim() === '') return;

    const description = prompt('Enter toy description:');
    if (description === null || description.trim() === '') return;

    const priceStr = prompt('Enter toy price:');
    const price = parseFloat(priceStr);
    if (isNaN(price)) {
      alert('Invalid price');
      return;
    }

    const categoryIdStr = prompt('Enter category ID (number):');
    const category_id = parseInt(categoryIdStr);
    if (isNaN(category_id)) {
      alert('Invalid category ID');
      return;
    }

    const token = localStorage.getItem('token');
    try {
      const response = await axios.post(
        'http://localhost:8081/user/toys',
        { name, description, price, category_id },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      setToys([...toys, response.data]);
    } catch (error) {
      const errorMsg = error.response?.data?.error || 'Failed to add toy';
      setMessage(errorMsg);
    }
  };

  if (loading) return <p>Loading toys...</p>;

  const filteredToys = toys.filter((toy) =>
    toy.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div style={{ maxWidth: 700, margin: 'auto' }}>
      <h2>Toy List</h2>
      {message && <p style={{ color: 'red' }}>{message}</p>}

      <input
        type="text"
        placeholder="Search toys by name..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        style={{ padding: '8px', marginBottom: '20px', width: '100%' }}
      />

      <button onClick={handleAddToy} style={{ marginBottom: 20 }}>
        Add Toy
      </button>
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {filteredToys.length > 0 ? (
          filteredToys.map((toy) => (
            <li
              key={toy.id}
              style={{
                marginBottom: 20,
                borderBottom: '1px solid #ccc',
                paddingBottom: 10,
              }}
            >
              <strong>{toy.name}</strong> — {toy.description}
              <br />
              <span>Бағасы: {toy.price} ₸</span>
              <br />
              <button
                onClick={() => handleEdit(toy.id)}
                style={{ marginRight: 10 }}
              >
                Өңдеу
              </button>
              <button onClick={() => handleDelete(toy.id)}>Өшіру</button>
            </li>
          ))
        ) : (
          <li>No toys found.</li>
        )}
      </ul>
    </div>
  );
}
