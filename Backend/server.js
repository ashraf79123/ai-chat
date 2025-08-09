require ("dotenv").config();
const app=require("./src/app");
const { createServer } = require("http");
const { Server } = require("socket.io");
const generateResponse=require("./src/server/ai.server");
const { text } = require("stream/consumers");
const httpServer = createServer(app);
const io = new Server(httpServer, { 
  cors: {
    origin: "http://localhost:5173", // ✅ remove the trailing slash
    methods: ["GET", "POST"],        // ✅ (optional, but good practice)
    credentials: true                // ✅ if you are using cookies/auth
  }
 });

 const chatHistry=[

 ]
io.on("connection", (socket) => {
  console.log("A user connection")

   socket.on("disconnect",()=>{
    console.log("A user disconnect")
   })
   socket.on("ai-message",async(data)=>{
    chatHistry.push({
      role:"user",
      parts:[{text:data}]
    })
    console.log(data)
    const mama=await generateResponse(chatHistry);
    console.log(mama)
    chatHistry.push({
      role:"model",
      parts:[{text:mama}]
    })
    socket.emit("ai-message-response",mama)
   })
});
httpServer.listen(3000,()=>{
  console.log("server is running port 3000")
})