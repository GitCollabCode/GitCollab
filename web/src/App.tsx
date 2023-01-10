import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './App.css'
import Footer from './components/Footer/Footer'
import LandingPage from './components/LandingPage/LandingPage'
import Navbar from './components/Navbar/Navbar'
import Profile from './components/Profile/Profile'
import NotFound from './components/Misc/NotFound'
import Modal from './components/Modal/Modal'


const App = () => {
  return (
    
    <Router>
      <div>
        <Modal></Modal>
        <Navbar />
        <Routes>
          <Route path="" element={<LandingPage/>}></Route>
          <Route path="/profile" element={<Profile />}></Route>
          {/* star catches any route that is not found */}
          <Route path="/*" element={<NotFound />}></Route>
          <Route path="/dashboard"></Route>
        </Routes>
      </div>
      <Footer />
    </Router>
  )
}

export default App
