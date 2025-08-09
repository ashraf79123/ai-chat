// App.js
import { useState, useEffect, useRef } from "react";
import "./App.css";
import { io } from "socket.io-client";

function App() {
  const [socket, setSocket] = useState(null);
  const [message, setMessage] = useState("");
  const [chatHistory, setChatHistory] = useState([]);
  const bottomRef = useRef(null);

  const handleSend = () => {
    if (message.trim() === "") return;
    
    const userMessage = {
      role: "user",
      text: message,
      id: Date.now(),
      timeStamp: new Date().toLocaleTimeString()
    };

    setChatHistory((prev) => [...prev, userMessage]);
    socket.emit("ai-message", message);
    setMessage("");
  };

  const handleKeyPress = (e) => {
    if (e.key === "Enter") {
      handleSend();
    }
  };

  useEffect(() => {
    const socketInstance = io("http://localhost:3000");
    setSocket(socketInstance);

    socketInstance.on("ai-message-response", (response) => {
      const botMessage = {
        id: Date.now() + 1,
        text: response,
        timeStamp: new Date().toLocaleTimeString(),
        role: "bot"
      };
      setChatHistory((prevHistory) => [...prevHistory, botMessage]);
    });

    return () => {
      socketInstance.disconnect();
    };
  }, []);

  // Auto scroll to latest message
  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [chatHistory]);

  return (
    <div className="chat-container">
      <h2 className="chat-header">Chat with AI</h2>
      <div className="chat-history">
        {chatHistory.map((msg, index) => (
          <div key={msg.id || index} className={`chat-message ${msg.role}`}>
            <div className="chat-text">{msg.text}</div>
          </div>
        ))}
        <div ref={bottomRef}></div>
      </div>
      <div className="chat-input-area">
        <input 
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="Type your message..."
          className="input-field"
        />
        <button onClick={handleSend}>Send</button>
      </div>
    </div>
  );
}

export default App;
