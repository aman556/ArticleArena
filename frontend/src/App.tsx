import React from "react";
import Profile from "./feature/Profile/pages/index";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./feature/Login/Login";
import Register from "./feature/Register/Register";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path= '/' element={<Login/>} />;
        <Route path= '/login' element={<Login/>} />;
        <Route path= '/register' element={<Register/>} />;
        <Route path= '/profile' element={<Profile/>} />;
      </Routes>
    </BrowserRouter>
  )
};

export default App;
