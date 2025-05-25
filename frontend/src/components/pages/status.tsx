export default function StatusPage() {
  const isLoggedIn = true; // Placeholder condition for user login status

  return (
    <>
      <h1>Login Info</h1>
      <div className="card">
        {isLoggedIn ? (
          <>
            <p>
              <strong>Username:</strong> johndoe
            </p>
            <p>
              <strong>Email:</strong> johndoe@example.com
            </p>
            <p>
              <strong>Status:</strong> Logged in
            </p>
          </>
        ) : (
          <p>Login required</p>
        )}
      </div>
    </>
  );
}