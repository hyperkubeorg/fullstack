import '@/App.css'
import { Routes, Route, BrowserRouter } from "react-router-dom";
import NotFoundPage from '@/components/pages/404';
import TodoPage from '@/components/pages/todo';
import StatusPage from '@/components/pages/status';
import LoginPage from '@/components/pages/login';
import SignupPage from '@/components/pages/signup';
import PrivacyPage from '@/components/pages/privacy';
import TermsPage from '@/components/pages/terms';

// // Get the URL search params. Example: ?id=123
// const searchParams = new URLSearchParams(window.location.search)

// // Get the URL hash. Example: #hash
// const hash = window.location.hash

// // Get the URL host. Example: localhost:8080
// const host = window.location.host

// // Get the URL protocol. Example: http: or https:
// const protocol = window.location.protocol

// // Get the URL port. Example: 8080
// const port = window.location.port

// // Get the URL origin. Example: http://localhost:8080
// const origin = window.location.origin

// // Get the URL href. Example: http://localhost:8080/test
// const href = window.location.href

// <Router location={window.location} navigator={window.history}>

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<TodoPage />} />

        {/* Auth module routes */}
        <Route path="/auth/status" element={<StatusPage />} />
        <Route path="/auth/login" element={<LoginPage />} />
        <Route path="/auth/signup" element={<SignupPage />} />

        {/* Static content paths */}
        <Route path="/privacy" element={<PrivacyPage />} />
        <Route path="/terms" element={<TermsPage />} />

        <Route path="/@/:username" element={<TodoPage />} />
        <Route path="/_/:boardname" element={<TodoPage />} />
        <Route path="/settings" element={<TodoPage />} />
        <Route path="/messages" element={<TodoPage />} />
        <Route path="/messages/:groupid" element={<TodoPage />} />
        <Route path="/contact" element={<TodoPage />} />
        <Route path="/about" element={<TodoPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}
