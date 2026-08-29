"use client";
import React, { useEffect, useState } from "react";

type FileRecord = {
  id: number;
  public_id: string;
  url: string;
  filename: string;
  bytes: number;
  created_at: string;
};

const API = process.env.NEXT_PUBLIC_API_URI || "http://127.0.0.1:3500";

export default function Home() {
  const [file, setFile] = useState<File | null>(null);
  const [uploads, setUploads] = useState<FileRecord[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchUploads();
  }, []);

  async function fetchUploads() {
    try {
      const res = await fetch(`${API}/api/uploads`);
      const j = await res.json();
      if (j.success) setUploads(j.data || []);
    } catch (err: any) {
      console.error(err);
    }
  }

  async function handleSubmit(e?: React.FormEvent) {
    e?.preventDefault();
    setError(null);
    if (!file) {
      setError("Please choose a file first");
      return;
    }

    setLoading(true);
    const fd = new FormData();
    fd.append("file", file);

    try {
      const res = await fetch(`${API}/api/upload`, {
        method: "POST",
        body: fd,
      });

      const j = await res.json();
      if (!res.ok) {
        setError(j.message || j.error || "Upload failed");
      } else {
        // refresh uploads and clear file
        await fetchUploads();
        setFile(null);
      }
    } catch (err: any) {
      setError(err.message || "Network error");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div style={{ padding: 24, maxWidth: 900, margin: "0 auto" }}>
      <h1>Upload to Cloudinary</h1>

      <form
        onSubmit={handleSubmit}
        style={{
          display: "flex",
          gap: 12,
          alignItems: "center",
          marginTop: 16,
        }}
      >
        <input
          type="file"
          onChange={(e) => setFile(e.target.files?.[0] ?? null)}
        />
        <button type="submit" disabled={loading}>
          {loading ? "Uploading..." : "Upload"}
        </button>
      </form>

      {error && <div style={{ color: "red", marginTop: 8 }}>{error}</div>}

      <h2 style={{ marginTop: 24 }}>Uploaded</h2>
      <div
        style={{
          display: "grid",
          gap: 12,
          gridTemplateColumns: "repeat(auto-fill, minmax(160px,1fr))",
        }}
      >
        {uploads.length === 0 && <div>No uploads yet</div>}
        {uploads.map((u, i) => (
          <div
            key={i}
            style={{
              padding: 8,
              border: "1px solid #eee",
              borderRadius: 8,
              textAlign: "center",
            }}
          >
            <div style={{ fontSize: 14, color: "#555" }}>{u.filename}</div>
            <a href={u.url} target="_blank" rel="noreferrer">
              <img
                src={u.url}
                alt={u.filename}
                style={{
                  width: "100%",
                  height: 120,
                  objectFit: "cover",
                  borderRadius: 6,
                  marginTop: 8,
                }}
              />
            </a>
            <div style={{ fontSize: 12, color: "#777" }}>
              {(u.bytes / 1024).toFixed(1)} KB
            </div>
            <div style={{ fontSize: 12, color: "#777" }}>
              {new Date(u.created_at).toLocaleString()}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
