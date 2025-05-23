import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';  

export default function CategoryPage() {
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [message, setMessage] = useState('');
  const [newCategoryName, setNewCategoryName] = useState('');
  const [editingId, setEditingId] = useState(null);
  const [editingName, setEditingName] = useState('');

  const token = localStorage.getItem('token');
  const navigate = useNavigate();

  const fetchCategories = async () => {
    if (!token) {
      setMessage('Авторизация қажет. Логин жасаңыз.');
      setLoading(false);
      return;
    }
    try {
      const res = await axios.get('http://localhost:8081/user/categories', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setCategories(res.data);
      setMessage('');
    } catch (err) {
      setMessage(err.response?.data?.message || 'Категорияларды алу қатесі');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCategories();
  }, []);

  const handleAddCategory = async () => {
    if (!newCategoryName.trim()) {
      alert('Категория атауы бос болмауы керек');
      return;
    }
    try {
      const res = await axios.post(
        'http://localhost:8081/user/categories',
        { name: newCategoryName },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      setCategories([...categories, res.data]);
      setNewCategoryName('');
      setMessage('Категория сәтті қосылды');

     
      setTimeout(() => {
        navigate('/categories');
      }, 2000);
    } catch (err) {
      setMessage(err.response?.data?.message || 'Категория қосуда қате');
    }
  };

  const startEditing = (id, name) => {
    setEditingId(id);
    setEditingName(name);
    setMessage('');
  };

  const cancelEditing = () => {
    setEditingId(null);
    setEditingName('');
    setMessage('');
  };

  const saveEditing = async () => {
    if (!editingName.trim()) {
      alert('Атау бос болмауы керек');
      return;
    }
    try {
      const res = await axios.put(
        `http://localhost:8081/user/categories/${editingId}`,
        { name: editingName },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      setCategories(categories.map(cat => (cat.id === editingId ? res.data : cat)));
      setEditingId(null);
      setEditingName('');
      setMessage('Категория сәтті өңделді');
    } catch (err) {
      setMessage(err.response?.data?.message || 'Категория өңдеуде қате');
    }
  };

  const handleDeleteCategory = async (id) => {
    if (!window.confirm('Бұл категорияны өшіргіңіз келе ме?')) return;
    try {
      await axios.delete(`http://localhost:8081/user/categories/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setCategories(categories.filter(cat => cat.id !== id));
      setMessage('Категория сәтті өшірілді');
    } catch (err) {
      setMessage(err.response?.data?.message || 'Категория өшіруде қате');
    }
  };

  if (loading) return <p>Жүктелуде...</p>;

  return (
    <div style={{ maxWidth: 600, margin: '30px auto', padding: 20, backgroundColor: '#f0f0f0', borderRadius: 8 }}>
      <h2>Категориялар тізімі</h2>
      {message && <p style={{ color: 'red' }}>{message}</p>}

      <div style={{ marginBottom: 20 }}>
        <input
          type="text"
          placeholder="Жаңа категорияның атауы"
          value={newCategoryName}
          onChange={(e) => setNewCategoryName(e.target.value)}
          style={{ padding: '6px 10px', width: '70%', marginRight: 10 }}
        />
        <button onClick={handleAddCategory} style={{ padding: '6px 14px' }}>Қосу</button>
      </div>

      <ul style={{ listStyle: 'none', padding: 0 }}>
        {categories.map((cat) => (
          <li
            key={cat.id}
            style={{
              padding: 12,
              marginBottom: 10,
              backgroundColor: 'white',
              borderRadius: 6,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'space-between',
            }}
          >
            {editingId === cat.id ? (
              <>
                <input
                  type="text"
                  value={editingName}
                  onChange={(e) => setEditingName(e.target.value)}
                  style={{ flexGrow: 1, marginRight: 10, padding: '6px 10px' }}
                />
                <button onClick={saveEditing} style={{ marginRight: 8, backgroundColor: '#007bff', color: 'white' }}>
                  Сақтау
                </button>
                <button onClick={cancelEditing} style={{ backgroundColor: '#6c757d', color: 'white' }}>
                  Болдырмау
                </button>
              </>
            ) : (
              <>
                <span>{cat.name}</span>
                <div>
                  <button
                    onClick={() => startEditing(cat.id, cat.name)}
                    style={{ marginRight: 10, backgroundColor: '#28a745', color: 'white' }}
                  >
                    Өңдеу
                  </button>
                  <button
                    onClick={() => handleDeleteCategory(cat.id)}
                    style={{ backgroundColor: '#dc3545', color: 'white' }}
                  >
                    Өшіру
                  </button>
                </div>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}
