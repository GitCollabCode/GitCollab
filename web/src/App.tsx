import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import './App.css'
import Footer from './components/Footer/Footer'
import LandingPage from './pages/LandingPage/LandingPage'
import Navbar from './components/Navbar/Navbar'
import Profile from './pages/Profile/Profile'
import NotFound from './components/Misc/NotFound'
import Modal from './components/Modal/Modal'
import Projects from './pages/Projects/Projects'
import NewProject from './pages/NewProject/NewProject'
const App = () => {
  return (
    <Router>
      <div>
        <Modal></Modal>
        <Navbar />
        <Routes>
          <Route path="" element={<LandingPage />} />
          <Route path="/profile/*" element={<Profile />} />
          <Route path="/projects" element={<Projects />} />
          <Route path="/new-project" element={<NewProject />} />
          {/* star catches any route that is not found */}
          <Route path="/*" element={<NotFound />} />
          <Route path="/dashboard" />
        </Routes>
      </div>
      <Footer />
    </Router>
  )
}

export default App
