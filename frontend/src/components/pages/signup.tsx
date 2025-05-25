import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";
import { useNavigate } from "react-router-dom";

export default function SignupPage() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
    terms: false,
    privacy: false,
  });

  const navigate = useNavigate();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === "checkbox" ? checked : value,
    });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      alert("Passwords do not match.");
      return;
    }

    try {
      const response = await fetch("/api/v1/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: formData.username,
          email: formData.email,
          password: formData.password,
          terms: formData.terms,
          privacy: formData.privacy,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        alert(errorData.message || "Signup failed.");
        return;
      }

      alert("Signup successful!");
      navigate("/");
    } catch (error) {
      console.error("Error during signup:", error);
      alert("An error occurred. Please try again later.");
    }
  };

  return (
    <>
      <h1 className="text-white">Sign Up Page</h1>
      <div className="card">
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="username" className="block text-sm font-medium text-white mb-1">Username</label>
            <input
              type="text"
              id="username"
              name="username"
              className="block w-full border border-white rounded-md shadow-sm"
              value={formData.username}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="email" className="block text-sm font-medium text-white mb-1">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              className="block w-full border border-white rounded-md shadow-sm"
              value={formData.email}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="password" className="block text-sm font-medium text-white mb-1">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              className="block w-full border border-white rounded-md shadow-sm"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="confirmPassword" className="block text-sm font-medium text-white mb-1">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              className="block w-full border border-white rounded-md shadow-sm"
              value={formData.confirmPassword}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-4">
            <label className="flex items-center text-white">
              <input
                type="checkbox"
                id="terms"
                name="terms"
                className="mr-2"
                checked={formData.terms}
                onChange={handleChange}
                required
              />
              I agree to the <a href="/auth/signup" className="text-blue-500 underline ml-2">Terms of Service</a>
            </label>
          </div>
          <div className="mb-4">
            <label className="flex items-center text-white">
              <input
                type="checkbox"
                id="privacy"
                name="privacy"
                className="mr-2"
                checked={formData.privacy}
                onChange={handleChange}
                required
              />
              I agree to the <a href="/privacy" className="text-blue-500 underline ml-2">Privacy Policy</a>
            </label>
          </div>
          <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded">Sign Up</button>
        </form>
        <p className="mt-4 text-sm text-white">
          Already have an account? <a href="/auth/login" className="text-blue-500 underline">Log in</a>
        </p>
      </div>
    </>
  );
}