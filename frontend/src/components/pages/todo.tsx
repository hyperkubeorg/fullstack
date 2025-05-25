import { useEffect, useState } from 'react'
import reactLogo from '@/assets/react.svg'
import viteLogo from '/vite.svg'

export default function TodoPage() {
  const [ count, setCount ] = useState(0);
  const [userStatus, setUserStatus] = useState("Not logged in.");

  useEffect(() => {
    async function fetchUserStatus() {
      try {
        const response = await fetch("/api/v1/auth/status");
        if (response.ok) {
          const data = await response.json();
          setUserStatus(`Logged in as, ${data.data.username}.`);
        } else {
          setUserStatus("Not logged in.");
        }
      } catch (error) {
        setUserStatus("Not logged in.");
      }
    }

    fetchUserStatus();
  }, []);

  return (
    <>
      <div className="flex items-center justify-center space-x-4 my-4">
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo w-24 h-24" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react w-24 h-24" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      <p className="read-the-docs">
        <span>Auth: </span>
        <a href="/auth/login">Login</a> |
        <a href="/auth/signup">Sign Up</a> |
        <a href="/auth/status">Status</a>
      </p>
      <div className="mt-4 p-4 border rounded-md">
        <p>{userStatus}</p>
      </div>
    </>
  );
}