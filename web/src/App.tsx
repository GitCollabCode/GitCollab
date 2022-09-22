import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './App.css'
import Footer from './components/Footer/Footer'
import LandingPage from './components/LandingPage/LandingPage'
import Login from './components/Login/Login'
import Navbar from './components/Navbar/Navbar'
import Profile from './components/Profile/Profile'
import SignUp from './components/Login/Login'

const App = () => {
  return (
    <Router>
      <div>
        <Navbar />

        <hr />
        <Routes>
          <Route path="" element={<LandingPage />}></Route>
          <Route path="/login" element={<Login />}></Route>
          <Route path="/sign-up" element={<SignUp />}></Route>
          <Route path="/profile" element={<Profile />}></Route>
          <Route path="/dashboard"></Route>
        </Routes>
      </div>
      <Footer />
    </Router>
  )
}

export default App
