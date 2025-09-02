import React, { useState } from "react";

const CopyableCodeBlock = ({ inline, className, children, ...props }) => {
  const [copied, setCopied] = useState(false);

  const code = String(children).replace(/\n$/, "");

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(code);
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch (e) {
      setCopied(false);
    }
  };

  if (inline) {
    return <code className={className} {...props}>{children}</code>;
  }

  return (
    <div style={{ position: "relative" }}>
      <button
        onClick={handleCopy}
        style={{
          position: "absolute",
          top: 8,
          right: 8,
          zIndex: 2,
          fontSize: "0.9rem",
          padding: "0.2rem 0.7rem",
          borderRadius: "0.3rem",
          border: "none",
          background: copied ? "#6cb4ff" : "#23272a",
          color: "#fff",
          cursor: "pointer",
        }}
      >
        {copied ? "Copied!" : "Copy"}
      </button>
      <pre className={className} {...props}>
        <code>{code}</code>
      </pre>
    </div>
  );
};

export default CopyableCodeBlock;