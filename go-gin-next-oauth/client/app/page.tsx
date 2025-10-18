"use client";
import { useEffect } from "react";

export default function LoginPage() {
  const handleGoogleLogin = () => {
    window.location.href = `${process.env.NEXT_PUBLIC_API_URI}/api/v1/auth/google/login`;
  };

  return (
    <div className="flex flex-col items-center mt-20">
      <h1 className="text-2xl font-bold mb-6">Login with Google</h1>
      <button
        onClick={handleGoogleLogin}
        className="bg-blue-500 text-white px-4 py-2 rounded"
      >
        Continue with Google
      </button>
    </div>
  );
}
