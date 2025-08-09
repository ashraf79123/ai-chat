// --------------socket.io-----------------

// const app=require("./src/app")
// const { createServer } = require("http");
// const { Server } = require("socket.io");  Yeh Socket.IO se Server class ko import karta hai, jo WebSocket 
                                              // communication ke   liye use hoti hai.
// const httpServer = createServer(app);  Socket.IO ko directly Express app ke saath nahi chalaya ja sakta,
                                             //  isliye httpServer create karte hain.

// io.on("connection", (socket) => {     🔹 Jab bhi koi client WebSocket se connect karega, ye function chalega.
// 🔹                                        socket ek object hai jo client aur server ke beech connection ko represent karta hai.
//   // ...
// });
// httpServer.listen(3000,()=>{
//   console.log("server is running port 3000")
// })

// -----------Socket.IO me HTTP kyu use hota hai?-------
// Jab bhi client aur server real-time chat ya data share karna chahte hain (jaise messaging app, live updates), to unhe ek connection chahiye jo:

// continuously bana rahe (disconnect na ho),

// dono taraf data bhej sake (client → server aur server → client).

// 💡 Lekin...
// 🟢 WebSocket hi ye kaam karta hai.
// 🔴 Lekin WebSocket directly start nahi ho sakta.

// 🧠 Isliye kya hota hai?
// Pehle HTTP se shuru hota hai

// Client server ko ek HTTP request bhejta hai:
// “Hello server, mujhe WebSocket chahiye!”

// Server kehta hai:

// “Theek hai! Ab hum WebSocket se baat karte hain.”

// Fir WebSocket shuru hota hai

// Ab dono taraf real-time data ja sakta hai — bina baar-baar HTTP request ke.

// ✅ To simple bhaasha me:
// HTTP ek darwaza hai, jisse WebSocket room me enter karta hai.


------------------------------------------------------------------------

// Socket.IO ek WebSocket-based library hai — jo WebSocket ko use karke aapko easy aur powerful real-time communication features deta hai.

// ----------------------------------------------------------------------
// socket io kon kon si cheezeino ka use karta hai internally 

// “Socket.IO kaun kaun si cheezein use karta hai?”
// Main simple bhaasha mein batata hoon.

// ✅ Socket.IO internally kya kya use karta hai?
// 1. WebSocket Protocol
// Sabse pehle try karta hai WebSocket connection banana.

// Ye fast, persistent, bidirectional hota hai.

// 2. Fallback Mechanisms
// Agar client ya server WebSocket support nahi karta, to ye fallback karta hai:

// 🔁 Fallback method	Use
// XHR polling (long polling)	Bar bar HTTP request bhejta hai server se update ke liye
// JSONP polling (rare)	Cross-domain fallback (older browsers ke liye)

// * socket io event based communication provide karta hai 
// ----------event based communication-----------------

// Socket.IO uses an event-driven model — meaning:

// 🔁 Client and Server can emit (send) and listen (receive) custom events.

// So instead of just sending raw messages, you send events like:

// "chat message"

// "user joined"

// "typing"

// "Socket.IO uses an event-driven model — meaning:

// 🔁 Client and Server can emit (send) and listen (receive) custom events.

// So instead of just sending raw messages, you send events like:

// "chat message"

// "user joined"

// "typing"

// "

---------------------------------------------------------------------------

io -> io pura server hai 
socket -> ek user/client ko represent karta hai?
*.on -> listen listen karna -> listener ye aye huai request ko sunta hai jo data ata hai usko accpt karta hai
emit=event fire/trigger karna  -> matlab req karna ya route ko hit karna 

*ex -> agr tumhne backend pe listener[on] lagaya hai to fire[emit] frontend se hoga aur agr frontend pe listener laagaya hai to fire backend see hoga

io = Entire Socket.IO server instance

To io represent karta hai:

const io = require('socket.io')(server);

saara server-side communication system

jitne bhi clients connect honge, un sab ko handle karega

aap broadcast kar sakte ho sab clients ko: io.emit(...)

 2. socket kya hota hai?
✅ socket = Ek particular client connection
Jab koi client (browser/app) server se connect hota hai, ek individual socket object ban jata hai uske liye.

Har socket ka unique ID hota hai.

Isse aap ek specific user/client ke sath baat kar sakte ho.

io.on("connection", (socket) => {     // job socket.io se request ayegi to connection banega ye connection ek event hai
  console.log("A user connection")
});

------2 type of event communication-----
1- in-built events  -> two 'connectio', 'disconnect'
2- custom event  -> jitne chahe ham event bana sakhte hai

io.on("connection", (socket) => {   in-built 
  console.log("A user connection")

   socket.on("disconnect",()=>{     in-built 
    console.log("A user disconnect")
   })
   socket.on("message",()=>{           custome event message
    console.log("Message received ")
   })
});

 socket.on("message",(data)=>{          frontend se data a raha hai data me recived hoga
    console.log(data)
   })

socket.emit("ai-message-response", { response }); // This line is using Socket.IO to send a message from the server to the client (or vice versa, depending on where it's written).

socket.emit(eventName, data)
socket.emit: This is how you send a real-time message.

"ai-message-response": This is the event name you're sending. The client/server will listen for this event.

{ response }: This is the data payload being sent with the event — in this case, an object containing a variable response.

commom event communication -
io.on('connection', (socket) => {
  console.log('User connected:', socket.id);

  // Listen for a custom event from client
  socket.on('chat message', (msg) => {
    console.log('Message from client:', msg);

    // Send message back to all clients
    io.emit('chat message', msg);
  });

  // On disconnect
  socket.on('disconnect', () => {
    console.log('User disconnected:', socket.id);
  });
});

-----------------------------------------------------------------------
koi bhi model use karo chahe cgatgpt google gemin by default kisi ke pass memory nhi hoti hai ki message yaad rakh sake 
2 type of memory in ai
1- short term memorr
2 - long term memooory