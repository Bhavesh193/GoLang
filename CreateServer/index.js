
const express = require("express")
const app = express() ;
const port = 3000

app.use(express.json());
app.use(express.urlencoded({extended :true}))

app.get("/" , (req, res) =>{
    res.status(200).send("Welcome the the Golang Server")
})
app.get("/get" , (req, res) =>{
    res.status(200).json( {message : "Welcome the the Golang Server"})
})
app.post("/post" , (req, res) =>{
    let myJson = req.body ; 
    res.status(200).send(myJson)
})
app.get("/postform" , (req, res) =>{
    res.status(200).send(JSON.stringify(req.body))
})
app.listen(port , () => {
    console.log(`Started server at PORT ${port}`);
    
})





