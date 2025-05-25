import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function LoginPage() {
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError("");

    const formData = new FormData(event.currentTarget);
    const usernameOrEmail = formData.get("username-email");
    const password = formData.get("password");

    try {
      const response = await fetch("/api/v1/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username_or_email: usernameOrEmail, password }),
      });

      if (response.ok) {
        navigate("/");
      } else {
        const data = await response.json();
        setError(data.message || "Login failed.");
      }
    } catch (err) {
      setError("An error occurred. Please try again later.");
    }
  };

  return (
    <>
      <h1 className="text-white">Login Page</h1>
      <div className="card">
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="username-email" className="block text-sm font-medium text-white mb-1">Username or Email</label>
            <input type="text" id="username-email" name="username-email" className="block w-full border border-white rounded-md shadow-sm" required />
          </div>
          <div className="mb-4">
            <label htmlFor="password" className="block text-sm font-medium text-white mb-1">Password</label>
            <input type="password" id="password" name="password" className="block w-full border border-white rounded-md shadow-sm" required />
          </div>
          {error && <p className="text-red-500 text-sm mb-4">{error}</p>}
          <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded">Login</button>
        </form>
        <p className="mt-4 text-sm text-white">
          Don't have an account? <a href="/auth/signup" className="text-blue-500 underline">Sign up</a>
        </p>
      </div>
    </>
  );
}