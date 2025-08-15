import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'
import './index.css'
import App from './App.tsx'
import { LoginForm } from './components/ui/custom/Login.tsx'
import { AddProject } from './components/ui/custom/Project.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
     <Routes>
      <Route path="/" element={<App />} />
      <Route path="/login" element={<LoginForm />} />
      <Route path="/projects" element={<AddProject />} />
    {/* <Route path="register" element={<Register />} /> */}
    </Routes>
    </BrowserRouter>
  </StrictMode>,
)
