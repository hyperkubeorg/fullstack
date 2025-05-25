import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export default function StatusPage() {
  const [userInfo, setUserInfo] = useState({ username: '', isAdmin: false, isBanned: false, isLoggedIn: false });

  useEffect(() => {
    async function fetchUserStatus() {
      try {
        const response = await fetch('/api/v1/auth/status');
        if (response.ok) {
          const data = await response.json();
          setUserInfo({
            username: data.data.username,
            isAdmin: data.data.is_admin,
            isBanned: data.data.is_banned,
            isLoggedIn: true,
          });
        } else {
          setUserInfo({ username: '', isAdmin: false, isBanned: false, isLoggedIn: false });
        }
      } catch (error) {
        console.error('Error fetching user status:', error);
        setUserInfo({ username: '', isAdmin: false, isBanned: false, isLoggedIn: false });
      }
    }

    fetchUserStatus();
  }, []);

  return (
    <>
      <h1>Login Info</h1>
      <div className="card">
        {userInfo.isLoggedIn ? (
          <>
            <p>
              <strong>Username:</strong> {userInfo.username}
            </p>
            <p>
              <strong>Admin:</strong> {userInfo.isAdmin ? 'Yes' : 'No'}
            </p>
            <p>
              <strong>Banned:</strong> {userInfo.isBanned ? 'Yes' : 'No'}
            </p>
            <p>
              <strong>Status:</strong> Logged in
            </p>
          </>
        ) : (
          <p>Login required</p>
        )}
      </div>
      <div style={{ marginTop: '20px' }}>
        <Link to="/">Back to Home</Link>
      </div>
    </>
  );
}