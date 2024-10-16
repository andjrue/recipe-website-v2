import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignIn from './components/SignIn'; // Adjust path if needed

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<SignIn />} />
        <Route path="/recipes" element={<div>Recipes Page (Coming Soon)</div>} />
      </Routes>
    </Router>
  );
}

export default App;

