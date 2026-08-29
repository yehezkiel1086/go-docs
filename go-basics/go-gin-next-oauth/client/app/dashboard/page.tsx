"use client";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";

export default function DashboardPage() {
  const searchParams = useSearchParams();
  const token = searchParams.get("token");

  useEffect(() => {
    if (token) {
      localStorage.setItem("jwt_token", token);
    }
  }, [token]);

  return (
    <div className="p-6">
      <h1 className="text-xl font-semibold">Welcome to your Dashboard</h1>
    </div>
  );
}
