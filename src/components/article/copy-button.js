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
    <div className="modern-code-block">
      <div className="code-block-header">
        <button
          type="button"
          className="modern-icon-copy"
          aria-label="Copy code sample"
          title="Copy code sample"
          onClick={handleCopy}
        >
          <span className="material-icons">content_copy</span>
        </button>
        {copied && <span className="modern-tooltip">Copied!</span>}
      </div>
      <pre className={className} {...props}>
        <code>{code}</code>
      </pre>
    </div>
  );
};

export default CopyableCodeBlock;